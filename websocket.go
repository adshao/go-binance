package binance

import (
	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

type wsConfig struct {
	endpoint string
}

func newWsConfig(endpoint string) *wsConfig {
	return &wsConfig{
		endpoint: endpoint,
	}
}

var wsServe = func(cfg *wsConfig, handler WsHandler, errHandler ErrHandler) (done chan struct{}, err error) {
	c, _, err := websocket.DefaultDialer.Dial(cfg.endpoint, nil)
	if err != nil {
		return nil, err
	}
	done = make(chan struct{})
	go func() {
		defer func() {
			cerr := c.Close()
			if cerr != nil {
				errHandler(cerr)
			}
		}()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				go errHandler(err)
				return
			}
			go handler(message)
		}
	}()
	return
}
