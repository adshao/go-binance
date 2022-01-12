package binance

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type rebateServiceTestSuite struct {
	baseTestSuite
}

func TestRebateService(t *testing.T) {
	suite.Run(t, new(rebateServiceTestSuite))
}

func (s *rebateServiceTestSuite) TestSpotRebateHistory() {
	data := []byte(`{
	   "status": "OK",
	   "type": "GENERAL",
	   "code": "000000000",
	   "data": {
			"page": 1,
			"totalRecords": 2,
			"totalPageNum": 1,
			"data": [
				{
					"asset": "USDT",
					"type": 1,
					"amount": "0.0001126", 
					"updateTime": 1637651320000
				},
				{
					"asset": "ETH",
					"type": 1,
					"amount": "0.00000056",
					"updateTime": 1637928379000
				}
			]
		}
	
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewSpotRebateHistoryService().Do(newContext())
	s.r().NoError(err)
	e := &SpotRebateHistory{
		Status: "OK",
		Type:   "GENERAL",
		Code:   "000000000",
		Data: SpotRebateHistoryData{
			Page:         1,
			TotalRecords: 2,
			TotalPageNum: 1,
			Data: []SpotRebateHistoryDataItem{
				{
					Asset:      "USDT",
					Type:       1,
					Amount:     "0.0001126",
					UpdateTime: 1637651320000,
				},
				{
					Asset:      "ETH",
					Type:       1,
					Amount:     "0.00000056",
					UpdateTime: 1637928379000,
				},
			},
		},
	}
	s.assertSpotRebateHistoryEqual(e, res)
}

func (s *rebateServiceTestSuite) assertSpotRebateHistoryEqual(e, a *SpotRebateHistory) {
	r := s.r()
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Code, a.Code, "Code")
	s.assertSpotRebateHistoryDataEqual(&e.Data, &a.Data)
}

func (s *rebateServiceTestSuite) assertSpotRebateHistoryDataEqual(e, a *SpotRebateHistoryData) {
	r := s.r()
	r.Equal(e.Page, a.Page, "Page")
	r.Equal(e.TotalRecords, a.TotalRecords, "TotalRecords")
	r.Equal(e.TotalPageNum, a.TotalPageNum, "TotalPageNum")

	r.Len(a.Data, len(e.Data))
	for i := 0; i < len(a.Data); i++ {
		s.assertSpotRebateHistoryDataItemEqual(&e.Data[i], &a.Data[i])
	}
}

func (s *rebateServiceTestSuite) assertSpotRebateHistoryDataItemEqual(e, a *SpotRebateHistoryDataItem) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
}
