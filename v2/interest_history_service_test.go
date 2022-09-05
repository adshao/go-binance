package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type interestHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestInterestHistoryService(t *testing.T) {
	suite.Run(t, new(interestHistoryServiceTestSuite))
}

func (s *interestHistoryServiceTestSuite) TestInterestHistory() {
	data := []byte(`
	[
		{
			"asset": "BUSD",
			"interest": "0.00006408",
			"lendingType": "DAILY",
			"productName": "BUSD",
			"time": 1577233578000
		},
		{
			"asset": "USDT",
			"interest": "0.00687654",
			"lendingType": "DAILY",
			"productName": "USDT",
			"time": 1577233562000
		}
	]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	lendingType := LendingTypeFlexible
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"lendingType": lendingType,
		})
		s.assertRequestEqual(e, r)
	})

	history, err := s.client.NewInterestHistoryService().
		LendingType(lendingType).
		Do(context.Background())
	r := s.r()
	r.NoError(err)

	s.Len(*history, 2)
	s.assertInterestHistoryElementEqual(&InterestHistoryElement{
		Asset:       "BUSD",
		Interest:    "0.00006408",
		LendingType: "DAILY",
		ProductName: "BUSD",
		Time:        1577233578000,
	}, &(*history)[0])
	s.assertInterestHistoryElementEqual(&InterestHistoryElement{
		Asset:       "USDT",
		Interest:    "0.00687654",
		LendingType: "DAILY",
		ProductName: "USDT",
		Time:        1577233562000,
	}, &(*history)[1])
}

func (s *interestHistoryServiceTestSuite) assertInterestHistoryElementEqual(e, a *InterestHistoryElement) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Interest, a.Interest, "Interest")
	r.Equal(e.LendingType, a.LendingType, "LendingType")
	r.Equal(e.ProductName, a.ProductName, "ProductName")
	r.Equal(e.Time, a.Time, "Time")
}
