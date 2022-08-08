package futures

import (
	"context"
	"fmt"
	"net/http"
)

// MarkPriceKlinesService list mark price klines
type MarkPriceKlinesService struct {
	c         *Client
	symbol    string
	interval  string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (mpks *MarkPriceKlinesService) Symbol(symbol string) *MarkPriceKlinesService {
	mpks.symbol = symbol
	return mpks
}

// Interval set interval
func (mpks *MarkPriceKlinesService) Interval(interval string) *MarkPriceKlinesService {
	mpks.interval = interval
	return mpks
}

// Limit set limit
func (mpks *MarkPriceKlinesService) Limit(limit int) *MarkPriceKlinesService {
	mpks.limit = &limit
	return mpks
}

// StartTime set startTime
func (mpks *MarkPriceKlinesService) StartTime(startTime int64) *MarkPriceKlinesService {
	mpks.startTime = &startTime
	return mpks
}

// EndTime set endTime
func (mpks *MarkPriceKlinesService) EndTime(endTime int64) *MarkPriceKlinesService {
	mpks.endTime = &endTime
	return mpks
}

// Do send request
func (mpks *MarkPriceKlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/markPriceKlines",
	}
	r.setParam("symbol", mpks.symbol)
	r.setParam("interval", mpks.interval)
	if mpks.limit != nil {
		r.setParam("limit", *mpks.limit)
	}
	if mpks.startTime != nil {
		r.setParam("startTime", *mpks.startTime)
	}
	if mpks.endTime != nil {
		r.setParam("endTime", *mpks.endTime)
	}
	data, _, err := mpks.c.callAPI(ctx, r, opts...)
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
			OpenTime:  item.GetIndex(0).MustInt64(),
			Open:      item.GetIndex(1).MustString(),
			High:      item.GetIndex(2).MustString(),
			Low:       item.GetIndex(3).MustString(),
			Close:     item.GetIndex(4).MustString(),
			CloseTime: item.GetIndex(6).MustInt64(),
		}
	}
	return res, nil
}
