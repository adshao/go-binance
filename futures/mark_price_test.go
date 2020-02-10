package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type premiumIndexServiceTestSuite struct {
	baseTestSuite
}

func TestPremiumIndexService(t *testing.T) {
	suite.Run(t, new(premiumIndexServiceTestSuite))
}

func (s *premiumIndexServiceTestSuite) TestGetPremiumIndex() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"markPrice": "11012.80409769",
		"lastFundingRate": "-0.03750000",
		"nextFundingTime": 1562569200000,
		"time": 1562566020000
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

	res, err := s.client.NewPremiumIndexService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	e := &PremiumIndex{
		Symbol:          symbol,
		MarkPrice:       "11012.80409769",
		LastFundingRate: "-0.03750000",
		NextFundingTime: int64(1562569200000),
		Time:            int64(1562566020000),
	}
	s.assertPremiumIndexEqual(e, res)
}

func (s *premiumIndexServiceTestSuite) assertPremiumIndexEqual(e, a *PremiumIndex) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.LastFundingRate, a.LastFundingRate, "LastFundingRate")
	r.Equal(e.NextFundingTime, a.NextFundingTime, "NextFundingTime")
	r.Equal(e.Time, a.Time, "Time")
}

type fundingRateServiceTestSuite struct {
	baseTestSuite
}

func TestFundingRateService(t *testing.T) {
	suite.Run(t, new(fundingRateServiceTestSuite))
}

func (s *fundingRateServiceTestSuite) TestGetFundingRate() {
	data := []byte(`[
		{
			"symbol": "BTCUSDT",
			"fundingRate": "-0.03750000",
			"fundingTime": 1570608000000,
			"time": 1576566020000
		},
		{
			"symbol": "BTCUSDT",
			"fundingRate": "0.00010000",
			"fundingTime": 1570636800000,
			"time": 1576566020000
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	startTime := int64(1576566020000)
	endTime := int64(1676566020000)
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewFundingRateService().Symbol(symbol).StartTime(startTime).
		EndTime(endTime).Limit(limit).Do(newContext())
	s.r().NoError(err)
	e := []*FundingRate{
		{
			Symbol:      symbol,
			FundingRate: "-0.03750000",
			FundingTime: int64(1570608000000),
			Time:        int64(1576566020000),
		},
		{
			Symbol:      symbol,
			FundingRate: "0.00010000",
			FundingTime: int64(1570636800000),
			Time:        int64(1576566020000),
		},
	}
	s.r().Len(res, len(e))
	for i := 0; i < len(res); i++ {
		s.assertFundingRateEqual(e[i], res[i])
	}
}

func (s *fundingRateServiceTestSuite) assertFundingRateEqual(e, a *FundingRate) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.FundingRate, a.FundingRate, "FundingRate")
	r.Equal(e.FundingTime, a.FundingTime, "FundingTime")
	r.Equal(e.Time, a.Time, "Time")
}
