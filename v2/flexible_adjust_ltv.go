package binance

import (
	"context"
	"net/http"
)

type FLAdjustLTVDirection string

const (
	FLAdjustLTVAdditional FLAdjustLTVDirection = "ADDITIONAL"
	FLAdjustLTVReduced    FLAdjustLTVDirection = "REDUCED"
)

type FlexibleLoanAdjustLTVService struct {
	c                *Client
	loanCoin         string
	collateralCoin   string
	adjustmentAmount float64
	direction        FLAdjustLTVDirection
}

type FLAdjustLTVStatus string

const (
	FLAdjustLTVSucceeds   FLAdjustLTVStatus = "Succeeds"
	FLAdjustLTVFailed     FLAdjustLTVStatus = "Failed"
	FLAdjustLTVProcessing FLAdjustLTVStatus = "Processing"
)

type FlexibleLoanAdjustLTVResp struct {
	LoanCoin         string            `json:"loanCoin"`
	CollateralCoin   string            `json:"collateralCoin"`
	Direction        string            `json:"direction"`
	AdjustmentAmount string            `json:"adjustmentAmount"`
	CurrentLTV       string            `json:"currentLTV"`
	Status           FLAdjustLTVStatus `json:"status"`
}

// LoanCoin set loanCoin
func (s *FlexibleLoanAdjustLTVService) LoanCoin(loanCoin string) *FlexibleLoanAdjustLTVService {
	s.loanCoin = loanCoin
	return s
}

// CollateralCoin set collateralCoin
func (s *FlexibleLoanAdjustLTVService) CollateralCoin(coll string) *FlexibleLoanAdjustLTVService {
	s.collateralCoin = coll
	return s
}

// AdjustmentAmount set adjustmentAmount
func (s *FlexibleLoanAdjustLTVService) AdjustmentAmount(amt float64) *FlexibleLoanAdjustLTVService {
	s.adjustmentAmount = amt
	return s
}

// Direction set direction
func (s *FlexibleLoanAdjustLTVService) Direction(direction FLAdjustLTVDirection) *FlexibleLoanAdjustLTVService {
	s.direction = direction
	return s
}

// Do send request
func (s *FlexibleLoanAdjustLTVService) Do(ctx context.Context, opts ...RequestOption) (res *FlexibleLoanAdjustLTVResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/loan/flexible/adjust/ltv",
		secType:  secTypeSigned,
	}

	r.setParam("loanCoin", s.loanCoin)
	r.setParam("collateralCoin", s.collateralCoin)
	r.setParam("adjustmentAmount", s.adjustmentAmount)
	r.setParam("direction", s.direction)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(FlexibleLoanAdjustLTVResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
