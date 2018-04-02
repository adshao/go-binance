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
                "status": 1,
                "TxID": "b3c6219639c8ae3f9cf010cdc24fw7f7yt8j1e063f9b4bd1a05cb44c4b6e2509"
            }
        ],
        "success": true
    }`)
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
	deposits, err := s.client.NewListDepositsService().Asset("BTC").
		Status(1).StartTime(1508198532000).EndTime(1508198532001).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(deposits, 1)
	e := &Deposit{
		InsertTime: 1508198532000,
		Amount:     0.04670582,
		Asset:      "ETH",
		Status:     1,
		TxID:       "b3c6219639c8ae3f9cf010cdc24fw7f7yt8j1e063f9b4bd1a05cb44c4b6e2509",
	}
	s.assertDepositEqual(e, deposits[0])
}

func (s *depositServiceTestSuite) assertDepositEqual(e, a *Deposit) {
	r := s.r()
	r.Equal(e.InsertTime, a.InsertTime, "InsertTime")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.InDelta(e.Amount, a.Amount, 0.0000000001, "Amount")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TxID, a.TxID, "TxID")
}
