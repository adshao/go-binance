package delivery

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type klineServiceTestSuite struct {
	baseTestSuite
}

func TestKlineService(t *testing.T) {
	suite.Run(t, new(klineServiceTestSuite))
}

// https://binance-docs.github.io/apidocs/delivery/en/#kline-candlestick-data
func (s *klineServiceTestSuite) TestKlines() {
	data := []byte(`[
        [
			1591258320000,
			"9640.7",
			"9642.4",
			"9640.6",
			"9642.0",
			"206",
			1591258379999,
			"2.13660389",
			48,
			"119",
			"1.23424865",
			"0"
        ]
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	interval := "15m"
	limit := 10
	startTime := int64(1499040000000)
	endTime := int64(1499040000001)
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol":    symbol,
			"interval":  interval,
			"limit":     limit,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})
	klines, err := s.client.NewKlinesService().Symbol(symbol).
		Interval(interval).Limit(limit).StartTime(startTime).
		EndTime(endTime).Do(newContext())
	s.r().NoError(err)
	s.Len(klines, 1)
	kline1 := &Kline{
		OpenTime:                 1591258320000,
		Open:                     "9640.7",
		High:                     "9642.4",
		Low:                      "9640.6",
		Close:                    "9642.0",
		Volume:                   "206",
		CloseTime:                1591258379999,
		QuoteAssetVolume:         "2.13660389",
		TradeNum:                 48,
		TakerBuyBaseAssetVolume:  "119",
		TakerBuyQuoteAssetVolume: "1.23424865",
	}
	s.assertKlineEqual(kline1, klines[0])
}

func (s *klineServiceTestSuite) assertKlineEqual(e, a *Kline) {
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
