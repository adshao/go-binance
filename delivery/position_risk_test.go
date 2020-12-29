package delivery

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
			"symbol": "BTCUSD_201225",
			"positionAmt": "0",
			"entryPrice": "0.0",
			"markPrice": "0.00000000",
			"unRealizedProfit": "0.00000000",
			"liquidationPrice": "0",
			"leverage": "125",
			"maxQty": "50",
			"marginType": "cross",
			"isolatedMargin": "0.00000000",
			"isAutoAddMargin": "false",
			"positionSide": "BOTH"
		},
		{
			"symbol": "BTCUSD_201225",
			"positionAmt": "1",
			"entryPrice": "11707.70000003",
			"markPrice": "11788.66626667",
			"unRealizedProfit": "0.00005866",
			"liquidationPrice": "11667.63509587",
			"leverage": "125",
			"maxQty": "50",
			"marginType": "cross",
			"isolatedMargin": "0.00012357",
			"isAutoAddMargin": "false",
			"positionSide": "LONG"
		},
		{
			"symbol": "BTCUSD_201225",
			"positionAmt": "0",
			"entryPrice": "0.0",
			"markPrice": "0.00000000",
			"unRealizedProfit": "0.00000000",
			"liquidationPrice": "0",
			"leverage": "125",
			"maxQty": "50",
			"marginType": "cross",
			"isolatedMargin": "0.00000000",
			"isAutoAddMargin": "false",
			"positionSide": "SHORT"
	  }
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"pair": "BTCUSD",
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetPositionRiskService().Pair("BTCUSD").Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 3)
	e := []PositionRisk{{
		Symbol:           "BTCUSD_201225",
		PositionAmt:      "0",
		EntryPrice:       "0.0",
		MarkPrice:        "0.00000000",
		UnRealizedProfit: "0.00000000",
		LiquidationPrice: "0",
		Leverage:         "125",
		MaxQuantity:      "50",
		MarginType:       "cross",
		IsolatedMargin:   "0.00000000",
		IsAutoAddMargin:  "false",
		PositionSide:     "BOTH",
	}, {
		Symbol:           "BTCUSD_201225",
		PositionAmt:      "1",
		EntryPrice:       "11707.70000003",
		MarkPrice:        "11788.66626667",
		UnRealizedProfit: "0.00005866",
		LiquidationPrice: "11667.63509587",
		Leverage:         "125",
		MaxQuantity:      "50",
		MarginType:       "cross",
		IsolatedMargin:   "0.00012357",
		IsAutoAddMargin:  "false",
		PositionSide:     "LONG",
	}, {
		Symbol:           "BTCUSD_201225",
		PositionAmt:      "0",
		EntryPrice:       "0.0",
		MarkPrice:        "0.00000000",
		UnRealizedProfit: "0.00000000",
		LiquidationPrice: "0",
		Leverage:         "125",
		MaxQuantity:      "50",
		MarginType:       "cross",
		IsolatedMargin:   "0.00000000",
		IsAutoAddMargin:  "false",
		PositionSide:     "SHORT",
	}}
	s.assertPositionRiskEqual(&e[0], res[0])
	s.assertPositionRiskEqual(&e[1], res[1])
	s.assertPositionRiskEqual(&e[2], res[2])
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
	r.Equal(e.MaxQuantity, a.MaxQuantity, "MaxQuantity")
	r.Equal(e.PositionAmt, a.PositionAmt, "PositionAmt")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.UnRealizedProfit, a.UnRealizedProfit, "UnRealizedProfit")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
}
