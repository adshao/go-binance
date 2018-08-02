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
}

func (s *exchangeInfoServiceTestSuite) assertExchangeInfoEqual(e, a *ExchangeInfo) {
	r := s.r()
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
