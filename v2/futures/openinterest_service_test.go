package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type openInterestServiceTestSuite struct {
	baseTestSuite
}

func TestGetOpenInterestService(t *testing.T) {
	suite.Run(t, new(openInterestServiceTestSuite))
}

func (s *openInterestServiceTestSuite) TestGetOpenInterest() {
	data := []byte(`{
		"openInterest": "10659.509", 
		"symbol": "BTCUSDT",
		"time": 1589437530011
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetOpenInterestService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)

	e := &OpenInterest{
		Symbol:       symbol,
		Time:         1589437530011,
		OpenInterest: "10659.509",
	}

	s.r().Equal(e.Symbol, res.Symbol, "Symbol")
	s.r().Equal(e.OpenInterest, res.OpenInterest, "OpenInterest")
	s.r().Equal(e.Time, res.Time, "Time")
}

func (s *openInterestServiceTestSuite) TestOpenInterestStatistics() {
	data := []byte(`[
		{ 
			"symbol":"BTCUSDT",
			"sumOpenInterest":"20403.63700000", 
			"sumOpenInterestValue": "150570784.07809979",
			"timestamp": 1583127900000
		},
		{
			"symbol":"BTCUSDT",
			"sumOpenInterest":"20401.36700000",
			"sumOpenInterestValue":"149940752.14464448",
			"timestamp": 1583128200000
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	period := "15m"
	limit := 10
	startTime := int64(1499040000000)
	endTime := int64(1499040000001)
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

	openInterests, err := s.client.NewOpenInterestStatisticsService().Symbol(symbol).
		Period(period).Limit(limit).StartTime(startTime).
		EndTime(endTime).Do(newContext())

	s.r().NoError(err)
	s.Len(openInterests, 2)

	openInterest1 := &OpenInterestStatistic{
		Symbol:               "BTCUSDT",
		SumOpenInterest:      "20403.63700000",
		SumOpenInterestValue: "150570784.07809979",
		Timestamp:            1583127900000,
	}
	openInterest2 := &OpenInterestStatistic{
		Symbol:               "BTCUSDT",
		SumOpenInterest:      "20401.36700000",
		SumOpenInterestValue: "149940752.14464448",
		Timestamp:            1583128200000,
	}
	s.assertOpenInterestStatisticEqual(openInterest1, openInterests[0])
	s.assertOpenInterestStatisticEqual(openInterest2, openInterests[1])
}

func (s *openInterestServiceTestSuite) assertOpenInterestStatisticEqual(e, a *OpenInterestStatistic) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.SumOpenInterest, a.SumOpenInterest, "SumOpenInterest")
	r.Equal(e.SumOpenInterestValue, a.SumOpenInterestValue, "SumOpenInterestValue")
}
