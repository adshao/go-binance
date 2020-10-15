package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type tradeServiceTestSuite struct {
	baseTestSuite
}

func TestTradeService(t *testing.T) {
	suite.Run(t, new(tradeServiceTestSuite))
}

func (s *tradeServiceTestSuite) TestAggregateTrades() {
	data := []byte(`[
        {
            "a": 26129,
            "p": "0.01633102",
            "q": "4.70443515",
            "f": 27781,
            "l": 27781,
            "T": 1498793709153,
            "m": true
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	fromID := int64(1)
	startTime := int64(1498793709153)
	endTime := int64(1498793709156)
	limit := 1
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol":    symbol,
			"fromId":    fromID,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	aggTrades, err := s.client.NewAggTradesService().Symbol(symbol).
		FromID(fromID).StartTime(startTime).EndTime(endTime).Limit(limit).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(aggTrades, 1)
	e := &AggTrade{
		AggTradeID:   26129,
		Price:        "0.01633102",
		Quantity:     "4.70443515",
		FirstTradeID: 27781,
		LastTradeID:  27781,
		Timestamp:    1498793709153,
		IsBuyerMaker: true,
	}
	s.assertAggTradeEqual(e, aggTrades[0])
}

func (s *tradeServiceTestSuite) assertAggTradeEqual(e, a *AggTrade) {
	r := s.r()
	r.Equal(e.AggTradeID, a.AggTradeID, "AggTradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.FirstTradeID, a.FirstTradeID, "FirstTradeID")
	r.Equal(e.LastTradeID, a.LastTradeID, "LastTradeID")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
}

func (s *tradeServiceTestSuite) TestHistoricalTrades() {
	data := []byte(`[
		{
		  "id": 28457,
		  "price": "4.00000100",
		  "qty": "12.00000000",
		  "quoteQty": "8000.00",
		  "time": 1499865549590,
		  "isBuyerMaker": true
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	limit := 3
	fromID := int64(1)
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol": symbol,
			"limit":  limit,
			"fromId": fromID,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewHistoricalTradesService().Symbol(symbol).
		Limit(limit).FromID(fromID).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(trades, 1)
	e := &Trade{
		ID:            28457,
		Price:         "4.00000100",
		Quantity:      "12.00000000",
		QuoteQuantity: "8000.00",
		Time:          1499865549590,
		IsBuyerMaker:  true,
	}
	s.assertTradeEqual(e, trades[0])
}

func (s *tradeServiceTestSuite) TestRecentTrades() {
	data := []byte(`[
		{
		  "id": 28457,
		  "price": "4.00000100",
		  "qty": "12.00000000",
		  "quoteQty": "8000.00",
		  "time": 1499865549590,
		  "isBuyerMaker": true
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	limit := 3
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol": symbol,
			"limit":  limit,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewRecentTradesService().Symbol(symbol).Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(trades, 1)
	e := &Trade{
		ID:            28457,
		Price:         "4.00000100",
		Quantity:      "12.00000000",
		QuoteQuantity: "8000.00",
		Time:          1499865549590,
		IsBuyerMaker:  true,
	}
	s.assertTradeEqual(e, trades[0])
}

func (s *tradeServiceTestSuite) assertTradeEqual(e, a *Trade) {
	r := s.r()
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.QuoteQuantity, a.QuoteQuantity, "QuoteQuantity")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
}

func (s *tradeServiceTestSuite) TestAccountTradeList() {
	data := []byte(`[
		{
			"buyer": false,
			"commission": "-0.07819010",
			"commissionAsset": "USDT",
			"id": 698759,
			"maker": false,
			"orderId": 25851813,
			"price": "7819.01",
			"qty": "0.002",
			"quoteQty": "15.63802",
			"realizedPnl": "-0.91539999",
			"side": "SELL",
			"positionSide": "SHORT",
			"symbol": "BTCUSDT",
			"time": 1569514978020
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	startTime := int64(1569514978020)
	endTime := int64(1569514978021)
	fromID := int64(698759)
	limit := 3
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"fromID":    fromID,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewListAccountTradeService().Symbol(symbol).
		StartTime(startTime).EndTime(endTime).FromID(fromID).Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(trades, 1)
	e := &AccountTrade{
		Buyer:           false,
		Commission:      "-0.07819010",
		CommissionAsset: "USDT",
		ID:              698759,
		Maker:           false,
		OrderID:         25851813,
		Price:           "7819.01",
		Quantity:        "0.002",
		QuoteQuantity:   "15.63802",
		RealizedPnl:     "-0.91539999",
		Side:            SideTypeSell,
		PositionSide:    PositionSideTypeShort,
		Symbol:          symbol,
		Time:            1569514978020,
	}
	s.assertAccountTradeEqual(e, trades[0])
}

func (s *tradeServiceTestSuite) assertAccountTradeEqual(e, a *AccountTrade) {
	r := s.r()
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.Buyer, a.Buyer, "Buyer")
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.Maker, a.Maker, "Maker")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.QuoteQuantity, a.QuoteQuantity, "QuoteQuantity")
	r.Equal(e.RealizedPnl, a.RealizedPnl, "RealizedPnl")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Time, a.Time, "Time")
}
