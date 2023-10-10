package delivery

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
            "symbol": "BTCUSD_PERP",
			"pair": "BTCUSD",
            "bidPrice": "40041.9",
            "bidQty": "3074",
            "askPrice": "40042.0",
            "askQty": "5286"
        },
        {
            "symbol": "ETHUSD_PERP",
			"pair": "ETHUSD",
            "bidPrice": "2525.41",
            "bidQty": "17054",
            "askPrice": "2525.42",
            "askQty": "7112"
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
		Symbol:      "BTCUSD_PERP",
		Pair:        "BTCUSD",
		BidPrice:    "40041.9",
		BidQuantity: "3074",
		AskPrice:    "40042.0",
		AskQuantity: "5286",
	}
	e2 := &BookTicker{
		Symbol:      "ETHUSD_PERP",
		Pair:        "ETHUSD",
		BidPrice:    "2525.41",
		BidQuantity: "17054",
		AskPrice:    "2525.42",
		AskQuantity: "7112",
	}
	s.assertBookTickerEqual(e1, tickers[0])
	s.assertBookTickerEqual(e2, tickers[1])
}

func (s *tickerServiceTestSuite) TestSingleBookTicker() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_210625",
			"pair": "BTCUSD",
			"bidPrice": "40078.9",
			"bidQty": "761",
			"askPrice": "40079.0",
			"askQty": "280"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_210625"

	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	tickers, err := s.client.NewListBookTickersService().Symbol("BTCUSD_210625").Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(tickers, 1)
	e := &BookTicker{
		Symbol:      "BTCUSD_210625",
		Pair:        "BTCUSD",
		BidPrice:    "40078.9",
		BidQuantity: "761",
		AskPrice:    "40079.0",
		AskQuantity: "280",
	}
	s.assertBookTickerEqual(e, tickers[0])
}

