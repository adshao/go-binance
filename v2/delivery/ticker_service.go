package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListBookTickersService list best price/qty on the order book for a symbol or symbols.
type ListBookTickersService struct {
	c      *Client
	symbol *string
	pair   *string
}

// Symbol set symbol.
func (s *ListBookTickersService) Symbol(symbol string) *ListBookTickersService {
	s.symbol = &symbol
	return s
}

// Pair set pair.
func (s *ListBookTickersService) Pair(pair string) *ListBookTickersService {
	s.pair = &pair
	return s
}

// Do send request.
func (s *ListBookTickersService) Do(ctx context.Context, opts ...RequestOption) (res []*BookTicker, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/ticker/bookTicker",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.pair != nil {
		r.setParam("pair", *s.pair)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*BookTicker{}, err
	}
	res = make([]*BookTicker, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*BookTicker{}, err
	}
	return res, nil
}

// BookTicker define book ticker info.
type BookTicker struct {
	Symbol      string `json:"symbol"`
	Pair        string `json:"pair"`
	BidPrice    string `json:"bidPrice"`
	BidQuantity string `json:"bidQty"`
	AskPrice    string `json:"askPrice"`
	AskQuantity string `json:"askQty"`
}

// ListPricesService list latest price for a symbol or symbols.
type ListPricesService struct {
	c      *Client
	symbol *string
	pair   *string
}

// Symbol set symbol.
func (s *ListPricesService) Symbol(symbol string) *ListPricesService {
	s.symbol = &symbol
	return s
}

// Pair set pair.
func (s *ListPricesService) Pair(pair string) *ListPricesService {
	s.pair = &pair
	return s
}

// Do send request.
func (s *ListPricesService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolPrice, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/ticker/price",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.pair != nil {
		r.setParam("pair", *s.pair)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*SymbolPrice{}, err
	}
	res = make([]*SymbolPrice, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*SymbolPrice{}, err
	}
	return res, nil
}

// SymbolPrice define symbol, price and pair.
type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Pair   string `json:"ps"`
	Price  string `json:"price"`
}

// ListPriceChangeStatsService show stats of price change in last 24 hours for single symbol, all symbols or pairs of symbols.
type ListPriceChangeStatsService struct {
	c      *Client
	symbol *string
	pair   *string
}

// Symbol set symbol.
func (s *ListPriceChangeStatsService) Symbol(symbol string) *ListPriceChangeStatsService {
	s.symbol = &symbol
	return s
}

// Pair set pair.
func (s *ListPriceChangeStatsService) Pair(pair string) *ListPriceChangeStatsService {
	s.pair = &pair
	return s
}

// Do send request.
func (s *ListPriceChangeStatsService) Do(ctx context.Context, opts ...RequestOption) (res []*PriceChangeStats, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/ticker/24hr",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.pair != nil {
		r.setParam("pair", *s.pair)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	res = make([]*PriceChangeStats, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PriceChangeStats define price change stats.
type PriceChangeStats struct {
	Symbol             string `json:"symbol"`
	Pair               string `json:"pair"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	LastPrice          string `json:"lastPrice"`
	LastQuantity       string `json:"lastQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	BaseVolume         string `json:"baseVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FristID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}
