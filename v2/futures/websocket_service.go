package futures

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Endpoints
const (
	baseWsMainUrl    = "wss://fstream.binance.com/ws"
	baseWsTestnetUrl = "wss://stream.binancefuture.com/ws"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
	// UseTestnet switch all the WS streams from production to the testnet
	UseTestnet = false
)

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint() string {
	if UseTestnet {
		return baseWsTestnetUrl
	}
	return baseWsMainUrl
}

// WsUserDataServe serve user data handler with listen key
// WsAggTradeEvent define websocket aggTrde event.
type WsAggTradeEvent struct {
	Event            string `json:"e"`
	Time             int64  `json:"E"`
	Symbol           string `json:"s"`
	AggregateTradeID int64  `json:"a"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int64  `json:"f"`
	LastTradeID      int64  `json:"l"`
	TradeTime        int64  `json:"T"`
	Maker            bool   `json:"m"`
}

// WsAggTradeHandler handle websocket that push trade information that is aggregated for a single taker order.
type WsAggTradeHandler func(event *WsAggTradeEvent)

// WsAggTradeServe serve websocket that push trade information that is aggregated for a single taker order.
func WsAggTradeServe(symbol string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsAggTradeEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsMarkPriceEvent define websocket markPriceUpdate event.
type WsMarkPriceEvent struct {
	Event           string `json:"e"`
	Time            int64  `json:"E"`
	Symbol          string `json:"s"`
	MarkPrice       string `json:"p"`
	IndexPrice      string `json:"i"`
	FundingRate     string `json:"r"`
	NextFundingTime int64  `json:"T"`
}

// WsMarkPriceHandler handle websocket that pushes price and funding rate for a single symbol.
type WsMarkPriceHandler func(event *WsMarkPriceEvent)

// WsMarkPriceServe serve websocket that pushes price and funding rate for a single symbol.
func WsMarkPriceServe(symbol string, handler WsMarkPriceHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@markPrice", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMarkPriceEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMarkPriceEvent defines an array of websocket markPriceUpdate events.
type WsAllMarkPriceEvent []*WsMarkPriceEvent

// WsAllMarkPriceHandler handle websocket that pushes price and funding rate for all symbol.
type WsAllMarkPriceHandler func(event WsAllMarkPriceEvent)

// WsAllMarkPriceServe serve websocket that pushes price and funding rate for all symbol.
func WsAllMarkPriceServe(handler WsAllMarkPriceHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!markPrice@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMarkPriceEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
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

// WsKlineHandler handle websocket kline event
type WsKlineHandler func(event *WsKlineEvent)

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func WsKlineServe(symbol string, interval string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@kline_%s", getWsEndpoint(), strings.ToLower(symbol), interval)
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

// WsMiniMarketTickerEvent define websocket mini market ticker event.
type WsMiniMarketTickerEvent struct {
	Event       string `json:"e"`
	Time        int64  `json:"E"`
	Symbol      string `json:"s"`
	ClosePrice  string `json:"c"`
	OpenPrice   string `json:"o"`
	HighPrice   string `json:"h"`
	LowPrice    string `json:"l"`
	Volume      string `json:"v"`
	QuoteVolume string `json:"q"`
}

// WsMiniMarketTickerHandler handle websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
type WsMiniMarketTickerHandler func(event *WsMiniMarketTickerEvent)

// WsMiniMarketTickerServe serve websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
func WsMiniMarketTickerServe(symbol string, handler WsMiniMarketTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@miniTicker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMiniMarketTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMiniMarketTickerEvent define an array of websocket mini market ticker events.
type WsAllMiniMarketTickerEvent []*WsMiniMarketTickerEvent

// WsAllMiniMarketTickerHandler handle websocket that pushes price and funding rate for all markets.
type WsAllMiniMarketTickerHandler func(event WsAllMiniMarketTickerEvent)

// WsAllMiniMarketTickerServe serve websocket that pushes price and funding rate for all markets.
func WsAllMiniMarketTickerServe(handler WsAllMiniMarketTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!miniTicker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMiniMarketTickerEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsMarketTickerEvent define websocket market ticker event.
type WsMarketTickerEvent struct {
	Event              string `json:"e"`
	Time               int64  `json:"E"`
	Symbol             string `json:"s"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	ClosePrice         string `json:"c"`
	CloseQty           string `json:"Q"`
	OpenPrice          string `json:"o"`
	HighPrice          string `json:"h"`
	LowPrice           string `json:"l"`
	BaseVolume         string `json:"v"`
	QuoteVolume        string `json:"q"`
	OpenTime           int64  `json:"O"`
	CloseTime          int64  `json:"C"`
	FirstID            int64  `json:"F"`
	LastID             int64  `json:"L"`
	TradeCount         int64  `json:"n"`
}

