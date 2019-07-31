package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type exchangeInfoServiceTestSuite struct {
	baseTestSuite
}

func TestExchangeInfoService(t *testing.T) {
	suite.Run(t, new(exchangeInfoServiceTestSuite))
}

func (s *exchangeInfoServiceTestSuite) TestExchangeInfo() {
	data := []byte(`{
  		"timezone": "UTC",
  		"serverTime": 1539281238296,
  		"rateLimits": [
    		{
      			"rateLimitType": "REQUEST_WEIGHT",
				"interval": "MINUTE",
				"limit": 1200
    		},
    		{
      			"rateLimitType": "ORDERS",
      			"interval": "SECOND",
      			"limit": 10
    		},
    		{
      			"rateLimitType": "ORDERS",
      			"interval": "DAY",
      			"limit": 100000
    		}
  		],
		"exchangeFilters": [],
		"symbols": [
			{
				"symbol": "ETHBTC",
				"status": "TRADING",
				"baseAsset": "ETH",
				"baseAssetPrecision": 8,
				"quoteAsset": "BTC",
				"quotePrecision": 8,
				"orderTypes":["LIMIT","LIMIT_MAKER","MARKET","STOP_LOSS_LIMIT","TAKE_PROFIT_LIMIT"],
				"icebergAllowed": true,
				"filters":[{"filterType":"PRICE_FILTER","minPrice":"0.00000100","maxPrice":"100000.00000000","tickSize":"0.00000100"},{"filterType":"LOT_SIZE","minQty":"0.00100000","maxQty":"100000.00000000","stepSize":"0.00100000"},{"filterType":"MIN_NOTIONAL","minNotional":"0.00100000"},{"filterType": "MAX_NUM_ALGO_ORDERS", "maxNumAlgoOrders": 5}]
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewExchangeInfoService().Do(newContext())
	s.r().NoError(err)
	ei := &ExchangeInfo{
		Timezone:   "UTC",
		ServerTime: 1539281238296,
		RateLimits: []RateLimit{
			{RateLimitType: "REQUEST_WEIGHT", Interval: "MINUTE", Limit: 1200},
			{RateLimitType: "ORDERS", Interval: "SECOND", Limit: 10},
			{RateLimitType: "ORDERS", Interval: "DAY", Limit: 100000},
		},
		ExchangeFilters: []interface{}{},
		Symbols: []Symbol{
			{
				Symbol:             "ETHBTC",
				Status:             "TRADING",
				BaseAsset:          "ETH",
				BaseAssetPrecision: 8,
				QuoteAsset:         "BTC",
				QuotePrecision:     8,
				OrderTypes:         []string{"LIMIT", "LIMIT_MAKER", "MARKET", "STOP_LOSS_LIMIT", "TAKE_PROFIT_LIMIT"},
				IcebergAllowed:     true,
				Filters: []map[string]interface{}{
					{"filterType": "PRICE_FILTER", "minPrice": "0.00000100", "maxPrice": "100000.00000000", "tickSize": "0.00000100"},
					{"filterType": "LOT_SIZE", "minQty": "0.00100000", "maxQty": "100000.00000000", "stepSize": "0.00100000"},
					{"filterType": "MIN_NOTIONAL", "minNotional": "0.00100000"},
					{"filterType": "MAX_NUM_ALGO_ORDERS", "maxNumAlgoOrders": 5},
				},
			},
		},
	}
	s.assertExchangeInfoEqual(ei, res)

	eLotSizeFilter := &LotSizeFilter{
		MaxQuantity: "100000.00000000",
		MinQuantity: "0.00100000",
		StepSize:    "0.00100000",
	}
	s.assertLotSizeFilterEqual(eLotSizeFilter, res.Symbols[0].LotSizeFilter())
	ePriceFilter := &PriceFilter{
		MaxPrice: "100000.00000000",
		MinPrice: "0.00000100",
		TickSize: "0.00000100",
	}
	s.assertPriceFilterEqual(ePriceFilter, res.Symbols[0].PriceFilter())
	eMinNotionalFilter := &MinNotionalFilter{
		MinNotional:      "0.00100000",
		AveragePriceMins: 0,
		ApplyToMarket:    false,
	}
	s.assertMinNotionalFilterEqual(eMinNotionalFilter, res.Symbols[0].MinNotionalFilter())
	eMaxNumAlgoOrdersFilter := &MaxNumAlgoOrdersFilter{
		MaxNumAlgoOrders: 5,
	}
	s.assertMaxNumAlgoOrdersFilterEqual(eMaxNumAlgoOrdersFilter, res.Symbols[0].MaxNumAlgoOrdersFilter())
}

func (s *exchangeInfoServiceTestSuite) assertExchangeInfoEqual(e, a *ExchangeInfo) {
	r := s.r()

	r.Equal(e.Timezone, a.Timezone, "Timezone")
	r.Equal(e.ServerTime, a.ServerTime, "ServerTime")

	for i := range a.RateLimits {
		r.Equal(e.RateLimits[i].RateLimitType, a.RateLimits[i].RateLimitType, "RateLimitType")
		r.Equal(e.RateLimits[i].Limit, a.RateLimits[i].Limit, "Limit")
		r.Equal(e.RateLimits[i].Interval, a.RateLimits[i].Interval, "Interval")
	}

	r.Equal(e.ExchangeFilters, a.ExchangeFilters, "ExchangeFilters")

	for i, currentSymbol := range a.Symbols {
		if a.Symbols[i].Symbol == e.Symbols[0].Symbol {
			r.Equal(e.Symbols[i].Status, currentSymbol.Status, "Status")
			r.Equal(e.Symbols[i].BaseAsset, currentSymbol.BaseAsset, "BaseAsset")
			r.Equal(e.Symbols[i].BaseAssetPrecision, currentSymbol.BaseAssetPrecision, "BaseAssetPrecision")
			r.Equal(e.Symbols[i].QuoteAsset, currentSymbol.QuoteAsset, "QuoteAsset")
			r.Equal(e.Symbols[i].QuotePrecision, currentSymbol.QuotePrecision, "QuotePrecision")
			r.Len(currentSymbol.OrderTypes, len(e.Symbols[i].OrderTypes))
			r.Equal(e.Symbols[i].OrderTypes, currentSymbol.OrderTypes, "OrderTypes")
			r.Equal(e.Symbols[i].IcebergAllowed, currentSymbol.IcebergAllowed, "IcebergAllowed")
			r.Len(currentSymbol.Filters, len(e.Symbols[i].Filters))
			for fi, currentFilter := range currentSymbol.Filters {
				r.Len(currentFilter, len(e.Symbols[i].Filters[fi]))
				switch currentFilter["filterType"] {
				case "PRICE_FILTER":
					r.Equal(e.Symbols[i].Filters[fi]["minPrice"], currentFilter["minPrice"], "minPrice")
					r.Equal(e.Symbols[i].Filters[fi]["maxPrice"], currentFilter["maxPrice"], "maxPrice")
					r.Equal(e.Symbols[i].Filters[fi]["tickSize"], currentFilter["tickSize"], "tickSize")
				case "LOT_SIZE":
					r.Equal(e.Symbols[i].Filters[fi]["minQty"], currentFilter["minQty"], "minQty")
					r.Equal(e.Symbols[i].Filters[fi]["maxQty"], currentFilter["maxQty"], "maxQty")
					r.Equal(e.Symbols[i].Filters[fi]["stepSize"], currentFilter["stepSize"], "stepSize")
				case "MIN_NOTIONAL":
					r.Equal(e.Symbols[i].Filters[fi]["minNotional"], currentFilter["minNotional"], "minNotional")
				}

			}

			return
		}

	}
	r.Fail("Symbol ETHBTC not found")
}

func (s *exchangeInfoServiceTestSuite) assertLotSizeFilterEqual(e, a *LotSizeFilter) {
	r := s.r()
	r.Equal(e.MaxQuantity, a.MaxQuantity, "MaxQuantity")
	r.Equal(e.MinQuantity, a.MinQuantity, "MinQuantity")
	r.Equal(e.StepSize, a.StepSize, "StepSize")
}

func (s *exchangeInfoServiceTestSuite) assertPriceFilterEqual(e, a *PriceFilter) {
	r := s.r()
	r.Equal(e.MaxPrice, a.MaxPrice, "MaxPrice")
	r.Equal(e.MinPrice, a.MinPrice, "MinPrice")
	r.Equal(e.TickSize, a.TickSize, "TickSize")
}

func (s *exchangeInfoServiceTestSuite) assertPercentPriceFilterEqual(e, a *PercentPriceFilter) {
	r := s.r()
	r.Equal(e.AveragePriceMins, a.AveragePriceMins, "AveragePriceMins")
	r.Equal(e.MultiplierUp, a.MultiplierUp, "MultiplierUp")
	r.Equal(e.MultiplierDown, a.MultiplierDown, "MultiplierDown")
}

func (s *exchangeInfoServiceTestSuite) assertMinNotionalFilterEqual(e, a *MinNotionalFilter) {
	r := s.r()
	r.Equal(e.MinNotional, a.MinNotional, "MinNotional")
	r.Equal(e.AveragePriceMins, a.AveragePriceMins, "AveragePriceMins")
	r.Equal(e.ApplyToMarket, a.ApplyToMarket, "ApplyToMarket")
}

func (s *exchangeInfoServiceTestSuite) assertIcebergPartsFilterEqual(e, a *IcebergPartsFilter) {
	r := s.r()
	r.Equal(e.Limit, a.Limit, "Limit")
}

func (s *exchangeInfoServiceTestSuite) assertMarketLotSizeFilterEqual(e, a *MarketLotSizeFilter) {
	r := s.r()
	r.Equal(e.MaxQuantity, a.MaxQuantity, "MaxQuantity")
	r.Equal(e.MinQuantity, a.MinQuantity, "MinQuantity")
	r.Equal(e.StepSize, a.StepSize, "StepSize")
}

func (s *exchangeInfoServiceTestSuite) assertMaxNumAlgoOrdersFilterEqual(e, a *MaxNumAlgoOrdersFilter) {
	r := s.r()
	r.Equal(e.MaxNumAlgoOrders, a.MaxNumAlgoOrders, "MaxNumAlgoOrders")
}