func (s *tickerServiceTestSuite) TestListBookTickersWithPair() {
	data := []byte(`[
		{
			"symbol": "ETHUSD_210625",
			"pair": "ETHUSD",
			"bidPrice": "2524.22",
			"bidQty": "5820",
			"askPrice": "2524.23",
			"askQty": "3575"
		},
		{
			"symbol": "ETHUSD_PERP",
			"pair": "ETHUSD",
			"bidPrice": "2518.08",
			"bidQty": "7137",
			"askPrice": "2518.09",
			"askQty": "42967"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	pair := "ETHUSD"

	s.assertReq(func(r *request) {
		e := newRequest().setParam("pair", pair)
		s.assertRequestEqual(e, r)
	})
	tickers, err := s.client.NewListBookTickersService().Pair("ETHUSD").Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(tickers, 2)
	e1 := &BookTicker{
		Symbol:      "ETHUSD_210625",
		Pair:        "ETHUSD",
		BidPrice:    "2524.22",
		BidQuantity: "5820",
		AskPrice:    "2524.23",
		AskQuantity: "3575",
	}
	e2 := &BookTicker{
		Symbol:      "ETHUSD_PERP",
		Pair:        "ETHUSD",
		BidPrice:    "2518.08",
		BidQuantity: "7137",
		AskPrice:    "2518.09",
		AskQuantity: "42967",
	}
	s.assertBookTickerEqual(e1, tickers[0])
	s.assertBookTickerEqual(e2, tickers[1])
}

func (s *tickerServiceTestSuite) assertBookTickerEqual(e, a *BookTicker) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.BidPrice, a.BidPrice, "BidPrice")
	r.Equal(e.BidQuantity, a.BidQuantity, "BidQuantity")
	r.Equal(e.AskPrice, a.AskPrice, "AskPrice")
	r.Equal(e.AskQuantity, a.AskQuantity, "AskQuantity")
}

func (s *tickerServiceTestSuite) TestListPrices() {
	data := []byte(`[
        {
            "symbol": "BTCUSD_PERP",
			"ps": "BTCUSD",
            "price": "40016.1"
        },
        {
            "symbol": "ETHUSD_PERP",
			"ps": "ETHUSD",
            "price": "2519.13"
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
		Symbol: "BTCUSD_PERP",
		Pair:   "BTCUSD",
		Price:  "40016.1",
	}
	e2 := &SymbolPrice{
		Symbol: "ETHUSD_PERP",
		Pair:   "ETHUSD",
		Price:  "2519.13",
	}
	s.assertSymbolPriceEqual(e1, prices[0])
	s.assertSymbolPriceEqual(e2, prices[1])
}

func (s *tickerServiceTestSuite) TestListSinglePrice() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_PERP",
			"ps": "BTCUSD",
			"price": "39984.3"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_PERP"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	prices, err := s.client.NewListPricesService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(prices, 1)
	e := &SymbolPrice{
		Symbol: "BTCUSD_PERP",
		Pair:   "BTCUSD",
		Price:  "39984.3",
	}
	s.assertSymbolPriceEqual(e, prices[0])
}

func (s *tickerServiceTestSuite) TestListPricesWithPair() {
	data := []byte(`[
		{
		"symbol": "BTCUSD_210924",
		"ps": "BTCUSD",
		"price": "40998.5"
		},
		{
		"symbol": "BTCUSD_210625",
		"ps": "BTCUSD",
		"price": "40049.3"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	pair := "BTCUSD"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("pair", pair)
		s.assertRequestEqual(e, r)
	})

	prices, err := s.client.NewListPricesService().Pair(pair).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(prices, 2)
	e1 := &SymbolPrice{
		Symbol: "BTCUSD_210924",
		Pair:   "BTCUSD",
		Price:  "40998.5",
	}
	e2 := &SymbolPrice{
		Symbol: "BTCUSD_210625",
		Pair:   "BTCUSD",
		Price:  "40049.3",
	}
	s.assertSymbolPriceEqual(e1, prices[0])
	s.assertSymbolPriceEqual(e2, prices[1])
}

func (s *tickerServiceTestSuite) assertSymbolPriceEqual(e, a *SymbolPrice) {
	r := s.r()
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}

func (s *tickerServiceTestSuite) TestListPriceChangeStats() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_PERP",
			"pair": "BTCUSD",
			"priceChange": "208.5",
			"priceChangePercent": "0.524",
			"weightedAvgPrice": "40186.36908209",
			"lastPrice": "40011.9",
			"lastQty": "2",
			"openPrice": "39803.4",
			"highPrice": "41327.2",
			"lowPrice": "39455.9",
			"volume": "55845297",
			"baseVolume": "138965.76942775",
			"openTime": 1623748920000,
			"closeTime": 1623835355736,
			"firstId": 172749700,
			"lastId": 173464362,
			"count": 714658
		},
		{
			"symbol": "ETHUSD_PERP",
			"pair": "ETHUSD",
			"priceChange": "-59.74",
			"priceChangePercent": "-2.316",
			"weightedAvgPrice": "2549.94814666",
			"lastPrice": "2519.36",
			"lastQty": "15",
			"openPrice": "2579.10",
			"highPrice": "2615.85",
			"lowPrice": "2498.62",
			"volume": "165282425",
			"baseVolume": "648179.55304919",
			"openTime": 1623748920000,
			"closeTime": 1623835355187,
			"firstId": 138575549,
			"lastId": 139103143,
			"count": 527595
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListPriceChangeStatsService().Do(newContext())
	r := s.r()
	r.NoError(err)

	e1 := &PriceChangeStats{
		Symbol:             "BTCUSD_PERP",
		Pair:               "BTCUSD",
		PriceChange:        "208.5",
		PriceChangePercent: "0.524",
		WeightedAvgPrice:   "40186.36908209",
		LastPrice:          "40011.9",
		LastQuantity:       "2",
		OpenPrice:          "39803.4",
		HighPrice:          "41327.2",
		LowPrice:           "39455.9",
		Volume:             "55845297",
		BaseVolume:         "138965.76942775",
		OpenTime:           1623748920000,
		CloseTime:          1623835355736,
		FirstID:            172749700,
		LastID:             173464362,
		Count:              714658,
	}
	e2 := &PriceChangeStats{
		Symbol:             "ETHUSD_PERP",
		Pair:               "ETHUSD",
		PriceChange:        "-59.74",
		PriceChangePercent: "-2.316",
		WeightedAvgPrice:   "2549.94814666",
		LastPrice:          "2519.36",
		LastQuantity:       "15",
		OpenPrice:          "2579.10",
		HighPrice:          "2615.85",
		LowPrice:           "2498.62",
		Volume:             "165282425",
		BaseVolume:         "648179.55304919",
		OpenTime:           1623748920000,
		CloseTime:          1623835355187,
		FirstID:            138575549,
		LastID:             139103143,
		Count:              527595,
	}

	s.r().Len(res, 2)
	s.assertPriceChangeStatsEqual(e1, res[0])
	s.assertPriceChangeStatsEqual(e2, res[1])
}

func (s *tickerServiceTestSuite) TestSinglePriceChangeStats() {
	data := []byte(`[
		{
		"symbol": "BTCUSD_PERP",
		"pair": "BTCUSD",
		"priceChange": "-15.2",
		"priceChangePercent": "-0.038",
		"weightedAvgPrice": "40185.26563935",
		"lastPrice": "39955.9",
		"lastQty": "50",
		"openPrice": "39971.1",
		"highPrice": "41327.2",
		"lowPrice": "39455.9",
		"volume": "55355578",
		"baseVolume": "137750.93213717",
		"openTime": 1623752520000,
		"closeTime": 1623838964257,
		"firstId": 172782522,
		"lastId": 173490102,
		"count": 707576
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_PERP"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	stats, err := s.client.NewListPriceChangeStatsService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(stats, 1)
	e := &PriceChangeStats{
		Symbol:             "BTCUSD_PERP",
		Pair:               "BTCUSD",
		PriceChange:        "-15.2",
		PriceChangePercent: "-0.038",
		WeightedAvgPrice:   "40185.26563935",
		LastPrice:          "39955.9",
		LastQuantity:       "50",
		OpenPrice:          "39971.1",
		HighPrice:          "41327.2",
		LowPrice:           "39455.9",
		Volume:             "55355578",
		BaseVolume:         "137750.93213717",
		OpenTime:           1623752520000,
		CloseTime:          1623838964257,
		FirstID:            172782522,
		LastID:             173490102,
		Count:              707576,
	}
	s.assertPriceChangeStatsEqual(e, stats[0])
}

func (s *tickerServiceTestSuite) TestPriceChangeStatsWithPair() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_210924",
			"pair": "BTCUSD",
			"priceChange": "-906.4",
			"priceChangePercent": "-2.204",
			"weightedAvgPrice": "41226.29913486",
			"lastPrice": "40221.8",
			"lastQty": "125",
			"openPrice": "41128.2",
			"highPrice": "42532.8",
			"lowPrice": "40052.0",
			"volume": "3421917",
			"baseVolume": "8300.32545198",
			"openTime": 1623755580000,
			"closeTime": 1623842010140,
			"firstId": 7516015,
			"lastId": 7591268,
			"count": 75254
		},
		{
			"symbol": "BTCUSD_210625",
			"pair": "BTCUSD",
			"priceChange": "-781.7",
			"priceChangePercent": "-1.948",
			"weightedAvgPrice": "40179.25736363",
			"lastPrice": "39337.4",
			"lastQty": "1",
			"openPrice": "40119.1",
			"highPrice": "41405.7",
			"lowPrice": "39153.8",
			"volume": "5479601",
			"baseVolume": "13637.88521626",
			"openTime": 1623755580000,
			"closeTime": 1623842010656,
			"firstId": 32157829,
			"lastId": 32307537,
			"count": 149709
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	pair := "BTCUSD"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("pair", pair)
		s.assertRequestEqual(e, r)
	})
	stats, err := s.client.NewListPriceChangeStatsService().Pair(pair).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(stats, 2)
	e1 := &PriceChangeStats{
		Symbol:             "BTCUSD_210924",
		Pair:               "BTCUSD",
		PriceChange:        "-906.4",
		PriceChangePercent: "-2.204",
		WeightedAvgPrice:   "41226.29913486",
		LastPrice:          "40221.8",
		LastQuantity:       "125",
		OpenPrice:          "41128.2",
		HighPrice:          "42532.8",
		LowPrice:           "40052.0",
		Volume:             "3421917",
		BaseVolume:         "8300.32545198",
		OpenTime:           1623755580000,
		CloseTime:          1623842010140,
		FirstID:            7516015,
		LastID:             7591268,
		Count:              75254,
	}
	e2 := &PriceChangeStats{
		Symbol:             "BTCUSD_210625",
		Pair:               "BTCUSD",
		PriceChange:        "-781.7",
		PriceChangePercent: "-1.948",
		WeightedAvgPrice:   "40179.25736363",
		LastPrice:          "39337.4",
		LastQuantity:       "1",
		OpenPrice:          "40119.1",
		HighPrice:          "41405.7",
		LowPrice:           "39153.8",
		Volume:             "5479601",
		BaseVolume:         "13637.88521626",
		OpenTime:           1623755580000,
		CloseTime:          1623842010656,
		FirstID:            32157829,
		LastID:             32307537,
		Count:              149709,
	}
	s.assertPriceChangeStatsEqual(e1, stats[0])
	s.assertPriceChangeStatsEqual(e2, stats[1])
}

func (s *tickerServiceTestSuite) assertPriceChangeStatsEqual(e, a *PriceChangeStats) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.LastQuantity, a.LastQuantity, "LastQuantity")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.BaseVolume, a.BaseVolume, "BaseVolume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FirstID, a.FirstID, "FirstID")
	r.Equal(e.LastID, a.LastID, "LastID")
	r.Equal(e.Count, a.Count, "Count")
}
