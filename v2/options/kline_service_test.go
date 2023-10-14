package options

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

func (s *klineServiceTestSuite) TestKlines() {
	/*
		OpenTime                 int64  `json:"openTime"`
		Open                     string `json:"open"`
		High                     string `json:"high"`
		Low                      string `json:"low"`
		Close                    string `json:"close"`
		CloseTime                int64  `json:"closeTime"`
		Amount                   string `json:"amount"`
		TakerAmount              string `json:"takerAmount"`
		Volume                   string `json:"volume"`
		TakerVolume              string `json:"takerVolume"`
		Interval                 string `json:"interval"`
		TradeCount               int64  `json:"tradeCount"`
	*/
	//amount:2.35 close:235 closeTime:1677931200000 high:235 interval:4h low:235 open:235 openTime:1677916800000 takerAmount:2.35 takerVolume:0.01 tradeCount:1 volume:0.01
	data := []byte(`[
        {
			"openTime":1499040000000,
			"open":"0.01634790",
			"high":"0.80000000",
			"low":"0.01575800",
			"close":"0.01577100",
			"closeTime":1499644799999,
			"amount":"34.66",
			"takerAmount":"7.6",
			"volume":"1.06",
			"takerVolume":"1.02",
			"interval":"15m",
			"tradeCount":14
		},
        {
			"openTime":1499040000001,
			"open":"0.01634790",
			"high":"0.80000000",
			"low":"0.01575800",
			"close":"0.01577101",
			"closeTime":1499644799999,
			"amount":"17.15",
			"takerAmount":"4.6",
			"volume":"0.06",
			"takerVolume":"0.02",
			"interval":"15m",
			"tradeCount":5
		}
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
	s.Len(klines, 2)
	kline1 := &Kline{
		OpenTime:    1499040000000,
		Open:        "0.01634790",
		High:        "0.80000000",
		Low:         "0.01575800",
		Close:       "0.01577100",
		CloseTime:   1499644799999,
		Amount:      "34.66",
		TakerAmount: "7.6",
		Volume:      "1.06",
		TakerVolume: "1.02",
		Interval:    "15m",
		TradeCount:  14,
	}
	kline2 := &Kline{
		OpenTime:    1499040000001,
		Open:        "0.01634790",
		High:        "0.80000000",
		Low:         "0.01575800",
		Close:       "0.01577101",
		CloseTime:   1499644799999,
		Amount:      "17.15",
		TakerAmount: "4.6",
		Volume:      "0.06",
		TakerVolume: "0.02",
		Interval:    "15m",
		TradeCount:  5,
	}
	s.assertKlineEqual(kline1, klines[0])
	s.assertKlineEqual(kline2, klines[1])
}
func (s *klineServiceTestSuite) assertKlineEqual(e, a *Kline) {
	r := s.r()
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.Open, a.Open, "Open")
	r.Equal(e.High, a.High, "High")
	r.Equal(e.Low, a.Low, "Low")
	r.Equal(e.Close, a.Close, "Close")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.TakerAmount, a.TakerAmount, "TakerAmount")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.TakerVolume, a.TakerVolume, "TakerVolume")
	r.Equal(e.Interval, a.Interval, "Interval")
	r.Equal(e.TradeCount, a.TradeCount, "TradeCount")
}
