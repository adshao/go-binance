package options

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
		"orderId": 4611875134427365377,
		"symbol": "BTC-200730-9000-C", 
		"price": "100",                
		"quantity": "1",               
		"executedQty": "0",            
		"fee": "0",                    
		"side": "BUY",                 
		"type": "LIMIT",               
		"timeInForce": "GTC",          
		"reduceOnly": false,           
		"postOnly": false,             
		"createTime": 1592465880683,   
		"updateTime": 1566818724722,   
		"status": "ACCEPTED",          
		"avgPrice": "0",               
		"clientOrderId": "testOrder",           
		"priceScale": 2,
		"quantityScale": 2,
		"optionSide": "CALL",
		"quoteAsset": "USDT",
		"mmp": false
	  }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTC-200730-9000-C"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	quantity := "1"
	price := "100"
	timeInForce := TimeInForceTypeGTC
	reduceOnly := false
	postOnly := false
	newOrderResponseType := NewOrderRespTypeRESULT
	clientOrderID := "testOrder"
	isMmp := false
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"quantity":         quantity,
			"price":            price,
			"timeInForce":      timeInForce,
			"reduceOnly":       reduceOnly,
			"postOnly":         postOnly,
			"newOrderRespType": newOrderResponseType,
			"clientOrderId":    clientOrderID,
			"isMmp":            isMmp,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).Price(price).TimeInForce(timeInForce).
		ReduceOnly(reduceOnly).PostOnly(postOnly).NewOrderResponseType(newOrderResponseType).
		ClientOrderID(clientOrderID).IsMmp(isMmp).
		Do(newContext())
	s.r().NoError(err)
	e := &CreateOrderResponse{
		OrderID:       4611875134427365377,
		Symbol:        symbol,
		Price:         price,
		Quantity:      quantity,
		ExecutedQty:   "0",
		Fee:           "0",
		Side:          side,
		Type:          orderType,
		TimeInForce:   timeInForce,
		ReduceOnly:    reduceOnly,
		PostOnly:      postOnly,
		CreateTime:    1592465880683,
		UpdateTime:    1566818724722,
		Status:        OrderStatusTypeAccepted,
		AvgPrice:      "0",
		ClientOrderID: clientOrderID,
		PriceScale:    2,
		QuantityScale: 2,
		OptionSide:    OptionSideTypeCall,
		QuoteAsset:    "USDT",
		Mmp:           isMmp,
	}
	s.assertCreateOrderResponseEqual(e, res)
}

func (s *baseOrderTestSuite) assertCreateOrderResponseEqual(e, a *CreateOrderResponse) {
	r := s.r()
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.ExecutedQty, a.ExecutedQty, "ExecutedQty")
	r.Equal(e.Fee, a.Fee, "Fee")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.PostOnly, a.PostOnly, "PostOnly")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.AvgPrice, a.AvgPrice, "AvgPrice")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.PriceScale, a.PriceScale, "PriceScale")
	r.Equal(e.QuantityScale, a.QuantityScale, "QuantityScale")
	r.Equal(e.OptionSide, a.OptionSide, "OptionSide")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.Mmp, a.Mmp, "Mmp")
}

