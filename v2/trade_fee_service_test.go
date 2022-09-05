package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type assetTradeFeeServiceSuite struct {
	baseTestSuite
}

func (a *assetTradeFeeServiceSuite) assertTradeFeeServiceEqual(expected, other *TradeFeeDetails) {
	r := a.r()

	r.Equal(expected.Symbol, other.Symbol)
	r.Equal(expected.MakerCommission, other.MakerCommission)
	r.Equal(expected.TakerCommission, other.TakerCommission)
}

func TestTradeFeeService(t *testing.T) {
	suite.Run(t, new(assetTradeFeeServiceSuite))
}

func (s *assetTradeFeeServiceSuite) TestListTradeFee() {
	data := []byte(`
	[
		{
			"symbol": "ADABNB",
			"makerCommission": "0.001",
			"takerCommission": "0.001"
		},
		{
			"symbol": "BNBBTC",
			"makerCommission": "0.001",
			"takerCommission": "0.001"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	fees, err := s.client.NewTradeFeeService().Do(context.Background())
	s.r().NoError(err)
	rows := fees

	s.Len(rows, 2)
	s.assertTradeFeeServiceEqual(&TradeFeeDetails{
		Symbol:          "ADABNB",
		MakerCommission: "0.001",
		TakerCommission: "0.001"},
		rows[0])
	s.assertTradeFeeServiceEqual(&TradeFeeDetails{
		Symbol:          "BNBBTC",
		MakerCommission: "0.001",
		TakerCommission: "0.001"},
		rows[1])
}

func (s *assetTradeFeeServiceSuite) TestSingleSymbolTradeFee() {
	data := []byte(`
	[
		{
			"symbol": "ADABNB",
			"makerCommission": "0.001",
			"takerCommission": "0.001"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	fees, err := s.client.NewTradeFeeService().Symbol("ADABNB").Do(context.Background())
	s.r().NoError(err)
	rows := fees

	s.Len(rows, 1)
	s.assertTradeFeeServiceEqual(&TradeFeeDetails{
		Symbol:          "ADABNB",
		MakerCommission: "0.001",
		TakerCommission: "0.001"},
		rows[0])
}
