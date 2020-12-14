package futures

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type websocketServiceTestSuite struct {
	baseTestSuite
	origWsServe func(*WsConfig, WsHandler, ErrHandler) (chan struct{}, chan struct{}, error)
	serveCount  int
}

func TestWebsocketService(t *testing.T) {
	suite.Run(t, new(websocketServiceTestSuite))
}

func (s *websocketServiceTestSuite) SetupTest() {
	s.origWsServe = wsServe
}

func (s *websocketServiceTestSuite) TearDownTest() {
	wsServe = s.origWsServe
	s.serveCount = 0
}

func (s *websocketServiceTestSuite) mockWsServe(data []byte, err error) {
	wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, innerErr error) {
		s.serveCount++
		doneC = make(chan struct{})
		stopC = make(chan struct{})
		go func() {
			<-stopC
			close(doneC)
		}()
		handler(data)
		if err != nil {
			errHandler(err)
		}
		return doneC, stopC, nil
	}
}

func (s *websocketServiceTestSuite) assertWsServe(count ...int) {
	e := 1
	if len(count) > 0 {
		e = count[0]
	}
	s.r().Equal(e, s.serveCount)
}

func (s *websocketServiceTestSuite) TestAggTradeServe() {
	data := []byte(`{
		"e": "aggTrade",
		"E": 123456789, 
		"s": "BTCUSDT",
		"a": 5933014,
		"p": "0.001",
		"q": "100",
		"f": 100,
		"l": 105,
		"T": 123456785,
		"m": true
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAggTradeServe("BTCUSDT", func(event *WsAggTradeEvent) {
		e := &WsAggTradeEvent{
			Event:            "aggTrade",
			Time:             123456789,
			Symbol:           "BTCUSDT",
			AggregateTradeID: 5933014,
			Price:            "0.001",
			Quantity:         "100",
			FirstTradeID:     100,
			LastTradeID:      105,
			TradeTime:        123456785,
			Maker:            true,
		}
		s.assertWsAggTradeEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsAggTradeEvent(e, a *WsAggTradeEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.AggregateTradeID, a.AggregateTradeID, "AggregateTradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.FirstTradeID, a.FirstTradeID, "FirstTradeID")
	r.Equal(e.LastTradeID, a.LastTradeID, "LastTradeID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.Maker, a.Maker, "Maker")
}

func (s *websocketServiceTestSuite) TestMarkPriceServe() {
	data := []byte(`{
		"e": "markPriceUpdate",
		"E": 1562305380000,
		"s": "BTCUSDT",
		"p": "11794.15000000",
		"i": "11784.62659091",
		"r": "0.00038167",
		"T": 1562306400000  
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarkPriceServe("BTCUSDT", func(event *WsMarkPriceEvent) {
		e := &WsMarkPriceEvent{
			Event:           "markPriceUpdate",
			Time:            1562305380000,
			Symbol:          "BTCUSDT",
			MarkPrice:       "11794.15000000",
			IndexPrice:      "11784.62659091",
			FundingRate:     "0.00038167",
			NextFundingTime: 1562306400000,
		}
		s.assertWsMarkPriceEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestAllMarkPriceServe() {
	data := []byte(`[{
		"e": "markPriceUpdate",
		"E": 1562305380000,
		"s": "BTCUSDT",
		"p": "11794.15000000",
		"i": "11784.62659091",
		"r": "0.00038167",
		"T": 1562306400000  
	  }]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMarkPriceServe(func(event WsAllMarkPriceEvent) {
		e := WsAllMarkPriceEvent{{
			Event:           "markPriceUpdate",
			Time:            1562305380000,
			Symbol:          "BTCUSDT",
			MarkPrice:       "11794.15000000",
			IndexPrice:      "11784.62659091",
			FundingRate:     "0.00038167",
			NextFundingTime: 1562306400000,
		}}
		s.assertWsMarkPriceEvent(e[0], event[0])
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMarkPriceEvent(e, a *WsMarkPriceEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.IndexPrice, a.IndexPrice, "IndexPrice")
	r.Equal(e.FundingRate, a.FundingRate, "FundingRate")
	r.Equal(e.NextFundingTime, a.NextFundingTime, "NextFundingTime")
}

func (s *websocketServiceTestSuite) TestKlineServe() {
	data := []byte(`{
		"e": "kline",
		"E": 123456789,
		"s": "BTCUSDT",
		"k": {
		  "t": 123400000,
		  "T": 123460000,
		  "s": "BTCUSDT",
		  "i": "1m",
		  "f": 100,
		  "L": 200,
		  "o": "0.0010",
		  "c": "0.0020",
		  "h": "0.0025",
		  "l": "0.0015",
		  "v": "1000",
		  "n": 100,
		  "x": false,
		  "q": "1.0000",
		  "V": "500",
		  "Q": "0.500"
		}
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsKlineServe("ETHBTC", "1m", func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   123456789,
			Symbol: "BTCUSDT",
			Kline: WsKline{
				StartTime:            123400000,
				EndTime:              123460000,
				Symbol:               "BTCUSDT",
				Interval:             "1m",
				FirstTradeID:         100,
				LastTradeID:          200,
				Open:                 "0.0010",
				Close:                "0.0020",
				High:                 "0.0025",
				Low:                  "0.0015",
				Volume:               "1000",
				TradeNum:             100,
				IsFinal:              false,
				QuoteVolume:          "1.0000",
				ActiveBuyVolume:      "500",
				ActiveBuyQuoteVolume: "0.500",
			},
		}
		s.assertWsKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsKlineEventEqual(e, a *WsKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.EndTime, ak.EndTime, "EndTime")
	r.Equal(ek.Symbol, ak.Symbol, "Symbol")
	r.Equal(ek.Interval, ak.Interval, "Interval")
	r.Equal(ek.FirstTradeID, ak.FirstTradeID, "FirstTradeID")
	r.Equal(ek.LastTradeID, ak.LastTradeID, "LastTradeID")
	r.Equal(ek.Open, ak.Open, "Open")
	r.Equal(ek.Close, ak.Close, "Close")
	r.Equal(ek.High, ak.High, "High")
	r.Equal(ek.Low, ak.Low, "Low")
	r.Equal(ek.Volume, ak.Volume, "Volume")
	r.Equal(ek.TradeNum, ak.TradeNum, "TradeNum")
	r.Equal(ek.IsFinal, ak.IsFinal, "IsFinal")
	r.Equal(ek.QuoteVolume, ak.QuoteVolume, "QuoteVolume")
	r.Equal(ek.ActiveBuyVolume, ak.ActiveBuyVolume, "ActiveBuyVolume")
	r.Equal(ek.ActiveBuyQuoteVolume, ak.ActiveBuyQuoteVolume, "ActiveBuyQuoteVolume")
}

func (s *websocketServiceTestSuite) TestMiniMarketTickerServe() {
	data := []byte(`{
		"e": "24hrMiniTicker", 
		"E": 123456789,  
		"s": "BTCUSDT", 
		"c": "0.0025",  
		"o": "0.0010", 
		"h": "0.0025",  
		"l": "0.0010", 
		"v": "10000", 
		"q": "18"  
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMiniMarketTickerServe("BTCUSDT", func(event *WsMiniMarketTickerEvent) {
		e := &WsMiniMarketTickerEvent{
			Event:       "24hrMiniTicker",
			Time:        123456789,
			Symbol:      "BTCUSDT",
			ClosePrice:  "0.0025",
			OpenPrice:   "0.0010",
			HighPrice:   "0.0025",
			LowPrice:    "0.0010",
			Volume:      "10000",
			QuoteVolume: "18",
		}
		s.assertWsMinMarketTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestAllMiniMarketTickerServe() {
	data := []byte(`[{
		"e": "24hrMiniTicker", 
		"E": 123456789,  
		"s": "BTCUSDT", 
		"c": "0.0025",  
		"o": "0.0010", 
		"h": "0.0025",  
		"l": "0.0010", 
		"v": "10000", 
		"q": "18"  
	  }]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMiniMarketTickerServe(func(event WsAllMiniMarketTickerEvent) {
		e := []*WsMiniMarketTickerEvent{{
			Event:       "24hrMiniTicker",
			Time:        123456789,
			Symbol:      "BTCUSDT",
			ClosePrice:  "0.0025",
			OpenPrice:   "0.0010",
			HighPrice:   "0.0025",
			LowPrice:    "0.0010",
			Volume:      "10000",
			QuoteVolume: "18",
		}}
		s.assertWsMinMarketTickerEvent(e[0], event[0])
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMinMarketTickerEvent(e, a *WsMiniMarketTickerEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.ClosePrice, a.ClosePrice, "ClosePrice")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
}

func (s *websocketServiceTestSuite) TestMarketTickerServe() {
	data := []byte(`{
		"e": "24hrTicker", 
		"E": 123456789, 
		"s": "BTCUSDT", 
		"p": "0.0015",
		"P": "250.00",
		"w": "0.0018",
		"c": "0.0025",
		"Q": "10", 
		"o": "0.0010",
		"h": "0.0025",
		"l": "0.0010",
		"v": "10000",
		"q": "18",
		"O": 0,
		"C": 86400000,
		"F": 0,
		"L": 18150,
		"n": 18151
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarketTickerServe("BTCUSDT", func(event *WsMarketTickerEvent) {
		e := &WsMarketTickerEvent{
			Event:              "24hrTicker",
			Time:               123456789,
			Symbol:             "BTCUSDT",
			PriceChange:        "0.0015",
			PriceChangePercent: "250.00",
			WeightedAvgPrice:   "0.0018",
			ClosePrice:         "0.0025",
			CloseQty:           "10",
			OpenPrice:          "0.0010",
			HighPrice:          "0.0025",
			LowPrice:           "0.0010",
			BaseVolume:         "10000",
			QuoteVolume:        "18",
			OpenTime:           0,
			CloseTime:          86400000,
			FirstID:            0,
			LastID:             18150,
			TradeCount:         18151,
		}
		s.assertWsMarketTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestAllMarketTickerServe() {
	data := []byte(`[{
		"e": "24hrTicker", 
		"E": 123456789, 
		"s": "BTCUSDT", 
		"p": "0.0015",
		"P": "250.00",
		"w": "0.0018",
		"c": "0.0025",
		"Q": "10", 
		"o": "0.0010",
		"h": "0.0025",
		"l": "0.0010",
		"v": "10000",
		"q": "18",
		"O": 0,
		"C": 86400000,
		"F": 0,
		"L": 18150,
		"n": 18151
	  }]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMarketTickerServe(func(event WsAllMarketTickerEvent) {
		e := WsAllMarketTickerEvent{{
			Event:              "24hrTicker",
			Time:               123456789,
			Symbol:             "BTCUSDT",
			PriceChange:        "0.0015",
			PriceChangePercent: "250.00",
			WeightedAvgPrice:   "0.0018",
			ClosePrice:         "0.0025",
			CloseQty:           "10",
			OpenPrice:          "0.0010",
			HighPrice:          "0.0025",
			LowPrice:           "0.0010",
			BaseVolume:         "10000",
			QuoteVolume:        "18",
			OpenTime:           0,
			CloseTime:          86400000,
			FirstID:            0,
			LastID:             18150,
			TradeCount:         18151,
		}}
		s.assertWsMarketTickerEvent(e[0], event[0])
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMarketTickerEvent(e, a *WsMarketTickerEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.ClosePrice, a.ClosePrice, "ClosePrice")
	r.Equal(e.CloseQty, a.CloseQty, "CloseQty")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.BaseVolume, a.BaseVolume, "BaseVolume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FirstID, a.FirstID, "FirstID")
	r.Equal(e.LastID, a.LastID, "LastID")
	r.Equal(e.TradeCount, a.TradeCount, "TradeCount")
}

func (s *websocketServiceTestSuite) TestBookTickerServe() {
	data := []byte(`{
		"e":"bookTicker",    
		"u":400900217,
		"E": 1568014460893,
		"T": 1568014460891,
		"s":"BNBUSDT",
		"b":"25.35190000",
		"B":"31.21000000",
		"a":"25.36520000",
		"A":"40.66000000"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsBookTickerServe("BNBUSDT", func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			Event:           "bookTicker",
			UpdateID:        400900217,
			Time:            1568014460893,
			TransactionTime: 1568014460891,
			Symbol:          "BNBUSDT",
			BestBidPrice:    "25.35190000",
			BestBidQty:      "31.21000000",
			BestAskPrice:    "25.36520000",
			BestAskQty:      "40.66000000",
		}
		s.assertWsBookTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestAllBookTickerServe() {
	data := []byte(`{
		"e":"bookTicker",    
		"u":400900217,
		"E": 1568014460893,
		"T": 1568014460891,
		"s":"BNBUSDT",
		"b":"25.35190000",
		"B":"31.21000000",
		"a":"25.36520000",
		"A":"40.66000000"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllBookTickerServe(func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			Event:           "bookTicker",
			UpdateID:        400900217,
			Time:            1568014460893,
			TransactionTime: 1568014460891,
			Symbol:          "BNBUSDT",
			BestBidPrice:    "25.35190000",
			BestBidQty:      "31.21000000",
			BestAskPrice:    "25.36520000",
			BestAskQty:      "40.66000000",
		}
		s.assertWsBookTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsBookTickerEvent(e, a *WsBookTickerEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.UpdateID, a.UpdateID, "UpdateID")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.BestBidPrice, a.BestBidPrice, "BestBidPrice")
	r.Equal(e.BestBidQty, a.BestBidQty, "BestBidQty")
	r.Equal(e.BestAskPrice, a.BestAskPrice, "BestAskPrice")
	r.Equal(e.BestAskQty, a.BestAskQty, "BestAskQty")
}

func (s *websocketServiceTestSuite) TestLiquidationOrderServe() {
	data := []byte(`{
		"e":"forceOrder", 
		"E":1568014460893,
		"o":{
			"s":"BTCUSDT",
			"S":"SELL",
			"o":"LIMIT",
			"f":"IOC", 
			"q":"0.014",
			"p":"9910",
			"ap":"9910",
			"X":"FILLED",
			"l":"0.014",
			"z":"0.014",
			"T":1568014460893
		}
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsLiquidationOrderServe("BTCUSDT", func(event *WsLiquidationOrderEvent) {
		e := &WsLiquidationOrderEvent{
			Event: "forceOrder",
			Time:  1568014460893,
			LiquidationOrder: WsLiquidationOrder{
				Symbol:               "BTCUSDT",
				Side:                 SideTypeSell,
				OrderType:            OrderTypeLimit,
				TimeInForce:          TimeInForceTypeIOC,
				OrigQuantity:         "0.014",
				Price:                "9910",
				AvgPrice:             "9910",
				OrderStatus:          OrderStatusTypeFilled,
				LastFilledQty:        "0.014",
				AccumulatedFilledQty: "0.014",
				TradeTime:            1568014460893,
			},
		}
		s.assertLiquidationOrderEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestAllLiquidationOrderServe() {
	data := []byte(`{
		"e":"forceOrder", 
		"E":1568014460893,
		"o":{
			"s":"BTCUSDT",
			"S":"SELL",
			"o":"LIMIT",
			"f":"IOC", 
			"q":"0.014",
			"p":"9910",
			"ap":"9910",
			"X":"FILLED",
			"l":"0.014",
			"z":"0.014",
			"T":1568014460893
		}
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllLiquidationOrderServe(func(event *WsLiquidationOrderEvent) {
		e := &WsLiquidationOrderEvent{
			Event: "forceOrder",
			Time:  1568014460893,
			LiquidationOrder: WsLiquidationOrder{
				Symbol:               "BTCUSDT",
				Side:                 SideTypeSell,
				OrderType:            OrderTypeLimit,
				TimeInForce:          TimeInForceTypeIOC,
				OrigQuantity:         "0.014",
				Price:                "9910",
				AvgPrice:             "9910",
				OrderStatus:          OrderStatusTypeFilled,
				LastFilledQty:        "0.014",
				AccumulatedFilledQty: "0.014",
				TradeTime:            1568014460893,
			},
		}
		s.assertLiquidationOrderEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertLiquidationOrderEvent(e, a *WsLiquidationOrderEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	elo, alo := e.LiquidationOrder, a.LiquidationOrder
	r.Equal(elo.Symbol, alo.Symbol, "Symbol")
	r.Equal(elo.Side, alo.Side, "Side")
	r.Equal(elo.OrderType, alo.OrderType, "OrderType")
	r.Equal(elo.TimeInForce, alo.TimeInForce, "TimeInForce")
	r.Equal(elo.OrigQuantity, alo.OrigQuantity, "OrigQuantity")
	r.Equal(elo.Price, alo.Price, "Price")
	r.Equal(elo.AvgPrice, alo.AvgPrice, "AvgPrice")
	r.Equal(elo.OrderStatus, alo.OrderStatus, "OrderStatus")
	r.Equal(elo.LastFilledQty, alo.LastFilledQty, "LastFilledQty")
	r.Equal(elo.AccumulatedFilledQty, alo.AccumulatedFilledQty, "AccumulatedFilledQty")
	r.Equal(elo.TradeTime, alo.TradeTime, "TradeTime")
}

func (s *websocketServiceTestSuite) TestPartialDepthServe() {
	data := []byte(`{
		"e": "depthUpdate", 
		"E": 1571889248277, 
		"T": 1571889248276,
		"s": "BTCUSDT",
		"U": 390497796,
		"u": 390497878,
		"pu": 390497794,
		"b": [
		  [
			"7403.89",
			"0.002"
		  ]
		],
		"a": [ 
		  [
			"7405.96",
			"3.340"
		  ]
		]
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsPartialDepthServe("BTCUSDT", 5, func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:            "depthUpdate",
			Time:             1571889248277,
			TransactionTime:  1571889248276,
			Symbol:           "BTCUSDT",
			FirstUpdateID:    390497796,
			LastUpdateID:     390497878,
			PrevLastUpdateID: 390497794,
			Bids:             []Bid{{Price: "7403.89", Quantity: "0.002"}},
			Asks:             []Ask{{Price: "7405.96", Quantity: "3.340"}},
		}
		s.assertDepthEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestDiffDepthServe() {
	data := []byte(`{
		"e": "depthUpdate",
		"E": 123456789,
		"T": 123456788,
		"s": "BTCUSDT",
		"U": 157,
		"u": 160,
		"pu": 149,
		"b": [
		  [
			"0.0024",
			"10"
		  ]
		],
		"a": [
		  [
			"0.0026",
			"100"
		  ]
		]
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsPartialDepthServe("BTCUSDT", 5, func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:            "depthUpdate",
			Time:             123456789,
			TransactionTime:  123456788,
			Symbol:           "BTCUSDT",
			FirstUpdateID:    157,
			LastUpdateID:     160,
			PrevLastUpdateID: 149,
			Bids:             []Bid{{Price: "0.0024", Quantity: "10"}},
			Asks:             []Ask{{Price: "0.0026", Quantity: "100"}},
		}
		s.assertDepthEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertDepthEvent(e, a *WsDepthEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.FirstUpdateID, a.FirstUpdateID, "FirstUpdateID")
	r.Equal(e.LastUpdateID, a.LastUpdateID, "LastUpdateID")
	r.Equal(e.PrevLastUpdateID, a.PrevLastUpdateID, "PrevLastUpdateID")
	for i, b := range e.Bids {
		r.Equal(b.Price, a.Bids[i].Price, "Price")
		r.Equal(b.Quantity, a.Bids[i].Quantity, "Quantity")
	}
	for i, b := range e.Asks {
		r.Equal(b.Price, a.Asks[i].Price, "Price")
		r.Equal(b.Quantity, a.Asks[i].Quantity, "Quantity")
	}
}

func (s *websocketServiceTestSuite) TestBLVTInfoServe() {
	data := []byte(`{
		"e":"nav",
		"E":1600245286355,
		"s":"TRXDOWN",
		"m":74164.75496502663,
		"b":[
			{
				"s":"TRXUSDT",
				"n":-87988261
			}
		],
		"n":14.78454447,
		"l":2.1786579638117898,
		"t":3,
		"f":-0.0048925
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsBLVTInfoServe("TRXDOWN", func(event *WsBLVTInfoEvent) {
		e := &WsBLVTInfoEvent{
			Event:          "nav",
			Time:           1600245286355,
			Symbol:         "TRXDOWN",
			Issued:         74164.75496502663,
			Baskets:        []WsBLVTBasket{{Symbol: "TRXUSDT", Position: -87988261}},
			Nav:            14.78454447,
			Leverage:       2.1786579638117898,
			TargetLeverage: 3,
			FundingRate:    -0.0048925,
		}
		s.assertBLVTInfoEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertBLVTInfoEvent(e, a *WsBLVTInfoEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Issued, a.Issued, "Issued")
	r.Equal(e.Nav, a.Nav, "Nav")
	r.Equal(e.Leverage, a.Leverage, "Leverage")
	r.Equal(e.TargetLeverage, a.TargetLeverage, "TargetLeverage")
	r.Equal(e.FundingRate, a.FundingRate, "FundingRate")
	for i, b := range e.Baskets {
		r.Equal(b.Position, a.Baskets[i].Position, "Position")
		r.Equal(b.Symbol, a.Baskets[i].Symbol, "Symbol")
	}
}

func (s *websocketServiceTestSuite) TestBLVTKlineServe() {
	data := []byte(`{
		"e":"kline",
		"E":1600243159447,
		"s":"TRXDOWN",
		"k":{
			"t":1600243140000,
			"T":1600243199999,
			"s":"TRXDOWN",
			"i":"1m", 
			"f":1600243140484, 
			"L":1600243159424,
			"o":"14.56800297", 
			"c":"14.59766270",
			"h":"14.63325437",
			"l":"14.56207102",
			"v":"2.22524220",
			"n":33
	   }
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsBLVTKlineServe("TRXDOWN", "1m", func(event *WsBLVTKlineEvent) {
		e := &WsBLVTKlineEvent{
			Event:  "kline",
			Time:   1600243159447,
			Symbol: "TRXDOWN",
			Kline: WsBLVTKline{
				StartTime:       1600243140000,
				CloseTime:       1600243199999,
				Symbol:          "TRXDOWN",
				Interval:        "1m",
				FirstUpdateTime: 1600243140484,
				LastUpdateTime:  1600243159424,
				OpenPrice:       "14.56800297",
				ClosePrice:      "14.59766270",
				HighPrice:       "14.63325437",
				LowPrice:        "14.56207102",
				Leverage:        "2.22524220",
				Count:           33,
			},
		}
		s.assertBLVTKlineEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertBLVTKlineEvent(e, a *WsBLVTKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.CloseTime, ak.CloseTime, "CloseTime")
	r.Equal(ek.Symbol, ak.Symbol, "Symbol")
	r.Equal(ek.Interval, ak.Interval, "Interval")
	r.Equal(ek.FirstUpdateTime, ak.FirstUpdateTime, "FirstUpdateTime")
	r.Equal(ek.LastUpdateTime, ak.LastUpdateTime, "LastUpdateTime")
	r.Equal(ek.OpenPrice, ak.OpenPrice, "OpenPrice")
	r.Equal(ek.ClosePrice, ak.ClosePrice, "ClosePrice")
	r.Equal(ek.HighPrice, ak.HighPrice, "HighPrice")
	r.Equal(ek.LowPrice, ak.LowPrice, "LowPrice")
	r.Equal(ek.Leverage, ak.Leverage, "Leverage")
	r.Equal(ek.Count, ak.Count, "Count")
}

func (s *websocketServiceTestSuite) TestWsCompositiveIndexServe() {
	data := []byte(`{
		"e":"compositeIndex",
		"E":1602310596000,
		"s":"DEFIUSDT",
		"p":"554.41604065",
		"c":[
		  {
			  "b":"BAL",
			  "w":"1.35038833",
			  "W":"0.03957100"
		  },
		  {
			"b":"BAND",
			"w":"3.53782729",
			"W":"0.03935200"
		  }
		]
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsCompositiveIndexServe("TRXDOWN", func(event *WsCompositeIndexEvent) {
		e := &WsCompositeIndexEvent{
			Event:  "compositeIndex",
			Time:   1602310596000,
			Symbol: "DEFIUSDT",
			Price:  "554.41604065",
			Composition: []WsComposition{
				{
					BaseAsset:    "BAL",
					WeightQty:    "1.35038833",
					WeighPercent: "0.03957100",
				},
				{
					BaseAsset:    "BAND",
					WeightQty:    "3.53782729",
					WeighPercent: "0.03935200",
				},
			},
		}
		s.assertCompositeIndexEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertCompositeIndexEvent(e, a *WsCompositeIndexEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
	for i, c := range e.Composition {
		r.Equal(c.BaseAsset, a.Composition[i].BaseAsset, "Position")
		r.Equal(c.WeightQty, a.Composition[i].WeightQty, "WeightQty")
		r.Equal(c.WeighPercent, a.Composition[i].WeighPercent, "WeighPercent")
	}
}
