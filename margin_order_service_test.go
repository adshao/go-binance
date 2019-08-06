package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginOrderServiceTestSuite struct {
	baseOrderTestSuite
}

func TestMarginOrderService(t *testing.T) {
	suite.Run(t, new(marginOrderServiceTestSuite))
}

func (s *marginOrderServiceTestSuite) TestCreateOrder() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"transactTime": 1499827319559,
		"price": "0.0001",
		"origQty": "12.00",
		"executedQty": "10.00",
		"cummulativeQuoteQty": "10.00",
		"status": "FILLED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	quantity := "12.00"
	price := "0.0001"
	newClientOrderID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"timeInForce":      timeInForce,
			"quantity":         quantity,
			"price":            price,
			"newClientOrderId": newClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).
		Price(price).NewClientOrderID(newClientOrderID).Do(newContext())
	s.r().NoError(err)
	e := &CreateOrderResponse{
		Symbol:                   "LTCBTC",
		OrderID:                  1,
		ClientOrderID:            "myOrder1",
		TransactTime:             1499827319559,
		Price:                    "0.0001",
		OrigQuantity:             "12.00",
		ExecutedQuantity:         "10.00",
		CummulativeQuoteQuantity: "10.00",
		Status:      OrderStatusTypeFilled,
		TimeInForce: TimeInForceTypeGTC,
		Type:        OrderTypeLimit,
		Side:        SideTypeBuy,
	}
	s.assertCreateOrderResponseEqual(e, res)
}

func (s *marginOrderServiceTestSuite) TestCreateOrderFull() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"transactTime": 1499827319559,
		"price": "0.0001",
		"origQty": "12.00",
		"executedQty": "10.00",
		"cummulativeQuoteQty": "10.00",
		"status": "FILLED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY",
		"fills": [
			{
				"price":"0.00002991",
				"qty":"344.00000000",
				"commission":"0.00332384",
				"commissionAsset":"BNB",
				"tradeId":1566397
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	quantity := "12.00"
	price := "0.0001"
	newClientOrderID := "myOrder1"
	newOrderRespType := NewOrderRespTypeFULL
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"timeInForce":      timeInForce,
			"quantity":         quantity,
			"price":            price,
			"newClientOrderId": newClientOrderID,
			"newOrderRespType": newOrderRespType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateMarginOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).
		Price(price).NewClientOrderID(newClientOrderID).
		NewOrderRespType(newOrderRespType).Do(newContext())
	s.r().NoError(err)
	e := &CreateOrderResponse{
		Symbol:                   "LTCBTC",
		OrderID:                  1,
		ClientOrderID:            "myOrder1",
		TransactTime:             1499827319559,
		Price:                    "0.0001",
		OrigQuantity:             "12.00",
		ExecutedQuantity:         "10.00",
		CummulativeQuoteQuantity: "10.00",
		Status:      OrderStatusTypeFilled,
		TimeInForce: TimeInForceTypeGTC,
		Type:        OrderTypeLimit,
		Side:        SideTypeBuy,
		Fills: []*Fill{
			&Fill{
				Price:           "0.00002991",
				Quantity:        "344.00000000",
				Commission:      "0.00332384",
				CommissionAsset: "BNB",
			},
		},
	}
	s.assertCreateOrderResponseEqual(e, res)
}

func (s *marginOrderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 28,
		"origClientOrderId": "myOrder1",
		"clientOrderId": "cancelMyOrder1",
		"transactTime": 1507725176595,
		"price": "1.00000000",
		"origQty": "10.00000000",
		"executedQty": "8.00000000",
		"cummulativeQuoteQty": "8.00000000",
		"status": "CANCELED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "SELL"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	orderID := int64(28)
	origClientOrderID := "myOrder1"
	newClientOrderID := "cancelMyOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":            symbol,
			"orderId":           orderID,
			"origClientOrderId": origClientOrderID,
			"newClientOrderId":  newClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelOrderService().Symbol(symbol).
		OrderID(orderID).OrigClientOrderID(origClientOrderID).
		NewClientOrderID(newClientOrderID).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CancelOrderResponse{
		Symbol:                   "LTCBTC",
		OrderID:                  28,
		OrigClientOrderID:        "myOrder1",
		ClientOrderID:            "cancelMyOrder1",
		TransactTime:             1507725176595,
		Price:                    "1.00000000",
		OrigQuantity:             "10.00000000",
		ExecutedQuantity:         "8.00000000",
		CummulativeQuoteQuantity: "8.00000000",
		Status:      OrderStatusTypeCanceled,
		TimeInForce: TimeInForceTypeGTC,
		Type:        OrderTypeLimit,
		Side:        SideTypeSell,
	}
	s.assertCancelOrderResponseEqual(e, res)
}

func (s *marginOrderServiceTestSuite) TestGetOrder() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"price": "0.1",
		"origQty": "1.0",
		"executedQty": "0.0",
		"cummulativeQuoteQty": "0.0",
		"status": "NEW",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY",
		"stopPrice": "0.0",
		"icebergQty": "0.0",
		"time": 1499827319559,
		"updateTime": 1499827319559,
		"isWorking": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	orderID := int64(1)
	origClientOrderID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"orderId":           orderID,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})
	order, err := s.client.NewGetMarginOrderService().Symbol(symbol).
		OrderID(orderID).OrigClientOrderID(origClientOrderID).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &Order{
		Symbol:                   "LTCBTC",
		OrderID:                  1,
		ClientOrderID:            "myOrder1",
		Price:                    "0.1",
		OrigQuantity:             "1.0",
		ExecutedQuantity:         "0.0",
		CummulativeQuoteQuantity: "0.0",
		Status:          OrderStatusTypeNew,
		TimeInForce:     TimeInForceTypeGTC,
		Type:            OrderTypeLimit,
		Side:            SideTypeBuy,
		StopPrice:       "0.0",
		IcebergQuantity: "0.0",
		Time:            1499827319559,
		UpdateTime:      1499827319559,
		IsWorking:       true,
	}
	s.assertOrderEqual(e, order)
}
