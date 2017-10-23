package binance

import (
	"context"
	"encoding/json"
)

type CreateOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	orderType        OrderType
	timeInForce      TimeInForce
	quantity         float64
	price            float64
	newClientOrderID *string
	stopPrice        *float64
	icebergQuantity  *float64
}

func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

func (s *CreateOrderService) Side(side SideType) *CreateOrderService {
	s.side = side
	return s
}

func (s *CreateOrderService) Type(orderType OrderType) *CreateOrderService {
	s.orderType = orderType
	return s
}

func (s *CreateOrderService) TimeInForce(timeInForce TimeInForce) *CreateOrderService {
	s.timeInForce = timeInForce
	return s
}

func (s *CreateOrderService) Quantity(quantity float64) *CreateOrderService {
	s.quantity = quantity
	return s
}

func (s *CreateOrderService) Price(price float64) *CreateOrderService {
	s.price = price
	return s
}

func (s *CreateOrderService) NewClientOrderID(newClientOrderID string) *CreateOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

func (s *CreateOrderService) StopPrice(stopPrice float64) *CreateOrderService {
	s.stopPrice = &stopPrice
	return s
}

func (s *CreateOrderService) IcebergQuantity(icebergQuantity float64) *CreateOrderService {
	s.icebergQuantity = &icebergQuantity
	return s
}

func (s *CreateOrderService) createOrder(endpoint string, ctx context.Context, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":      s.symbol,
		"side":        s.side,
		"type":        s.orderType,
		"timeInForce": s.timeInForce,
		"quantity":    s.quantity,
		"price":       s.price,
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
	r.SetFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	return
}

func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, err := s.createOrder("/api/v3/order", ctx, opts...)
	if err != nil {
		return
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return
}

func (s *CreateOrderService) Test(ctx context.Context, opts ...RequestOption) (err error) {
	_, err = s.createOrder("/api/v3/order/test", ctx, opts...)
	return
}

type CreateOrderResponse struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
	TransactTime  int64  `json:"transactTime"`
}

type ListOpenOrdersService struct {
	c      *Client
	symbol string
}

func (s *ListOpenOrdersService) Symbol(symbol string) *ListOpenOrdersService {
	s.symbol = symbol
	return s
}

func (s *ListOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/openOrders",
		secType:  secTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

type GetOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

func (s *GetOrderService) Symbol(symbol string) *GetOrderService {
	s.symbol = symbol
	return s
}

func (s *GetOrderService) OrderID(orderID int64) *GetOrderService {
	s.orderID = &orderID
	return s
}

func (s *GetOrderService) OrigClientOrderID(origClientOrderID string) *GetOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/order",
		secType:  secTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	if s.orderID != nil {
		r.SetParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.SetParam("origClientOrderId", *s.origClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return
}

type Order struct {
	Symbol           string `json:"symbol"`
	OrderID          int64  `json:"orderId"`
	ClientOrderID    string `json:"clientOrderId"`
	Price            string `json:"price"`
	OrigQuantity     string `json:"origQty"`
	ExecutedQuantity string `json:"executedQty"`
	Status           string `json:"status"`
	TimeInForce      string `json:"timeInForce"`
	Type             string `json:"type"`
	Side             string `json:"side"`
	StopPrice        string `json:"stopPrice"`
	IcebergQuantity  string `json:"icebergQty"`
	Time             int64  `json:"time"`
}

type ListOrdersService struct {
	c       *Client
	symbol  string
	orderID *int64
	limit   *int
}

func (s *ListOrdersService) Symbol(symbol string) *ListOrdersService {
	s.symbol = symbol
	return s
}

func (s *ListOrdersService) OrderID(orderID int64) *ListOrdersService {
	s.orderID = &orderID
	return s
}

func (s *ListOrdersService) Limit(limit int) *ListOrdersService {
	s.limit = &limit
	return s
}

func (s *ListOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/allOrders",
		secType:  secTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	if s.orderID != nil {
		r.SetParam("orderId", *s.orderID)
	}
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

type CancelOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	newClientOrderID  *string
}

func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

func (s *CancelOrderService) OrderID(orderID int64) *CancelOrderService {
	s.orderID = &orderID
	return s
}

func (s *CancelOrderService) OrigClientOrderID(origClientOrderID string) *CancelOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

func (s *CancelOrderService) NewClientOrderID(newClientOrderID string) *CancelOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/api/v3/order",
		secType:  secTypeSigned,
	}
	r.SetFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.SetFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.SetFormParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.newClientOrderID != nil {
		r.SetFormParam("newClientOrderId", *s.newClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return
}

type CancelOrderResponse struct {
	Symbol            string `json:"symbol"`
	OrigClientOrderID string `json:"origClientOrderId"`
	OrderID           int64  `json:"orderId"`
	ClientOrderID     string `json:"clientOrderId"`
}
