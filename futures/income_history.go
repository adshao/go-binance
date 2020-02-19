package futures

import (
	"context"
	"encoding/json"
)

// GetIncomeHistoryService get position margin history service
type GetIncomeHistoryService struct {
	c          *Client
	symbol     string
	incomeType string
	startTime  *int64
	endTime    *int64
	limit      *int64
}

// Symbol set symbol
func (s *GetIncomeHistoryService) Symbol(symbol string) *GetIncomeHistoryService {
	s.symbol = symbol
	return s
}

// IncomeType set income type
func (s *GetIncomeHistoryService) IncomeType(incomeType string) *GetIncomeHistoryService {
	s.incomeType = incomeType
	return s
}

// StartTime set startTime
func (s *GetIncomeHistoryService) StartTime(startTime int64) *GetIncomeHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetIncomeHistoryService) EndTime(endTime int64) *GetIncomeHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetIncomeHistoryService) Limit(limit int64) *GetIncomeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetIncomeHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*IncomeHistory, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/income",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.incomeType != "" {
		r.setParam("incomeType", s.incomeType)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*IncomeHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IncomeHistory define position margin history info
type IncomeHistory struct {
	Asset      string `json:"asset"`
	Income     string `json:"income"`
	IncomeType string `json:"incomeType"`
	Info       string `json:"info"`
	Symbol     string `json:"symbol"`
	Time       int64  `json:"time"`
}
