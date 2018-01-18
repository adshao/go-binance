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
	res, err := s.client.NewExchangeInfoService().Do(newContext())
	s.r().NoError(err)
	ei := &ExchangeInfo{
		Symbols: []Symbol{
			{
				Symbol:             "ETHBTC",
				BaseAsset:          "ETH",
				BaseAssetPrecision: 8,
				QuoteAsset:         "BTC",
				QuotePrecision:     8,
			},
		},
	}
	s.assertExchangeInfoEqual(ei, res)
}

func (s *exchangeInfoServiceTestSuite) assertExchangeInfoEqual(e, a *ExchangeInfo) {
	r := s.r()
	for i := 0; i < len(a.Symbols); i++ {
		if a.Symbols[i].Symbol == e.Symbols[0].Symbol {
			r.Equal(e.Symbols[i].BaseAsset, a.Symbols[i].BaseAsset, "BaseAsset")
			r.Equal(e.Symbols[i].BaseAssetPrecision, a.Symbols[i].BaseAssetPrecision, "BaseAssetPrecision")
			r.Equal(e.Symbols[i].QuoteAsset, a.Symbols[i].QuoteAsset, "QuoteAsset")
			r.Equal(e.Symbols[i].QuotePrecision, a.Symbols[i].QuotePrecision, "QuotePrecision")
			return
		}

	}
	r.Fail("Symbol ETHBTC not found")
}
