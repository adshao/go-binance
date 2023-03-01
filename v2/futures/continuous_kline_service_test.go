package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ContinuousklineServiceTestSuite struct {
	baseTestSuite
}

func TestContinuousKlineService(t *testing.T) {
	suite.Run(t, new(ContinuousklineServiceTestSuite))
}

func (s *ContinuousklineServiceTestSuite) TestContinuousKlines() {
	data := []byte(`[
        [
            1499040000000,
            "0.01634790",
            "0.80000000",
            "0.01575800",
            "0.01577100",
            "148976.11427815",
            1499644799999,
            "2434.19055334",
            308,
            "1756.87402397",
            "28.46694368",
            "17928899.62484339"
        ],
        [
            1499040000001,
            "0.01634790",
            "0.80000000",
            "0.01575800",
            "0.01577101",
            "148976.11427815",
            1499644799999,
            "2434.19055334",
            308,
            "1756.87402397",
            "28.46694368",
            "17928899.62484339"
        ]
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	pair := "LTCBTC"
	contractType := "PERPETUAL"
	interval := "15m"
	limit := 10
	startTime := int64(1499040000000)
	endTime := int64(1499040000001)
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"pair":         pair,
			"contractType": contractType,
			"interval":     interval,
			"limit":        limit,
			"startTime":    startTime,
			"endTime":      endTime,
		})
		s.assertRequestEqual(e, r)
	})
	klines, err := s.client.NewContinuousKlinesService().Pair(pair).
		ContractType(contractType).Interval(interval).Limit(limit).
		StartTime(startTime).EndTime(endTime).Do(newContext())
	s.r().NoError(err)
	s.Len(klines, 2)
	kline1 := &ContinuousKline{
		OpenTime:                 1499040000000,
		Open:                     "0.01634790",
		High:                     "0.80000000",
		Low:                      "0.01575800",
		Close:                    "0.01577100",
		Volume:                   "148976.11427815",
		CloseTime:                1499644799999,
		QuoteAssetVolume:         "2434.19055334",
		TradeNum:                 308,
		TakerBuyBaseAssetVolume:  "1756.87402397",
		TakerBuyQuoteAssetVolume: "28.46694368",
	}
	kline2 := &ContinuousKline{
		OpenTime:                 1499040000001,
		Open:                     "0.01634790",
		High:                     "0.80000000",
		Low:                      "0.01575800",
		Close:                    "0.01577101",
		Volume:                   "148976.11427815",
		CloseTime:                1499644799999,
		QuoteAssetVolume:         "2434.19055334",
		TradeNum:                 308,
		TakerBuyBaseAssetVolume:  "1756.87402397",
		TakerBuyQuoteAssetVolume: "28.46694368",
	}
	s.assertContinuousKlineEqual(kline1, klines[0])
	s.assertContinuousKlineEqual(kline2, klines[1])
}

func (s *ContinuousklineServiceTestSuite) assertContinuousKlineEqual(e, a *ContinuousKline) {
	r := s.r()
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.Open, a.Open, "Open")
	r.Equal(e.High, a.High, "High")
	r.Equal(e.Low, a.Low, "Low")
	r.Equal(e.Close, a.Close, "Close")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.QuoteAssetVolume, a.QuoteAssetVolume, "QuoteAssetVolume")
	r.Equal(e.TradeNum, a.TradeNum, "TradeNum")
	r.Equal(e.TakerBuyBaseAssetVolume, a.TakerBuyBaseAssetVolume, "TakerBuyBaseAssetVolume")
	r.Equal(e.TakerBuyQuoteAssetVolume, a.TakerBuyQuoteAssetVolume, "TakerBuyQuoteAssetVolume")
}
