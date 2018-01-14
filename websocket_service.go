package binance

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	baseURL = "wss://stream.binance.com:9443/ws"
)

// WsDepthHandler handle websocket depth event
type WsDepthHandler func(event *WsDepthEvent)

// WsDepthServe serve websocket depth handler with a symbol
func WsDepthServe(symbol string, handler WsDepthHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s@depth", baseURL, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			// TODO: callback if there is an error
			return
		}
		event := new(WsDepthEvent)
		event.Event = j.Get("e").MustString()
		event.Time = j.Get("E").MustInt64()
		event.Symbol = j.Get("s").MustString()
		event.UpdateID = j.Get("u").MustInt64()
		bidsLen := len(j.Get("b").MustArray())
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := j.Get("b").GetIndex(i)
			event.Bids[i] = Bid{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		asksLen := len(j.Get("a").MustArray())
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {
			item := j.Get("a").GetIndex(i)
			event.Asks[i] = Ask{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler)
}

// WsDepthEvent define websocket depth event
type WsDepthEvent struct {
	Event    string `json:"e"`
	Time     int64  `json:"E"`
	Symbol   string `json:"s"`
	UpdateID int64  `json:"u"`
	Bids     []Bid  `json:"b"`
	Asks     []Ask  `json:"a"`
}

// WsKlineHandler handle websocket kline event
type WsKlineHandler func(event *WsKlineEvent)

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func WsKlineServe(symbol string, interval string, handler WsKlineHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s@kline_%s", baseURL, strings.ToLower(symbol), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler)
}

// WsKlineEvent define websocket kline event
type WsKlineEvent struct {
	Event  string  `json:"e"`
	Time   int64   `json:"E"`
	Symbol string  `json:"s"`
	Kline  WsKline `json:"k"`
}

// WsKline define websocket kline
type WsKline struct {
	StartTime            int64  `json:"t"`
	EndTime              int64  `json:"T"`
	Symbol               string `json:"s"`
	Interval             string `json:"i"`
	FirstTradeID         int64  `json:"f"`
	LastTradeID          int64  `json:"L"`
	Open                 string `json:"o"`
	Close                string `json:"c"`
	High                 string `json:"h"`
	Low                  string `json:"l"`
	Volume               string `json:"v"`
	TradeNum             int64  `json:"n"`
	IsFinal              bool   `json:"x"`
	QuoteVolume          string `json:"q"`
	ActiveBuyVolume      string `json:"V"`
	ActiveBuyQuoteVolume string `json:"Q"`
}

// WsAggTradeHandler handle websocket aggregate trade event
type WsAggTradeHandler func(event *WsAggTradeEvent)

// WsAggTradeServe serve websocket aggregate handler with a symbol
func WsAggTradeServe(symbol string, handler WsAggTradeHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", baseURL, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsAggTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler)
}

// WsAggTradeEvent define websocket aggregate trade event
type WsAggTradeEvent struct {
	Event                 string `json:"e"`
	Time                  int64  `json:"E"`
	Symbol                string `json:"s"`
	AggTradeID            int64  `json:"a"`
	Price                 string `json:"p"`
	Quantity              string `json:"q"`
	FirstBreakdownTradeID int64  `json:"f"`
	LastBreakdownTradeID  int64  `json:"l"`
	TradeTime             int64  `json:"T"`
	IsBuyerMaker          bool   `json:"m"`
	Placeholder           bool   `json:"M"` // add this field to avoid case insensitive unmarshaling
}

// WsUserDataServe serve user data handler with listen key
func WsUserDataServe(listenKey string, handler WsHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s", baseURL, listenKey)
	cfg := newWsConfig(endpoint)
	return wsServe(cfg, handler)
}

// WsMarketStatHandler handle websocket that push single market statistics for 24hr
type WsMarketStatHandler func(event *WsMarketStatEvent)

// WsMarketStatServe serve websocket that push 24hr statistics for single market every second
func WsMarketStatServe(symbol string, handler WsMarketStatHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s@ticker", baseURL, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsMarketStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			return
		}
		handler(&event)
	}
	return wsServe(cfg, wsHandler)
}

// WsAllMarketsStatHandler handle websocket that push all markets statistics for 24hr
type WsAllMarketsStatHandler func(event WsAllMarketsStatEvent)

// WsAllMarketsStatServe serve websocket that push 24hr statistics for all market every second
func WsAllMarketsStatServe(handler WsAllMarketsStatHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/!ticker@arr", baseURL)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMarketsStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler)
}

// WsAllMarketsStatEvent define array of websocket market statistics events
type WsAllMarketsStatEvent []*WsMarketStatEvent

// WsMarketStatEvent define websocket market statistics event
type WsMarketStatEvent struct {
	Event              string `json:"e"`
	Time               int64  `json:"E"`
	Symbol             string `json:"s"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	PrevClosePrice     string `json:"x"`
	LastPrice          string `json:"c"`
	CloseQty           string `json:"Q"`
	BidPrice           string `json:"b"`
	BidQty             string `json:"B"`
	AskPrice           string `json:"a"`
	AskQty             string `json:"A"`
	OpenPrice          string `json:"o"`
	HighPrice          string `json:"h"`
	LowPrice           string `json:"l"`
	BaseVolume         string `json:"v"`
	QuoteVolume        string `json:"q"`
	OpenTime           int64  `json:"O"`
	CloseTime          int64  `json:"C"`
	FirstID            int64  `json:"F"`
	LastID             int64  `json:"L"`
	Count              int64  `json:"n"`
}
