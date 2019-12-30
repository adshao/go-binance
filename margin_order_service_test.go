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
			"sideEffectType":   SideEffectTypeNoSideEffect,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateMarginOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).
		Price(price).NewClientOrderID(newClientOrderID).SideEffectType(SideEffectTypeNoSideEffect).
		Do(newContext())
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
		Status:                   OrderStatusTypeFilled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeBuy,
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
		Status:                   OrderStatusTypeFilled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeBuy,
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
		Status:                   OrderStatusTypeCanceled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeSell,
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
		Status:                   OrderStatusTypeNew,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeBuy,
		StopPrice:                "0.0",
		IcebergQuantity:          "0.0",
		Time:                     1499827319559,
		UpdateTime:               1499827319559,
		IsWorking:                true,
	}
	s.assertOrderEqual(e, order)
}

func (s *marginOrderServiceTestSuite) TestListMarginOpenOrders() {
	data := []byte(`[
		{
			"clientOrderId": "qhcZw71gAkCCTv0t0k8LUK",
			"cummulativeQuoteQty": "0.00000000",
			"executedQty": "0.00000000",
			"icebergQty": "0.00000000",
			"isWorking": true,
			"orderId": 211842552,
			"origQty": "0.30000000",
			"price": "0.00475010",
			"side": "SELL",
			"status": "NEW",
			"stopPrice": "0.00000000",
			"symbol": "BNBBTC",
			"time": 1562040170089,
			"timeInForce": "GTC",
			"type": "LIMIT",
			"updateTime": 1562040170089
		   }
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	orders, err := s.client.NewListMarginOpenOrdersService().Symbol(symbol).
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &Order{
		Symbol:                   "BNBBTC",
		OrderID:                  211842552,
		ClientOrderID:            "qhcZw71gAkCCTv0t0k8LUK",
		CummulativeQuoteQuantity: "0.00000000",
		Price:                    "0.00475010",
		OrigQuantity:             "0.30000000",
		ExecutedQuantity:         "0.00000000",
		Status:                   OrderStatusTypeNew,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeSell,
		StopPrice:                "0.00000000",
		IcebergQuantity:          "0.00000000",
		Time:                     1562040170089,
		UpdateTime:               1562040170089,
		IsWorking:                true,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *marginOrderServiceTestSuite) TestListMarginOrders() {
	data := []byte(`[
		{
			"id": 43123876,
			"price": "0.00395740",
			"qty": "4.06000000",
			"quoteQty": "0.01606704",
			"symbol": "BNBBTC",
			"time": 1556089977693
		},
		{
			"id": 43123877,
			"price": "0.00395740",
			"qty": "0.77000000",
			"quoteQty": "0.00304719",
			"symbol": "BNBBTC",
			"time": 1556089977693
		},
		{
			"id": 43253549,
			"price": "0.00428930",
			"qty": "23.30000000",
			"quoteQty": "0.09994069",
			"symbol": "BNBBTC",
			"time": 1556163963504
		}
  	]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BNBBTC"
	limit := 3
	startTime := int64(1556089977693)
	endTime := int64(1556163963504)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewListMarginOrdersService().Symbol(symbol).
		StartTime(startTime).EndTime(endTime).
		Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(orders, 3)
	e := []*MarginAllOrder{
		{
			ID:            43123876,
			Price:         "0.00395740",
			Quantity:      "4.06000000",
			QuoteQuantity: "0.01606704",
			Symbol:        "BNBBTC",
			Time:          1556089977693,
		},
		{
			ID:            43123877,
			Price:         "0.00395740",
			Quantity:      "0.77000000",
			QuoteQuantity: "0.00304719",
			Symbol:        "BNBBTC",
			Time:          1556089977693,
		},
		{
			ID:            43253549,
			Price:         "0.00428930",
			Quantity:      "23.30000000",
			QuoteQuantity: "0.09994069",
			Symbol:        "BNBBTC",
			Time:          1556163963504,
		},
	}
	s.r().Len(orders, len(e))
	for i := 0; i < len(orders); i++ {
		s.assertMarginAllOrderEqual(e[i], orders[i])
	}
}

func (s *marginOrderServiceTestSuite) assertMarginAllOrderEqual(e, a *MarginAllOrder) {
	r := s.r()
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.QuoteQuantity, a.QuoteQuantity, "QuoteQuantity")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Time, a.Time, "Time")
}
