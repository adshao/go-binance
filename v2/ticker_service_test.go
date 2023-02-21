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

func (s *tickerServiceTestSuite) TestListPricesForMultipleSymbols() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "price": "4.00000200"
        },
        {
            "symbol": "ETHUSDT",
            "price": "2856.76"
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	symbol1, symbol2 := "ETHUSDT", "LTCBTC"
	symbols := make([]string, 2)

	symbols[0] = symbol1
	symbols[1] = symbol2

	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbols", `["ETHUSDT","LTCBTC"]`)
		s.assertRequestEqual(e, r)
	})

	prices, err := s.client.NewListPricesService().Symbols(symbols).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(prices, 2)
	e1 := &SymbolPrice{
		Symbol: "LTCBTC",
		Price:  "4.00000200",
	}
	e2 := &SymbolPrice{
		Symbol: "ETHUSDT",
		Price:  "2856.76",
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
        "count": 76,
        "bidQty": "300.00000000",
        "askQty": "400.00000000"
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
		BidQty:             "300.00000000",
		AskQty:             "400.00000000",
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

func (s *tickerServiceTestSuite) TestMultiplePriceChangeStats() {
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
  	},{
    	"symbol": "ETHBTC",
    	"priceChange": "-194.99999800",
    	"priceChangePercent": "-195.960",
    	"weightedAvgPrice": "10.29628482",
    	"prevClosePrice": "10.10002000",
    	"lastPrice": "14.00000200",
    	"lastQty": "1200.00000000",
    	"bidPrice": "14.00000000",
    	"askPrice": "14.00000200",
    	"openPrice": "199.00000000",
    	"highPrice": "1100.00000000",
    	"lowPrice": "10.10000000",
    	"volume": "18913.30000000",
    	"quoteVolume": "115.30000000",
    	"openTime": 1499783499041,
    	"closeTime": 1499869899041,
    	"firstId": 28381,
    	"lastId": 28461,
    	"count": 71
  	}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbols", `["BNBBTC","ETHBTC"]`)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListPriceChangeStatsService().Symbols([]string{"BNBBTC", "ETHBTC"}).Do(newContext())
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
		{
			Symbol:             "ETHBTC",
			PriceChange:        "-194.99999800",
			PriceChangePercent: "-195.960",
			WeightedAvgPrice:   "10.29628482",
			PrevClosePrice:     "10.10002000",
			LastPrice:          "14.00000200",
			BidPrice:           "14.00000000",
			AskPrice:           "14.00000200",
			OpenPrice:          "199.00000000",
			HighPrice:          "1100.00000000",
			LowPrice:           "10.10000000",
			Volume:             "18913.30000000",
			QuoteVolume:        "115.30000000",
			OpenTime:           1499783499041,
			CloseTime:          1499869899041,
			FristID:            28381,
			LastID:             28461,
			Count:              71,
		},
	}
	s.assertListPriceChangeStatsEqual(e, res)
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

func (s *tickerServiceTestSuite) TestListSymbolTicker() {
	data := []byte(`[
		{
			"symbol": "ETHBTC",
			"priceChange": "0.00004700",
			"priceChangePercent": "0.066",
			"weightedAvgPrice": "0.07168666",
			"openPrice": "0.07093500",
			"highPrice": "0.07321800",
			"lowPrice": "0.07054200",
			"lastPrice": "0.07098200",
			"volume": "86992.33370000",
			"quoteVolume": "6236.18963157",
			"openTime": 1659097380000,
			"closeTime": 1659183780986,
			"firstId": 359930693,
			"lastId": 360209854,
			"count": 279162
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "ETHBTC"
	windowSize := "1m" // 1 minute
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol).setParam("windowSize", windowSize)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewListSymbolTickerService().Symbol(symbol).WindowSize(windowSize).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := make([]*SymbolTicker, 0)
	e = append(e, &SymbolTicker{
		Symbol:             "ETHBTC",
		PriceChange:        "0.00004700",
		PriceChangePercent: "0.066",
		WeightedAvgPrice:   "0.07168666",
		OpenPrice:          "0.07093500",
		HighPrice:          "0.07321800",
		LowPrice:           "0.07054200",
		LastPrice:          "0.07098200",
		Volume:             "86992.33370000",
		QuoteVolume:        "6236.18963157",
		OpenTime:           1659097380000,
		CloseTime:          1659183780986,
		FirstId:            359930693,
		LastId:             360209854,
		Count:              279162,
	})
	s.assertSymbolTicker(e, res)
}

func (s *tickerServiceTestSuite) assertSymbolTicker(e, st []*SymbolTicker) {
	for i := range e {
		s.r().Equal(e[i].Symbol, st[i].Symbol, "Symbol")
		s.r().Equal(e[i].PriceChange, st[i].PriceChange, "PriceChange")
		s.r().Equal(e[i].PriceChangePercent, st[i].PriceChangePercent, "PriceChangePercent")
		s.r().Equal(e[i].WeightedAvgPrice, st[i].WeightedAvgPrice, "WeightedAvgPrice")
		s.r().Equal(e[i].OpenPrice, st[i].OpenPrice, "OpenPrice")
		s.r().Equal(e[i].HighPrice, st[i].HighPrice, "HighPrice")
		s.r().Equal(e[i].LowPrice, st[i].LowPrice, "LowPrice")
		s.r().Equal(e[i].LastPrice, st[i].LastPrice, "LastPrice")
		s.r().Equal(e[i].Volume, st[i].Volume, "Volume")
		s.r().Equal(e[i].QuoteVolume, st[i].QuoteVolume, "QuoteVolume")
		s.r().Equal(e[i].OpenTime, st[i].OpenTime, "OpenTime")
		s.r().Equal(e[i].CloseTime, st[i].CloseTime, "CloseTime")
		s.r().Equal(e[i].FirstId, st[i].FirstId, "FirstId")
		s.r().Equal(e[i].LastId, st[i].LastId, "LastId")
		s.r().Equal(e[i].Count, st[i].Count, "Count")
	}
}
