package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type lendingInterestServiceTestSuite struct {
	baseTestSuite
}

func TestLendingInterestService(t *testing.T) {
	suite.Run(t, new(lendingInterestServiceTestSuite))
}

func (s *lendingInterestServiceTestSuite) TestListLendingInterestService() {
	data := []byte(`
	[
		{
			"asset": "USDT",
			"interest": "0.00006408",
			"lendingType": "DAILY",
			"productName": "USDT",
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

	asset := `USDT`
	startTime := int64(1508198532000)
	endTime := int64(1508198532001)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			`asset`:       asset,
			`size`:        2,
			`current`:     1,
			`lendingType`: `DAILY`,
			`startTime`:   startTime,
			`endTime`:     endTime,
			`recvWindow`:  60000,
		})
		s.assertRequestEqual(e, r)
	})

	results, err := s.client.NewListLendingInterestService().
		Asset(asset).
		Size(2).
		Current(1).
		LendingType(`DAILY`).
		StartTime(startTime).
		EndTime(endTime).
		Do(context.Background())
	r := s.r()
	r.NoError(err)
	resultArr := *results

	s.Len(resultArr, 2)
	s.assertInterestEqual(&LendingInterestResponse{
		Time:        1577233578000,
		Asset:       `USDT`,
		ProductName: `USDT`,
		Interest:    `0.00006408`,
		LendingType: `DAILY`,
	}, &resultArr[0])
	s.assertInterestEqual(&LendingInterestResponse{
		Time:        1577233562000,
		Asset:       `USDT`,
		ProductName: `USDT`,
		Interest:    `0.00687654`,
		LendingType: `DAILY`,
	}, &resultArr[1])
}

func (s *lendingInterestServiceTestSuite) assertInterestEqual(e, a *LendingInterestResponse) {
	r := s.r()
	r.Equal(e.Interest, a.Interest, `Interest`)
	r.Equal(e.LendingType, a.LendingType, `LendingType`)
	r.Equal(e.Asset, a.Asset, `Asset`)
	r.Equal(e.Time, a.Time, `Time`)
	r.Equal(e.ProductName, a.ProductName, `ProductName`)
}
