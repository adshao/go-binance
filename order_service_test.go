package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type orderServiceTestSuite struct {
	baseTestSuite
}

func TestOrderService(t *testing.T) {
	suite.Run(t, new(orderServiceTestSuite))
}

func (s *orderServiceTestSuite) TestCreateOrder() {
	data := []byte(`{
        "symbol":"LTCBTC",
        "orderId": 1,
        "clientOrderId": "myOrder1",
        "transactTime": 1499827319559
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	timeInForce := TimeInForceGTC
	quantity := "12.00"
	price := "0.0001"
	newClientOrderID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().SetFormParams(params{
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
		Symbol:        "LTCBTC",
		OrderID:       1,
		ClientOrderID: "myOrder1",
		TransactTime:  1499827319559,
	}
	s.assertCreateOrderResponseEqual(e, res)

	err = s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).
		Price(price).NewClientOrderID(newClientOrderID).Test(newContext())
	s.r().NoError(err)
}

func (s *orderServiceTestSuite) assertCreateOrderResponseEqual(e, a *CreateOrderResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.TransactTime, a.TransactTime, "TransactTime")
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
            "time": 1499827319559
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().SetParams(params{
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
		Status:           "NEW",
		TimeInForce:      "GTC",
		Type:             "LIMIT",
		Side:             "BUY",
		StopPrice:        "0.0",
		IcebergQuantity:  "0.0",
		Time:             1499827319559,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *orderServiceTestSuite) assertOrderEqual(e, a *Order) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.IcebergQuantity, a.IcebergQuantity, "IcebergQuantity")
	r.Equal(e.Time, e.Time, "Time")
}

func (s *orderServiceTestSuite) TestGetOrder() {
	data := []byte(`{
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
        "time": 1499827319559
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	orderID := int64(1)
	origClientOrderID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().SetParams(params{
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
		Symbol:           "LTCBTC",
		OrderID:          1,
		ClientOrderID:    "myOrder1",
		Price:            "0.1",
		OrigQuantity:     "1.0",
		ExecutedQuantity: "0.0",
		Status:           "NEW",
		TimeInForce:      "GTC",
		Type:             "LIMIT",
		Side:             "BUY",
		StopPrice:        "0.0",
		IcebergQuantity:  "0.0",
		Time:             1499827319559,
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
            "time": 1499827319559
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	orderID := int64(1)
	limit := 3
	s.assertReq(func(r *request) {
		e := newSignedRequest().SetParams(params{
			"symbol":  symbol,
			"orderId": orderID,
			"limit":   limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewListOrdersService().Symbol(symbol).
		OrderID(orderID).Limit(limit).Do(newContext())
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
		Status:           "NEW",
		TimeInForce:      "GTC",
		Type:             "LIMIT",
		Side:             "BUY",
		StopPrice:        "0.0",
		IcebergQuantity:  "0.0",
		Time:             1499827319559,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *orderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
        "symbol": "LTCBTC",
        "origClientOrderId": "myOrder1",
        "orderId": 1,
        "clientOrderId": "cancelMyOrder1"
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	orderID := int64(1)
	origClientOrderID := "myOrder1"
	newClientOrderID := "cancelMyOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().SetFormParams(params{
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
		Symbol:            "LTCBTC",
		OrderID:           1,
		OrigClientOrderID: "myOrder1",
		ClientOrderID:     "cancelMyOrder1",
	}
	s.assertCancelOrderResponseEqual(e, res)
}

func (s *orderServiceTestSuite) assertCancelOrderResponseEqual(e, a *CancelOrderResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrigClientOrderID, a.OrigClientOrderID, "OrigClientOrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
}
