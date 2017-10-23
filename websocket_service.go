package binance

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	baseURL = "wss://stream.binance.com:9443/ws"
)

type WsDepthHandler func(event *WsDepthEvent)

func WsDepthServe(symbol string, handler WsDepthHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s@depth", baseURL, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			return
		}
		event := new(WsDepthEvent)
		event.Event = j.Get("e").MustString()
		event.Time = j.Get("E").MustInt64()
		event.Symbol = j.Get("s").MustString()
		event.UpdateID = j.Get("u").MustInt64()
		bidsLen := len(j.Get("b").MustArray())
		event.Bids = make([]Bid, 0)
		for i := 0; i < bidsLen; i++ {
			item := j.Get("b").GetIndex(i)
			event.Bids[i] = Bid{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		asksLen := len(j.Get("a").MustArray())
		event.Asks = make([]Ask, 0)
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

type WsDepthEvent struct {
	Event    string `json:"e"`
	Time     int64  `json:"E"`
	Symbol   string `json:"s"`
	UpdateID int64  `json:"u"`
	Bids     []Bid  `json:"b"`
	Asks     []Ask  `json:"a"`
}

type WsKlineHandler func(event *WsKlineEvent)

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

type WsKlineEvent struct {
	Event  string `json:"e"`
	Time   int64  `json:"E"`
	Symbol string `json:"s"`
	Kline  struct {
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
	} `json:"k"`
}

type WsAggTradeHandler func(event *WsAggTradeEvent)

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

type WsAggTradeEvent struct {
	Event                 string `json:"e"`
	Time                  int64  `json:"E"`
	Symbol                string `json:"s"`
	AggregateTradeID      int64  `json:"a"`
	Price                 string `json:"p"`
	Quantity              string `json:"q"`
	FirstBreakdownTradeID int64  `json:"f"`
	LastBreakdownTradeID  int64  `json:"l"`
	TradeTime             int64  `json:"T"`
	IsBuyerMaker          bool   `json:"m"`
}

func WsUserDataServe(listenKey string, handler WsHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s", baseURL, listenKey)
	cfg := newWsConfig(endpoint)
	return wsServe(cfg, handler)
}
