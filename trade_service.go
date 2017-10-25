package binance

import (
	"context"
	"encoding/json"
)

// ListTradesService list trades
type ListTradesService struct {
	c      *Client
	symbol string
	limit  *int
	fromID *int64
}

func (s *ListTradesService) Symbol(symbol string) *ListTradesService {
	s.symbol = symbol
	return s
}

func (s *ListTradesService) Limit(limit int) *ListTradesService {
	s.limit = &limit
	return s
}

func (s *ListTradesService) FromID(fromID int64) *ListTradesService {
	s.fromID = &fromID
	return s
}

func (s *ListTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*Trade, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/myTrades",
		secType:  secTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}
	if s.fromID != nil {
		r.SetParam("fromId", *s.fromID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*Trade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// Trade define trade info
type Trade struct {
	ID              int64  `json:"id"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

// AggTradesService list aggregate trades
type AggTradesService struct {
	c         *Client
	symbol    string
	fromID    *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

func (s *AggTradesService) Symbol(symbol string) *AggTradesService {
	s.symbol = symbol
	return s
}

func (s *AggTradesService) FromID(fromID int64) *AggTradesService {
	s.fromID = &fromID
	return s
}

func (s *AggTradesService) StartTime(startTime int64) *AggTradesService {
	s.startTime = &startTime
	return s
}

func (s *AggTradesService) EndTime(endTime int64) *AggTradesService {
	s.endTime = &endTime
	return s
}

func (s *AggTradesService) Limit(limit int) *AggTradesService {
	s.limit = &limit
	return s
}

func (s *AggTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*AggTrade, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/aggTrades",
	}
	r.SetParam("symbol", s.symbol)
	if s.fromID != nil {
		r.SetParam("fromId", *s.fromID)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*AggTrade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// AggTrade define aggregate trade info
type AggTrade struct {
	AggTradeID       int64  `json:"a"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int64  `json:"f"`
	LastTradeID      int64  `json:"l"`
	Timestamp        int64  `json:"T"`
	IsBuyerMaker     bool   `json:"m"`
	IsBestPriceMatch bool   `json:"M"`
}
