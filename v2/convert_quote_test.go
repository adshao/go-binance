package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type convertGetQuoteTestSuite struct {
	baseTestSuite
}

func TestConvertGetQuoteService(t *testing.T) {
	suite.Run(t, new(convertGetQuoteTestSuite))
}

func (s *convertGetQuoteTestSuite) TestConvertGetQuote() {
	data := []byte(`{
		"quoteId":"12415572564",
		"ratio":"38163.7",
		"inverseRatio":"0.0000262",
		"validTimestamp":1623319461670,
		"toAmount":"3816.37",
		"fromAmount":"0.1"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"fromAsset":  "BNB",
			"toAsset":    "USDT",
			"fromAmount": 1,
			"validTime":  "10s",
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewConvertGetQuoteService().
		FromAsset("BNB").
		ToAsset("USDT").
		FromAmount(1).
		ValidTime("10s").
		Do(newContext())
	s.r().NoError(err)
	e := &ConvertGetQuote{
		QuoteID:        "12415572564",
		Ratio:          "38163.7",
		InverseRatio:   "0.0000262",
		ValidTimestamp: 1623319461670,
		ToAmount:       "3816.37",
		FromAmount:     "0.1",
	}
	s.assertConvertGetQuuoteEqual(e, res)
}

func (s *convertGetQuoteTestSuite) assertConvertGetQuuoteEqual(e, a *ConvertGetQuote) {
	r := s.r()

	r.Equal(e.QuoteID, a.QuoteID, "QuoteID")
	r.Equal(e.Ratio, a.Ratio, "Ratio")
	r.Equal(e.InverseRatio, a.InverseRatio, "InverseRatio")
	r.Equal(e.ValidTimestamp, a.ValidTimestamp, "ValidTimestamp")
	r.Equal(e.ToAmount, a.ToAmount, "ToAmount")
	r.Equal(e.FromAmount, a.FromAmount, "FromAmount")
}

type convertAcceptQuoteTestSuite struct {
	baseTestSuite
}

func TestConvertAcceptQuoteService(t *testing.T) {
	suite.Run(t, new(convertAcceptQuoteTestSuite))
}

func (s *convertAcceptQuoteTestSuite) TestConvertAcceptQuote() {
	data := []byte(`{
		"orderId":"933256278426274426",
  	"createTime":1623381330472,
  	"orderStatus":"PROCESS"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"quoteId": "12415572564",
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewConvertAcceptQuoteService().
		QuoteID("12415572564").
		Do(newContext())
	s.r().NoError(err)
	e := &ConvertAcceptQuote{
		OrderId:     "933256278426274426",
		CreateTime:  1623381330472,
		OrderStatus: "PROCESS",
	}
	s.assertConvertAcceptQuuoteEqual(e, res)
}

func (s *convertAcceptQuoteTestSuite) assertConvertAcceptQuuoteEqual(e, a *ConvertAcceptQuote) {
	r := s.r()

	r.Equal(e.OrderId, a.OrderId, "OrderId")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.OrderStatus, a.OrderStatus, "OrderStatus")
}

type convertOrderStatusTestSuite struct {
	baseTestSuite
}

func TestConvertOrderStatusService(t *testing.T) {
	suite.Run(t, new(convertOrderStatusTestSuite))
}

func (s *convertOrderStatusTestSuite) TestConvertOrderStatus() {
	data := []byte(`{
		"orderId":933256278426274426,
		"orderStatus":"SUCCESS",
		"fromAsset":"BTC",
		"fromAmount":"0.00054414",
		"toAsset":"USDT",
		"toAmount":"20",
		"ratio":"36755",
		"inverseRatio":"0.00002721",
		"createTime":1623381330472
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"orderId": "933256278426274426",
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewConvertOrderStatusService().
		OrderID("933256278426274426").
		Do(newContext())
	s.r().NoError(err)
	e := &ConvertOrderStatus{
		OrderId:      933256278426274426,
		OrderStatus:  "SUCCESS",
		FromAsset:    "BTC",
		FromAmount:   "0.00054414",
		ToAsset:      "USDT",
		ToAmount:     "20",
		Ratio:        "36755",
		InverseRatio: "0.00002721",
		CreateTime:   1623381330472,
	}
	s.assertConvertOrderStatusEqual(e, res)
}

func (s *convertOrderStatusTestSuite) assertConvertOrderStatusEqual(e, a *ConvertOrderStatus) {
	r := s.r()

	r.Equal(e.OrderId, a.OrderId, "OrderId")
	r.Equal(e.OrderStatus, a.OrderStatus, "OrderId")
	r.Equal(e.FromAsset, a.FromAsset, "OrderId")
	r.Equal(e.FromAmount, a.FromAmount, "OrderId")
	r.Equal(e.ToAsset, a.ToAsset, "OrderId")
	r.Equal(e.ToAmount, a.ToAmount, "OrderId")
	r.Equal(e.Ratio, a.Ratio, "OrderId")
	r.Equal(e.InverseRatio, a.InverseRatio, "OrderId")
	r.Equal(e.CreateTime, a.CreateTime, "OrderId")
}
