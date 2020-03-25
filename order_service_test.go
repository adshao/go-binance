package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseOrderTestSuite struct {
	baseTestSuite
}

type orderServiceTestSuite struct {
	baseOrderTestSuite
}

func TestOrderService(t *testing.T) {
	suite.Run(t, new(orderServiceTestSuite))
}

func (s *orderServiceTestSuite) TestCreateOrder() {
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
	quoteOrderQty := "10.00"
	price := "0.0001"
	newClientOrderID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"timeInForce":      timeInForce,
			"quantity":         quantity,
			"quoteOrderQty":    quoteOrderQty,
			"price":            price,
			"newClientOrderId": newClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
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
		Status:                   OrderStatusTypeFilled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeBuy,
	}
	s.assertCreateOrderResponseEqual(e, res)

	err = s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
		Price(price).NewClientOrderID(newClientOrderID).Test(newContext())
	s.r().NoError(err)
}

func (s *orderServiceTestSuite) TestCreateOrderFull() {
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
	quoteOrderQty := "10.00"
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
			"quoteOrderQty":    quoteOrderQty,
			"price":            price,
			"newClientOrderId": newClientOrderID,
			"newOrderRespType": newOrderRespType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
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

	err = s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
		Price(price).NewClientOrderID(newClientOrderID).
		NewOrderRespType(newOrderRespType).Test(newContext())
	s.r().NoError(err)
}

func (s *baseOrderTestSuite) assertCreateOrderResponseEqual(e, a *CreateOrderResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.TransactTime, a.TransactTime, "TransactTime")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")

	r.Len(a.Fills, len(e.Fills))
	for idx, fill := range e.Fills {
		s.assertFillEqual(fill, a.Fills[idx])
	}
}

func (s *baseOrderTestSuite) assertFillEqual(e, a *Fill) {
	r := s.r()
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
}

