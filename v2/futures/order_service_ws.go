package futures

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/adshao/go-binance/v2/common"
)

// WsApiMethodType define method name for websocket API
type WsApiMethodType string

// WsApiRequest define common websocket API request
type WsApiRequest struct {
	Id     string          `json:"id"`
	Method WsApiMethodType `json:"method"`
	Params params          `json:"params"`
}

const (
	// apiKey define key for websocket API parameters
	apiKey = "apiKey"

	// OrderPlaceWsApiMethod define method for creation order via websocket API
	OrderPlaceWsApiMethod WsApiMethodType = "order.place"

	// CancelWsApiMethod define method for cancel order via websocket API
	CancelWsApiMethod WsApiMethodType = "order.cancel"
)

var (
	// ErrorRequestIDNotSet defines that request ID is not set
	ErrorRequestIDNotSet = errors.New("ws service: request id is not set")
)

// OrderPlaceWsService creates order
type OrderPlaceWsService struct {
	c *ClientWs
}

// NewOrderPlaceWsService init OrderPlaceWsService
func NewOrderPlaceWsService(apiKey, secretKey string) (*OrderPlaceWsService, error) {
	client, err := NewClientWs(apiKey, secretKey)
	if err != nil {
		return nil, err
	}

	return &OrderPlaceWsService{c: client}, nil
}

// OrderPlaceWsRequest parameters for 'order.place' websocket API
type OrderPlaceWsRequest struct {
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                string
	reduceOnly              *bool
	price                   *string
	newClientOrderID        *string
	stopPrice               *string
	workingType             *WorkingType
	activationPrice         *string
	callbackRate            *string
	priceProtect            *bool
	newOrderRespType        NewOrderRespType
	closePosition           *bool
	selfTradePreventionMode *string
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

// SelfTradePreventionMode set selfTradePreventionMode
func (s *OrderPlaceWsRequest) SelfTradePreventionMode(selfTradePreventionMode string) *OrderPlaceWsRequest {
	s.selfTradePreventionMode = &selfTradePreventionMode
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
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}

	return m
}

// Do - sends 'order.place' request
func (s *OrderPlaceWsService) Do(requestID string, request *OrderPlaceWsRequest) error {
	if requestID == "" {
		return ErrorRequestIDNotSet
	}

	params := request.buildParams()
	params[apiKey] = s.c.APIKey
	params[timestampKey] = currentTimestamp() - s.c.TimeOffset

	signature, err := getSignature(s.c.SecretKey, params)
	if err != nil {
		return err
	}
	params[signatureKey] = signature

	req := WsApiRequest{
		Id:     requestID,
		Method: OrderPlaceWsApiMethod,
		Params: params,
	}

	rawData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	if err := s.c.Write(requestID, rawData); err != nil {
		return err
	}

	return nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *OrderPlaceWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderPlaceWsService) GetReadChannel() <-chan []byte {
	return s.c.readC
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderPlaceWsService) GetReadErrorChannel() <-chan error {
	return s.c.readErrChan
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderPlaceWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

// getSignature creates signature for params
func getSignature(secretKey string, params params) (string, error) {
	queryValues := url.Values{}
	for key, value := range params {
		queryValues.Add(key, fmt.Sprintf("%v", value))
	}
	queryString := queryValues.Encode()

	raw := fmt.Sprintf("%s", queryString)
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(raw))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", mac.Sum(nil)), nil
}

// NewCancelOrderRequest init CancelOrderRequest
func NewCancelOrderRequest() *CancelOrderRequest {
	return &CancelOrderRequest{}
}

// CancelOrderRequest parameters for 'order.cancel' websocket API
type CancelOrderRequest struct {
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *CancelOrderRequest) Symbol(symbol string) *CancelOrderRequest {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelOrderRequest) OrderID(orderID int64) *CancelOrderRequest {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelOrderRequest) OrigClientOrderID(origClientOrderID string) *CancelOrderRequest {
	s.origClientOrderID = &origClientOrderID
	return s
}

// buildParams builds params
func (s *CancelOrderRequest) buildParams() params {
	m := params{
		"symbol": s.symbol,
	}

	if s.orderID != nil {
		m["orderId"] = *s.orderID
	}

	if s.origClientOrderID != nil {
		m["origClientOrderId"] = *s.origClientOrderID
	}

	return m
}

// CancelOrderResult define order cancel result
type CancelOrderResult struct {
	CancelOrderResponse
}

// CancelOrderWsResponse define 'order.cancel' websocket API response
type CancelOrderWsResponse struct {
	Id     string            `json:"id"`
	Status int               `json:"status"`
	Result CancelOrderResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}

// OrderCancelWsService cancel order
type OrderCancelWsService struct {
	c *ClientWs
}

// NewOrderCancelWsService init OrderCancelWsService
func NewOrderCancelWsService(apiKey, secretKey string) (*OrderCancelWsService, error) {
	client, err := NewClientWs(apiKey, secretKey)
	if err != nil {
		return nil, err
	}

	return &OrderCancelWsService{c: client}, nil
}

// Do - sends 'order.cancel' request
func (s *OrderCancelWsService) Do(requestID string, request *CancelOrderRequest) error {
	if requestID == "" {
		return ErrorRequestIDNotSet
	}

	params := request.buildParams()
	params[apiKey] = s.c.APIKey
	params[timestampKey] = currentTimestamp() - s.c.TimeOffset

	signature, err := getSignature(s.c.SecretKey, params)
	if err != nil {
		return err
	}
	params[signatureKey] = signature

	req := WsApiRequest{
		Id:     requestID,
		Method: CancelWsApiMethod,
		Params: params,
	}

	rawData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	if err := s.c.Write(requestID, rawData); err != nil {
		return err
	}

	return nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *OrderCancelWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderCancelWsService) GetReadChannel() <-chan []byte {
	return s.c.readC
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderCancelWsService) GetReadErrorChannel() <-chan error {
	return s.c.readErrChan
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderCancelWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}
