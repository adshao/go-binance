package binance

import (
	"context"
	"encoding/json"
)

// CreateMarginOrderService create order
type CreateMarginOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	orderType        OrderType
	quantity         string
	price            *string
	stopPrice        *string
	newClientOrderID *string
	icebergQuantity  *string
	newOrderRespType *NewOrderRespType
	sideEffectType   *SideEffectType
	timeInForce      *TimeInForceType
}

// Symbol set symbol
func (s *CreateMarginOrderService) Symbol(symbol string) *CreateMarginOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateMarginOrderService) Side(side SideType) *CreateMarginOrderService {
	s.side = side
	return s
}

// Type set type
func (s *CreateMarginOrderService) Type(orderType OrderType) *CreateMarginOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateMarginOrderService) TimeInForce(timeInForce TimeInForceType) *CreateMarginOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateMarginOrderService) Quantity(quantity string) *CreateMarginOrderService {
	s.quantity = quantity
	return s
}

// Price set price
func (s *CreateMarginOrderService) Price(price string) *CreateMarginOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateMarginOrderService) NewClientOrderID(newClientOrderID string) *CreateMarginOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *CreateMarginOrderService) StopPrice(stopPrice string) *CreateMarginOrderService {
	s.stopPrice = &stopPrice
	return s
}

// IcebergQuantity set icebergQuantity
func (s *CreateMarginOrderService) IcebergQuantity(icebergQuantity string) *CreateMarginOrderService {
	s.icebergQuantity = &icebergQuantity
	return s
}

// NewOrderRespType set icebergQuantity
func (s *CreateMarginOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateMarginOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SideEffectType set sideEffectType
func (s *CreateMarginOrderService) SideEffectType(sideEffectType SideEffectType) *CreateMarginOrderService {
	s.sideEffectType = &sideEffectType
	return s
}

// Do send request
func (s *CreateMarginOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/margin/order",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"type":     s.orderType,
		"quantity": s.quantity,
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
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
	if s.icebergQuantity != nil {
		m["icebergQty"] = *s.icebergQuantity
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.sideEffectType != nil {
		m["sideEffectType"] = *s.sideEffectType
	}
	r.setFormParams(m)
	res = new(CreateOrderResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelMarginOrderService cancel an order
type CancelMarginOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	newClientOrderID  *string
}

// Symbol set symbol
func (s *CancelMarginOrderService) Symbol(symbol string) *CancelMarginOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelMarginOrderService) OrderID(orderID int64) *CancelMarginOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelMarginOrderService) OrigClientOrderID(origClientOrderID string) *CancelMarginOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CancelMarginOrderService) NewClientOrderID(newClientOrderID string) *CancelMarginOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// Do send request
func (s *CancelMarginOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelMarginOrderResponse, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/sapi/v1/margin/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.newClientOrderID != nil {
		r.setFormParam("newClientOrderId", *s.newClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelMarginOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMarginOrderService get an order
type GetMarginOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *GetMarginOrderService) Symbol(symbol string) *GetMarginOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetMarginOrderService) OrderID(orderID int64) *GetMarginOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetMarginOrderService) OrigClientOrderID(origClientOrderID string) *GetMarginOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetMarginOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListMarginOpenOrdersService list margin open orders
type ListMarginOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *ListMarginOpenOrdersService) Symbol(symbol string) *ListMarginOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *ListMarginOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Order{}, err
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Order{}, err
	}
	return res, nil
}

// ListMarginOrdersService all account orders; active, canceled, or filled
type ListMarginOrdersService struct {
	c         *Client
	symbol    string
	orderID   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListMarginOrdersService) Symbol(symbol string) *ListMarginOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *ListMarginOrdersService) OrderID(orderID int64) *ListMarginOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set starttime
func (s *ListMarginOrdersService) StartTime(startTime int64) *ListMarginOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListMarginOrdersService) EndTime(endTime int64) *ListMarginOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListMarginOrdersService) Limit(limit int) *ListMarginOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListMarginOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAllOrder, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/allOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAllOrder{}, err
	}
	res = make([]*MarginAllOrder, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAllOrder{}, err
	}
	return res, nil
}

// MarginAllOrder define item of margin all orders
type MarginAllOrder struct {
	ID            int64  `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"qty"`
	QuoteQuantity string `json:"quoteQty"`
	Symbol        string `json:"symbol"`
	Time          int64  `json:"time"`
}

// CancelMarginOrderResponse define response of canceling order
type CancelMarginOrderResponse struct {
	Symbol                   string          `json:"symbol"`
	OrigClientOrderID        string          `json:"origClientOrderId"`
	OrderID                  string          `json:"orderId"`
	ClientOrderID            string          `json:"clientOrderId"`
	TransactTime             int64           `json:"transactTime"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
}
