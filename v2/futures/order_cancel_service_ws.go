package futures

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
)

// NewOrderCancelRequest init OrderCancelRequest
func NewOrderCancelRequest() *OrderCancelRequest {
	return &OrderCancelRequest{}
}

// OrderCancelRequest parameters for 'order.cancel' websocket API
type OrderCancelRequest struct {
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *OrderCancelRequest) Symbol(symbol string) *OrderCancelRequest {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *OrderCancelRequest) OrderID(orderID int64) *OrderCancelRequest {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *OrderCancelRequest) OrigClientOrderID(origClientOrderID string) *OrderCancelRequest {
	s.origClientOrderID = &origClientOrderID
	return s
}

// buildParams builds params
func (s *OrderCancelRequest) buildParams() params {
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

// OrderCancelWsResponse define 'order.cancel' websocket API response
type OrderCancelWsResponse struct {
	Id     string            `json:"id"`
	Status int               `json:"status"`
	Result CancelOrderResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}

// OrderCancelWsService cancel order
type OrderCancelWsService struct {
	c wsClient
}

// Do - sends 'order.cancel' request
func (s *OrderCancelWsService) Do(requestID string, request *OrderCancelRequest) error {
	rawData, err := createWsRequest(requestID, s.c, CancelWsApiMethod, request.buildParams())
	if err != nil {
		return err
	}

	if err := s.c.Write(requestID, rawData); err != nil {
		return err
	}

	return nil
}

// SyncDo - sends 'order.cancel' request and receives response
func (s *OrderCancelWsService) SyncDo(requestID string, request *OrderCancelRequest) (*OrderCancelWsResponse, error) {
	rawData, err := createWsRequest(requestID, s.c, CancelWsApiMethod, request.buildParams())
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	cancelOrderWsResponse := &OrderCancelWsResponse{}
	if err := json.Unmarshal(response, cancelOrderWsResponse); err != nil {
		return nil, err
	}

	return cancelOrderWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *OrderCancelWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderCancelWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderCancelWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderCancelWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}
