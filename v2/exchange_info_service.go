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
		endpoint: "/api/v3/exchangeInfo",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ExchangeInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ExchangeInfo exchange info
type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	RateLimits      []RateLimit   `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []Symbol      `json:"symbols"`
}

// RateLimit struct
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	Limit         int64  `json:"limit"`
}

// Symbol market symbol
type Symbol struct {
	Symbol                 string                   `json:"symbol"`
	Status                 string                   `json:"status"`
	BaseAsset              string                   `json:"baseAsset"`
	BaseAssetPrecision     int                      `json:"baseAssetPrecision"`
	QuoteAsset             string                   `json:"quoteAsset"`
	QuotePrecision         int                      `json:"quotePrecision"`
	OrderTypes             []string                 `json:"orderTypes"`
	IcebergAllowed         bool                     `json:"icebergAllowed"`
	OcoAllowed             bool                     `json:"ocoAllowed"`
	IsSpotTradingAllowed   bool                     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed bool                     `json:"isMarginTradingAllowed"`
	Filters                []map[string]interface{} `json:"filters"`
}

// LotSizeFilter define lot size filter of symbol
type LotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

// PriceFilter define price filter of symbol
type PriceFilter struct {
	MaxPrice string `json:"maxPrice"`
	MinPrice string `json:"minPrice"`
	TickSize string `json:"tickSize"`
}

// PercentPriceFilter define percent price filter of symbol
type PercentPriceFilter struct {
	AveragePriceMins int    `json:"avgPriceMins"`
	MultiplierUp     string `json:"multiplierUp"`
	MultiplierDown   string `json:"multiplierDown"`
}

// MinNotionalFilter define min notional filter of symbol
type MinNotionalFilter struct {
	MinNotional      string `json:"minNotional"`
	AveragePriceMins int    `json:"avgPriceMins"`
	ApplyToMarket    bool   `json:"applyToMarket"`
}

// IcebergPartsFilter define iceberg part filter of symbol
type IcebergPartsFilter struct {
	Limit int `json:"limit"`
}

// MarketLotSizeFilter define market lot size filter of symbol
type MarketLotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

// MaxNumAlgoOrdersFilter define max num algo orders filter of symbol
type MaxNumAlgoOrdersFilter struct {
	MaxNumAlgoOrders int `json:"maxNumAlgoOrders"`
}

// LotSizeFilter return lot size filter of symbol
func (s *Symbol) LotSizeFilter() *LotSizeFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeLotSize) {
			f := &LotSizeFilter{}
			if i, ok := filter["maxQty"]; ok {
				f.MaxQuantity = i.(string)
			}
			if i, ok := filter["minQty"]; ok {
				f.MinQuantity = i.(string)
			}
			if i, ok := filter["stepSize"]; ok {
				f.StepSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// PriceFilter return price filter of symbol
func (s *Symbol) PriceFilter() *PriceFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypePriceFilter) {
			f := &PriceFilter{}
			if i, ok := filter["maxPrice"]; ok {
				f.MaxPrice = i.(string)
			}
			if i, ok := filter["minPrice"]; ok {
				f.MinPrice = i.(string)
			}
			if i, ok := filter["tickSize"]; ok {
				f.TickSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// PercentPriceFilter return percent price filter of symbol
func (s *Symbol) PercentPriceFilter() *PercentPriceFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypePercentPrice) {
			f := &PercentPriceFilter{}
			if i, ok := filter["avgPriceMins"]; ok {
				f.AveragePriceMins = int(i.(float64))
			}
			if i, ok := filter["multiplierUp"]; ok {
				f.MultiplierUp = i.(string)
			}
			if i, ok := filter["multiplierDown"]; ok {
				f.MultiplierDown = i.(string)
			}
			return f
		}
	}
	return nil
}

// MinNotionalFilter return min notional filter of symbol
func (s *Symbol) MinNotionalFilter() *MinNotionalFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMinNotional) {
			f := &MinNotionalFilter{}
			if i, ok := filter["minNotional"]; ok {
				f.MinNotional = i.(string)
			}
			if i, ok := filter["avgPriceMins"]; ok {
				f.AveragePriceMins = int(i.(float64))
			}
			if i, ok := filter["applyToMarket"]; ok {
				f.ApplyToMarket = i.(bool)
			}
			return f
		}
	}
	return nil
}

// IcebergPartsFilter return iceberg part filter of symbol
func (s *Symbol) IcebergPartsFilter() *IcebergPartsFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeIcebergParts) {
			f := &IcebergPartsFilter{}
			if i, ok := filter["limit"]; ok {
				f.Limit = int(i.(float64))
			}
			return f
		}
	}
	return nil
}

// MarketLotSizeFilter return market lot size filter of symbol
func (s *Symbol) MarketLotSizeFilter() *MarketLotSizeFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMarketLotSize) {
			f := &MarketLotSizeFilter{}
			if i, ok := filter["maxQty"]; ok {
				f.MaxQuantity = i.(string)
			}
			if i, ok := filter["minQty"]; ok {
				f.MinQuantity = i.(string)
			}
			if i, ok := filter["stepSize"]; ok {
				f.StepSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// MaxNumAlgoOrdersFilter return max num algo orders filter of symbol
func (s *Symbol) MaxNumAlgoOrdersFilter() *MaxNumAlgoOrdersFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMaxNumAlgoOrders) {
			f := &MaxNumAlgoOrdersFilter{}
			if i, ok := filter["maxNumAlgoOrders"]; ok {
				f.MaxNumAlgoOrders = int(i.(float64))
			}
			return f
		}
	}
	return nil
}
