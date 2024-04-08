package options

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateOrderService create order
type CreateOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	orderType        OrderType
	quantity         string
	price            *string
	timeInForce      *TimeInForceType
	reduceOnly       *bool
	postOnly         *bool
	newOrderRespType NewOrderRespType
	clientOrderID    *string
	isMmp            *bool
}

// Symbol set symbol
func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateOrderService) Side(side SideType) *CreateOrderService {
	s.side = side
	return s
}

// Type set type
func (s *CreateOrderService) Type(orderType OrderType) *CreateOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateOrderService) TimeInForce(timeInForce TimeInForceType) *CreateOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateOrderService) Quantity(quantity string) *CreateOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateOrderService) ReduceOnly(reduceOnly bool) *CreateOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// PostOnly set postOnly
func (s *CreateOrderService) PostOnly(postOnly bool) *CreateOrderService {
	s.postOnly = &postOnly
	return s
}

// Price set price
func (s *CreateOrderService) Price(price string) *CreateOrderService {
	s.price = &price
	return s
}

// ClientOrderID set clientOrderID
func (s *CreateOrderService) ClientOrderID(ClientOrderID string) *CreateOrderService {
	s.clientOrderID = &ClientOrderID
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// IsMmp set isMmp
func (s *CreateOrderService) IsMmp(isMmp bool) *CreateOrderService {
	s.isMmp = &isMmp
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {

	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.postOnly != nil {
		m["postOnly"] = *s.postOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.clientOrderID != nil {
		m["clientOrderId"] = *s.clientOrderID
	}
	if s.isMmp != nil {
		m["isMmp"] = *s.isMmp
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/eapi/v1/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOrderResponse define create order response
type CreateOrderResponse struct {
	OrderID           int64           `json:"orderId"`
	Symbol            string          `json:"symbol"`
	Price             string          `json:"price"`
	Quantity          string          `json:"quantity"`
	ExecutedQty       string          `json:"executedQty"`
	Fee               string          `json:"fee"`
	Side              SideType        `json:"side"`
	Type              OrderType       `json:"type"`
	TimeInForce       TimeInForceType `json:"timeInForce"`
	ReduceOnly        bool            `json:"reduceOnly"`
	PostOnly          bool            `json:"postOnly"`
	CreateTime        int64           `json:"createTime"`
	UpdateTime        int64           `json:"updateTime"`
	Status            OrderStatusType `json:"status"`
	AvgPrice          string          `json:"avgPrice"`
	ClientOrderID     string          `json:"clientOrderId"`
	PriceScale        int             `json:"priceScale"`
	QuantityScale     int             `json:"quantityScale"`
	OptionSide        OptionSideType  `json:"optionSide"`
	QuoteAsset        string          `json:"quoteAsset"`
	Mmp               bool            `json:"mmp"`
	RateLimitOrder10s string          `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string          `json:"rateLimitOrder1m,omitempty"`
}

// ListOpenOrdersService list opened orders
type ListOpenOrdersService struct {
	c         *Client
	symbol    string
	orderId   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListOpenOrdersService) Symbol(symbol string) *ListOpenOrdersService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *ListOpenOrdersService) OrderId(orderId int64) *ListOpenOrdersService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *ListOpenOrdersService) StartTime(startTime int64) *ListOpenOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListOpenOrdersService) EndTime(endTime int64) *ListOpenOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListOpenOrdersService) Limit(limit int) *ListOpenOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	if s.orderId != nil {
		r.setParam("orderId", s.orderId)
	}
	if s.startTime != nil {
		r.setParam("startTime", s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", s.limit)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
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

// GetOrderService get an order
type GetOrderService struct {
	c             *Client
	symbol        string
	orderID       *int64
	clientOrderID *string
}

// Symbol set symbol
func (s *GetOrderService) Symbol(symbol string) *GetOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetOrderService) OrderID(orderID int64) *GetOrderService {
	s.orderID = &orderID
	return s
}

// ClientOrderID set clientOrderID
func (s *GetOrderService) ClientOrderID(clientOrderID string) *GetOrderService {
	s.clientOrderID = &clientOrderID
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.clientOrderID != nil {
		r.setParam("clientOrderID", *s.clientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
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

// Order define order info
type Order struct {
	OrderID       int64           `json:"orderId"`
	Symbol        string          `json:"symbol"`
	Price         string          `json:"price"`
	Quantity      string          `json:"quantity"`
	ExecutedQty   string          `json:"executedQty"`
	Fee           string          `json:"fee"`
	Side          SideType        `json:"side"`
	Type          OrderType       `json:"type"`
	TimeInForce   TimeInForceType `json:"timeInForce"`
	ReduceOnly    bool            `json:"reduceOnly"`
	PostOnly      bool            `json:"postOnly"`
	CreateTime    int64           `json:"createTime"`
	UpdateTime    int64           `json:"updateTime"`
	Status        OrderStatusType `json:"status"`
	AvgPrice      string          `json:"avgPrice"`
	Source        string          `json:"source"`
	ClientOrderID string          `json:"clientOrderId"`
	PriceScale    int             `json:"priceScale"`
	QuantityScale int             `json:"quantityScale"`
	OptionSide    OptionSideType  `json:"optionSide"`
	QuoteAsset    string          `json:"quoteAsset"`
	Mmp           bool            `json:"mmp"`
}

// CancelOrderService cancel an order
type CancelOrderService struct {
	c             *Client
	symbol        string
	orderID       *int64
	clientOrderID *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelOrderService) OrderID(orderID int64) *CancelOrderService {
	s.orderID = &orderID
	return s
}

// ClientOrderID set clientOrderID
func (s *CancelOrderService) ClientOrderID(clientOrderID string) *CancelOrderService {
	s.clientOrderID = &clientOrderID
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/eapi/v1/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.clientOrderID != nil {
		r.setFormParam("clientOrderID", *s.clientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelOrderResponse define response of canceling order
type CancelOrderResponse struct {
	OrderID       int64           `json:"orderId"`
	Symbol        string          `json:"symbol"`
	Price         string          `json:"price"`
	Quantity      string          `json:"quantity"`
	ExecutedQty   string          `json:"executedQty"`
	Fee           string          `json:"fee"`
	Side          SideType        `json:"side"`
	Type          OrderType       `json:"type"`
	TimeInForce   TimeInForceType `json:"timeInForce"`
	ReduceOnly    bool            `json:"reduceOnly"`
	PostOnly      bool            `json:"postOnly"`
	CreateTime    int64           `json:"createTime"`
	UpdateTime    int64           `json:"updateTime"`
	Status        OrderStatusType `json:"status"`
	AvgPrice      string          `json:"avgPrice"`
	Source        string          `json:"source"`
	ClientOrderID string          `json:"clientOrderId"`
	PriceScale    int             `json:"priceScale"`
	QuantityScale int             `json:"quantityScale"`
	OptionSide    OptionSideType  `json:"optionSide"`
	QuoteAsset    string          `json:"quoteAsset"`
	Mmp           bool            `json:"mmp"`
}

// CancelAllOpenOrdersService cancel all open orders
type CancelAllOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelAllOpenOrdersService) Symbol(symbol string) *CancelAllOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/eapi/v1/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CancelMultiplesOrdersService cancel a list of orders
type CancelMultiplesOrdersService struct {
	c                 *Client
	symbol            string
	orderIDList       []int64
	clientOrderIDList []string
}

// CancelSingleOrderResponse define element of response of canceling multiple order
type CancelSingleOrderResponse struct {
	OrderID       int64           `json:"orderId"`
	Symbol        string          `json:"symbol"`
	Price         string          `json:"price"`
	Quantity      string          `json:"quantity"`
	ExecutedQty   string          `json:"executedQty"`
	Fee           string          `json:"fee"`
	Side          SideType        `json:"side"`
	Type          OrderType       `json:"type"`
	TimeInForce   TimeInForceType `json:"timeInForce"`
	CreateTime    int64           `json:"createTime"`
	Status        OrderStatusType `json:"status"`
	AvgPrice      string          `json:"avgPrice"`
	ReduceOnly    bool            `json:"reduceOnly"`
	ClientOrderID string          `json:"clientOrderId"`
	UpdateTime    int64           `json:"updateTime"`
}

// Symbol set symbol
func (s *CancelMultiplesOrdersService) Symbol(symbol string) *CancelMultiplesOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelMultiplesOrdersService) OrderIDList(orderIDList []int64) *CancelMultiplesOrdersService {
	s.orderIDList = orderIDList
	return s
}

// ClientOrderIDList set clientOrderIDList
func (s *CancelMultiplesOrdersService) ClientOrderIDList(clientOrderIDList []string) *CancelMultiplesOrdersService {
	s.clientOrderIDList = clientOrderIDList
	return s
}

// Do send request
func (s *CancelMultiplesOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*CancelSingleOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/eapi/v1/batchOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderIDList != nil {
		// convert a slice of integers to a string e.g. [1 2 3] => "[1,2,3]"
		orderIDListString := strings.Join(strings.Fields(fmt.Sprint(s.orderIDList)), ",")
		r.setFormParam("orderIdList", orderIDListString)
	}
	if s.clientOrderIDList != nil {
		r.setFormParam("clientOrderIdList", s.clientOrderIDList)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*CancelSingleOrderResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CancelSingleOrderResponse{}, err
	}
	return res, nil
}

type CreateBatchOrdersService struct {
	c      *Client
	orders []*CreateOrderService
}

type CreateBatchOrdersResponse struct {
	Orders []*SingleOrder
}

// SingleOrder define single order info returned by batch create request
type SingleOrder struct {
	OrderID       int64     `json:"orderId"`
	Symbol        string    `json:"symbol"`
	Price         string    `json:"price"`
	Quantity      string    `json:"quantity"`
	Side          SideType  `json:"side"`
	Type          OrderType `json:"type"`
	ReduceOnly    bool      `json:"reduceOnly"`
	PostOnly      bool      `json:"postOnly"`
	ClientOrderID string    `json:"clientOrderId"`
	Mmp           bool      `json:"mmp"`
}

func (s *CreateBatchOrdersService) OrderList(orders []*CreateOrderService) *CreateBatchOrdersService {
	s.orders = orders
	return s
}

func (s *CreateBatchOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *CreateBatchOrdersResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/eapi/v1/batchOrders",
		secType:  secTypeSigned,
	}

	orders := []params{}
	for _, order := range s.orders {
		m := params{
			"symbol":           order.symbol,
			"side":             order.side,
			"type":             order.orderType,
			"quantity":         order.quantity,
			"newOrderRespType": order.newOrderRespType,
		}

		if order.timeInForce != nil {
			m["timeInForce"] = *order.timeInForce
		}
		if order.reduceOnly != nil {
			m["reduceOnly"] = *order.reduceOnly
		}
		if order.postOnly != nil {
			m["postOnly"] = *order.postOnly
		}
		if order.price != nil {
			m["price"] = *order.price
		}
		if order.clientOrderID != nil {
			m["clientOrderId"] = *order.clientOrderID
		}
		if order.isMmp != nil {
			m["isMmp"] = *order.isMmp
		}
		orders = append(orders, m)
	}
	b, err := json.Marshal(orders)
	if err != nil {
		return &CreateBatchOrdersResponse{}, err
	}
	m := params{
		"batchOrders": string(b),
	}

	r.setFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return &CreateBatchOrdersResponse{}, err
	}

	rawMessages := make([]*json.RawMessage, 0)

	err = json.Unmarshal(data, &rawMessages)

	if err != nil {
		return &CreateBatchOrdersResponse{}, err
	}

	batchCreateOrdersResponse := new(CreateBatchOrdersResponse)

	for _, j := range rawMessages {
		o := new(SingleOrder)
		if err := json.Unmarshal(*j, o); err != nil {
			return &CreateBatchOrdersResponse{}, err
		}

		/*
			TODO: the following 4 lines are copied from futures package, not sure why there is
			such condition check and it looks wrong to me. Need to double confirm.

			if o.ClientOrderID != "" {
				batchCreateOrdersResponse.Orders = append(batchCreateOrdersResponse.Orders, o)
				continue
			}
		*/
		batchCreateOrdersResponse.Orders = append(batchCreateOrdersResponse.Orders, o)

	}

	return batchCreateOrdersResponse, nil

}
