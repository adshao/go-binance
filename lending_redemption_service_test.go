package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type lendingRedemptionServiceTestSuite struct {
	baseTestSuite
}

func TestLendingRedemptionService(t *testing.T) {
	suite.Run(t, new(lendingRedemptionServiceTestSuite))
}

func (s *lendingRedemptionServiceTestSuite) TestListLendingRedemptionService() {
	data := []byte(`
	[
		{
			"amount": "10.54000000",
			"asset": "USDT",
			"createTime": 1577257222000,
			"principal": "10.54000000",
			"projectId": "USDT001",
			"projectName": "USDT",
			"status": "PAID",
			"type": "FAST"
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

	results, err := s.client.NewListLendingRedemptionService().
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
	s.assertRedemptionEqual(&LendingRedemptionResponse{
		Amount:      `10.54000000`,
		Asset:       `USDT`,
		CreateTime:  1577257222000,
		Principal:   `10.54000000`,
		ProjectID:   `USDT001`,
		ProjectName: `USDT`,
		Type:        `FAST`,
		Status:      `PAID`,
	}, &resultArr[0])
}

func (s *lendingRedemptionServiceTestSuite) assertRedemptionEqual(e, a *LendingRedemptionResponse) {
	r := s.r()
	r.Equal(e.Amount, a.Amount, `Amount`)
	r.Equal(e.Asset, a.Asset, `Asset`)
	r.Equal(e.CreateTime, a.CreateTime, `CreateTime`)
	r.Equal(e.Principal, a.Principal, `Principal`)
	r.Equal(e.ProjectID, a.ProjectID, `ProjectID`)
	r.Equal(e.ProjectName, a.ProjectName, `ProjectName`)
	r.Equal(e.Type, a.Type, `Type`)
	r.Equal(e.Status, a.Status, `Status`)
}
