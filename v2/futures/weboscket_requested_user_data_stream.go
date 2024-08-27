package futures

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	userDataStreamMethodRequest = "REQUEST"

	userDataStreamRequestBalance      = "@balance"
	userDataStreamRequestTypePosition = "@position"
	userDataStreamRequestTypeAccount  = "@account"

	// userDataStreamRateLimit represent number of messages that can be sent per second
	userDataStreamRateLimit = 4
)

type WsRequestedUserDataStreamService struct {
	listenKey string
	apiKey    string
	keepAlive bool

	conn        *websocket.Conn
	connMutex   sync.Mutex
	rateLimiter *time.Ticker

	doneC        chan struct{}
	stopC        chan struct{}
	eventHandler WsUserDataHandler
	errHandler   ErrHandler

	respHandlers      map[int64]userDataStreamResultHandler
	respHandlersMutex sync.Mutex
}

func NewWsRequestedUserDataStreamService(listenKey, apiKey string, keepAlive bool, errHandler ErrHandler) (service *WsRequestedUserDataStreamService, doneC, stopC chan struct{}) {
	service = &WsRequestedUserDataStreamService{
		listenKey:    listenKey,
		apiKey:       apiKey,
		keepAlive:    keepAlive,
		rateLimiter:  time.NewTicker(time.Second / userDataStreamRateLimit),
		doneC:        make(chan struct{}),
		stopC:        make(chan struct{}),
		errHandler:   errHandler,
		respHandlers: make(map[int64]userDataStreamResultHandler),
	}
	doneC = service.doneC
	stopC = service.stopC
	return
}

func (s *WsRequestedUserDataStreamService) APIKey(apiKey string) *WsRequestedUserDataStreamService {
	s.apiKey = apiKey
	return s
}

func (s *WsRequestedUserDataStreamService) KeepAlive(keepAlive bool) *WsRequestedUserDataStreamService {
	s.keepAlive = keepAlive
	return s
}

func (s *WsRequestedUserDataStreamService) ListenKey(listenKey string) *WsRequestedUserDataStreamService {
	s.listenKey = listenKey
	return s
}

func (s *WsRequestedUserDataStreamService) StartUserDataStream(handler WsUserDataHandler) error {
	if err := s.connect(); err != nil {
		return err
	}
	s.eventHandler = handler
	return nil
}

type WsUserDataPositionResponse struct {
	Positions []WsUserDataPosition `json:"positions"`
}

type WsUserDataPosition struct {
	EntryPrice       string `json:"entryPrice"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  bool   `json:"isAutoAddMargin"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Leverage         int    `json:"leverage"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	MaxQty           string `json:"maxQty"`
	Amount           string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnrealizedPnL    string `json:"unRealizedProfit"`
	Side             string `json:"positionSide"`
}
type WsUserDataPositionResponseHandler func(response *WsUserDataPositionResponse)

func (s *WsRequestedUserDataStreamService) RequestUserPosition(handler WsUserDataPositionResponseHandler) error {
	wrappedHandler := userDataStreamResultHandlerWrapper(userDataStreamRequestTypePosition, handler, s.errHandler)
	return s.sendStreamRequest(userDataStreamRequestTypePosition, wrappedHandler)
}

type WsUserDataBalanceResponse struct {
	AccountAlias string              `json:"accountAlias"`
	Balances     []WsUserDataBalance `json:"balances"`
}

type WsUserDataBalance struct {
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
}

type WsUserDataBalanceResponseHandler func(response *WsUserDataBalanceResponse)

func (s *WsRequestedUserDataStreamService) RequestUserBalance(handler WsUserDataBalanceResponseHandler) error {
	wrappedHandler := userDataStreamResultHandlerWrapper(userDataStreamRequestBalance, handler, s.errHandler)
	return s.sendStreamRequest(userDataStreamRequestBalance, wrappedHandler)
}

type WsUserDataBalanceAccountInfoResponse struct {
	FeeTier      int    `json:"feeTier"`
	CanTrade     bool   `json:"canTrade"`
	CanDeposit   bool   `json:"canDeposit"`
	CanWithdraw  bool   `json:"canWithdraw"`
	AccountAlias string `json:"accountAlias"`
}

type WsUserDataBalanceAccountInfoResponseHandler func(response *WsUserDataBalanceAccountInfoResponse)

func (s *WsRequestedUserDataStreamService) RequestAccountInformation(handler WsUserDataBalanceAccountInfoResponseHandler) error {
	wrappedHandler := userDataStreamResultHandlerWrapper(userDataStreamRequestTypeAccount, handler, s.errHandler)
	return s.sendStreamRequest(userDataStreamRequestTypeAccount, wrappedHandler)
}

func (s *WsRequestedUserDataStreamService) connect() error {
	s.connMutex.Lock()
	defer s.connMutex.Unlock()

	if s.conn != nil {
		return fmt.Errorf("websocket already connected")
	}

	endpoint := fmt.Sprintf("%s/%s", getWsEndpoint(), s.listenKey)
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		return fmt.Errorf("unable to dial websocket, endpoint: %s: %w", endpoint, err)
	}
	conn.SetPingHandler(func(pingPayload string) error {
		s.connMutex.Lock()
		defer s.connMutex.Unlock()
		<-s.rateLimiter.C
		return conn.WriteControl(websocket.PongMessage, []byte(pingPayload), time.Now().Add(WebsocketTimeout))
	})
	s.conn = conn

	go s.readMessages()
	return nil
}

func (s *WsRequestedUserDataStreamService) readMessages() {
	silent := false
	go func() {
		select {
		case <-s.stopC:
			silent = true
		case <-s.doneC:
		}
		s.conn.Close()
	}()
	for {
		messageType, message, err := s.conn.ReadMessage()
		if err != nil {
			if !silent {
				s.errHandler(fmt.Errorf("unable to read message: %w", err))
			}
			return
		}
		if messageType == websocket.TextMessage {
			s.handleMessage(message)
		}
	}
}

func (s *WsRequestedUserDataStreamService) handleMessage(message []byte) {
	j, err := newJSON(message)
	if err != nil {
		s.errHandler(fmt.Errorf("unable to unmarshal JSON response: %w", err))
		return
	}

	errorJSON, ok := j.CheckGet("error")
	if ok {
		code, _ := errorJSON.Get("code").Int()
		msg, _ := errorJSON.Get("msg").String()
		s.errHandler(fmt.Errorf("error response: code: %d, message: %s", code, msg))
		return
	}

	// if 'id` is not present, the message is not a response but stream event
	_, ok = j.CheckGet("id")
	if !ok {
		s.handleStreamEvent(message)
	} else {
		s.handleResponse(message)
	}
}

