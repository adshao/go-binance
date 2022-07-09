package futures

import (
	"context"
	"fmt"
	"net/http"
)

// ContinuousKlinesService list klines
type ContinuousKlinesService struct {
	c         *Client
	symbol    string
	contractType ContractType
	interval  string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *ContinuousKlinesService) Symbol(symbol string) *ContinuousKlinesService {
	s.symbol = symbol
	return s
}

// ContractType set contractType
func (s *ContinuousKlinesService) ContractType(contractType ContractType) *ContinuousKlinesService {
	s.contractType = contractType
	return s
}

// Interval set interval
func (s *ContinuousKlinesService) Interval(interval string) *ContinuousKlinesService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *ContinuousKlinesService) Limit(limit int) *ContinuousKlinesService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *ContinuousKlinesService) StartTime(startTime int64) *ContinuousKlinesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ContinuousKlinesService) EndTime(endTime int64) *ContinuousKlinesService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ContinuousKlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/klines",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Kline{}, err
	}
	j, err := newJSON(data)
	if err != nil {
		return []*Kline{}, err
	}
	num := len(j.MustArray())
	res = make([]*Kline, num)
	for i := 0; i < num; i++ {
		item := j.GetIndex(i)
		if len(item.MustArray()) < 11 {
			err = fmt.Errorf("invalid kline response")
			return []*Kline{}, err
		}
		res[i] = &Kline{
			OpenTime:                 item.GetIndex(0).MustInt64(),
			Open:                     item.GetIndex(1).MustString(),
			High:                     item.GetIndex(2).MustString(),
			Low:                      item.GetIndex(3).MustString(),
			Close:                    item.GetIndex(4).MustString(),
			Volume:                   item.GetIndex(5).MustString(),
			CloseTime:                item.GetIndex(6).MustInt64(),
			QuoteAssetVolume:         item.GetIndex(7).MustString(),
			TradeNum:                 item.GetIndex(8).MustInt64(),
			TakerBuyBaseAssetVolume:  item.GetIndex(9).MustString(),
			TakerBuyQuoteAssetVolume: item.GetIndex(10).MustString(),
		}
	}
	return res, nil
}
