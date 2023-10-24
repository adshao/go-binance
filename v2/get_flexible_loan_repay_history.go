package binance

import (
	"context"
	"net/http"
)

type GetFlexibleLoanRepayHistoryService struct {
	c              *Client
	loanCoin       *string
	collateralCoin *string
	startTime      *int64
	endTime        *int64
	current        *int64
	limit          *int64
}

type GetFlexibleLoanRepayHistoryResp struct {
	Rows  []FlexibleLoanRepayHistory `json:"rows"`
	Total int64                      `json:"total"`
}

type FlexibleLoanRepayHistory struct {
	LoanCoin         string                  `json:"loanCoin"`
	RepayAmount      string                  `json:"repayAmount"`
	CollateralCoin   string                  `json:"collateralCoin"`
	CollateralReturn string                  `json:"collateralReturn"`
	RepayTime        interface{}             `json:"repayTime"`
	RepayStatus      FlexibleLoanRepayStatus `json:"repayStatus"`
}

// LoanCoin set loanCoin
func (s *GetFlexibleLoanRepayHistoryService) LoanCoin(loanCoin string) *GetFlexibleLoanRepayHistoryService {
	s.loanCoin = &loanCoin
	return s
}

// CollateralCoin set collateralCoin
func (s *GetFlexibleLoanRepayHistoryService) CollateralCoin(coll string) *GetFlexibleLoanRepayHistoryService {
	s.collateralCoin = &coll
	return s
}

// Current set current
func (s *GetFlexibleLoanRepayHistoryService) Current(current int64) *GetFlexibleLoanRepayHistoryService {
	s.current = &current
	return s
}

// Limit set limit
func (s *GetFlexibleLoanRepayHistoryService) Limit(limit int64) *GetFlexibleLoanRepayHistoryService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *GetFlexibleLoanRepayHistoryService) StartTime(startTime int64) *GetFlexibleLoanRepayHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetFlexibleLoanRepayHistoryService) EndTime(endTime int64) *GetFlexibleLoanRepayHistoryService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *GetFlexibleLoanRepayHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *GetFlexibleLoanRepayHistoryResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/flexible/repay/history",
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

	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}

	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(GetFlexibleLoanRepayHistoryResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
