package futures

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
)

const (
	// reconnectMinInterval define reconnect min interval
	reconnectMinInterval = 100 * time.Millisecond

	// reconnectMaxInterval define reconnect max interval
	reconnectMaxInterval = 10 * time.Second
)

var (
	// ErrorWsConnectionClosed defines that connection closed
	ErrorWsConnectionClosed = errors.New("ws error: connection closed")

	// ErrorWsIdAlreadySent defines that request with the same id was already sent
	ErrorWsIdAlreadySent = errors.New("ws error: request with same id already sent")
)

// messageId define id field of request/response
type messageId struct {
	Id string `json:"id"`
}

// ClientWs define API websocket client
type ClientWs struct {
	APIKey                      string
	SecretKey                   string
	Debug                       bool
	Logger                      *log.Logger
	Conn                        *websocket.Conn
	TimeOffset                  int64
	mu                          sync.Mutex
	reconnectSignal             chan struct{}
	connectionEstablishedSignal chan struct{}
	requestsList                RequestList
	readC                       chan []byte
	readErrChan                 chan error
	reconnectCount              atomic.Int64
}

func (c *ClientWs) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Println(fmt.Sprintf(format, v...))
	}
}

// NewClientWs init ClientWs
func NewClientWs(apiKey, secretKey string) (*ClientWs, error) {
	conn, err := WsApiInitReadWriteConn()
	if err != nil {
		return nil, err
	}

	client := &ClientWs{
		APIKey:                      apiKey,
		SecretKey:                   secretKey,
		Logger:                      log.New(os.Stderr, "Binance-golang ", log.LstdFlags),
		Conn:                        conn,
		mu:                          sync.Mutex{},
		reconnectSignal:             make(chan struct{}, 1),
		connectionEstablishedSignal: make(chan struct{}, 1),
		requestsList:                NewRequestList(),
		readErrChan:                 make(chan error),
		readC:                       make(chan []byte),
	}

	go client.handleReconnect()
	go client.read()

	return client, nil
}

// Write sends data into websocket connection
func (c *ClientWs) Write(id string, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.requestsList.IsAlreadyInList(id) {
		return ErrorWsIdAlreadySent
	}

	if err := c.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
		c.debug("write: unable to write message into websocket conn '%v'", err)
		return err
	}

	c.requestsList.Add(id)

	return nil
}

// read data from connection
func (c *ClientWs) read() {
	defer func() {
		// reading from closed connection 1000 times caused panic
		// prevent panic for any case
		if r := recover(); r != nil {
		}
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			c.debug("read: error reading message '%v'")
			c.reconnectSignal <- struct{}{}
			c.readErrChan <- errors.Join(err, ErrorWsConnectionClosed)

			c.debug("read: wait to get connected")
			<-c.connectionEstablishedSignal

			c.debug("read: connection established")
			continue
		}

		msg := messageId{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			c.readErrChan <- err
			continue
		}

		c.debug("read: got new message")
		c.readC <- message
		c.requestsList.Remove(msg.Id)
	}
}

// wait until all responses received
// make sure that you are not sending requests
func (c *ClientWs) wait(timeout time.Duration) {
	doneC := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if c.requestsList.Len() == 0 {
					doneC <- struct{}{}
					return
				}
			}
		}
	}()

	t := time.After(timeout)
	select {
	case <-t:
	case <-doneC:
	}

	cancel()
}

// handleReconnect waits for reconnect signal and starts reconnect
func (c *ClientWs) handleReconnect() {
	for _ = range c.reconnectSignal {
		c.debug("reconnect: received signal")

		b := &backoff.Backoff{
			Min:    reconnectMinInterval,
			Max:    reconnectMaxInterval,
			Factor: 1.8,
			Jitter: false,
		}

		conn := c.startReconnect(b)

		b.Reset()

		c.mu.Lock()
		c.Conn = conn
		c.mu.Unlock()

		c.debug("reconnect: connected")
		c.connectionEstablishedSignal <- struct{}{}
	}
}

// startReconnect starts reconnect loop with increasing delay
func (c *ClientWs) startReconnect(b *backoff.Backoff) *websocket.Conn {
	for {
		c.reconnectCount.Add(1)
		conn, err := WsApiInitReadWriteConn()
		if err != nil {
			delay := b.Duration()
			c.debug("reconnect: error while reconnecting. try in %s", delay.Round(time.Millisecond))
			time.Sleep(delay)
			continue
		}

		return conn
	}
}

// GetReconnectCount returns reconnect counter value (useful for metrics outside)
func (c *ClientWs) GetReconnectCount() int64 {
	return c.reconnectCount.Load()
}

// NewRequestList creates request list
func NewRequestList() RequestList {
	return RequestList{
		mu:       sync.Mutex{},
		requests: make(map[string]struct{}), // TODO preallocate buckets
	}
}

// RequestList state of requests that were sent/received
type RequestList struct {
	mu       sync.Mutex
	requests map[string]struct{}
}

// Add adds request into list
func (l *RequestList) Add(id string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.requests[id] = struct{}{}
}

// RecreateList creates new request list
func (l *RequestList) RecreateList() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.requests = make(map[string]struct{})
}

// Remove adds request from list
func (l *RequestList) Remove(id string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.requests, id)
}

// Len get list length
func (l *RequestList) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.requests)
}

// IsAlreadyInList checks if id is presented in requests list
func (l *RequestList) IsAlreadyInList(id string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.requests[id]; ok {
		return true
	}

	return false
}
