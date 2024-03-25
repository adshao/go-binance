package futures

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type fundingRateHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestFundingRateHistoryService(t *testing.T) {
	suite.Run(t, new(fundingRateHistoryServiceTestSuite))
}

func (s *fundingRateHistoryServiceTestSuite) TestGetFundingRateHistory() {
	data := []byte(`[{
		"symbol": "BTCUSDT",
		"fundingTime": 1698768000000,
		"fundingRate": "0.00010000",
		"markPrice": "34287.54619963"
	}]`)
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

	res, err := s.client.NewFundingRateHistoryService().Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Do(newContext())
	s.r().NoError(err)
	e := []*FundingRateHistory{
		{
			Symbol:      symbol,
			FundingTime: 1698768000000,
			FundingRate: "0.00010000",
			MarkPrice:   "34287.54619963",
		},
	}
	s.assertFundingRateHistoryEqual(e, res)
}

func (s *fundingRateHistoryServiceTestSuite) assertFundingRateHistoryEqual(e, a []*FundingRateHistory) {
	r := s.r()
	r.Equal(e[0].Symbol, a[0].Symbol, "Symbol")
	r.Equal(e[0].FundingRate, a[0].FundingRate, "FundingRate")
	r.Equal(e[0].FundingTime, a[0].FundingTime, "FundingTime")
	r.Equal(e[0].MarkPrice, a[0].MarkPrice, "MarkPrice")
}

func TestName(t *testing.T) {
	client := NewClient("", "")

	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancelFunc()

	rateHistory, err := client.NewFundingRateHistoryService().Do(ctx)
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		fmt.Printf("%+v\n", rateHistory[i])
	}
}
