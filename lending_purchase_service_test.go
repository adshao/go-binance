package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type lendingPurchaseServiceTestSuite struct {
	baseTestSuite
}

func TestLendingPurchaseService(t *testing.T) {
	suite.Run(t, new(lendingPurchaseServiceTestSuite))
}

func (s *lendingPurchaseServiceTestSuite) TestListLendingPurchaseService() {
	data := []byte(`
	[
		{
			"amount": "100.00000000",
			"asset": "USDT",
			"createTime": 1575018510000,
			"lendingType": "DAILY",
			"productName": "USDT",
			"purchaseId": 26055,
			"status": "SUCCESS"
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
			`size`:        1,
			`current`:     1,
			`lendingType`: `DAILY`,
			`startTime`:   startTime,
			`endTime`:     endTime,
			`recvWindow`:  60000,
		})
		s.assertRequestEqual(e, r)
	})

	results, err := s.client.NewListLendingPurchaseService().
		Asset(asset).
		Size(1).
		Current(1).
		LendingType(`DAILY`).
		StartTime(startTime).
		EndTime(endTime).
		Do(context.Background())
	r := s.r()
	r.NoError(err)
	resultArr := *results

	s.Len(resultArr, 1)
	s.assertPurchaseEqual(&LendingPurchaseResponse{
		Amount:      `100.00000000`,
		Asset:       `USDT`,
		CreateTime:  1575018510000,
		LendingType: `DAILY`,
		ProductName: `USDT`,
		PurchaseID:  26055,
		Status:      `SUCCESS`,
	}, &resultArr[0])
}

func (s *lendingPurchaseServiceTestSuite) assertPurchaseEqual(e, a *LendingPurchaseResponse) {
	r := s.r()
	r.Equal(e.Amount, a.Amount, `Amount`)
	r.Equal(e.Asset, a.Asset, `Asset`)
	r.Equal(e.CreateTime, a.CreateTime, `CreateTime`)
	r.Equal(e.LendingType, a.LendingType, `LendingType`)
	r.Equal(e.ProductName, a.ProductName, `ProductName`)
	r.Equal(e.PurchaseID, a.PurchaseID, `PurchaseID`)
	r.Equal(e.Status, a.Status, `Status`)
}
