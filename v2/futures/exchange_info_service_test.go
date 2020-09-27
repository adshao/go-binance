package futures

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
		"exchangeFilters": [],
		"rateLimits": [
			{
				"interval": "MINUTE",
				"intervalNum": 1,
				"limit": 6000,
				"rateLimitType": "REQUEST_WEIGHT"
			},
			{
				"interval": "MINUTE",
				"intervalNum": 1,
				"limit": 6000,
				"rateLimitType": "ORDERS"
			}
		],
		"serverTime": 1565613908500,
		"symbols": [
			{
				"filters": [
					{
						"filterType": "PRICE_FILTER",
						"maxPrice": "10000000",
						"minPrice": "0.00000100",
						"tickSize": "0.00000100"
					},
					{
						"filterType": "LOT_SIZE",
						"maxQty": "10000000",
						"minQty": "0.00100000",
						"stepSize": "0.00100000"
					},
					{
						"filterType": "MARKET_LOT_SIZE",
						"maxQty": "10000000",
						"minQty": "0.00100000",
						"stepSize": "0.00100000"
					},
					{
						"filterType": "MAX_NUM_ORDERS",
						"limit": 100
					},
					{
						"filterType": "PERCENT_PRICE",
						"multiplierUp": "1.1500",
						"multiplierDown": "0.8500",
						"multiplierDecimal": 4
					}
				],
				"maintMarginPercent": "2.5000",
				"pricePrecision": 2,
				"quantityPrecision": 3,
				"requiredMarginPercent": "5.0000",
				"status": "TRADING",
				"OrderType": [
					"LIMIT", 
					"MARKET", 
					"STOP",
					"TAKE_PROFIT"
				],
				"symbol": "BTCUSDT",
				"quoteAsset": "USDT",
				"baseAsset": "BTC",
				"timeInForce": [
					"GTC",
					"IOC",
					"FOK",
					"GTX"
				]
			}
		],
		"timezone": "UTC"
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
		ServerTime: 1565613908500,
		RateLimits: []RateLimit{
			{RateLimitType: "REQUEST_WEIGHT", Interval: "MINUTE", IntervalNum: 1, Limit: 6000},
			{RateLimitType: "ORDERS", Interval: "MINUTE", IntervalNum: 1, Limit: 6000},
		},
		ExchangeFilters: []interface{}{},
		Symbols: []Symbol{
			{
				Symbol:                "BTCUSDT",
				BaseAsset:             "BTC",
				QuoteAsset:            "USDT",
				Status:                "TRADING",
				MaintMarginPercent:    "2.5000",
				PricePrecision:        2,
				QuantityPrecision:     3,
				RequiredMarginPercent: "5.0000",
				OrderType:             []OrderType{OrderTypeLimit, OrderTypeMarket, OrderTypeStop, OrderTypeTakeProfit},
				TimeInForce:           []TimeInForceType{TimeInForceTypeGTC, TimeInForceTypeIOC, TimeInForceTypeFOK, TimeInForceTypeGTX},
				Filters: []map[string]interface{}{
					{"filterType": "PRICE_FILTER", "minPrice": "0.00000100", "maxPrice": "10000000", "tickSize": "0.00000100"},
					{"filterType": "LOT_SIZE", "minQty": "0.00100000", "maxQty": "10000000", "stepSize": "0.00100000"},
					{"filterType": "MARKET_LOT_SIZE", "maxQty": "10000000", "minQty": "0.00100000", "stepSize": "0.00100000"},
					{"filterType": "MAX_NUM_ORDERS", "limit": 100},
					{"filterType": "PERCENT_PRICE", "multiplierUp": "1.1500", "multiplierDown": "0.8500", "multiplierDecimal": 4},
				},
			},
		},
	}
	s.assertExchangeInfoEqual(ei, res)
	s.r().Len(ei.Symbols[0].Filters, 5, "Filters")
	ePriceFilter := &PriceFilter{
		MaxPrice: "10000000",
		MinPrice: "0.00000100",
		TickSize: "0.00000100",
	}
	s.assertPriceFilterEqual(ePriceFilter, res.Symbols[0].PriceFilter())
	eLotSizeFilter := &LotSizeFilter{
		MaxQuantity: "10000000",
		MinQuantity: "0.00100000",
		StepSize:    "0.00100000",
	}
	s.assertLotSizeFilterEqual(eLotSizeFilter, res.Symbols[0].LotSizeFilter())
	eMarketLotSizeFilter := &MarketLotSizeFilter{
		MaxQuantity: "10000000",
		MinQuantity: "0.00100000",
		StepSize:    "0.00100000",
	}
	s.assertMarketLotSizeFilterEqual(eMarketLotSizeFilter, res.Symbols[0].MarketLotSizeFilter())
	eMaxNumOrdersFilter := &MaxNumOrdersFilter{
		Limit: 100,
	}
	s.assertMaxNumOrdersFilterEqual(eMaxNumOrdersFilter, res.Symbols[0].MaxNumOrdersFilter())
	ePercentPriceFilter := &PercentPriceFilter{
		MultiplierDecimal: 4,
		MultiplierUp:      "1.1500",
		MultiplierDown:    "0.8500",
	}
	s.assertPercentPriceFilterEqual(ePercentPriceFilter, res.Symbols[0].PercentPriceFilter())
}

