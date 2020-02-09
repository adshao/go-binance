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
