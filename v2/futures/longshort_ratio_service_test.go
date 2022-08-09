package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type longShortRatioServiceTestSuite struct {
	baseTestSuite
}

func TestLongShortRatioService(t *testing.T) {
	suite.Run(t, new(longShortRatioServiceTestSuite))
}

func (s *longShortRatioServiceTestSuite) TestOpenInterestStatistics() {
	data := []byte(`[
		{ 
			"symbol":"BTCUSDT",
			"longShortRatio":"1.8105",
			"longAccount": "0.6442", 
			"shortAccount":"0.3558", 
			"timestamp":1583139600000
		},
		{
			"symbol":"BTCUSDT",
			"longShortRatio":"0.5576",
			"longAccount": "0.3580", 
			"shortAccount":"0.6420",                  
			"timestamp":1583139900000
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	period := "15m"
	limit := 10
	startTime := int64(1583139600000)
	endTime := int64(1583139900000)
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol":    symbol,
			"period":    period,
			"limit":     limit,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	longShortRatios, err := s.client.NewLongShortRatioService().Symbol(symbol).
		Period(period).Limit(limit).StartTime(startTime).
		EndTime(endTime).Do(newContext())

	s.r().NoError(err)
	s.Len(longShortRatios, 2)

	longShortRatio1 := &LongShortRatio{
		Symbol:         "BTCUSDT",
		LongShortRatio: "1.8105",
		ShortAccount:   "0.3558",
		LongAccount:    "0.6442",
		Timestamp:      1583139600000,
	}
	longShortRatio2 := &LongShortRatio{
		Symbol:         "BTCUSDT",
		LongShortRatio: "0.5576",
		ShortAccount:   "0.6420",
		LongAccount:    "0.3580",
		Timestamp:      1583139900000,
	}
	s.assertLongShortRatioEqual(longShortRatio1, longShortRatios[0])
	s.assertLongShortRatioEqual(longShortRatio2, longShortRatios[1])
}

func (s *longShortRatioServiceTestSuite) assertLongShortRatioEqual(e, a *LongShortRatio) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.LongShortRatio, a.LongShortRatio, "LongShortRatio")
	r.Equal(e.LongAccount, a.LongAccount, "LongAccount")
	r.Equal(e.ShortAccount, a.ShortAccount, "ShortAccount")
}