func (s *exchangeInfoServiceTestSuite) assertExchangeInfoEqual(e, a *ExchangeInfo) {
	r := s.r()

	r.Equal(e.Timezone, a.Timezone, "Timezone")
	r.Equal(e.ServerTime, a.ServerTime, "ServerTime")
	r.Len(a.RateLimits, len(e.RateLimits), "RateLimits")
	for i := range a.RateLimits {
		r.Equal(e.RateLimits[i].RateLimitType, a.RateLimits[i].RateLimitType, "RateLimitType")
		r.Equal(e.RateLimits[i].Limit, a.RateLimits[i].Limit, "Limit")
		r.Equal(e.RateLimits[i].Interval, a.RateLimits[i].Interval, "Interval")
		r.Equal(e.RateLimits[i].IntervalNum, a.RateLimits[i].IntervalNum, "IntervalNum")
	}
	r.Equal(e.ExchangeFilters, a.ExchangeFilters, "ExchangeFilters")
	r.Len(a.Symbols, len(e.Symbols), "Symbols")

	for i := range a.Symbols {
		r.Equal(e.Symbols[i].Symbol, a.Symbols[i].Symbol, "Symbol")
		r.Equal(e.Symbols[i].BaseAsset, a.Symbols[i].BaseAsset, "BaseAsset")
		r.Equal(e.Symbols[i].QuoteAsset, a.Symbols[i].QuoteAsset, "QuoteAsset")
		r.Equal(e.Symbols[i].Status, a.Symbols[i].Status, "Status")
		r.Equal(e.Symbols[i].MaintMarginPercent, a.Symbols[i].MaintMarginPercent, "MaintMarginPercent")
		r.Equal(e.Symbols[i].PricePrecision, a.Symbols[i].PricePrecision, "PricePrecision")
		r.Equal(e.Symbols[i].QuantityPrecision, a.Symbols[i].QuantityPrecision, "QuantityPrecision")
		r.Equal(e.Symbols[i].RequiredMarginPercent, a.Symbols[i].RequiredMarginPercent, "RequiredMarginPercent")
		r.Len(a.Symbols[i].OrderType, len(e.Symbols[i].OrderType))
		for j, orderType := range e.Symbols[i].OrderType {
			r.Equal(orderType, a.Symbols[i].OrderType[j], "OrderType")
		}
		r.Len(a.Symbols[i].TimeInForce, len(e.Symbols[i].TimeInForce), "TimeInForce")
		for j, timeInForce := range e.Symbols[i].TimeInForce {
			r.Equal(timeInForce, a.Symbols[i].TimeInForce[j], "TimeInForce")
		}
	}
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
	r.Equal(e.MultiplierDecimal, a.MultiplierDecimal, "MultiplierDecimal")
	r.Equal(e.MultiplierUp, a.MultiplierUp, "MultiplierUp")
	r.Equal(e.MultiplierDown, a.MultiplierDown, "MultiplierDown")
}

func (s *exchangeInfoServiceTestSuite) assertMarketLotSizeFilterEqual(e, a *MarketLotSizeFilter) {
	r := s.r()
	r.Equal(e.MaxQuantity, a.MaxQuantity, "MaxQuantity")
	r.Equal(e.MinQuantity, a.MinQuantity, "MinQuantity")
	r.Equal(e.StepSize, a.StepSize, "StepSize")
}

func (s *exchangeInfoServiceTestSuite) assertMaxNumOrdersFilterEqual(e, a *MaxNumOrdersFilter) {
	r := s.r()
	r.Equal(e.Limit, a.Limit, "Limit")
}
