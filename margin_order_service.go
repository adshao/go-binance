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
func (s *CancelMarginOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
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
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
