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
	data := []byte(`[{
		"symbol": "BTCUSDT",
		"markPrice": "11012.80409769",
		"indexPrice": "11781.80495970",
		"estimatedSettlePrice": "11781.80495970",
		"lastFundingRate": "-0.03750000",
		"nextFundingTime": 1562569200000,
		"interestRate": "0.00010000",
		"time": 1562566020000
	}]`)
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
	e := []*PremiumIndex{&PremiumIndex{
		Symbol:               symbol,
		MarkPrice:            "11012.80409769",
		IndexPrice:           "11781.80495970",
		EstimatedSettlePrice: "11781.80495970",
		LastFundingRate:      "-0.03750000",
		NextFundingTime:      int64(1562569200000),
		InterestRate:         "0.00010000",
		Time:                 int64(1562566020000),
	},
	}
	s.assertPremiumIndexEqual(e, res)
}

func (s *premiumIndexServiceTestSuite) assertPremiumIndexEqual(e, a []*PremiumIndex) {
	r := s.r()
	r.Equal(e[0].Symbol, a[0].Symbol, "Symbol")
	r.Equal(e[0].MarkPrice, a[0].MarkPrice, "MarkPrice")
	r.Equal(e[0].IndexPrice, a[0].IndexPrice, "IndexPrice")
	r.Equal(e[0].EstimatedSettlePrice, a[0].EstimatedSettlePrice, "EstimatedSettlePrice")
	r.Equal(e[0].LastFundingRate, a[0].LastFundingRate, "LastFundingRate")
	r.Equal(e[0].NextFundingTime, a[0].NextFundingTime, "NextFundingTime")
	r.Equal(e[0].InterestRate, a[0].InterestRate, "InterestRate")
	r.Equal(e[0].Time, a[0].Time, "Time")
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
			"markPrice": "1576566020000"
		},
		{
			"symbol": "BTCUSDT",
			"fundingRate": "0.00010000",
			"fundingTime": 1570636800000,
			"markPrice": "1576566020000"
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
			MarkPrice:   "1576566020000",
		},
		{
			Symbol:      symbol,
			FundingRate: "0.00010000",
			FundingTime: int64(1570636800000),
			MarkPrice:   "1576566020000",
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
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
}

type getLeverageBracketServiceTestSuite struct {
	baseTestSuite
}

func TestGetLeverageBracketService(t *testing.T) {
	suite.Run(t, new(getLeverageBracketServiceTestSuite))
}

func (s *getLeverageBracketServiceTestSuite) TestGetLeverageBracket() {
	data := []byte(`{
		"symbol": "ETHUSDT",
		"brackets": [
			{
				"bracket": 1,
				"initialLeverage": 75,
				"notionalCap": 10000,
				"notionalFloor": 0,
				"maintMarginRatio": 0.0065,
				"cum": 1.2345
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "ETHUSDT"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetLeverageBracketService().Symbol(symbol).Do(newContext())

	s.r().NoError(err)

	e := []*LeverageBracket{
		{
			Symbol: symbol,
			Brackets: []Bracket{
				{
					Bracket:          1,
					InitialLeverage:  75,
					NotionalCap:      10000,
					NotionalFloor:    0,
					MaintMarginRatio: 0.0065,
					Cum:              1.2345,
				},
			},
		},
	}
	s.r().Len(res, len(e))
	for i := 0; i < len(res); i++ {
		s.assertLeverageBracketEqual(e[i], res[i])
	}
}

func (s *getLeverageBracketServiceTestSuite) assertLeverageBracketEqual(e, a *LeverageBracket) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Brackets[0].Bracket, a.Brackets[0].Bracket, "Bracket")
	r.Equal(e.Brackets[0].InitialLeverage, a.Brackets[0].InitialLeverage, "InitialLeverage")
	r.Equal(e.Brackets[0].NotionalCap, a.Brackets[0].NotionalCap, "NotionalCap")
	r.Equal(e.Brackets[0].NotionalFloor, a.Brackets[0].NotionalFloor, "NotionalFloor")
	r.Equal(e.Brackets[0].MaintMarginRatio, a.Brackets[0].MaintMarginRatio, "MaintMarginRatio")
	r.Equal(e.Brackets[0].Cum, a.Brackets[0].Cum, "Cum")
}
