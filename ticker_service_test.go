package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type tickerServiceTestSuite struct {
	baseTestSuite
}

func TestTickerService(t *testing.T) {
	suite.Run(t, new(tickerServiceTestSuite))
}

func (s *tickerServiceTestSuite) TestListBookTickers() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "bidPrice": "4.00000000",
            "bidQty": "431.00000000",
            "askPrice": "4.00000200",
            "askQty": "9.00000000"
        },
        {
            "symbol": "ETHBTC",
            "bidPrice": "0.07946700",
            "bidQty": "9.00000000",
            "askPrice": "100000.00000000",
            "askQty": "1000.00000000"
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	tickers, err := s.client.NewListBookTickersService().Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(tickers, 2)
	e1 := &BookTicker{
		Symbol:      "LTCBTC",
		BidPrice:    "4.00000000",
		BidQuantity: "431.00000000",
		AskPrice:    "4.00000200",
		AskQuantity: "9.00000000",
	}
	e2 := &BookTicker{
		Symbol:      "ETHBTC",
		BidPrice:    "0.07946700",
		BidQuantity: "9.00000000",
		AskPrice:    "100000.00000000",
		AskQuantity: "1000.00000000",
	}
	s.assertBookTickerEqual(e1, tickers[0])
	s.assertBookTickerEqual(e2, tickers[1])
}

func (s *tickerServiceTestSuite) TestSingleBookTicker() {
	data := []byte(`{
            "symbol": "LTCBTC",
            "bidPrice": "4.00000000",
            "bidQty": "431.00000000",
            "askPrice": "4.00000200",
            "askQty": "9.00000000"
        }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"

	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	tickers, err := s.client.NewListBookTickersService().Symbol("LTCBTC").Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(tickers, 1)
	e := &BookTicker{
		Symbol:      "LTCBTC",
		BidPrice:    "4.00000000",
		BidQuantity: "431.00000000",
		AskPrice:    "4.00000200",
		AskQuantity: "9.00000000",
	}
	s.assertBookTickerEqual(e, tickers[0])
}

func (s *tickerServiceTestSuite) assertBookTickerEqual(e, a *BookTicker) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.BidPrice, a.BidPrice, "BidPrice")
	r.Equal(e.BidQuantity, a.BidQuantity, "BidQuantity")
	r.Equal(e.AskPrice, a.AskPrice, "AskPrice")
	r.Equal(e.AskQuantity, a.AskQuantity, "AskQuantity")
}

func (s *tickerServiceTestSuite) TestListPrices() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "price": "4.00000200"
        },
        {
            "symbol": "ETHBTC",
            "price": "0.07946600"
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	prices, err := s.client.NewListPricesService().Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(prices, 2)
	e1 := &SymbolPrice{
		Symbol: "LTCBTC",
		Price:  "4.00000200",
	}
	e2 := &SymbolPrice{
		Symbol: "ETHBTC",
		Price:  "0.07946600",
	}
	s.assertSymbolPriceEqual(e1, prices[0])
	s.assertSymbolPriceEqual(e2, prices[1])
}

func (s *tickerServiceTestSuite) TestListSinglePrice() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"price": "4.00000200"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	prices, err := s.client.NewListPricesService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(prices, 1)
	e1 := &SymbolPrice{
		Symbol: "LTCBTC",
		Price:  "4.00000200",
	}
	s.assertSymbolPriceEqual(e1, prices[0])
}

func (s *tickerServiceTestSuite) assertSymbolPriceEqual(e, a *SymbolPrice) {
	r := s.r()
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}

func (s *tickerServiceTestSuite) TestPriceChangeStats() {
	data := []byte(`{
		"symbol": "BNBBTC",
        "priceChange": "-94.99999800",
        "priceChangePercent": "-95.960",
        "weightedAvgPrice": "0.29628482",
        "prevClosePrice": "0.10002000",
		"lastPrice": "4.00000200",
		"lastQty": "200.00000000",
        "bidPrice": "4.00000000",
        "askPrice": "4.00000200",
        "openPrice": "99.00000000",
        "highPrice": "100.00000000",
        "lowPrice": "0.10000000",
        "volume": "8913.30000000",
        "openTime": 1499783499040,
        "closeTime": 1499869899040,
        "firstId": 28385,
        "lastId": 28460,
        "count": 76
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	stats, err := s.client.NewListPriceChangeStatsService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(stats, 1)
	e := &PriceChangeStats{
		Symbol:             "BNBBTC",
		PriceChange:        "-94.99999800",
		PriceChangePercent: "-95.960",
		WeightedAvgPrice:   "0.29628482",
		PrevClosePrice:     "0.10002000",
		LastPrice:          "4.00000200",
		LastQty:            "200.00000000",
		BidPrice:           "4.00000000",
		AskPrice:           "4.00000200",
		OpenPrice:          "99.00000000",
		HighPrice:          "100.00000000",
		LowPrice:           "0.10000000",
		Volume:             "8913.30000000",
		OpenTime:           1499783499040,
		CloseTime:          1499869899040,
		FristID:            28385,
		LastID:             28460,
		Count:              76,
	}
	s.assertPriceChangeStatsEqual(e, stats[0])
}

func (s *tickerServiceTestSuite) assertPriceChangeStatsEqual(e, a *PriceChangeStats) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.PrevClosePrice, a.PrevClosePrice, "PrevClosePrice")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.LastQty, a.LastQty, "LastQty")
	r.Equal(e.BidPrice, a.BidPrice, "BidPrice")
	r.Equal(e.AskPrice, a.AskPrice, "AskPrice")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FristID, a.FristID, "FristID")
	r.Equal(e.LastID, a.LastID, "LastID")
	r.Equal(e.Count, a.Count, "Count")
}

