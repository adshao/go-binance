package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type websocketServiceTestSuite struct {
	baseTestSuite
	origWsServe wsServeFunc
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

func (s *websocketServiceTestSuite) mockWsServe(data []byte) {
	wsServe = func(cfg *wsConfig, handler WsHandler) (done chan struct{}, err error) {
		s.serveCount++
		done = make(chan struct{})
		defer close(done)
		handler(data)
		return
	}
}

func (s *websocketServiceTestSuite) assertWsServe(count ...int) {
	e := 1
	if len(count) > 0 {
		e = count[0]
	}
	s.r().Equal(e, s.serveCount)
}

func (s *websocketServiceTestSuite) TestDepthServe() {
	data := []byte(`{
        "e": "depthUpdate",
        "E": 1499404630606,
        "s": "ETHBTC",
        "u": 7913455,
        "b": [
            [
                "0.10376590",
                "59.15767010",
                []
            ]
        ],
        "a": [
            [
                "0.10376586",
                "159.15767010",
                []
            ],
            [
                "0.10383109",
                "345.86845230",
                []
            ],
            [
                "0.10490700",
                "0.00000000",
                []
            ]
        ]
    }`)
	s.mockWsServe(data)
	defer s.assertWsServe()
	_, err := WsDepthServe("ETHBTC", func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:    "depthUpdate",
			Time:     1499404630606,
			Symbol:   "ETHBTC",
			UpdateID: 7913455,
			Bids: []Bid{
				{
					Price:    "0.10376590",
					Quantity: "59.15767010",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.10376586",
					Quantity: "159.15767010",
				},
				{
					Price:    "0.10383109",
					Quantity: "345.86845230",
				},
				{
					Price:    "0.10490700",
					Quantity: "0.00000000",
				},
			},
		}
		s.assertWsDepthEventEqual(e, event)
	})
	s.r().NoError(err)
}

func (s *websocketServiceTestSuite) assertWsDepthEventEqual(e, a *WsDepthEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.UpdateID, a.UpdateID, "UpdateID")
	for i := 0; i < len(e.Bids); i++ {
		r.Equal(e.Bids[i].Price, a.Bids[i].Price, "Price")
		r.Equal(e.Bids[i].Quantity, a.Bids[i].Quantity, "Quantity")
	}
	for i := 0; i < len(e.Asks); i++ {
		r.Equal(e.Asks[i].Price, a.Asks[i].Price, "Price")
		r.Equal(e.Asks[i].Quantity, a.Asks[i].Quantity, "Quantity")
	}
}

func (s *websocketServiceTestSuite) TestKlineServe() {
	data := []byte(`{
        "e": "kline",
        "E": 1499404907056,
        "s": "ETHBTC",
        "k": {
            "t": 1499404860000,
            "T": 1499404919999,
            "s": "ETHBTC",
            "i": "1m",
            "f": 77462,
            "L": 77465,
            "o": "0.10278577",
            "c": "0.10278645",
            "h": "0.10278712",
            "l": "0.10278518",
            "v": "17.47929838",
            "n": 4,
            "x": false,
            "q": "1.79662878",
            "V": "2.34879839",
            "Q": "0.24142166",
            "B": "13279784.01349473"
        }
    }`)
	s.mockWsServe(data)
	defer s.assertWsServe()

	_, err := WsKlineServe("ETHBTC", "1m", func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   1499404907056,
			Symbol: "ETHBTC",
			Kline: WsKline{
				StartTime:            1499404860000,
				EndTime:              1499404919999,
				Symbol:               "ETHBTC",
				Interval:             "1m",
				FirstTradeID:         77462,
				LastTradeID:          77465,
				Open:                 "0.10278577",
				Close:                "0.10278645",
				High:                 "0.10278712",
				Low:                  "0.10278518",
				Volume:               "17.47929838",
				TradeNum:             4,
				IsFinal:              false,
				QuoteVolume:          "1.79662878",
				ActiveBuyVolume:      "2.34879839",
				ActiveBuyQuoteVolume: "0.24142166",
			},
		}
		s.assertWsKlineEventEqual(e, event)
	})
	s.r().NoError(err)
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

func (s *websocketServiceTestSuite) TestWsAggTradeServe() {
	data := []byte(`{
        "e": "aggTrade",
        "E": 1499405254326,
        "s": "ETHBTC",
        "a": 70232,
        "p": "0.10281118",
        "q": "8.15632997",
        "f": 77489,
        "l": 77489,
        "T": 1499405254324,
        "m": false,
        "M": true
    }`)
	s.mockWsServe(data)
	defer s.assertWsServe()

	_, err := WsAggTradeServe("ETHBTC", func(event *WsAggTradeEvent) {
		e := &WsAggTradeEvent{
			Event:                 "aggTrade",
			Time:                  1499405254326,
			Symbol:                "ETHBTC",
			AggTradeID:            70232,
			Price:                 "0.10281118",
			Quantity:              "8.15632997",
			FirstBreakdownTradeID: 77489,
			LastBreakdownTradeID:  77489,
			TradeTime:             1499405254324,
			IsBuyerMaker:          false,
		}
		s.assertWsAggTradeServeEqual(e, event)
	})
	s.r().NoError(err)
}

func (s *websocketServiceTestSuite) assertWsAggTradeServeEqual(e, a *WsAggTradeEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.AggTradeID, a.AggTradeID, "AggTradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.FirstBreakdownTradeID, a.FirstBreakdownTradeID, "FirstBreakdownTradeID")
	r.Equal(e.LastBreakdownTradeID, a.LastBreakdownTradeID, "LastBreakdownTradeID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
}

func (s *websocketServiceTestSuite) testWsUserDataServe(data []byte) {
	s.mockWsServe(data)
	defer s.assertWsServe()

	_, err := WsUserDataServe("listenKey", func(event []byte) {
		s.r().Equal(data, event)
	})
	s.r().NoError(err)
}

func (s *websocketServiceTestSuite) TestWsUserDataServe() {
	s.testWsUserDataServe([]byte(`{
        "e": "outboundAccountInfo",
        "E": 1499405658849,
        "m": 0,
        "t": 0,
        "b": 0,
        "s": 0,
        "T": true,
        "W": true,
        "D": true,
        "B": [
            {
                "a": "LTC",
                "f": "17366.18538083",
                "l": "0.00000000"
            },
            {
                "a": "BTC",
                "f": "10537.85314051",
                "l": "2.19464093"
            },
            {
                "a": "ETH",
                "f": "17902.35190619",
                "l": "0.00000000"
            },
            {
                "a": "BNC",
                "f": "1114503.29769312",
                "l": "0.00000000"
            },
            {
                "a": "NEO",
                "f": "0.00000000",
                "l": "0.00000000"
            }
        ]
    }`))
}
