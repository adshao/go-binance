/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Tuesday 21 July 2020 - 14:49:10
** @Filename:				asset_dividend_service_test copy.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type listDustLogServiceTestSuite struct {
	baseTestSuite
}

func TestListDustLogService(t *testing.T) {
	suite.Run(t, new(listDustLogServiceTestSuite))
}

func (s *listDustLogServiceTestSuite) TestListDustLog() {
	data := []byte(`
	{
			"total": 2,
			"userAssetDribblets": [
				{
					"totalTransferedAmount": "0.00132256",
					"totalServiceChargeAmount": "0.00002699",
					"transId": 4359321,
					"userAssetDribbletDetails": [
						{
							"transId": 4359321,
							"serviceChargeAmount": "0.000009",
							"uid": "10000015",
							"amount": "0.0009",
							"operateTime": 1618413101,
							"transferedAmount": "0.000441",
							"fromAsset": "USDT"
						},
						{
							"transId": 4359321,
							"serviceChargeAmount": "0.00001799",
							"uid": "10000015",
							"amount": "0.0009",
							"operateTime": 1618413101,
							"transferedAmount": "0.00088156",
							"fromAsset": "ETH"
						}
					],
					"operateTime": 1618413101
				},
				{
					"totalTransferedAmount": "0.00058795",
					"totalServiceChargeAmount": "0.000012",
					"transId": 4357015,
					"userAssetDribbletDetails": [
						{
							"transId": 4357015,
							"serviceChargeAmount": "0.00001",
							"uid": "10000015",
							"amount": "0.001",
							"operateTime": 1618413101,
							"transferedAmount": "0.00049",
							"fromAsset": "USDT"
						},
						{
							"transId": 4357015,
							"serviceChargeAmount": "0.000002",
							"uid": "10000015",
							"amount": "0.0001",
							"operateTime": 1618413101,
							"transferedAmount": "0.00009795",
							"fromAsset": "ETH"
						}
					],
					"operateTime": 1618413101
				}
			]		
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	dustLog, err := s.client.NewListDustLogService().Do(context.Background())
	r := s.r()
	r.NoError(err)
	rows := dustLog.Results.Rows
	s.Len(rows, 2)
	s.Len(rows[0].Logs, 2)
	s.Len(rows[1].Logs, 2)

	s.assertDustRowEqual(&DustRow{
		TotalTransferedAmount:    "0.00132256",
		TotalServiceChargeAmount: "0.00002699",
		TranID:                   4359321,
		Logs: []DustLog{
			{
				TranID:              4359321,
				ServiceChargeAmount: "0.000009",
				UID:                 "10000015",
				Amount:              "0.0009",
				OperateTime:         1618413101,
				TransferedAmount:    "0.000441",
				FromAsset:           "USDT",
			},
			{
				TranID:              4359321,
				ServiceChargeAmount: "0.00001799",
				UID:                 "10000015",
				Amount:              "0.0009",
				OperateTime:         1618413101,
				TransferedAmount:    "0.00088156",
				FromAsset:           "ETH",
			},
		},
		OperateTime: 1618413101,
	}, &rows[0])
	s.assertDustRowEqual(&DustRow{
		TotalTransferedAmount:    "0.00058795",
		TotalServiceChargeAmount: "0.000012",
		TranID:                   4357015,
		Logs: []DustLog{
			{
				TranID:              4357015,
				ServiceChargeAmount: "0.00001",
				UID:                 "10000015",
				Amount:              "0.001",
				OperateTime:         1618413101,
				TransferedAmount:    "0.00049",
				FromAsset:           "USDT",
			},
			{
				TranID:              4357015,
				ServiceChargeAmount: "0.000002",
				UID:                 "10000015",
				Amount:              "0.0001",
				OperateTime:         1618413101,
				TransferedAmount:    "0.00009795",
				FromAsset:           "ETH",
			},
		},
		OperateTime: 1618413101,
	}, &rows[1])
}

func (s *listDustLogServiceTestSuite) assertDustRowEqual(e, a *DustRow) {
	r := s.r()
	r.Equal(e.TotalTransferedAmount, a.TotalTransferedAmount, `TransferedTotal`)
	r.Equal(e.TotalServiceChargeAmount, a.TotalServiceChargeAmount, `ServiceChargeTotal`)
	r.Equal(e.TranID, a.TranID, `TranID`)
	s.assertDustLogEqual(&e.Logs[0], &a.Logs[0])
	s.assertDustLogEqual(&e.Logs[1], &a.Logs[1])
	r.Equal(e.OperateTime, a.OperateTime, `OperateTime`)
}

func (s *listDustLogServiceTestSuite) assertDustLogEqual(e, a *DustLog) {
	r := s.r()
	r.Equal(e.TranID, a.TranID, `TranID`)
	r.Equal(e.ServiceChargeAmount, a.ServiceChargeAmount, `ServiceChargeAmount`)
	r.Equal(e.UID, a.UID, `UID`)
	r.Equal(e.Amount, a.Amount, `Amount`)
	r.Equal(e.OperateTime, a.OperateTime, `OperateTime`)
	r.Equal(e.TransferedAmount, a.TransferedAmount, `TransferedAmount`)
	r.Equal(e.FromAsset, a.FromAsset, `FromAsset`)
}

type dustTransferTestSuite struct {
	baseTestSuite
}

func TestDustTransferService(t *testing.T) {
	suite.Run(t, new(dustTransferTestSuite))
}

func (s *dustTransferTestSuite) TestTransfer() {
	data := []byte(`{
		"totalServiceCharge":"0.02102542",
		"totalTransfered":"1.05127099",
		"transferResult":[
			{
				"amount":"0.03000000",
				"fromAsset":"ETH",
				"operateTime":1563368549307,
				"serviceChargeAmount":"0.00500000",
				"tranId":2970932918,
				"transferedAmount":"0.25000000"
			},
			{
				"amount":"0.09000000",
				"fromAsset":"LTC",
				"operateTime":1563368549404,
				"serviceChargeAmount":"0.01548000",
				"tranId":2970932918,
				"transferedAmount":"0.77400000"
			},
			{
				"amount":"248.61878453",
				"fromAsset":"TRX",
				"operateTime":1563368549489,
				"serviceChargeAmount":"0.00054542",
				"tranId":2970932918,
				"transferedAmount":"0.02727099"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := []string{"ETH", "LTC", "TRX"}
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		for _, a := range asset {
			e.addParam("asset", a)
		}
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewDustTransferService().Asset(asset).Do(newContext())
	s.r().NoError(err)
	e := &DustTransferResponse{
		TotalServiceCharge: "0.02102542",
		TotalTransfered:    "1.05127099",
		TransferResult: []*DustTransferResult{
			{
				Amount:              "0.03000000",
				FromAsset:           "ETH",
				OperateTime:         1563368549307,
				ServiceChargeAmount: "0.00500000",
				TranID:              2970932918,
				TransferedAmount:    "0.25000000",
			},
			{
				Amount:              "0.09000000",
				FromAsset:           "LTC",
				OperateTime:         1563368549404,
				ServiceChargeAmount: "0.01548000",
				TranID:              2970932918,
				TransferedAmount:    "0.77400000",
			},
			{
				Amount:              "248.61878453",
				FromAsset:           "TRX",
				OperateTime:         1563368549489,
				ServiceChargeAmount: "0.00054542",
				TranID:              2970932918,
				TransferedAmount:    "0.02727099",
			},
		},
	}
	s.assertTransferResponse(e, res)
}

func (s *dustTransferTestSuite) assertTransferResponse(e, a *DustTransferResponse) {
	r := s.r()
	r.Equal(e.TotalServiceCharge, a.TotalServiceCharge, "TotalServiceCharge")
	r.Equal(e.TotalTransfered, a.TotalTransfered, "TotalTransfered")
	for i, etr := range e.TransferResult {
		r.Equal(etr.Amount, a.TransferResult[i].Amount, "Amount")
		r.Equal(etr.FromAsset, a.TransferResult[i].FromAsset, "FromAsset")
		r.Equal(etr.OperateTime, a.TransferResult[i].OperateTime, "OperateTime")
		r.Equal(etr.ServiceChargeAmount, a.TransferResult[i].ServiceChargeAmount, "ServiceChargeAmount")
		r.Equal(etr.TranID, a.TransferResult[i].TranID, "TranID")
		r.Equal(etr.TransferedAmount, a.TransferResult[i].TransferedAmount, "TransferedAmount")
	}
}