func (s *tickerServiceTestSuite) TestListPriceChangeStats() {
	data := []byte(`[{
    	"symbol": "BNBBTC",
    	"priceChange": "-94.99999800",
    	"priceChangePercent": "-95.960",
    	"weightedAvgPrice": "0.29628482",
    	"prevClosePrice": "0.10002000",
    	"lastPrice": "4.00000200",
    	"lastQty": "200.00000000",
    	"bidPrice": "4.00000000",
    	"askPrice": "4.00000200",
    	"openPrice": "99.00000000",
    	"highPrice": "100.00000000",
    	"lowPrice": "0.10000000",
    	"volume": "8913.30000000",
    	"quoteVolume": "15.30000000",
    	"openTime": 1499783499040,
    	"closeTime": 1499869899040,
    	"firstId": 28385,
    	"lastId": 28460,
    	"count": 76
  	}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListPriceChangeStatsService().Do(newContext())
	r := s.r()
	r.NoError(err)
	e := []*PriceChangeStats{
		{
			Symbol:             "BNBBTC",
			PriceChange:        "-94.99999800",
			PriceChangePercent: "-95.960",
			WeightedAvgPrice:   "0.29628482",
			PrevClosePrice:     "0.10002000",
			LastPrice:          "4.00000200",
			BidPrice:           "4.00000000",
			AskPrice:           "4.00000200",
			OpenPrice:          "99.00000000",
			HighPrice:          "100.00000000",
			LowPrice:           "0.10000000",
			Volume:             "8913.30000000",
			QuoteVolume:        "15.30000000",
			OpenTime:           1499783499040,
			CloseTime:          1499869899040,
			FristID:            28385,
			LastID:             28460,
			Count:              76,
		},
	}
	s.assertListPriceChangeStatsEqual(e, res)
}

func (s *tickerServiceTestSuite) assertListPriceChangeStatsEqual(e, a []*PriceChangeStats) {
	r := s.r()
	for i := range e {
		r.Equal(e[i].Symbol, a[i].Symbol, "Symbol")
		r.Equal(e[i].PriceChange, a[i].PriceChange, "PriceChange")
		r.Equal(e[i].PriceChangePercent, a[i].PriceChangePercent, "PriceChangePercent")
		r.Equal(e[i].WeightedAvgPrice, a[i].WeightedAvgPrice, "WeightedAvgPrice")
		r.Equal(e[i].PrevClosePrice, a[i].PrevClosePrice, "PrevClosePrice")
		r.Equal(e[i].LastPrice, a[i].LastPrice, "LastPrice")
		r.Equal(e[i].BidPrice, a[i].BidPrice, "BidPrice")
		r.Equal(e[i].AskPrice, a[i].AskPrice, "AskPrice")
		r.Equal(e[i].OpenPrice, a[i].OpenPrice, "OpenPrice")
		r.Equal(e[i].HighPrice, a[i].HighPrice, "HighPrice")
		r.Equal(e[i].LowPrice, a[i].LowPrice, "LowPrice")
		r.Equal(e[i].Volume, a[i].Volume, "Volume")
		r.Equal(e[i].OpenTime, a[i].OpenTime, "OpenTime")
		r.Equal(e[i].CloseTime, a[i].CloseTime, "CloseTime")
		r.Equal(e[i].FristID, a[i].FristID, "FristID")
		r.Equal(e[i].LastID, a[i].LastID, "LastID")
		r.Equal(e[i].Count, a[i].Count, "Count")
	}

}

func (s *tickerServiceTestSuite) TestAveragePrice() {
	data := []byte(`{
		"mins": 5,
		"price": "9.35751834"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewAveragePriceService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &AvgPrice{
		Mins:  5,
		Price: "9.35751834",
	}
	s.assertAvgPrice(e, res)
}

func (s *tickerServiceTestSuite) assertAvgPrice(e, a *AvgPrice) {
	s.r().Equal(e.Mins, a.Mins, "Mins")
	s.r().Equal(e.Price, a.Price, "Price")
}