// WsMarketTickerHandler handle websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
type WsMarketTickerHandler func(event *WsMarketTickerEvent)

// WsMarketTickerServe serve websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
func WsMarketTickerServe(symbol string, handler WsMarketTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@ticker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMarketTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMarketTickerEvent define an array of websocket mini ticker events.
type WsAllMarketTickerEvent []*WsMarketTickerEvent

// WsAllMarketTickerHandler handle websocket that pushes price and funding rate for all markets.
type WsAllMarketTickerHandler func(event WsAllMarketTickerEvent)

// WsAllMarketTickerServe serve websocket that pushes price and funding rate for all markets.
func WsAllMarketTickerServe(handler WsAllMarketTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!ticker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMarketTickerEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsBookTickerEvent define websocket best book ticker event.
type WsBookTickerEvent struct {
	Event           string `json:"e"`
	UpdateID        int64  `json:"u"`
	Time            int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	Symbol          string `json:"s"`
	BestBidPrice    string `json:"b"`
	BestBidQty      string `json:"B"`
	BestAskPrice    string `json:"a"`
	BestAskQty      string `json:"A"`
}

// WsBookTickerHandler handle websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
type WsBookTickerHandler func(event *WsBookTickerEvent)

// WsBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
func WsBookTickerServe(symbol string, handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@bookTicker", getWsEndpoint(), strings.ToLower(symbol))
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
func WsAllBookTickerServe(handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!bookTicker", getWsEndpoint())
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

// WsLiquidationOrderEvent define websocket liquidation order event.
type WsLiquidationOrderEvent struct {
	Event            string             `json:"e"`
	Time             int64              `json:"E"`
	LiquidationOrder WsLiquidationOrder `json:"o"`
}

// WsLiquidationOrder define websocket liquidation order.
type WsLiquidationOrder struct {
	Symbol               string          `json:"s"`
	Side                 SideType        `json:"S"`
	OrderType            OrderType       `json:"o"`
	TimeInForce          TimeInForceType `json:"f"`
	OrigQuantity         string          `json:"q"`
	Price                string          `json:"p"`
	AvgPrice             string          `json:"ap"`
	OrderStatus          OrderStatusType `json:"X"`
	LastFilledQty        string          `json:"l"`
	AccumulatedFilledQty string          `json:"z"`
	TradeTime            int64           `json:"T"`
}

// WsLiquidationOrderHandler handle websocket that pushes force liquidation order information for specific symbol.
type WsLiquidationOrderHandler func(event *WsLiquidationOrderEvent)

// WsLiquidationOrderServe serve websocket that pushes force liquidation order information for specific symbol.
func WsLiquidationOrderServe(symbol string, handler WsLiquidationOrderHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@forceOrder", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsLiquidationOrderEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllLiquidationOrderServe serve websocket that pushes force liquidation order information for all symbols.
func WsAllLiquidationOrderServe(handler WsLiquidationOrderHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!forceOrder@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsLiquidationOrderEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsDepthEvent define websocket depth book event
type WsDepthEvent struct {
	Event            string `json:"e"`
	Time             int64  `json:"E"`
	TransactionTime  int64  `json:"T"`
	Symbol           string `json:"s"`
	FirstUpdateID    int64  `json:"U"`
	LastUpdateID     int64  `json:"u"`
	PrevLastUpdateID int64  `json:"pu"`
	Bids             []Bid  `json:"b"`
	Asks             []Ask  `json:"a"`
}

// WsDepthHandler handle websocket depth event
type WsDepthHandler func(event *WsDepthEvent)

// WsPartialDepthServe serve websocket partial depth handler.
func WsPartialDepthServe(symbol string, levels int, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth%d", getWsEndpoint(), strings.ToLower(symbol), levels)
	cfg := newWsConfig(endpoint)
	return wsDepthServe(cfg, handler, errHandler)
}

// WsDiffDepthServe serve websocket diff. depth handler.
func WsDiffDepthServe(symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	return wsDepthServe(cfg, handler, errHandler)
}

func wsDepthServe(cfg *WsConfig, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsDepthEvent)
		event.Event = j.Get("e").MustString()
		event.Time = j.Get("E").MustInt64()
		event.TransactionTime = j.Get("T").MustInt64()
		event.Symbol = j.Get("s").MustString()
		event.FirstUpdateID = j.Get("U").MustInt64()
		event.LastUpdateID = j.Get("u").MustInt64()
		event.PrevLastUpdateID = j.Get("pu").MustInt64()
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

// WsBLVTInfoEvent define websocket BLVT info event
type WsBLVTInfoEvent struct {
	Event          string         `json:"e"`
	Time           int64          `json:"E"`
	Symbol         string         `json:"s"`
	Issued         float64        `json:"m"`
	Baskets        []WsBLVTBasket `json:"b"`
	Nav            float64        `json:"n"`
	Leverage       float64        `json:"l"`
	TargetLeverage int64          `json:"t"`
	FundingRate    float64        `json:"f"`
}

// WsBLVTBasket define websocket BLVT basket
type WsBLVTBasket struct {
	Symbol   string `json:"s"`
	Position int64  `json:"n"`
}

// WsBLVTInfoHandler handle websocket BLVT event
type WsBLVTInfoHandler func(event *WsBLVTInfoEvent)

// WsBLVTInfoServe serve BLVT info stream
func WsBLVTInfoServe(name string, handler WsBLVTInfoHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@tokenNav", getWsEndpoint(), strings.ToUpper(name))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBLVTInfoEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsBLVTKlineEvent define BLVT kline event
type WsBLVTKlineEvent struct {
	Event  string      `json:"e"`
	Time   int64       `json:"E"`
	Symbol string      `json:"s"`
	Kline  WsBLVTKline `json:"k"`
}

// WsBLVTKline BLVT kline
type WsBLVTKline struct {
	StartTime       int64  `json:"t"`
	CloseTime       int64  `json:"T"`
	Symbol          string `json:"s"`
	Interval        string `json:"i"`
	FirstUpdateTime int64  `json:"f"`
	LastUpdateTime  int64  `json:"L"`
	OpenPrice       string `json:"o"`
	ClosePrice      string `json:"c"`
	HighPrice       string `json:"h"`
	LowPrice        string `json:"l"`
	Leverage        string `json:"v"`
	Count           int64  `json:"n"`
}

// WsBLVTKlineHandler BLVT kline handler
type WsBLVTKlineHandler func(event *WsBLVTKlineEvent)

// WsBLVTKlineServe serve BLVT kline stream
func WsBLVTKlineServe(name string, interval string, handler WsBLVTKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@nav_Kline_%s", getWsEndpoint(), strings.ToUpper(name), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBLVTKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCompositeIndexEvent websocket composite index event
type WsCompositeIndexEvent struct {
	Event       string          `json:"e"`
	Time        int64           `json:"E"`
	Symbol      string          `json:"s"`
	Price       string          `json:"p"`
	Composition []WsComposition `json:"c"`
}

// WsComposition websocket composite index event composition
type WsComposition struct {
	BaseAsset    string `json:"b"`
	WeightQty    string `json:"w"`
	WeighPercent string `json:"W"`
}

// WsCompositeIndexHandler websocket composite index handler
type WsCompositeIndexHandler func(event *WsCompositeIndexEvent)

// WsCompositiveIndexServe serve composite index information for index symbols
func WsCompositiveIndexServe(symbol string, handler WsCompositeIndexHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@compositeIndex", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsCompositeIndexEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}
