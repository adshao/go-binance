package binance

import (
	"context"
	"encoding/json"
)

// ExchangeInfoService exchange info service
type ExchangeInfoService struct {
	c *Client
}

// Do send request
func (s *ExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeInfo, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/exchangeInfo",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(ExchangeInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return
}

// ExchangeInfo exchange info
type ExchangeInfo struct {
	Symbols []Symbol `json:"symbols"`
}

// Symbol market symbol
type Symbol struct {
	Symbol             string `json:"symbol"`
	BaseAsset          string `json:"baseAsset"`
	BaseAssetPrecision int    `json:"baseAssetPrecision"`
	QuoteAsset         string `json:"quoteAsset"`
	QuotePrecision     int    `json:"quotePrecision"`
}
