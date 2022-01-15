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
	data := []byte(`[
    {
        "amount":"0.00999800",
        "coin":"PAXG",
        "network":"ETH",
        "status":1,
        "address":"0x788cabe9236ce061e5a892e1a59395a81fc8d62c",
        "addressTag":"",
        "txId":"0xaad4654a3234aa6118af9b4b335f5ae81c360b2394721c019b5d1e75328b09f3",
        "insertTime":1599621997000,
        "transferType":0,
        "unlockConfirm":"12/12",
        "confirmTimes":"12/12"
    },
    {
        "amount":"0.50000000",
        "coin":"IOTA",
        "network":"IOTA",
        "status":1,
        "address":"SIZ9VLMHWATXKV99LH99CIGFJFUMLEHGWVZVNNZXRJJVWBPHYWPPBOSDORZ9EQSHCZAMPVAPGFYQAUUV9DROOXJLNW",
        "addressTag":"",
        "txId":"ESBFVQUTPIWQNJSPXFNHNYHSQNTGKRVKPRABQWTAXCDWOAKDKYWPTVG9BGXNVNKTLEJGESAVXIKIZ9999",
        "insertTime":1599620082000,
        "transferType":0,
        "unlockConfirm":"1/12",
        "confirmTimes":"1/1"
    }
]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":      "BTC",
			"status":    1,
			"startTime": 1508198532000,
			"endTime":   1508198532001,
			"offset":    0,
			"limit":     1000,
		})
		s.assertRequestEqual(e, r)
	})

	deposits, err := s.client.NewListDepositsService().
		Coin("BTC").
		Status(1).
		StartTime(1508198532000).
		EndTime(1508198532001).
		Offset(0).
		Limit(1000).
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(deposits, 2)
	s.assertDepositEqual(&Deposit{
		Amount:        "0.00999800",
		Coin:          "PAXG",
		Network:       "ETH",
		Status:        1,
		Address:       "0x788cabe9236ce061e5a892e1a59395a81fc8d62c",
		AddressTag:    "",
		TxID:          "0xaad4654a3234aa6118af9b4b335f5ae81c360b2394721c019b5d1e75328b09f3",
		InsertTime:    1599621997000,
		TransferType:  0,
		UnlockConfirm: "12/12",
		ConfirmTimes:  "12/12",
	}, deposits[0])
	s.assertDepositEqual(&Deposit{
		Amount:        "0.50000000",
		Coin:          "IOTA",
		Network:       "IOTA",
		Status:        1,
		Address:       "SIZ9VLMHWATXKV99LH99CIGFJFUMLEHGWVZVNNZXRJJVWBPHYWPPBOSDORZ9EQSHCZAMPVAPGFYQAUUV9DROOXJLNW",
		AddressTag:    "",
		TxID:          "ESBFVQUTPIWQNJSPXFNHNYHSQNTGKRVKPRABQWTAXCDWOAKDKYWPTVG9BGXNVNKTLEJGESAVXIKIZ9999",
		InsertTime:    1599620082000,
		TransferType:  0,
		UnlockConfirm: "1/12",
		ConfirmTimes:  "1/1",
	}, deposits[1])
}

func (s *depositServiceTestSuite) assertDepositEqual(e, a *Deposit) {
	r := s.r()
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.Network, a.Network, "Network")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.Address, a.Address, "Address")
	r.Equal(e.AddressTag, a.AddressTag, "AddressTag")
	r.Equal(e.TxID, a.TxID, "TxID")
	r.Equal(e.InsertTime, a.InsertTime, "InsertTime")
	r.Equal(e.TransferType, a.TransferType, "TransferType")
	r.Equal(e.UnlockConfirm, a.UnlockConfirm, "UnlockConfirm")
	r.Equal(e.ConfirmTimes, a.ConfirmTimes, "ConfirmTimes")
}

func (s *depositServiceTestSuite) TestGetDepositAddress() {
	data := []byte(`
	{
		"address": "1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv",
		"coin": "BTC",
		"tag": "",
		"url": "https://btc.com/1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "BTC"
	network := "BTC"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":    coin,
			"network": network,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetDepositAddressService().
		Coin(coin).
		Network(network).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv", res.Address)
	r.Equal("", res.Tag)
	r.Equal("BTC", res.Coin)
	r.Equal("https://btc.com/1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv", res.URL)
}
