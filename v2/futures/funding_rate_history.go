package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// FundingRateHistoryService gets funding rate history
type FundingRateHistoryService struct {
	c         *Client
	symbol    *string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *FundingRateHistoryService) Symbol(symbol string) *FundingRateHistoryService {
	s.symbol = &symbol
	return s
}

// StartTime set startTime
func (s *FundingRateHistoryService) StartTime(startTime int64) *FundingRateHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set startTime
func (s *FundingRateHistoryService) EndTime(endTime int64) *FundingRateHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *FundingRateHistoryService) Limit(limit int) *FundingRateHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *FundingRateHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*FundingRateHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/fundingRate",
		secType:  secTypeNone,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
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
	data, _, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return []*FundingRateHistory{}, err
	}
	res = make([]*FundingRateHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*FundingRateHistory{}, err
	}
	return res, nil
}

// FundingRateHistory defines funding rate history data
type FundingRateHistory struct {
	Symbol      string `json:"symbol"`
	FundingTime int64  `json:"fundingTime"`
	FundingRate string `json:"fundingRate"`
	MarkPrice   string `json:"markPrice"`
}
