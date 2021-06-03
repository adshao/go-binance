/*
* go-binance
* date: 2021-06-03 14:44
 */

package delivery

import (
	"context"
	"encoding/json"
)

//Symbol Order Book Ticker
type TickBookTickerService struct {
	c      *Client
	symbol string
	pair   string
}

func (s *TickBookTickerService) Symbol(symbol string) *TickBookTickerService {
	s.symbol = symbol
	return s
}

func (s *TickBookTickerService) Pair(pair string) *TickBookTickerService {
	s.pair = pair
	return s
}

func (s *TickBookTickerService) Do(ctx context.Context, opts ...RequestOption) (res []BookTicker, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/dapi/v1/ticker/bookTicker",
		secType:  secTypeSigned,
	}
	r.setParams(params{
		"symbol": s.symbol,
		"pair":   s.pair,
	})
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]BookTicker, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//{
//"symbol": "BTCUSD_200626",  // 交易对
//"pair": "BTCUSD",           // 标的交易对
//"bidPrice": "9650.1",       //最优买单价
//"bidQty": "16",             //最优买单挂单量
//"askPrice": "9650.3",       //最优卖单价
//"askQty": "7",              //最优卖单挂单量
//"time": 1591257300345
//}

type BookTicker struct {
	Symbol   string `json:"symbol"`
	Pair     string `json:"pair"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     int64  `json:"time"`
}
