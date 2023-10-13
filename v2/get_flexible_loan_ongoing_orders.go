package binance

import (
	"context"
	"net/http"
)

type GetFlexibleLoanOngoingOrdersService struct {
	c              *Client
	loanCoin       *string
	collateralCoin *string
	current        *int64
	limit          *int64
}

type GetFlexibleLoanOngoingOrdersResp struct {
	Rows  []FlexibleLoanOngoingOrder `json:"rows"`
	Total int64                      `json:"total"`
}

type FlexibleLoanOngoingOrder struct {
	LoanCoin         string `json:"loanCoin"`
	TotalDebt        string `json:"totalDebt"`
	CollateralCoin   string `json:"collateralCoin"`
	CollateralAmount string `json:"collateralAmount"`
	CurrentLTV       string `json:"currentLTV"`
}

// LoanCoin set loanCoin
func (s *GetFlexibleLoanOngoingOrdersService) LoanCoin(loanCoin string) *GetFlexibleLoanOngoingOrdersService {
	s.loanCoin = &loanCoin
	return s
}

// CollateralCoin set collateralCoin
func (s *GetFlexibleLoanOngoingOrdersService) CollateralCoin(coll string) *GetFlexibleLoanOngoingOrdersService {
	s.collateralCoin = &coll
	return s
}

// Current set current
func (s *GetFlexibleLoanOngoingOrdersService) Current(current int64) *GetFlexibleLoanOngoingOrdersService {
	s.current = &current
	return s
}

// Limit set limit
func (s *GetFlexibleLoanOngoingOrdersService) Limit(limit int64) *GetFlexibleLoanOngoingOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetFlexibleLoanOngoingOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *GetFlexibleLoanOngoingOrdersResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/flexible/ongoing/orders",
		secType:  secTypeSigned,
	}

	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}

	if s.collateralCoin != nil {
		r.setParam("collateralCoin", *s.collateralCoin)
	}

	if s.current != nil {
		r.setParam("current", *s.current)
	}

	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(GetFlexibleLoanOngoingOrdersResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
