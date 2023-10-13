package binance

import (
	"context"
	"net/http"
)

type FlexibleLoanRepayService struct {
	c                *Client
	loanCoin         string
	collateralCoin   string
	repayAmount      float64
	collateralReturn *bool
	fullRepayment    *bool
}

type FlexibleLoanRepayStatus string

const (
	FlexibleLoanRepaid   FlexibleBorrowStatus = "Repaid"
	FlexibleLoanRepaying FlexibleBorrowStatus = "Repaying"
	FlexibleLoanFailed   FlexibleBorrowStatus = "Failed"
)

type FlexibleLoanRepayResp struct {
	LoanCoin            string                  `json:"loanCoin"`
	CollateralCoin      string                  `json:"collateralCoin"`
	RemainingDebt       string                  `json:"remainingDebt"`
	RemainingCollateral string                  `json:"remainingCollateral"`
	FullRepayment       bool                    `json:"fullRepayment"`
	CurrentLTV          string                  `json:"currentLTV"`
	RepayStatus         FlexibleLoanRepayStatus `json:"repayStatus"`
}

// LoanCoin set loanCoin
func (s *FlexibleLoanRepayService) LoanCoin(loanCoin string) *FlexibleLoanRepayService {
	s.loanCoin = loanCoin
	return s
}

// CollateralCoin set collateralCoin
func (s *FlexibleLoanRepayService) CollateralCoin(coll string) *FlexibleLoanRepayService {
	s.collateralCoin = coll
	return s
}

// RepayAmount set repayAmount
func (s *FlexibleLoanRepayService) RepayAmount(amt float64) *FlexibleLoanRepayService {
	s.repayAmount = amt
	return s
}

// CollateralReturn set collateralReturn
func (s *FlexibleLoanRepayService) CollateralReturn(collReturn bool) *FlexibleLoanRepayService {
	s.collateralReturn = &collReturn
	return s
}

// FullRepayment set fullRepayment
func (s *FlexibleLoanRepayService) FullRepayment(fullRepayment bool) *FlexibleLoanRepayService {
	s.fullRepayment = &fullRepayment
	return s
}

// Do send request
func (s *FlexibleLoanRepayService) Do(ctx context.Context, opts ...RequestOption) (res *FlexibleLoanRepayResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/loan/flexible/repay",
		secType:  secTypeSigned,
	}

	r.setParam("loanCoin", s.loanCoin)
	r.setParam("collateralCoin", s.collateralCoin)
	r.setParam("repayAmount", s.repayAmount)

	if s.collateralReturn != nil {
		r.setParam("collateralReturn", *s.collateralReturn)
	}

	if s.fullRepayment != nil {
		r.setParam("fullRepayment", *s.fullRepayment)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(FlexibleLoanRepayResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
