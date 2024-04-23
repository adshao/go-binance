package binance

import (
	"context"
	"net/http"
)

type GetFlexibleLoanBorrowHistoryService struct {
	c              *Client
	loanCoin       *string
	collateralCoin *string
	startTime      *int64
	endTime        *int64
	current        *int64
	limit          *int64
}

type GetFlexibleLoanBorrowHistoryResp struct {
	Rows  []FlexibleLoanBorrowHistory `json:"rows"`
	Total int64                       `json:"total"`
}

type FlexibleLoanBorrowHistory struct {
	LoanCoin                string               `json:"loanCoin"`
	InitialLoanAmount       string               `json:"initialLoanAmount"`
	CollateralCoin          string               `json:"collateralCoin"`
	InitialCollateralAmount string               `json:"initialCollateralAmount"`
	BorrowTime              interface{}          `json:"borrowTime"`
	Status                  FlexibleBorrowStatus `json:"status"`
}

// LoanCoin set loanCoin
func (s *GetFlexibleLoanBorrowHistoryService) LoanCoin(loanCoin string) *GetFlexibleLoanBorrowHistoryService {
	s.loanCoin = &loanCoin
	return s
}

// CollateralCoin set collateralCoin
func (s *GetFlexibleLoanBorrowHistoryService) CollateralCoin(coll string) *GetFlexibleLoanBorrowHistoryService {
	s.collateralCoin = &coll
	return s
}

// Current set current
func (s *GetFlexibleLoanBorrowHistoryService) Current(current int64) *GetFlexibleLoanBorrowHistoryService {
	s.current = &current
	return s
}

// Limit set limit
func (s *GetFlexibleLoanBorrowHistoryService) Limit(limit int64) *GetFlexibleLoanBorrowHistoryService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *GetFlexibleLoanBorrowHistoryService) StartTime(startTime int64) *GetFlexibleLoanBorrowHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetFlexibleLoanBorrowHistoryService) EndTime(endTime int64) *GetFlexibleLoanBorrowHistoryService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *GetFlexibleLoanBorrowHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *GetFlexibleLoanBorrowHistoryResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v2/loan/flexible/borrow/history",
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

	res = new(GetFlexibleLoanBorrowHistoryResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
