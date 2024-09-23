package futures

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
)

// OrderPlaceWsService creates order
type OrderPlaceWsService struct {
	c      wsClient
	signFn func(string, string) (*string, error)
}

// NewOrderPlaceWsService init OrderPlaceWsService
func NewOrderPlaceWsService(apiKey, secretKey string) (*OrderPlaceWsService, error) {
	conn, err := newConnection()
	if err != nil {
		return nil, err
	}

	client, err := NewClientWs(conn, apiKey, secretKey)
	if err != nil {
		return nil, err
	}

	return &OrderPlaceWsService{c: client}, nil
}

// OrderPlaceWsRequest parameters for 'order.place' websocket API
type OrderPlaceWsRequest struct {
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         string
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	stopPrice        *string
	workingType      *WorkingType
	activationPrice  *string
	callbackRate     *string
	priceProtect     *bool
	newOrderRespType NewOrderRespType
	closePosition    *bool
}

// NewOrderPlaceWsRequest init OrderPlaceWsRequest
func NewOrderPlaceWsRequest() *OrderPlaceWsRequest {
	return &OrderPlaceWsRequest{}
}

// Symbol set symbol
func (s *OrderPlaceWsRequest) Symbol(symbol string) *OrderPlaceWsRequest {
	s.symbol = symbol
	return s
}

// Side set side
func (s *OrderPlaceWsRequest) Side(side SideType) *OrderPlaceWsRequest {
	s.side = side
	return s
}

// PositionSide set side
func (s *OrderPlaceWsRequest) PositionSide(positionSide PositionSideType) *OrderPlaceWsRequest {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *OrderPlaceWsRequest) Type(orderType OrderType) *OrderPlaceWsRequest {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *OrderPlaceWsRequest) TimeInForce(timeInForce TimeInForceType) *OrderPlaceWsRequest {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *OrderPlaceWsRequest) Quantity(quantity string) *OrderPlaceWsRequest {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *OrderPlaceWsRequest) ReduceOnly(reduceOnly bool) *OrderPlaceWsRequest {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *OrderPlaceWsRequest) Price(price string) *OrderPlaceWsRequest {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *OrderPlaceWsRequest) NewClientOrderID(newClientOrderID string) *OrderPlaceWsRequest {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *OrderPlaceWsRequest) StopPrice(stopPrice string) *OrderPlaceWsRequest {
	s.stopPrice = &stopPrice
	return s
}

// WorkingType set workingType
func (s *OrderPlaceWsRequest) WorkingType(workingType WorkingType) *OrderPlaceWsRequest {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *OrderPlaceWsRequest) ActivationPrice(activationPrice string) *OrderPlaceWsRequest {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *OrderPlaceWsRequest) CallbackRate(callbackRate string) *OrderPlaceWsRequest {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *OrderPlaceWsRequest) PriceProtect(priceProtect bool) *OrderPlaceWsRequest {
	s.priceProtect = &priceProtect
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *OrderPlaceWsRequest) NewOrderResponseType(newOrderResponseType NewOrderRespType) *OrderPlaceWsRequest {
	s.newOrderRespType = newOrderResponseType
	return s
}

// ClosePosition set closePosition
func (s *OrderPlaceWsRequest) ClosePosition(closePosition bool) *OrderPlaceWsRequest {
	s.closePosition = &closePosition
	return s
}

// CreateOrderResult define order creation result
type CreateOrderResult struct {
	CreateOrderResponse
}

// CreateOrderWsResponse define 'order.place' websocket API response
type CreateOrderWsResponse struct {
	Id     string            `json:"id"`
	Status int               `json:"status"`
	Result CreateOrderResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}

// buildParams builds params
func (s *OrderPlaceWsRequest) buildParams() params {
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.quantity != "" {
		m["quantity"] = s.quantity
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.closePosition != nil {
		m["closePosition"] = *s.closePosition
	}

	return m
}

// Do - sends 'order.place' request
func (s *OrderPlaceWsService) Do(requestID string, request *OrderPlaceWsRequest) error {
	rawData, err := createWsRequest(requestID, s.c, OrderPlaceWsApiMethod, request.buildParams())
	if err != nil {
		return err
	}

	if err := s.c.Write(requestID, rawData); err != nil {
		return err
	}

	return nil
}

// SyncDo - sends 'order.place' request and receives response
func (s *OrderPlaceWsService) SyncDo(requestID string, request *OrderPlaceWsRequest) (*CreateOrderWsResponse, error) {
	rawData, err := createWsRequest(requestID, s.c, OrderPlaceWsApiMethod, request.buildParams())
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	createOrderWsResponse := &CreateOrderWsResponse{}
	if err := json.Unmarshal(response, createOrderWsResponse); err != nil {
		return nil, err
	}

	return createOrderWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *OrderPlaceWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderPlaceWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderPlaceWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderPlaceWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}
