package binance

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type convertTradeTestSuite struct {
	baseTestSuite
}

func TestConvertTradeService(t *testing.T) {
	suite.Run(t, new(convertTradeTestSuite))
}

func (s *convertTradeTestSuite) TestConvertTradeHistory() {
	data := []byte(`{
	   "list": [
			{
				"quoteId": "f3b91c525b2644c7bc1e1cd31b6e1aa6",
				"orderId": 940708407462087195,
				"orderStatus": "SUCCESS",
				"fromAsset": "USDT",
				"fromAmount": "20",
				"toAsset": "BNB",
				"toAmount": "0.06154036",
				"ratio": "0.00307702",
				"inverseRatio": "324.99",
				"createTime": 1624248872184
			}
	   ],
		"startTime": 1623824139000,
		"endTime": 1626416139000,
		"limit": 100,
		"moreData": false
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	startTime := time.Now().AddDate(0, 0, -7).Unix() * 1000
	endTime := time.Now().Unix() * 1000
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewConvertTradeHistoryService().
		StartTime(startTime).
		EndTime(endTime).
		Do(newContext())
	s.r().NoError(err)
	e := &ConvertTradeHistory{
		List: []ConvertTradeHistoryItem{
			{
				QuoteId:      "f3b91c525b2644c7bc1e1cd31b6e1aa6",
				OrderId:      940708407462087195,
				OrderStatus:  "SUCCESS",
				FromAsset:    "USDT",
				FromAmount:   "20",
				ToAsset:      "BNB",
				ToAmount:     "0.06154036",
				Ratio:        "0.00307702",
				InverseRatio: "324.99",
				CreateTime:   1624248872184,
			},
		},
		StartTime: 1623824139000,
		EndTime:   1626416139000,
		Limit:     100,
		MoreData:  false,
	}
	s.assertConvertTradeHistoryEqual(e, res)
}

func (s *convertTradeTestSuite) assertConvertTradeHistoryEqual(e, a *ConvertTradeHistory) {
	r := s.r()

	r.Len(a.List, len(e.List))
	for i := 0; i < len(a.List); i++ {
		s.assertConvertTradeHistoryItemEqual(&e.List[i], &a.List[i])
	}

	r.Equal(e.StartTime, a.StartTime, "StartTime")
	r.Equal(e.EndTime, a.EndTime, "EndTime")
	r.Equal(e.Limit, a.Limit, "Limit")
	r.Equal(e.MoreData, a.MoreData, "MoreData")
}

func (s *convertTradeTestSuite) assertConvertTradeHistoryItemEqual(e, a *ConvertTradeHistoryItem) {
	r := s.r()
	r.Equal(e.QuoteId, a.QuoteId, "QuoteId")
	r.Equal(e.OrderId, a.OrderId, "OrderId")
	r.Equal(e.OrderStatus, a.OrderStatus, "OrderStatus")
	r.Equal(e.FromAsset, a.FromAsset, "FromAsset")
	r.Equal(e.FromAmount, a.FromAmount, "FromAmount")
	r.Equal(e.ToAsset, a.ToAsset, "ToAsset")
	r.Equal(e.ToAmount, a.ToAmount, "ToAmount")
	r.Equal(e.Ratio, a.Ratio, "Ratio")
	r.Equal(e.InverseRatio, a.InverseRatio, "InverseRatio")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
}