func (s *orderServiceTestSuite) TestCreateOCO() {
	data := []byte(`{
		"orderListId": 0,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "C3wyj4WVEktd7u9aVBRXcN",
		"transactionTime": 1574040868128,
		"symbol": "LTCBTC",
		"orders": [
		  {
			"symbol": "LTCBTC",
			"orderId": 2,
			"clientOrderId": "pO9ufTiFGg3nw2fOdgeOXa"
		  },
		  {
			"symbol": "LTCBTC",
			"orderId": 3,
			"clientOrderId": "TXOvglzXuaubXAaENpaRCB"
		  }
		],
		"orderReports": [
		  {
			"symbol": "LTCBTC",
			"origClientOrderId": "pO9ufTiFGg3nw2fOdgeOXa",
			"orderId": 2,
			"orderListId": 0,
			"clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
			"price": "1.00000000",
			"origQty": "10.00000000",
			"executedQty": "0.00000000",
			"cummulativeQuoteQty": "0.00000000",
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "STOP_LOSS",
			"side": "SELL",
			"stopPrice": "1.00000000"
		  },
		  {
			"symbol": "LTCBTC",
			"origClientOrderId": "TXOvglzXuaubXAaENpaRCB",
			"orderId": 3,
			"orderListId": 0,
			"clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
			"price": "3.00000000",
			"origQty": "10.00000000",
			"executedQty": "0.00000000",
			"cummulativeQuoteQty": "0.00000000",
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "LIMIT_MAKER",
			"side": "SELL"
		  }
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	timeInForce := TimeInForceTypeGTC
	quantity := "10"
	price := "3"
	stopPrice := "3.1"
	stopLimitPrice := "3.2"
	limitClientOrderID := "myOrder1"
	newOrderRespType := NewOrderRespTypeFULL
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":               symbol,
			"side":                 side,
			"quantity":             quantity,
			"price":                price,
			"stopPrice":            stopPrice,
			"stopLimitPrice":       stopLimitPrice,
			"stopLimitTimeInForce": timeInForce,
			"limitClientOrderId":   limitClientOrderID,
			"newOrderRespType":     newOrderRespType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOCOService().
		Symbol(symbol).
		Side(side).
		Quantity(quantity).
		Price(price).
		StopPrice(stopPrice).
		StopLimitPrice(stopLimitPrice).
		StopLimitTimeInForce(timeInForce).
		LimitClientOrderID(limitClientOrderID).
		NewOrderRespType(newOrderRespType).
		Do(newContext())

	s.r().NoError(err)
	e := &CreateOCOResponse{
		OrderListID:       0,
		ContingencyType:   "OCO",
		ListStatusType:    "EXEC_STARTED",
		ListOrderStatus:   "EXECUTING",
		ListClientOrderID: "C3wyj4WVEktd7u9aVBRXcN",
		TransactionTime:   1574040868128,
		Symbol:            "LTCBTC",
		Orders: []*OCOOrder{
			&OCOOrder{
				Symbol:        "LTCBTC",
				OrderID:       2,
				ClientOrderID: "pO9ufTiFGg3nw2fOdgeOXa",
			},
			&OCOOrder{
				Symbol:        "LTCBTC",
				OrderID:       3,
				ClientOrderID: "TXOvglzXuaubXAaENpaRCB",
			},
		},
		OrderReports: []*OCOOrderReport{
			&OCOOrderReport{
				Symbol:                   "LTCBTC",
				OrderID:                  2,
				OrderListID:              0,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
				Price:                    "1.00000000",
				OrigQuantity:             "10.00000000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeNew,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeStopLoss,
				Side:                     SideTypeSell,
				StopPrice:                "1.00000000",
			},
			&OCOOrderReport{
				Symbol:                   "LTCBTC",
				OrderID:                  3,
				OrderListID:              0,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
				Price:                    "3.00000000",
				OrigQuantity:             "10.00000000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeNew,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeLimitMaker,
				Side:                     SideTypeSell,
			},
		},
	}
	s.assertCreateOCOResponseEqual(e, res)
}

func (s *baseOrderTestSuite) assertCreateOCOResponseEqual(e, a *CreateOCOResponse) {
	r := s.r()
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")

	r.Len(a.OrderReports, len(e.OrderReports))
	for idx, orderReport := range e.OrderReports {
		s.assertOCOOrderReportEqual(orderReport, a.OrderReports[idx])
	}

	r.Len(a.Orders, len(e.Orders))
	for idx, order := range e.Orders {
		s.assertOCOOrderEqual(order, a.Orders[idx])
	}
}

func (s *baseOrderTestSuite) assertOCOOrderReportEqual(e, a *OCOOrderReport) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	// r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
}

func (s *baseOrderTestSuite) assertOCOOrderEqual(e, a *OCOOrder) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}

func (s *orderServiceTestSuite) TestListOpenOrders() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "orderId": 1,
            "clientOrderId": "myOrder1",
            "price": "0.1",
            "origQty": "1.0",
            "executedQty": "0.0",
            "status": "NEW",
            "timeInForce": "GTC",
            "type": "LIMIT",
            "side": "BUY",
            "stopPrice": "0.0",
            "icebergQty": "0.0",
			"time": 1499827319559,
			"updateTime": 1499827319559,
    		"isWorking": true
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	orders, err := s.client.NewListOpenOrdersService().Symbol(symbol).
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &Order{
		Symbol:           "LTCBTC",
		OrderID:          1,
		ClientOrderID:    "myOrder1",
		Price:            "0.1",
		OrigQuantity:     "1.0",
		ExecutedQuantity: "0.0",
		Status:           OrderStatusTypeNew,
		TimeInForce:      TimeInForceTypeGTC,
		Type:             OrderTypeLimit,
		Side:             SideTypeBuy,
		StopPrice:        "0.0",
		IcebergQuantity:  "0.0",
		Time:             1499827319559,
		UpdateTime:       1499827319559,
		IsWorking:        true,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *baseOrderTestSuite) assertOrderEqual(e, a *Order) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.IcebergQuantity, a.IcebergQuantity, "IcebergQuantity")
	r.Equal(e.Time, e.Time, "Time")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.IsWorking, a.IsWorking, "IsWorking")
}

func (s *orderServiceTestSuite) TestGetOrder() {
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
	order, err := s.client.NewGetOrderService().Symbol(symbol).
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

func (s *orderServiceTestSuite) TestListOrders() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "orderId": 1,
            "clientOrderId": "myOrder1",
            "price": "0.1",
            "origQty": "1.0",
            "executedQty": "0.0",
            "status": "NEW",
            "timeInForce": "GTC",
            "type": "LIMIT",
            "side": "BUY",
            "stopPrice": "0.0",
            "icebergQty": "0.0",
			"time": 1499827319559,
			"updateTime": 1499827319559,
			"isWorking": true
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	orderID := int64(1)
	limit := 3
	startTime := int64(1499827319559)
	endTime := int64(1499827319560)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"orderId":   orderID,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewListOrdersService().Symbol(symbol).
		OrderID(orderID).StartTime(startTime).EndTime(endTime).
		Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &Order{
		Symbol:           "LTCBTC",
		OrderID:          1,
		ClientOrderID:    "myOrder1",
		Price:            "0.1",
		OrigQuantity:     "1.0",
		ExecutedQuantity: "0.0",
		Status:           OrderStatusTypeNew,
		TimeInForce:      TimeInForceTypeGTC,
		Type:             OrderTypeLimit,
		Side:             SideTypeBuy,
		StopPrice:        "0.0",
		IcebergQuantity:  "0.0",
		Time:             1499827319559,
		UpdateTime:       1499827319559,
		IsWorking:        true,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *orderServiceTestSuite) TestCancelOrder() {
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

func (s *baseOrderTestSuite) assertCancelOrderResponseEqual(e, a *CancelOrderResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrigClientOrderID, a.OrigClientOrderID, "OrigClientOrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.TransactTime, a.TransactTime, "TransactTime")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
}
