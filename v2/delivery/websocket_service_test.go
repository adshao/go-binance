package delivery

import (
	"errors"
	"math/rand"
	"testing"
	"time"

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

// https://binance-docs.github.io/apidocs/delivery/en/#aggregate-trade-streams
func (s *websocketServiceTestSuite) TestAggTradeServe() {
	data := []byte(`{
		  "e":"aggTrade",
		  "E":1591261134288,
		  "a":424951,
		  "s":"BTCUSD_200626",
		  "p":"9643.5",
		  "q":"2",
		  "f":606073,
		  "l":606073,
		  "T":1591261134199,
		  "m":false
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAggTradeServe("BTCUSD_200626", func(event *WsAggTradeEvent) {
		e := &WsAggTradeEvent{
			Event:            "aggTrade",
			Time:             1591261134288,
			AggregateTradeID: 424951,
			Symbol:           "BTCUSD_200626",
			Price:            "9643.5",
			Quantity:         "2",
			FirstTradeID:     606073,
			LastTradeID:      606073,
			TradeTime:        1591261134199,
			Maker:            false,
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

// https://binance-docs.github.io/apidocs/delivery/en/#index-price-stream
func (s *websocketServiceTestSuite) TestIndexPriceServe() {
	data := []byte(`{
		"e": "indexPriceUpdate",
		"E": 1591261236000,
		"i": "BTCUSD",
		"p": "9636.57860000"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsIndexPriceServe("BTCUSD", func(event *WsIndexPriceEvent) {
		e := &WsIndexPriceEvent{
			Event:      "indexPriceUpdate",
			Time:       1591261236000,
			Pair:       "BTCUSD",
			IndexPrice: "9636.57860000",
		}
		s.assertWsIndexPriceEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsIndexPriceEvent(e, a *WsIndexPriceEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.IndexPrice, a.IndexPrice, "IndexPrice")
}

// https://binance-docs.github.io/apidocs/delivery/en/#mark-price-stream
func (s *websocketServiceTestSuite) TestMarkPriceServe() {
	data := []byte(`{
    	"e":"markPriceUpdate",
    	"E":1596095725000,
    	"s":"BTCUSD_201225",
    	"p":"10934.62615417",
    	"P":"10962.17178236",
    	"r":"",
    	"T":0
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarkPriceServe("BTCUSD_201225", func(event *WsMarkPriceEvent) {
		e := &WsMarkPriceEvent{
			Event:                "markPriceUpdate",
			Time:                 1596095725000,
			Symbol:               "BTCUSD_201225",
			MarkPrice:            "10934.62615417",
			EstimatedSettlePrice: "10962.17178236",
			FundingRate:          "",
			NextFundingTime:      0,
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

// https://binance-docs.github.io/apidocs/delivery/en/#mark-price-of-all-symbols-of-a-pair
func (s *websocketServiceTestSuite) TestPairMarkPriceServe() {
	data := []byte(`[
	  {
		"e":"markPriceUpdate",
		"E":1596095725000,
		"s":"BTCUSD_201225",
		"p":"10934.62615417",
		"P":"10962.17178236",
		"r":"",
		"T":0
	  },
	  {
		"e":"markPriceUpdate",
		"E":1596095725000,
		"s":"BTCUSD_PERP",
		"p":"11012.31359011",
		"P":"10962.17178236",
		"r":"0.00000000",
		"T":1596096000000
  }]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsPairMarkPriceServe(func(event WsPairMarkPriceEvent) {
		e := WsPairMarkPriceEvent{{
			Event:                "markPriceUpdate",
			Time:                 1596095725000,
			Symbol:               "BTCUSD_201225",
			MarkPrice:            "10934.62615417",
			EstimatedSettlePrice: "10962.17178236",
			FundingRate:          "",
			NextFundingTime:      0,
		}, {
			Event:                "markPriceUpdate",
			Time:                 1596095725000,
			Symbol:               "BTCUSD_PERP",
			MarkPrice:            "11012.31359011",
			EstimatedSettlePrice: "10962.17178236",
			FundingRate:          "0.00000000",
			NextFundingTime:      1596096000000,
		}}
		s.assertWsMarkPriceEvent(e[0], event[0])
		s.assertWsMarkPriceEvent(e[1], event[1])
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
	r.Equal(e.EstimatedSettlePrice, a.EstimatedSettlePrice, "EstimatedSettlePrice")
	r.Equal(e.FundingRate, a.FundingRate, "FundingRate")
	r.Equal(e.NextFundingTime, a.NextFundingTime, "NextFundingTime")
}

// https://binance-docs.github.io/apidocs/delivery/en/#kline-candlestick-streams
func (s *websocketServiceTestSuite) TestKlineServe() {
	data := []byte(`{
	  "e":"kline",
	  "E":1591261542539,
	  "s":"BTCUSD_200626",
	  "k":{
		"t":1591261500000,
		"T":1591261559999,
		"s":"BTCUSD_200626",
		"i":"1m",
		"f":606400,
		"L":606430,
		"o":"9638.9",
		"c":"9639.8",
		"h":"9639.8",
		"l":"9638.6",
		"v":"156",
		"n":31,
		"x":false,
		"q":"1.61836886",
		"V":"73",
		"Q":"0.75731156",
		"B":"0"
      }
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsKlineServe("BTCUSD_200626", "1m", func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   1591261542539,
			Symbol: "BTCUSD_200626",
			Kline: WsKline{
				StartTime:            1591261500000,
				EndTime:              1591261559999,
				Symbol:               "BTCUSD_200626",
				Interval:             "1m",
				FirstTradeID:         606400,
				LastTradeID:          606430,
				Open:                 "9638.9",
				Close:                "9639.8",
				High:                 "9639.8",
				Low:                  "9638.6",
				Volume:               "156",
				TradeNum:             31,
				IsFinal:              false,
				QuoteVolume:          "1.61836886",
				ActiveBuyVolume:      "73",
				ActiveBuyQuoteVolume: "0.75731156",
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

// https://binance-docs.github.io/apidocs/delivery/en/#continues-contract-kline-candlestick-streams
func (s *websocketServiceTestSuite) TestContinuousKlineServe() {
	data := []byte(`{
	  "e":"continuous_kline",
	  "E":1591261542539,
	  "ps":"BTCUSD",
      "ct":"NEXT_QUARTER",
	  "k":{
		"t":1591261500000,
		"T":1591261559999,
		"i":"1m",
		"f":606400,
		"L":606430,
		"o":"9638.9",
		"c":"9639.8",
		"h":"9639.8",
		"l":"9638.6",
		"v":"156",
		"n":31,
		"x":false,
		"q":"1.61836886",
		"V":"73",
		"Q":"0.75731156",
		"B":"0"
      }
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsContinuousKlineServe("BTCUSD", "NEXT_QUARTER", "1m", func(event *WsContinuousKlineEvent) {
		e := &WsContinuousKlineEvent{
			Event:        "continuous_kline",
			Time:         1591261542539,
			Pair:         "BTCUSD",
			ContractType: "NEXT_QUARTER",
			Kline: WsContinuousKline{
				StartTime:            1591261500000,
				EndTime:              1591261559999,
				Interval:             "1m",
				FirstTradeID:         606400,
				LastTradeID:          606430,
				Open:                 "9638.9",
				Close:                "9639.8",
				High:                 "9639.8",
				Low:                  "9638.6",
				Volume:               "156",
				TradeNum:             31,
				IsFinal:              false,
				QuoteVolume:          "1.61836886",
				ActiveBuyVolume:      "73",
				ActiveBuyQuoteVolume: "0.75731156",
			},
		}
		s.assertWsContinuousKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsContinuousKlineEventEqual(e, a *WsContinuousKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.ContractType, a.ContractType, "ContractType")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.EndTime, ak.EndTime, "EndTime")
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

// https://binance-docs.github.io/apidocs/delivery/en/#index-kline-candlestick-streams
func (s *websocketServiceTestSuite) TestIndexKlineServe() {
	data := []byte(`{
	  "e":"indexPrice_kline",
	  "E":1591267070033,
	  "ps":"BTCUSD",
	  "k":{
		"t":1591267020000,
		"T":1591267079999,
		"s":"0",
		"i":"1m",
		"f":1591267020000,
		"L":1591267070000,
		"o":"9542.21900000",
		"c":"9542.50440000",
		"h":"9542.71640000",
		"l":"9542.21040000",
		"v":"0",
		"n":51,
		"x":false,
		"q":"0",
		"V":"0",
		"Q":"0",
		"B":"0"
	  }
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsIndexPriceKlineServe("BTCUSD", "1m", func(event *WsIndexPriceKlineEvent) {
		e := &WsIndexPriceKlineEvent{
			Event: "indexPrice_kline",
			Time:  1591267070033,
			Pair:  "BTCUSD",
			Kline: WsIndexPriceKline{
				StartTime: 1591267020000,
				EndTime:   1591267079999,
				Interval:  "1m",
				Open:      "9542.21900000",
				Close:     "9542.50440000",
				High:      "9542.71640000",
				Low:       "9542.21040000",
				TradeNum:  51,
				IsFinal:   false,
			},
		}
		s.assertWsIndexKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsIndexKlineEventEqual(e, a *WsIndexPriceKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Pair, a.Pair, "Pair")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.EndTime, ak.EndTime, "EndTime")
	r.Equal(ek.Interval, ak.Interval, "Interval")
	r.Equal(ek.Open, ak.Open, "Open")
	r.Equal(ek.Close, ak.Close, "Close")
	r.Equal(ek.High, ak.High, "High")
	r.Equal(ek.Low, ak.Low, "Low")
	r.Equal(ek.TradeNum, ak.TradeNum, "TradeNum")
	r.Equal(ek.IsFinal, ak.IsFinal, "IsFinal")
}

// https://binance-docs.github.io/apidocs/delivery/en/#mark-price-kline-candlestick-streams
func (s *websocketServiceTestSuite) TestMarkPriceKlineServe() {
	data := []byte(`{
	  "e":"markPrice_kline",
	  "E":1591267398004,
	  "ps":"BTCUSD",
	  "k":{
		"t":1591267380000,
		"T":1591267439999,
		"s":"BTCUSD_200626",
		"i":"1m",
		"f":1591267380000,
		"L":1591267398000,
		"o":"9539.67161333",
		"c":"9540.82761333",
		"h":"9540.82761333",
		"l":"9539.66961333",
		"v":"0",
		"n":19,
		"x":false,
		"q":"0",
		"V":"0",
		"Q":"0",
		"B":"0"
	  }
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarkPriceKlineServe("BTCUSD_200626", "1m", func(event *WsMarkPriceKlineEvent) {
		e := &WsMarkPriceKlineEvent{
			Event: "markPrice_kline",
			Time:  1591267398004,
			Pair:  "BTCUSD",
			Kline: WsMarkPriceKline{
				StartTime: 1591267380000,
				EndTime:   1591267439999,
				Symbol:    "BTCUSD_200626",
				Interval:  "1m",
				Open:      "9539.67161333",
				Close:     "9540.82761333",
				High:      "9540.82761333",
				Low:       "9539.66961333",
				TradeNum:  19,
				IsFinal:   false,
			},
		}
		s.assertWsMarkPriceKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMarkPriceKlineEventEqual(e, a *WsMarkPriceKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Pair, a.Pair, "Pair")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.EndTime, ak.EndTime, "EndTime")
	r.Equal(ek.Interval, ak.Interval, "Interval")
	r.Equal(ek.Symbol, ak.Symbol, "Symbol")
	r.Equal(ek.Open, ak.Open, "Open")
	r.Equal(ek.Close, ak.Close, "Close")
	r.Equal(ek.High, ak.High, "High")
	r.Equal(ek.Low, ak.Low, "Low")
	r.Equal(ek.TradeNum, ak.TradeNum, "TradeNum")
	r.Equal(ek.IsFinal, ak.IsFinal, "IsFinal")
}

// https://binance-docs.github.io/apidocs/delivery/en/#individual-symbol-mini-ticker-stream
func (s *websocketServiceTestSuite) TestMiniMarketTickerServe() {
	data := []byte(`{
	    "e":"24hrMiniTicker",
	    "E":1591267704450,
	    "s":"BTCUSD_200626",
	    "ps":"BTCUSD",
	    "c":"9561.7",
	    "o":"9580.9",
	    "h":"10000.0",
	    "l":"7000.0",
	    "v":"487476",
	    "q":"33264343847.22378500"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMiniMarketTickerServe("BTCUSD_200626", func(event *WsMiniMarketTickerEvent) {
		e := &WsMiniMarketTickerEvent{
			Event:       "24hrMiniTicker",
			Time:        1591267704450,
			Symbol:      "BTCUSD_200626",
			Pair:        "BTCUSD",
			ClosePrice:  "9561.7",
			OpenPrice:   "9580.9",
			HighPrice:   "10000.0",
			LowPrice:    "7000.0",
			Volume:      "487476",
			QuoteVolume: "33264343847.22378500",
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

// https://binance-docs.github.io/apidocs/delivery/en/#all-market-mini-tickers-stream
func (s *websocketServiceTestSuite) TestAllMiniMarketTickerServe() {
	data := []byte(`[{
		"e":"24hrMiniTicker",
		"E":1591267704450,
		"s":"BTCUSD_200626",
		"ps":"BTCUSD",
		"c":"9561.7",
		"o":"9580.9",
		"h":"10000.0",
		"l":"7000.0",
		"v":"487476",
		"q":"33264343847.22378500"
	  }]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMiniMarketTickerServe(func(event WsAllMiniMarketTickerEvent) {
		e := []*WsMiniMarketTickerEvent{{
			Event:       "24hrMiniTicker",
			Time:        1591267704450,
			Symbol:      "BTCUSD_200626",
			Pair:        "BTCUSD",
			ClosePrice:  "9561.7",
			OpenPrice:   "9580.9",
			HighPrice:   "10000.0",
			LowPrice:    "7000.0",
			Volume:      "487476",
			QuoteVolume: "33264343847.22378500",
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
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.ClosePrice, a.ClosePrice, "ClosePrice")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
}

// https://binance-docs.github.io/apidocs/delivery/en/#individual-symbol-ticker-streams
func (s *websocketServiceTestSuite) TestMarketTickerServe() {
	data := []byte(`{
		"e":"24hrTicker",
		"E":1591268262453,
		"s":"BTCUSD_200626",
		"ps":"BTCUSD",
		"p":"-43.4",
		"P":"-0.452",
		"w":"0.00147974",
		"c":"9548.5",
		"Q":"2",
		"o":"9591.9",
		"h":"10000.0",
		"l":"7000.0",
		"v":"487850",
		"q":"32968676323.46222700",
		"O":1591181820000,
		"C":1591268262442,
		"F":512014,
		"L":615289,
		"n":103272
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarketTickerServe("BTCUSDT", func(event *WsMarketTickerEvent) {
		e := &WsMarketTickerEvent{
			Event:              "24hrTicker",
			Time:               1591268262453,
			Symbol:             "BTCUSD_200626",
			Pair:               "BTCUSD",
			PriceChange:        "-43.4",
			PriceChangePercent: "-0.452",
			WeightedAvgPrice:   "0.00147974",
			ClosePrice:         "9548.5",
			CloseQty:           "2",
			OpenPrice:          "9591.9",
			HighPrice:          "10000.0",
			LowPrice:           "7000.0",
			BaseVolume:         "487850",
			QuoteVolume:        "32968676323.46222700",
			OpenTime:           1591181820000,
			CloseTime:          1591268262442,
			FirstID:            512014,
			LastID:             615289,
			TradeCount:         103272,
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

// https://binance-docs.github.io/apidocs/delivery/en/#all-market-tickers-streams
func (s *websocketServiceTestSuite) TestAllMarketTickerServe() {
	data := []byte(`[{
		"e":"24hrTicker",
		"E":1591268262453,
		"s":"BTCUSD_200626",
		"ps":"BTCUSD",
		"p":"-43.4",
		"P":"-0.452",
		"w":"0.00147974",
		"c":"9548.5",
		"Q":"2",
		"o":"9591.9",
		"h":"10000.0",
		"l":"7000.0",
		"v":"487850",
		"q":"32968676323.46222700",
		"O":1591181820000,
		"C":1591268262442,
		"F":512014,
		"L":615289,
		"n":103272
	  }]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMarketTickerServe(func(event WsAllMarketTickerEvent) {
		e := WsAllMarketTickerEvent{{
			Event:              "24hrTicker",
			Time:               1591268262453,
			Symbol:             "BTCUSD_200626",
			Pair:               "BTCUSD",
			PriceChange:        "-43.4",
			PriceChangePercent: "-0.452",
			WeightedAvgPrice:   "0.00147974",
			ClosePrice:         "9548.5",
			CloseQty:           "2",
			OpenPrice:          "9591.9",
			HighPrice:          "10000.0",
			LowPrice:           "7000.0",
			BaseVolume:         "487850",
			QuoteVolume:        "32968676323.46222700",
			OpenTime:           1591181820000,
			CloseTime:          1591268262442,
			FirstID:            512014,
			LastID:             615289,
			TradeCount:         103272,
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
	r.Equal(e.Pair, a.Pair, "Pair")
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

// https://binance-docs.github.io/apidocs/delivery/en/#individual-symbol-book-ticker-streams
func (s *websocketServiceTestSuite) TestBookTickerServe() {
	data := []byte(`{
  		"e":"bookTicker",
  		"u":17242169,
  		"s":"BTCUSD_200626",
  		"ps":"BTCUSD",
  		"b":"9548.1",
  		"B":"52",
  		"a":"9548.5",
  		"A":"11",
  		"T":1591268628155,
  		"E":1591268628166
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsBookTickerServe("BTCUSD_200626", func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			Event:           "bookTicker",
			UpdateID:        17242169,
			Symbol:          "BTCUSD_200626",
			Pair:            "BTCUSD",
			BestBidPrice:    "9548.1",
			BestBidQty:      "52",
			BestAskPrice:    "9548.5",
			BestAskQty:      "11",
			TransactionTime: 1591268628155,
			Time:            1591268628166,
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

// https://binance-docs.github.io/apidocs/delivery/en/#all-book-tickers-stream
func (s *websocketServiceTestSuite) TestAllBookTickerServe() {
	data := []byte(`{
  		"e":"bookTicker",
  		"u":17242169,
  		"s":"BTCUSD_200626",
  		"ps":"BTCUSD",
  		"b":"9548.1",
  		"B":"52",
  		"a":"9548.5",
  		"A":"11",
  		"T":1591268628155,
  		"E":1591268628166
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllBookTickerServe(func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			Event:           "bookTicker",
			UpdateID:        17242169,
			Symbol:          "BTCUSD_200626",
			Pair:            "BTCUSD",
			BestBidPrice:    "9548.1",
			BestBidQty:      "52",
			BestAskPrice:    "9548.5",
			BestAskQty:      "11",
			TransactionTime: 1591268628155,
			Time:            1591268628166,
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
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.BestBidPrice, a.BestBidPrice, "BestBidPrice")
	r.Equal(e.BestBidQty, a.BestBidQty, "BestBidQty")
	r.Equal(e.BestAskPrice, a.BestAskPrice, "BestAskPrice")
	r.Equal(e.BestAskQty, a.BestAskQty, "BestAskQty")
}

// https://binance-docs.github.io/apidocs/delivery/en/#liquidation-order-streams
func (s *websocketServiceTestSuite) TestLiquidationOrderServe() {
	data := []byte(`{
	  "e":"forceOrder",
	  "E": 1591154240950,
	  "o":{
	    "s":"BTCUSD_200925",
	    "ps": "BTCUSD",
	    "S":"SELL",
	    "o":"LIMIT",
	    "f":"IOC",
	    "q":"1",
	    "p":"9425.5",
	    "ap":"9496.5",
	    "X":"FILLED",
	    "l":"1",
	    "z":"1",
	    "T": 1591154240949
	  }
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsLiquidationOrderServe("BTCUSD_200925", func(event *WsLiquidationOrderEvent) {
		e := &WsLiquidationOrderEvent{
			Event: "forceOrder",
			Time:  1591154240950,
			LiquidationOrder: WsLiquidationOrder{
				Symbol:               "BTCUSD_200925",
				Pair:                 "BTCUSD",
				Side:                 SideTypeSell,
				OrderType:            OrderTypeLimit,
				TimeInForce:          TimeInForceTypeIOC,
				OrigQuantity:         "1",
				Price:                "9425.5",
				AvgPrice:             "9496.5",
				OrderStatus:          OrderStatusTypeFilled,
				LastFilledQty:        "1",
				AccumulatedFilledQty: "1",
				TradeTime:            1591154240949,
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

// https://binance-docs.github.io/apidocs/delivery/en/#all-market-liquidation-order-streams
func (s *websocketServiceTestSuite) TestAllLiquidationOrderServe() {
	data := []byte(`{
	  "e":"forceOrder",
	  "E": 1591154240950,
	  "o":{
		"s":"BTCUSD_200925",
		"ps": "BTCUSD",
		"S":"SELL",
		"o":"LIMIT",
		"f":"IOC",
		"q":"1",
		"p":"9425.5",
		"ap":"9496.5",
		"X":"FILLED",
		"l":"1",
		"z":"1",
		"T": 1591154240949
	  }
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllLiquidationOrderServe(func(event *WsLiquidationOrderEvent) {
		e := &WsLiquidationOrderEvent{
			Event: "forceOrder",
			Time:  1591154240950,
			LiquidationOrder: WsLiquidationOrder{
				Symbol:               "BTCUSD_200925",
				Pair:                 "BTCUSD",
				Side:                 SideTypeSell,
				OrderType:            OrderTypeLimit,
				TimeInForce:          TimeInForceTypeIOC,
				OrigQuantity:         "1",
				Price:                "9425.5",
				AvgPrice:             "9496.5",
				OrderStatus:          OrderStatusTypeFilled,
				LastFilledQty:        "1",
				AccumulatedFilledQty: "1",
				TradeTime:            1591154240949,
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
	r.Equal(elo.Pair, alo.Pair, "Pair")
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

// https://binance-docs.github.io/apidocs/delivery/en/#partial-book-depth-streams
func (s *websocketServiceTestSuite) testPartialDepthServe(levels int, rate *time.Duration, expectedErr error, expectedServeCnt int) {
	data := []byte(`{
	  "e":"depthUpdate",     
	  "E":1591269996801,     
	  "T":1591269996646,     
	  "s":"BTCUSD_200626",   
	  "ps":"BTCUSD",         
	  "U":17276694,
	  "u":17276701,
	  "pu":17276678,
	  "b":[                  
		["9523.0","5"],
		["9522.8","8"],
		["9522.6","2"],
		["9522.4","1"],
		["9522.0","5"]
	  ],
	  "a":[
		["9524.6","2"],
		["9524.7","3"],
		["9524.9","16"],
		["9525.1","10"],
		["9525.3","6"]
	  ]
	}`)
	s.mockWsServe(data, expectedErr)
	defer s.assertWsServe(expectedServeCnt)

	// doneC, stopC, err := WsPartialDepthServe("BTCUSD_200626", 5,
	handler := func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:            "depthUpdate",
			Time:             1591269996801,
			TransactionTime:  1591269996646,
			Symbol:           "BTCUSD_200626",
			Pair:             "BTCUSD",
			FirstUpdateID:    17276694,
			LastUpdateID:     17276701,
			PrevLastUpdateID: 17276678,
			Bids: []Bid{
				{Price: "9523.0", Quantity: "5"},
				{Price: "9522.8", Quantity: "8"},
				{Price: "9522.6", Quantity: "2"},
				{Price: "9522.4", Quantity: "1"},
				{Price: "9522.0", Quantity: "5"},
			},
			Asks: []Ask{
				{Price: "9524.6", Quantity: "2"},
				{Price: "9524.7", Quantity: "3"},
				{Price: "9524.9", Quantity: "16"},
				{Price: "9525.1", Quantity: "10"},
				{Price: "9525.3", Quantity: "6"},
			},
		}
		s.assertDepthEvent(e, event)
	}
	errHandler := func(err error) {
	}

	var doneC, stopC chan struct{}
	var err error
	if rate == nil {
		doneC, stopC, err = WsPartialDepthServe("BTCUSD_200626", levels, handler, errHandler)
	} else {
		doneC, stopC, err = WsPartialDepthServeWithRate("BTCUSD_200626", levels, rate, handler, errHandler)
	}

	if expectedErr == nil {
		s.r().NoError(err)
	} else {
		s.r().EqualError(err, expectedErr.Error())
	}

	if stopC != nil {
		stopC <- struct{}{}
	}
	if doneC != nil {
		<-doneC
	}
}

func (s *websocketServiceTestSuite) TestPartialDepthServe() {
	s.testPartialDepthServe(5, nil, nil, 1)
}

func (s *websocketServiceTestSuite) TestPartialDepthServeWithInvalidLevels() {
	s.testPartialDepthServe(8, nil, errors.New("Invalid levels"), 0)
}

func (s *websocketServiceTestSuite) TestPartialDepthServeWithValidRate() {
	rate := 250 * time.Millisecond
	s.testPartialDepthServe(5, &rate, nil, 1)
	rate = 500 * time.Millisecond
	s.testPartialDepthServe(5, &rate, nil, 2)
	rate = 100 * time.Millisecond
	s.testPartialDepthServe(5, &rate, nil, 3)
}

func (s *websocketServiceTestSuite) TestPartialDepthServeWithInvalidRate() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)
	for {
		rate := time.Duration(rand.Intn(100)*10) * time.Millisecond
		switch rate {
		case 250 * time.Millisecond:
		case 500 * time.Millisecond:
		case 100 * time.Millisecond:
		default:
			s.testPartialDepthServe(5, &rate, errors.New("Invalid rate"), 0)
			return
		}
	}
}

// https://binance-docs.github.io/apidocs/delivery/en/#diff-book-depth-streams
func (s *websocketServiceTestSuite) TestDiffDepthServe() {
	data := []byte(`{
	  "e": "depthUpdate",
	  "E": 1591270260907,
	  "T": 1591270260891,
	  "s": "BTCUSD_200626",
	  "ps": "BTCUSD",
	  "U": 17285681,
	  "u": 17285702,
	  "pu": 17285675,
	  "b": [
		["9517.6","10"]
	  ],
	  "a": [
		["9518.5","45"]
	  ]
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsDiffDepthServe("BTCUSD_200626", func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:            "depthUpdate",
			Time:             1591270260907,
			TransactionTime:  1591270260891,
			Symbol:           "BTCUSD_200626",
			Pair:             "BTCUSD",
			FirstUpdateID:    17285681,
			LastUpdateID:     17285702,
			PrevLastUpdateID: 17285675,
			Bids:             []Bid{{Price: "9517.6", Quantity: "10"}},
			Asks:             []Ask{{Price: "9518.5", Quantity: "45"}},
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
	r.Equal(e.Pair, a.Pair, "Pair")
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

func (s *websocketServiceTestSuite) testWsUserDataServe(data []byte, expectedEvent *WsUserDataEvent) {
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsUserDataServe("fakeListenKey", func(event *WsUserDataEvent) {
		s.assertUserDataEvent(expectedEvent, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestWsUserDataServeStreamExpired() {
	data := []byte(`{
		"e": "listenKeyExpired",
		"E": 1576653824250
	}`)
	expectedEvent := &WsUserDataEvent{
		Event: "listenKeyExpired",
		Time:  1576653824250,
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeMarginCall() {
	data := []byte(`{
	  "e":"MARGIN_CALL",
	  "E":1587727187525,
	  "i": "SfsR",
	  "cw":"3.16812045",
	  "p":[{
	    "s":"BTCUSD_200925",
	    "ps":"LONG",
	    "pa":"132",
	    "mt":"CROSSED",
	    "iw":"0",
	    "mp":"9187.17127000",
	    "up":"-1.166074",
	    "mm":"1.614445"
	  }]
    }`)
	expectedEvent := &WsUserDataEvent{
		Event:              "MARGIN_CALL",
		Time:               1587727187525,
		Alias:              "SfsR",
		CrossWalletBalance: "3.16812045",
		MarginCallPositions: []WsPosition{
			{
				Symbol:                    "BTCUSD_200925",
				Side:                      "LONG",
				Amount:                    "132",
				MarginType:                "CROSSED",
				IsolatedWallet:            "0",
				MarkPrice:                 "9187.17127000",
				UnrealizedPnL:             "-1.166074",
				MaintenanceMarginRequired: "1.614445",
			},
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeAccountUpdate() {
	data := []byte(`{
		"e": "ACCOUNT_UPDATE",
		"E": 1564745798939,
		"T": 1564745798938,
		"i": "SfsR",
		"a":
		  {
			"m":"ORDER",
			"B":[
			  {
				"a":"BTC",
				"wb":"122624.12345678",
				"cw":"100.12345678",
				"bc":"50.12345678"
			  },
			  {
				"a":"ETH",           
				"wb":"1.00000000",
				"cw":"0.00000000",
				"bc":"-49.12345678"
			}
			],
			"P":[
			  {
				"s":"BTCUSD_200925",
				"pa":"0",
				"ep":"0.00000",
				"cr":"200",
				"up":"0",
				"mt":"isolated",
				"iw":"0.00000000",
				"ps":"BOTH"
			  },
			  {
				  "s":"BTCUSD_200925",
				  "pa":"20",
				  "ep":"6563.6",
				  "cr":"0",
				  "up":"2850.21200000",
				  "mt":"isolated",
				  "iw":"13200.70726908",
				  "ps":"LONG"
			   },
			  {
				  "s":"BTCUSD_200925",
				  "pa":"-10",
				  "ep":"6563.8",
				  "cr":"-45.04000000",
				  "up":"-1423.15600000",
				  "mt":"isolated",
				  "iw":"6570.42511771",
				  "ps":"SHORT"
			  }
			]
		  }
	}`)
	expectedEvent := &WsUserDataEvent{
		Event:           "ACCOUNT_UPDATE",
		Time:            1564745798939,
		TransactionTime: 1564745798938,
		Alias:           "SfsR",
		AccountUpdate: WsAccountUpdate{
			Reason: "ORDER",
			Balances: []WsBalance{
				{
					Asset:              "BTC",
					Balance:            "122624.12345678",
					CrossWalletBalance: "100.12345678",
					BalanceChange:      "50.12345678",
				},
				{
					Asset:              "ETH",
					Balance:            "1.00000000",
					CrossWalletBalance: "0.00000000",
					BalanceChange:      "-49.12345678",
				},
			},
			Positions: []WsPosition{
				{
					Symbol:              "BTCUSD_200925",
					Amount:              "0",
					EntryPrice:          "0.00000",
					AccumulatedRealized: "200",
					UnrealizedPnL:       "0",
					MarginType:          "isolated",
					IsolatedWallet:      "0.00000000",
					Side:                "BOTH",
				},
				{
					Symbol:              "BTCUSD_200925",
					Amount:              "20",
					EntryPrice:          "6563.6",
					AccumulatedRealized: "0",
					UnrealizedPnL:       "2850.21200000",
					MarginType:          "isolated",
					IsolatedWallet:      "13200.70726908",
					Side:                "LONG",
				},
				{
					Symbol:              "BTCUSD_200925",
					Amount:              "-10",
					EntryPrice:          "6563.8",
					AccumulatedRealized: "-45.04000000",
					UnrealizedPnL:       "-1423.15600000",
					MarginType:          "isolated",
					IsolatedWallet:      "6570.42511771",
					Side:                "SHORT",
				},
			},
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeOrderTradeUpdate() {
	data := []byte(`{
		"e":"ORDER_TRADE_UPDATE",
		"E":1568879465651,
		"T":1568879465650,
		"i": "SfsR",
		"o":{
		  "s":"BTCUSD_200925",
		  "c":"TEST",
		  "S":"SELL",
		  "o":"TRAILING_STOP_MARKET",
		  "f":"GTC",
		  "q":"2",
		  "p":"0",
		  "ap":"0",
		  "sp":"9103.1",
		  "x":"NEW",
		  "X":"NEW",
		  "i":8888888,
		  "l":"0",
		  "z":"0",
		  "L":"0",
		  "ma": "BTC",
		  "N":"BTC",
		  "n":"0",
		  "T":1591274595442,
		  "t":0,
		  "rp": "0",
		  "b":"0",
		  "a":"0",
		  "m":false,
		  "R":false,
		  "wt":"CONTRACT_PRICE",
		  "ot":"TRAILING_STOP_MARKET",
		  "ps":"LONG",
		  "cp":false,
		  "AP":"9476.8",
		  "cr":"5.0",
		  "pP": false
		}
	}`)
	expectedEvent := &WsUserDataEvent{
		Event:           "ORDER_TRADE_UPDATE",
		Time:            1568879465651,
		TransactionTime: 1568879465650,
		Alias:           "SfsR",
		OrderTradeUpdate: WsOrderTradeUpdate{
			Symbol:               "BTCUSD_200925",
			ClientOrderID:        "TEST",
			Side:                 "SELL",
			Type:                 "TRAILING_STOP_MARKET",
			TimeInForce:          "GTC",
			OriginalQty:          "2",
			OriginalPrice:        "0",
			AveragePrice:         "0",
			StopPrice:            "9103.1",
			ExecutionType:        "NEW",
			Status:               "NEW",
			ID:                   8888888,
			LastFilledQty:        "0",
			AccumulatedFilledQty: "0",
			LastFilledPrice:      "0",
			MarginAsset:          "BTC",
			CommissionAsset:      "BTC",
			Commission:           "0",
			TradeTime:            1591274595442,
			TradeID:              0,
			RealizedPnL:          "0",
			BidsNotional:         "0",
			AsksNotional:         "0",
			IsMaker:              false,
			IsReduceOnly:         false,
			WorkingType:          "CONTRACT_PRICE",
			OriginalType:         "TRAILING_STOP_MARKET",
			PositionSide:         "LONG",
			IsClosingPosition:    false,
			ActivationPrice:      "9476.8",
			CallbackRate:         "5.0",
			IsProtected:          false,
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) assertUserDataEvent(e, a *WsUserDataEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.CrossWalletBalance, a.CrossWalletBalance, "CrossWalletBalance")
	for i, e := range e.MarginCallPositions {
		a := a.MarginCallPositions[i]
		s.assertPosition(e, a)
	}
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	s.assertAccountUpdate(e.AccountUpdate, a.AccountUpdate)
	s.assertOrderTradeUpdate(e.OrderTradeUpdate, a.OrderTradeUpdate)
}

func (s *websocketServiceTestSuite) assertPosition(e, a WsPosition) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.MarginType, a.MarginType, "MarginType")
	r.Equal(e.IsolatedWallet, a.IsolatedWallet, "IsolatedWallet")
	r.Equal(e.EntryPrice, a.EntryPrice, "EntryPrice")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.UnrealizedPnL, a.UnrealizedPnL, "UnrealizedPnL")
	r.Equal(e.AccumulatedRealized, a.AccumulatedRealized, "AccumulatedRealized")
	r.Equal(e.MaintenanceMarginRequired, a.MaintenanceMarginRequired, "MaintenanceMarginRequired")
}

func (s *websocketServiceTestSuite) assertAccountUpdate(e, a WsAccountUpdate) {
	r := s.r()
	r.Equal(e.Reason, a.Reason, "Reason")
	for i, e := range e.Balances {
		a := a.Balances[i]
		r.Equal(e.Asset, a.Asset, "Asset")
		r.Equal(e.Balance, a.Balance, "Balance")
		r.Equal(e.CrossWalletBalance, a.CrossWalletBalance, "CrossWalletBalance")
	}
	for i, e := range e.Positions {
		a := a.Positions[i]
		s.assertPosition(e, a)
	}
}

func (s *websocketServiceTestSuite) assertOrderTradeUpdate(e, a WsOrderTradeUpdate) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.OriginalQty, a.OriginalQty, "OriginalQty")
	r.Equal(e.OriginalPrice, a.OriginalPrice, "OriginalPrice")
	r.Equal(e.AveragePrice, a.AveragePrice, "AveragePrice")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.ExecutionType, a.ExecutionType, "ExecutionType")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.LastFilledQty, a.LastFilledQty, "LastFilledQty")
	r.Equal(e.AccumulatedFilledQty, a.AccumulatedFilledQty, "AccumulatedFilledQty")
	r.Equal(e.LastFilledPrice, a.LastFilledPrice, "LastFilledPrice")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.TradeID, a.TradeID, "TradeID")
	r.Equal(e.BidsNotional, a.BidsNotional, "BidsNotional")
	r.Equal(e.AsksNotional, a.AsksNotional, "AsksNotional")
	r.Equal(e.IsMaker, a.IsMaker, "IsMaker")
	r.Equal(e.IsReduceOnly, a.IsReduceOnly, "IsReduceOnly")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.OriginalType, a.OriginalType, "OriginalType")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.IsClosingPosition, a.IsClosingPosition, "IsClosingPosition")
	r.Equal(e.ActivationPrice, a.ActivationPrice, "ActivationPrice")
	r.Equal(e.CallbackRate, a.CallbackRate, "CallbackRate")
	r.Equal(e.RealizedPnL, a.RealizedPnL, "RealizedPnL")
}
