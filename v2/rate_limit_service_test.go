package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type assetRateLimitServiceSuite struct {
	baseTestSuite
}

func (a *assetRateLimitServiceSuite) assertRateLimitServiceEqual(expected, other *RateLimitFull) {
	r := a.r()

	r.EqualValues(expected, other)
}

func TestRateLimitService(t *testing.T) {
	suite.Run(t, new(assetRateLimitServiceSuite))
}

func (s *assetRateLimitServiceSuite) TestListRateLimit() {
	data := []byte(`
	[
		{
			"rateLimitType": "ORDERS",
			"interval": "SECOND",
			"intervalNum": 10,
			"limit": 10000,
			"count": 0
		},
		{
			"rateLimitType": "RAW_REQUESTS",
			"interval": "MINUTE",
			"intervalNum": 5,
			"limit": 5000,
			"count": 100
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	limits, err := s.client.NewRateLimitService().Do(context.Background())
	s.r().NoError(err)
	rows := limits

	s.Len(rows, 2)
	s.assertRateLimitServiceEqual(&RateLimitFull{
		RateLimitType: "ORDERS",
		Interval:      "SECOND",
		IntervalNum:   10,
		Limit:         10000,
		Count:         0,
	},
		rows[0])
	s.assertRateLimitServiceEqual(&RateLimitFull{
		RateLimitType: "RAW_REQUESTS",
		Interval:      "MINUTE",
		IntervalNum:   5,
		Limit:         5000,
		Count:         100,
	},
		rows[1])
}
