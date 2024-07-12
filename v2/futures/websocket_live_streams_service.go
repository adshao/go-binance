package futures

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"

	"github.com/aiviaio/go-binance/v2/common"
)

type WsLiveStreamsService struct {
	conn      *websocket.Conn
	connMutex sync.Mutex

	doneC           chan struct{}
	errHandler      ErrHandler
	wsHandlersMutex sync.Mutex
	wsHandlers      map[string]WsHandler

	ops      map[int64]*liveStreamOp
	opsMutex sync.Mutex
}

type liveStreamOp struct {
	ID     int64    `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"params"`
	// stopC is used for common.LiveMethodSubscribe to which a signal is sent when the stream should be stopped.
	// If stopC was closed, nothing will happen.
	// If stopC is nil, the stream represents a common.LiveMethodUnsubscribe operation.
	stopC chan struct{}
}

func NewWsLiveStreamsService(errHandler ErrHandler) (service *WsLiveStreamsService, doneC chan struct{}) {
	doneC = make(chan struct{})
	service = &WsLiveStreamsService{
		doneC:      doneC,
		errHandler: errHandler,
		wsHandlers: make(map[string]WsHandler),
		ops:        make(map[int64]*liveStreamOp),
	}
	return
}

// WsAllMarketsStatServe serve websocket that push 24hr statistics for all market every second
func (s *WsLiveStreamsService) WsAllMarketsStatServe(handler WsAllMarketTickerHandler) (stopC chan<- struct{}, err error) {
	if err = s.connectIfNotConnected(); err != nil {
		return
	}
	wsHandler := func(message []byte) {
		var event WsAllMarketTickerEvent
		err = json.Unmarshal(message, &event)
		if err != nil {
			s.errHandler(fmt.Errorf("unable to unmarshal event: %w", err))
			return
		}
		handler(event)
	}
	if err = s.addWsHandler(common.LiveStreamAllMarketTickers, wsHandler); err != nil {
		return
	}

	return s.subscribe(common.LiveStreamAllMarketTickers)
}

// WsAggTradeServe serve websocket aggregate handler with a symbol
func (s *WsLiveStreamsService) WsAggTradeServe(symbol string, handler WsAggTradeHandler) (stopC chan<- struct{}, err error) {
	if err = s.connectIfNotConnected(); err != nil {
		return
	}
	wsHandler := wsHandlerWrapper(handler, s.errHandler)
	streamName := strings.ToLower(symbol) + common.LiveStreamAggTrade
	if err = s.addWsHandler(streamName, wsHandler); err != nil {
		return
	}

	return s.subscribe(streamName)
}

// WsDepthServe serve websocket depth handler with a symbol, using 1sec updates
func (s *WsLiveStreamsService) WsDepthServe(symbol string, handler WsDepthHandler) (stopC chan<- struct{}, err error) {
	if err = s.connectIfNotConnected(); err != nil {
		return
	}
	streamName := strings.ToLower(symbol) + common.LiveStreamDepth
	wsHandler := wsDepthHandlerWrapper(handler, s.errHandler)
	if err = s.addWsHandler(streamName, wsHandler); err != nil {
		return
	}

	return s.subscribe(streamName)
}

// WsAvgPriceServe serve websocket that pushes 1m average price for a specified symbol.
func (s *WsLiveStreamsService) WsAvgPriceServe(symbol string, handler WsAvgPriceHandler) (stopC chan<- struct{}, err error) {
	if err = s.connectIfNotConnected(); err != nil {
		return
	}
	streamName := strings.ToLower(symbol) + common.LiveStreamAvgPrice
	wsHandler := wsHandlerWrapper(handler, s.errHandler)
	if err = s.addWsHandler(streamName, wsHandler); err != nil {
		return
	}

	return s.subscribe(streamName)
}

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func (s *WsLiveStreamsService) WsKlineServe(symbol string, interval string, handler WsKlineHandler) (stopC chan struct{}, err error) {
	if err = s.connectIfNotConnected(); err != nil {
		return
	}
	streamName := strings.ToLower(symbol) + common.LiveStreamKline + interval
	wsHandler := wsHandlerWrapper(handler, s.errHandler)
	if err = s.addWsHandler(streamName, wsHandler); err != nil {
		return
	}

	return s.subscribe(streamName)
}

// WsMarkPriceServe serve websocket that pushes price and funding rate for a single symbol.
func (s *WsLiveStreamsService) WsMarkPriceServe(symbol string, handler WsMarkPriceHandler) (stopC chan struct{}, err error) {
	if err = s.connectIfNotConnected(); err != nil {
		return
	}
	streamName := strings.ToLower(symbol) + common.LiveStreamMarkPrice
	wsHandler := wsHandlerWrapper(handler, s.errHandler)
	if err = s.addWsHandler(streamName, wsHandler); err != nil {
		return
	}

	return s.subscribe(streamName)
}

