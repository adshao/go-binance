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
