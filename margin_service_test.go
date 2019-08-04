package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginTestSuite struct {
	baseTestSuite
}

func TestMarginAccountService(t *testing.T) {
	suite.Run(t, new(marginTestSuite))
}

func (s *marginTestSuite) TestTransfer() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	transferType := MarginTransferTypeToMargin
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
			"type":   transferType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginTransferService().Asset(asset).
		Amount(amount).Type(transferType).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) assertTransactionResponseEqual(a, e *TransactionResponse) {
	s.r().Equal(a.TranID, e.TranID, "TranID")
}

func (s *marginTestSuite) TestLoan() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginLoanService().Asset(asset).
		Amount(amount).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) TestRepay() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginRepayService().Asset(asset).
		Amount(amount).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) TestListMarginLoans() {
	data := []byte(`{
		"rows": [
		  {
			"asset": "BNB",
			"principal": "0.84624403",
			"timestamp": 1555056425000,
			"status": "CONFIRMED"
		  }
		],
		"total": 1
	  }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BNB"
	txID := int64(1)
	startTime := int64(1555056425000)
	endTime := int64(1555056425001)
	current := int64(1)
	size := int64(10)
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"asset":     asset,
			"txId":      txID,
			"startTime": startTime,
			"endTime":   endTime,
			"current":   current,
			"size":      size,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListMarginLoansService().Asset(asset).
		TxID(txID).StartTime(startTime).EndTime(endTime).
		Current(current).Size(size).Do(newContext())
	s.r().NoError(err)
	e := &MarginLoanResponse{
		Rows: []MarginLoan{
			{
				Asset:     asset,
				Principal: "0.84624403",
				Timestamp: 1555056425000,
				Status:    MarginLoanStatusTypeConfirmed,
			},
		},
		Total: 1,
	}
	s.assertMarginLoanResponseEqual(e, res)
}

func (s *marginTestSuite) assertMarginLoanResponseEqual(e, a *MarginLoanResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	r.Len(a.Rows, len(e.Rows), "Rows")
	for i := 0; i < len(e.Rows); i++ {
		s.assertMarginLoanEqual(&e.Rows[i], &a.Rows[i])
	}
}

func (s *marginTestSuite) assertMarginLoanEqual(e, a *MarginLoan) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Principal, a.Principal, "Principal")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.Status, a.Status, "Status")
}