func (s *orderServiceTestSuite) TestListOpenOrders() {
	data := []byte(`[
		{
		  "orderId": 4611875134427365377,
		  "symbol": "BTC-200730-9000-C",
		  "price": "100",
		  "quantity": "1",
		  "executedQty": "0",
		  "fee": "0",
		  "side": "BUY",
		  "type": "LIMIT",
		  "timeInForce": "GTC",
		  "reduceOnly": false,
		  "postOnly": false,
		  "createTime": 1592465880683,
		  "updateTime": 1592465880683,
		  "status": "ACCEPTED",
		  "avgPrice": "0",
		  "clientOrderId": "",
		  "priceScale": 2,
		  "quantityScale": 2,
		  "optionSide": "CALL",
		  "quoteAsset": "USDT",
		  "mmp": false
		}
	  ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTC-200730-9000-C"
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
		OrderID:       4611875134427365377,
		Symbol:        symbol,
		Price:         "100",
		Quantity:      "1",
		ExecutedQty:   "0",
		Fee:           "0",
		Side:          SideTypeBuy,
		Type:          OrderTypeLimit,
		TimeInForce:   TimeInForceTypeGTC,
		ReduceOnly:    false,
		PostOnly:      false,
		CreateTime:    1592465880683,
		UpdateTime:    1592465880683,
		Status:        OrderStatusTypeAccepted,
		AvgPrice:      "0",
		ClientOrderID: "",
		PriceScale:    2,
		QuantityScale: 2,
		OptionSide:    OptionSideTypeCall,
		QuoteAsset:    "USDT",
		Mmp:           false,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *baseOrderTestSuite) assertOrderEqual(e, a *Order) {
	r := s.r()
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.ExecutedQty, a.ExecutedQty, "ExecutedQty")
	r.Equal(e.Fee, a.Fee, "Fee")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.PostOnly, a.PostOnly, "PostOnly")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.AvgPrice, a.AvgPrice, "AvgPrice")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.PriceScale, a.PriceScale, "PriceScale")
	r.Equal(e.QuantityScale, a.QuantityScale, "QuantityScale")
	r.Equal(e.OptionSide, a.OptionSide, "OptionSide")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.Mmp, a.Mmp, "Mmp")
}

func (s *orderServiceTestSuite) TestGetOrder() {
	data := []byte(`{
		"orderId": 4611875134427365377,
		"symbol": "BTC-200730-9000-C",
		"price": "100",
		"quantity": "1",
		"executedQty": "0",
		"fee": "0",
		"side": "BUY",
		"type": "LIMIT",
		"timeInForce": "GTC",
		"reduceOnly": false,
		"postOnly": false,
		"createTime": 1592465880683,
		"updateTime": 1566818724722,
		"status": "ACCEPTED",
		"avgPrice": "0",
		"source": "API",
		"clientOrderId": "",
		"priceScale": 2,
		"quantityScale": 2,
		"optionSide": "CALL",
		"quoteAsset": "USDT",
		"mmp": false
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTC-200730-9000-C"
	orderID := int64(4611875134427365377)
	clientOrderID := ""
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":        symbol,
			"orderId":       orderID,
			"clientOrderId": clientOrderID,
		})
		s.assertRequestEqual(e, r)
	})
	order, err := s.client.NewGetOrderService().Symbol(symbol).
		OrderID(orderID).ClientOrderID(clientOrderID).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &Order{
		OrderID:       orderID,
		Symbol:        symbol,
		Price:         "100",
		Quantity:      "1",
		ExecutedQty:   "0",
		Fee:           "0",
		Side:          SideTypeBuy,
		Type:          OrderTypeLimit,
		TimeInForce:   TimeInForceTypeGTC,
		ReduceOnly:    false,
		PostOnly:      false,
		CreateTime:    1592465880683,
		UpdateTime:    1566818724722,
		Status:        OrderStatusTypeAccepted,
		AvgPrice:      "0",
		Source:        "API",
		ClientOrderID: clientOrderID,
		PriceScale:    2,
		QuantityScale: 2,
		OptionSide:    OptionSideTypeCall,
		QuoteAsset:    "USDT",
		Mmp:           false,
	}
	s.assertOrderEqual(e, order)
}

func (s *orderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"orderId": 4611875134427365377,
		"symbol": "BTC-200730-9000-C",
		"price": "100",
		"quantity": "1",
		"executedQty": "0",
		"fee": "0",
		"side": "BUY",
		"type": "LIMIT",
		"timeInForce": "GTC",
		"reduceOnly": false,
		"postOnly": false, 
		"CreateTime": 1592465880683,
		"updateTime": 1566818724722,
		"status": "ACCEPTED",
		"avgPrice": "0",
		"source": "API",
		"clientOrderId": "",
		"priceScale": 4,
		"quantityScale": 4,
		"optionSide": "CALL",
		"quoteAsset": "USDT",
		"mmp": false
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTC-200730-9000-C"
	orderID := int64(4611875134427365377)
	clientOrderID := ""
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":        symbol,
			"orderId":       orderID,
			"clientOrderId": clientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelOrderService().Symbol(symbol).
		OrderID(orderID).ClientOrderID(clientOrderID).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CancelOrderResponse{
		OrderID:       orderID,
		Symbol:        symbol,
		Price:         "100",
		Quantity:      "1",
		ExecutedQty:   "0",
		Fee:           "0",
		Side:          SideTypeBuy,
		Type:          OrderTypeLimit,
		TimeInForce:   TimeInForceTypeGTC,
		ReduceOnly:    false,
		PostOnly:      false,
		CreateTime:    1592465880683,
		UpdateTime:    1566818724722,
		Status:        OrderStatusTypeAccepted,
		AvgPrice:      "0",
		Source:        "API",
		ClientOrderID: clientOrderID,
		PriceScale:    4,
		QuantityScale: 4,
		OptionSide:    OptionSideTypeCall,
		QuoteAsset:    "USDT",
		Mmp:           false,
	}
	s.assertCancelOrderResponseEqual(e, res)
}

func (s *orderServiceTestSuite) assertCancelOrderResponseEqual(e, a *CancelOrderResponse) {
	r := s.r()
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.ExecutedQty, a.ExecutedQty, "ExecutedQty")
	r.Equal(e.Fee, a.Fee, "Fee")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.PostOnly, a.PostOnly, "PostOnly")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.AvgPrice, a.AvgPrice, "AvgPrice")
	r.Equal(e.Source, a.Source, "Source")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.PriceScale, a.PriceScale, "PriceScale")
	r.Equal(e.QuantityScale, a.QuantityScale, "QuantityScale")
	r.Equal(e.OptionSide, a.OptionSide, "OptionSide")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.Mmp, a.Mmp, "Mmp")
}

func (s *orderServiceTestSuite) TestCancelAllOpenOrders() {
	data := []byte(`{
		"code": 0,
		"msg": "success"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTC-200730-9000-C"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	err := s.client.NewCancelAllOpenOrdersService().Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
}