func (s *WsRequestedUserDataStreamService) handleStreamEvent(eventRaw []byte) {
	event := new(WsUserDataEvent)
	if err := json.Unmarshal(eventRaw, event); err != nil {
		s.errHandler(fmt.Errorf("unable to unmarshal stream event: %w", err))
		return
	}
	s.eventHandler(event)
}

func (s *WsRequestedUserDataStreamService) handleResponse(responseRaw []byte) {
	response := new(userDataStreamResponse)
	if err := json.Unmarshal(responseRaw, response); err != nil {
		s.errHandler(fmt.Errorf("unable to unmarshal stream event: %w", err))
		return
	}

	var respHandler userDataStreamResultHandler
	var ok bool

	s.respHandlersMutex.Lock()
	respHandler, ok = s.respHandlers[response.ID]
	delete(s.respHandlers, response.ID)
	s.respHandlersMutex.Unlock()

	if !ok {
		s.errHandler(fmt.Errorf("response handler not found, responseID: %d", response.ID))
		return
	}
	if response.Result == nil || len(response.Result) == 0 {
		s.errHandler(fmt.Errorf("response result is empty, handlerTag: %s, responseID: %d", respHandler.tag, response.ID))
		return
	}
	respHandler.handler(&response.Result[0])
}

type userDataStreamRequest struct {
	ID     int64       `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type userDataStreamResponse struct {
	ID     int64                  `json:"id"`
	Result []userDataStreamResult `json:"result"`
}

type userDataStreamResult struct {
	Req string          `json:"req"`
	Res json.RawMessage `json:"res"`
}

type userDataStreamResultHandler struct {
	tag     string
	handler func(result *userDataStreamResult)
}

func (s *WsRequestedUserDataStreamService) sendStreamRequest(requestType string, handler userDataStreamResultHandler) error {
	if s.eventHandler == nil {
		return fmt.Errorf("data stream is not started")
	}

	endpoint := s.listenKey + requestType
	op := &userDataStreamRequest{ID: time.Now().UnixNano(), Method: userDataStreamMethodRequest, Params: []string{endpoint}}

	s.respHandlersMutex.Lock()
	s.respHandlers[op.ID] = handler
	s.respHandlersMutex.Unlock()

	<-s.rateLimiter.C
	s.connMutex.Lock()
	defer s.connMutex.Unlock()
	return s.conn.WriteJSON(op)
}

func userDataStreamResultHandlerWrapper[T any](tag string, handler func(*T), errHandler ErrHandler) userDataStreamResultHandler {
	return userDataStreamResultHandler{
		tag: tag,
		handler: func(result *userDataStreamResult) {
			if handler == nil {
				return
			}
			response := new(T)
			err := json.Unmarshal(result.Res, &response)
			if err != nil {
				errHandler(fmt.Errorf("unable to unmarshal response: %w", err))
				return
			}
			handler(response)
		},
	}
}
