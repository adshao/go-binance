package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

// HistoricalTradesService trades
type HistoricalTradesService struct {
	c      *Client
	symbol string
	limit  *int
	fromID *int64
}

// Symbol set symbol
func (s *HistoricalTradesService) Symbol(symbol string) *HistoricalTradesService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *HistoricalTradesService) Limit(limit int) *HistoricalTradesService {
	s.limit = &limit
	return s
}

// FromID set fromID
func (s *HistoricalTradesService) FromID(fromID int64) *HistoricalTradesService {
	s.fromID = &fromID
	return s
}

// Do send request
func (s *HistoricalTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*Trade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/historicalTrades",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
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
	ID            int64  `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"qty"`
	QuoteQuantity string `json:"quoteQty"`
	Time          int64  `json:"time"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
}

// TradeV3 define v3 trade info
type TradeV3 struct {
	ID              int64  `json:"id"`
	Symbol          string `json:"symbol"`
	OrderID         int64  `json:"orderId"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	QuoteQuantity   string `json:"quoteQty"`
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

// Symbol set symbol
func (s *AggTradesService) Symbol(symbol string) *AggTradesService {
	s.symbol = symbol
	return s
}

// FromID set fromID
func (s *AggTradesService) FromID(fromID int64) *AggTradesService {
	s.fromID = &fromID
	return s
}

// StartTime set startTime
func (s *AggTradesService) StartTime(startTime int64) *AggTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *AggTradesService) EndTime(endTime int64) *AggTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *AggTradesService) Limit(limit int) *AggTradesService {
	s.limit = &limit
	return s
}

// Do send request
func (s *AggTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*AggTrade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/aggTrades",
	}
	r.setParam("symbol", s.symbol)
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
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
		return []*AggTrade{}, err
	}
	res = make([]*AggTrade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*AggTrade{}, err
	}
	return res, nil
}

// AggTrade define aggregate trade info
type AggTrade struct {
	AggTradeID   int64  `json:"a"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	FirstTradeID int64  `json:"f"`
	LastTradeID  int64  `json:"l"`
	Timestamp    int64  `json:"T"`
	IsBuyerMaker bool   `json:"m"`
}

// RecentTradesService list recent trades
type RecentTradesService struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *RecentTradesService) Symbol(symbol string) *RecentTradesService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *RecentTradesService) Limit(limit int) *RecentTradesService {
	s.limit = &limit
	return s
}

// Do send request
func (s *RecentTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*Trade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/trades",
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Trade{}, err
	}
	res = make([]*Trade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Trade{}, err
	}
	return res, nil
}

// ListAccountTradeService define account trade list service
type ListAccountTradeService struct {
	c         *Client
	symbol    string
	startTime *int64
	endTime   *int64
	fromID    *int64
	limit     *int
}

// Symbol set symbol
func (s *ListAccountTradeService) Symbol(symbol string) *ListAccountTradeService {
	s.symbol = symbol
	return s
}

// StartTime set startTime
func (s *ListAccountTradeService) StartTime(startTime int64) *ListAccountTradeService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListAccountTradeService) EndTime(endTime int64) *ListAccountTradeService {
	s.endTime = &endTime
	return s
}

// FromID set fromID
func (s *ListAccountTradeService) FromID(fromID int64) *ListAccountTradeService {
	s.fromID = &fromID
	return s
}

// Limit set limit
func (s *ListAccountTradeService) Limit(limit int) *ListAccountTradeService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListAccountTradeService) Do(ctx context.Context, opts ...RequestOption) (res []*AccountTrade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/userTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromID", *s.fromID)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*AccountTrade{}, err
	}
	res = make([]*AccountTrade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*AccountTrade{}, err
	}
	return res, nil
}

// AccountTrade define account trade
type AccountTrade struct {
	Buyer           bool             `json:"buyer"`
	Commission      string           `json:"commission"`
	CommissionAsset string           `json:"commissionAsset"`
	ID              int64            `json:"id"`
	Maker           bool             `json:"maker"`
	OrderID         int64            `json:"orderId"`
	Price           string           `json:"price"`
	Quantity        string           `json:"qty"`
	QuoteQuantity   string           `json:"quoteQty"`
	RealizedPnl     string           `json:"realizedPnl"`
	Side            SideType         `json:"side"`
	PositionSide    PositionSideType `json:"positionSide"`
	Symbol          string           `json:"symbol"`
	Time            int64            `json:"time"`
}
