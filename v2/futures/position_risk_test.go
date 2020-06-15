package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type positionRiskServiceTestSuite struct {
	baseTestSuite
}

func TestPositionRiskTestService(t *testing.T) {
	suite.Run(t, new(positionRiskServiceTestSuite))
}

func (s *positionRiskServiceTestSuite) TestGetPositionRisk() {
	data := []byte(`[
		{
			"entryPrice": "10359.38000",
			"marginType": "isolated",
			"isAutoAddMargin": "false",
			"isolatedMargin": "3.15899368",
			"leverage": "125",
			"liquidationPrice": "9332.61",
			"markPrice": "10348.27548846",
			"maxNotionalValue": "50000",
			"positionAmt": "0.003",
			"symbol": "BTCUSDT",
			"unRealizedProfit": "-0.03331353",
			"positionSide": "BOTH"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetPositionRiskService().Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 1)
	e := &PositionRisk{
		EntryPrice:       "10359.38000",
		MarginType:       "isolated",
		IsAutoAddMargin:  "false",
		IsolatedMargin:   "3.15899368",
		Leverage:         "125",
		LiquidationPrice: "9332.61",
		MarkPrice:        "10348.27548846",
		MaxNotionalValue: "50000",
		PositionAmt:      "0.003",
		Symbol:           "BTCUSDT",
		UnRealizedProfit: "-0.03331353",
		PositionSide:     "BOTH",
	}
	s.assertPositionRiskEqual(e, res[0])
}

func (s *positionRiskServiceTestSuite) assertPositionRiskEqual(e, a *PositionRisk) {
	r := s.r()
	r.Equal(e.EntryPrice, a.EntryPrice, "EntryPrice")
	r.Equal(e.MarginType, a.MarginType, "MarginType")
	r.Equal(e.IsAutoAddMargin, a.IsAutoAddMargin, "IsAutoAddMargin")
	r.Equal(e.IsolatedMargin, a.IsolatedMargin, "IsolatedMargin")
	r.Equal(e.Leverage, a.Leverage, "Leverage")
	r.Equal(e.LiquidationPrice, a.LiquidationPrice, "LiquidationPrice")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.MaxNotionalValue, a.MaxNotionalValue, "MaxNotionalValue")
	r.Equal(e.PositionAmt, a.PositionAmt, "PositionAmt")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.UnRealizedProfit, a.UnRealizedProfit, "UnRealizedProfit")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
}
