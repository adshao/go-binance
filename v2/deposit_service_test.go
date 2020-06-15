package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type depositServiceTestSuite struct {
	baseTestSuite
}

func TestDepositService(t *testing.T) {
	suite.Run(t, new(depositServiceTestSuite))
}

func (s *depositServiceTestSuite) TestListDeposits() {
	data := []byte(`
	{
		"depositList": [
			{
				"insertTime": 1508198532000,
				"amount": 0.04670582,
				"asset": "ETH",
				"address": "0x6915f16f8791d0a1cc2bf47c13a6b2a92000504b",
				"txId": "0xdf33b22bdb2b28b1f75ccd201a4a4m6e7g83jy5fc5d5a9d1340961598cfcb0a1",
				"status": 1
			},
			{
				"insertTime": 1508298532000,
				"amount": 1000,
				"asset": "XMR",
				"address": "463tWEBn5XZJSxLU34r6g7h8jtxuNcDbjLSjkn3XAXHCbLrTTErJrBWYgHJQyrCwkNgYvyV3z8zctJLPCZy24jvb3NiTcTJ",
				"addressTag": "342341222",
				"txId": "b3c6219639c8ae3f9cf010cdc24fw7f7yt8j1e063f9b4bd1a05cb44c4b6e2509",
				"status": 1
			}
		],
		"success": true
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":     "BTC",
			"status":    1,
			"startTime": 1508198532000,
			"endTime":   1508198532001,
		})
		s.assertRequestEqual(e, r)
	})

	deposits, err := s.client.NewListDepositsService().
		Asset("BTC").
		Status(1).
		StartTime(1508198532000).
		EndTime(1508198532001).
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(deposits, 2)
	s.assertDepositEqual(&Deposit{
		InsertTime: 1508198532000,
		Amount:     0.04670582,
		Asset:      "ETH",
		Address:    "0x6915f16f8791d0a1cc2bf47c13a6b2a92000504b",
		AddressTag: "",
		TxID:       "0xdf33b22bdb2b28b1f75ccd201a4a4m6e7g83jy5fc5d5a9d1340961598cfcb0a1",
		Status:     1,
	}, deposits[0])
	s.assertDepositEqual(&Deposit{
		InsertTime: 1508298532000,
		Amount:     1000.0,
		Asset:      "XMR",
		Address:    "463tWEBn5XZJSxLU34r6g7h8jtxuNcDbjLSjkn3XAXHCbLrTTErJrBWYgHJQyrCwkNgYvyV3z8zctJLPCZy24jvb3NiTcTJ",
		AddressTag: "342341222",
		TxID:       "b3c6219639c8ae3f9cf010cdc24fw7f7yt8j1e063f9b4bd1a05cb44c4b6e2509",
		Status:     1,
	}, deposits[1])
}

func (s *depositServiceTestSuite) assertDepositEqual(e, a *Deposit) {
	r := s.r()
	r.Equal(e.InsertTime, a.InsertTime, "InsertTime")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.InDelta(e.Amount, a.Amount, 0.0000000001, "Amount")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TxID, a.TxID, "TxID")
}

func (s *depositServiceTestSuite) TestGetDepositAddress() {
	data := []byte(`
	{
		"address": "0xbf1f86b3c8ff4f8cbfc195e9713b6f0000000000",
		"success": true,
		"addressTag": "1231212",
		"asset": "ETH",
		"url": "https://etherscan.io/address/0xbf1f86b3c8ff4f8cbfc195e9713b6f0000000000"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "ETH"
	status := true
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":  asset,
			"status": status,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetDepositAddressService().
		Asset(asset).
		Status(status).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.True(res.Success)
	r.Equal("0xbf1f86b3c8ff4f8cbfc195e9713b6f0000000000", res.Address)
	r.Equal("1231212", res.AddressTag)
	r.Equal("ETH", res.Asset)
	r.Equal("https://etherscan.io/address/0xbf1f86b3c8ff4f8cbfc195e9713b6f0000000000", res.URL)
}
