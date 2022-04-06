package binance

import (
	"encoding/json"
	"fmt"
	"strings"
)

type WsService interface {
	WsPartialDepthServe(symbol string, levels string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsPartialDepthServe100Ms(symbol string, levels string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsCombinedPartialDepthServe(symbolLevels map[string]string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsDepthServe(symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsDepthServe100Ms(symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsCombinedDepthServe(symbols []string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsCombinedDepthServe100Ms(symbols []string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsCombinedKlineServe(symbolIntervalPair map[string]string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsKlineServe(symbol string, interval string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsAggTradeServe(symbol string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsCombinedAggTradeServe(symbols []string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsTradeServe(symbol string, handler WsTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsUserDataServe(listenKey string, handler WsUserDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsCombinedMarketStatServe(symbols []string, handler WsMarketStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsMarketStatServe(symbol string, handler WsMarketStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsAllMarketsStatServe(handler WsAllMarketsStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsAllMiniMarketsStatServe(handler WsAllMiniMarketsStatServeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsBookTickerServe(symbol string, handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
	WsAllBookTickerServe(handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error)
}

type wsService struct {
	baseUrl         string
	combinedBaseUrl string
}

func NewWsService(baseUrl, combinedBaseUrl string) WsService {
	return &wsService{
		baseUrl:         baseUrl,
		combinedBaseUrl: combinedBaseUrl,
	}
}

// WsPartialDepthServe serve websocket partial depth handler with a symbol, using 1sec updates
func (s wsService) WsPartialDepthServe(symbol string, levels string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth%s", s.baseUrl, strings.ToLower(symbol), levels)
	return wsPartialDepthServe(endpoint, symbol, handler, errHandler)
}

// WsPartialDepthServe100Ms serve websocket partial depth handler with a symbol, using 100msec updates
func (s wsService) WsPartialDepthServe100Ms(symbol string, levels string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth%s@100ms", s.baseUrl, strings.ToLower(symbol), levels)
	return wsPartialDepthServe(endpoint, symbol, handler, errHandler)
}

// WsPartialDepthServe serve websocket partial depth handler with a symbol
func (s wsService) wsPartialDepthServe(endpoint string, symbol string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsPartialDepthEvent)
		event.Symbol = symbol
		event.LastUpdateID = j.Get("lastUpdateId").MustInt64()
		bidsLen := len(j.Get("bids").MustArray())
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := j.Get("bids").GetIndex(i)
			event.Bids[i] = Bid{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		asksLen := len(j.Get("asks").MustArray())
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {
			item := j.Get("asks").GetIndex(i)
			event.Asks[i] = Ask{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCombinedPartialDepthServe is similar to WsPartialDepthServe, but it for multiple symbols
func (s wsService) WsCombinedPartialDepthServe(symbolLevels map[string]string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := s.combinedBaseUrl
	for s, l := range symbolLevels {
		endpoint += fmt.Sprintf("%s@depth%s", strings.ToLower(s), l) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsPartialDepthEvent)
		stream := j.Get("stream").MustString()
		symbol := strings.Split(stream, "@")[0]
		event.Symbol = strings.ToUpper(symbol)
		data := j.Get("data").MustMap()
		event.LastUpdateID, _ = data["lastUpdateId"].(json.Number).Int64()
		bidsLen := len(data["bids"].([]interface{}))
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := data["bids"].([]interface{})[i].([]interface{})
			event.Bids[i] = Bid{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		asksLen := len(data["asks"].([]interface{}))
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {

			item := data["asks"].([]interface{})[i].([]interface{})
			event.Asks[i] = Ask{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsDepthServe serve websocket depth handler with a symbol, using 1sec updates
func (s wsService) WsDepthServe(symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth", s.baseUrl, strings.ToLower(symbol))
	return wsDepthServe(endpoint, handler, errHandler)
}

// WsDepthServe100Ms serve websocket depth handler with a symbol, using 100msec updates
func (s wsService) WsDepthServe100Ms(symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth@100ms", s.baseUrl, strings.ToLower(symbol))
	return wsDepthServe(endpoint, handler, errHandler)
}

// WsDepthServe serve websocket depth handler with an arbitrary endpoint address
func (s wsService) wsDepthServe(endpoint string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsDepthEvent)
		event.Event = j.Get("e").MustString()
		event.Time = j.Get("E").MustInt64()
		event.Symbol = j.Get("s").MustString()
		event.LastUpdateID = j.Get("u").MustInt64()
		event.FirstUpdateID = j.Get("U").MustInt64()
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
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCombinedDepthServe is similar to WsDepthServe, but it for multiple symbols
func (s wsService) WsCombinedDepthServe(symbols []string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := s.combinedBaseUrl
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@depth", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return wsCombinedDepthServe(endpoint, handler, errHandler)
}

func (s wsService) WsCombinedDepthServe100Ms(symbols []string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := s.combinedBaseUrl
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@depth@100ms", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return wsCombinedDepthServe(endpoint, handler, errHandler)
}

func (s wsService) wsCombinedDepthServe(endpoint string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsDepthEvent)
		stream := j.Get("stream").MustString()
		symbol := strings.Split(stream, "@")[0]
		event.Symbol = strings.ToUpper(symbol)
		data := j.Get("data").MustMap()
		event.Time, _ = data["E"].(json.Number).Int64()
		event.LastUpdateID, _ = data["u"].(json.Number).Int64()
		event.FirstUpdateID, _ = data["U"].(json.Number).Int64()
		bidsLen := len(data["b"].([]interface{}))
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := data["b"].([]interface{})[i].([]interface{})
			event.Bids[i] = Bid{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		asksLen := len(data["a"].([]interface{}))
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {

			item := data["a"].([]interface{})[i].([]interface{})
			event.Asks[i] = Ask{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCombinedKlineServe is similar to WsKlineServe, but it handles multiple symbols with it interval
func (s wsService) WsCombinedKlineServe(symbolIntervalPair map[string]string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := s.combinedBaseUrl
	for symbol, interval := range symbolIntervalPair {
		endpoint += fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		data := j.Get("data").MustMap()

		symbol := strings.Split(stream, "@")[0]

		jsonData, _ := json.Marshal(data)

		event := new(WsKlineEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}
		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func (s wsService) WsKlineServe(symbol string, interval string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@kline_%s", s.baseUrl, strings.ToLower(symbol), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAggTradeServe serve websocket aggregate handler with a symbol
func (s wsService) WsAggTradeServe(symbol string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", s.baseUrl, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsAggTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCombinedAggTradeServe is similar to WsAggTradeServe, but it handles multiple symbolx
func (s wsService) WsCombinedAggTradeServe(symbols []string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := s.combinedBaseUrl
	for s := range symbols {
		endpoint += fmt.Sprintf("%s@aggTrade", strings.ToLower(symbols[s])) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		data := j.Get("data").MustMap()

		symbol := strings.Split(stream, "@")[0]

		jsonData, _ := json.Marshal(data)

		event := new(WsAggTradeEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}

		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsTradeServe serve websocket handler with a symbol
func (s wsService) WsTradeServe(symbol string, handler WsTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@trade", s.baseUrl, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsUserDataServe serve user data handler with listen key
func (s wsService) WsUserDataServe(listenKey string, handler WsUserDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s", s.baseUrl, listenKey)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		event := new(WsUserDataEvent)

		err = json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}

		switch UserDataEventType(j.Get("e").MustString()) {
		case UserDataEventTypeOutboundAccountPosition:

		case UserDataEventTypeBalanceUpdate:
			err = json.Unmarshal(message, &event.BalanceUpdate)
			if err != nil {
				errHandler(err)
				return
			}
		case UserDataEventTypeExecutionReport:
			err = json.Unmarshal(message, &event.OrderUpdate)
			if err != nil {
				errHandler(err)
				return
			}
			// Unmarshal has case sensitive problem
			event.TransactionTime = j.Get("T").MustInt64()
			event.OrderUpdate.TransactionTime = j.Get("T").MustInt64()
			event.OrderUpdate.Id = j.Get("i").MustInt64()
			event.OrderUpdate.TradeId = j.Get("t").MustInt64()
			event.OrderUpdate.FeeAsset = j.Get("N").MustString()
		case UserDataEventTypeListStatus:
			err = json.Unmarshal(message, &event.OCOUpdate)
			if err != nil {
				errHandler(err)
				return
			}
		}

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCombinedMarketStatServe is similar to WsMarketStatServe, but it handles multiple symbolx
func (s wsService) WsCombinedMarketStatServe(symbols []string, handler WsMarketStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := s.combinedBaseUrl
	for s := range symbols {
		endpoint += fmt.Sprintf("%s@ticker", strings.ToLower(symbols[s])) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)

	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		data := j.Get("data").MustMap()

		symbol := strings.Split(stream, "@")[0]

		jsonData, _ := json.Marshal(data)

		event := new(WsMarketStatEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}

		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsMarketStatServe serve websocket that push 24hr statistics for single market every second
func (s wsService) WsMarketStatServe(symbol string, handler WsMarketStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@ticker", s.baseUrl, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsMarketStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(&event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMarketsStatServe serve websocket that push 24hr statistics for all market every second
func (s wsService) WsAllMarketsStatServe(handler WsAllMarketsStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!ticker@arr", s.baseUrl)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMarketsStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMiniMarketsStatServe serve websocket that push mini version of 24hr statistics for all market every second
func (s wsService) WsAllMiniMarketsStatServe(handler WsAllMiniMarketsStatServeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!miniTicker@arr", s.baseUrl)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMiniMarketsStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
func (s wsService) WsBookTickerServe(symbol string, handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@bookTicker", s.baseUrl, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBookTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for all symbols.
func (s wsService) WsAllBookTickerServe(handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!bookTicker", s.baseUrl)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBookTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}