// subscribe sends a subscription message to the WebSocket
func (s *WsLiveStreamsService) subscribe(streams ...string) (stopC chan struct{}, err error) {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()

	stopC = make(chan struct{})
	op := &liveStreamOp{ID: time.Now().UnixNano(), Method: common.LiveMethodSubscribe, Params: streams, stopC: stopC}
	s.opsMutex.Lock()
	s.ops[op.ID] = op
	s.opsMutex.Unlock()
	go s.waitForStop(op)
	if err = s.conn.WriteJSON(op); err != nil {
		return
	}

	return
}

// unsubscribe sends an unsubscription message to the WebSocket
func (s *WsLiveStreamsService) unsubscribe(streams ...string) error {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()

	op := &liveStreamOp{ID: time.Now().UnixMilli(), Method: common.LiveMethodUnsubscribe, Params: streams}
	s.opsMutex.Lock()
	s.ops[op.ID] = op
	s.opsMutex.Unlock()
	return s.conn.WriteJSON(op)
}

// readMessages listens for incoming messages and handles them
func (s *WsLiveStreamsService) readMessages() {
	defer s.conn.Close()

	for {
		select {
		case <-s.doneC:
			return
		default:
			messageType, message, err := s.conn.ReadMessage()
			if err != nil {
				s.errHandler(fmt.Errorf("unable to read message: %w", err))
				return
			}
			switch messageType {
			case websocket.TextMessage:
				s.handleMessage(message)
			case websocket.PongMessage:
				// ignore
			}
		}
	}
}

// handleMessage processes incoming WebSocket messages
func (s *WsLiveStreamsService) handleMessage(message []byte) {
	j, err := newJSON(message)
	if err != nil {
		s.errHandler(fmt.Errorf("unable to unmarshal JSON response: %w", err))
		return
	}

	// Check message type and process accordingly
	if errorMsg, ok := j.CheckGet("error"); ok {
		// if 'error' is not nil, then the server has sent an error message
		id, _ := j.Get("id").Int64()
		s.opsMutex.Lock()
		op := s.ops[id]
		if op != nil && op.stopC != nil {
			close(op.stopC) // close the stop channel to prevent the operation from waiting indefinitely
		}
		delete(s.ops, id) // delete the operation from the map as it was not successful
		s.opsMutex.Unlock()
		if op != nil {
			s.errHandler(fmt.Errorf("operation was not successful, opMethod: %s, opParams: %v, error: %v",
				op.Method, op.Params, errorMsg))
		} else {
			s.errHandler(fmt.Errorf("error was received: %v", errorMsg))
		}
	} else if _, ok := j.CheckGet("result"); ok {
		// if 'result' is nil, then the server has sent a message about successful operation
		id, _ := j.Get("id").Int64()
		s.opsMutex.Lock()
		if op, ok := s.ops[id]; ok && op.stopC == nil {
			delete(s.ops, id) // delete the operation from the map as it was successful, and we don't need it anymore
		}
		s.opsMutex.Unlock()
	} else if streamRaw, ok := j.CheckGet("stream"); ok {
		// if 'stream' is not nil, then the server has sent a stream message
		s.handlerStreamMessage(streamRaw.MustString(), j.Get("data"))
	} else {
		// if none of the above, then the server has sent an unexpected message
		s.errHandler(fmt.Errorf("unexpected message: %v", j))
	}
}

func (s *WsLiveStreamsService) connectIfNotConnected() error {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()
	if s.conn == nil {
		err := s.connect()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *WsLiveStreamsService) connect() error {
	endpoint := strings.TrimSuffix(getCombinedEndpoint(), "?streams=")
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		return fmt.Errorf("unable to dial websocket, endpoint: %s: %w", endpoint, err)
	}
	s.conn = conn

	go s.readMessages()
	return nil
}

func (s *WsLiveStreamsService) addWsHandler(streamName string, handler WsHandler) error {
	s.wsHandlersMutex.Lock()
	defer s.wsHandlersMutex.Unlock()
	if _, ok := s.wsHandlers[streamName]; ok {
		return fmt.Errorf("stream %s already connected", streamName)
	}
	s.wsHandlers[streamName] = handler
	return nil
}

func (s *WsLiveStreamsService) handlerStreamMessage(stream string, data *simplejson.Json) {
	s.wsHandlersMutex.Lock()
	defer s.wsHandlersMutex.Unlock()
	if handler, ok := s.wsHandlers[stream]; ok {
		encodedData, err := data.Encode()
		if err != nil {
			s.errHandler(fmt.Errorf("unable to encode data: %w", err))
			return
		}
		handler(encodedData)
	} else {
		s.errHandler(fmt.Errorf("handler for stream %s not found", stream))
	}
}

func (s *WsLiveStreamsService) waitForStop(op *liveStreamOp) {
	defer close(op.stopC)
	if _, ok := <-op.stopC; ok {
		if err := s.unsubscribe(op.Params...); err != nil {
			s.errHandler(fmt.Errorf("unable to unsubscribe: %w", err))
		}
	}
}

func wsHandlerWrapper[T any](handler func(t *T), errHandler ErrHandler) WsHandler {
	return func(message []byte) {
		event := new(T)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
}
