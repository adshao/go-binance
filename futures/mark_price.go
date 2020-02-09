package futures

import (
	"context"
	"encoding/json"
)

// PremiumIndexService get premium index
type PremiumIndexService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *PremiumIndexService) Symbol(symbol string) *PremiumIndexService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *PremiumIndexService) Do(ctx context.Context, opts ...RequestOption) (res *PremiumIndex, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/premiumIndex",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(PremiumIndex)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PremiumIndex define premium index of mark price
type PremiumIndex struct {
	Symbol          string `json:"symbol"`
	MarkPrice       string `json:"markPrice"`
	LastFundingRate string `json:"lastFundingRate"`
	NextFundingTime int64  `json:"nextFundingTime"`
	Time            int64  `json:"time"`
}

// FundingRateService get funding rate
type FundingRateService struct {
	c         *Client
	symbol    string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *FundingRateService) Symbol(symbol string) *FundingRateService {
	s.symbol = symbol
	return s
}

// StartTime set startTime
func (s *FundingRateService) StartTime(startTime int64) *FundingRateService {
	s.startTime = &startTime
	return s
}

// EndTime set startTime
func (s *FundingRateService) EndTime(endTime int64) *FundingRateService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *FundingRateService) Limit(limit int) *FundingRateService {
	s.limit = &limit
	return s
}

// Do send request
func (s *FundingRateService) Do(ctx context.Context, opts ...RequestOption) (res []*FundingRate, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/fundingRate",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
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
		return []*FundingRate{}, err
	}
	res = make([]*FundingRate, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*FundingRate{}, err
	}
	return res, nil
}

// FundingRate define funding rate of mark price
type FundingRate struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64  `json:"fundingTime"`
	Time        int64  `json:"time"`
}
