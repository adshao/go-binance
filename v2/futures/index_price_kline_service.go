package futures

import (
	"context"
	"fmt"
	"net/http"
)

// IndexPriceKlinesService list klines
type IndexPriceKlinesService struct {
	c         *Client
	pair      string
	interval  string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Pair sets pair
func (ipks *IndexPriceKlinesService) Pair(pair string) *IndexPriceKlinesService {
	ipks.pair = pair
	return ipks
}

// Interval set interval
func (ipks *IndexPriceKlinesService) Interval(interval string) *IndexPriceKlinesService {
	ipks.interval = interval
	return ipks
}

// Limit set limit
func (ipks *IndexPriceKlinesService) Limit(limit int) *IndexPriceKlinesService {
	ipks.limit = &limit
	return ipks
}

// StartTime set startTime
func (ipks *IndexPriceKlinesService) StartTime(startTime int64) *IndexPriceKlinesService {
	ipks.startTime = &startTime
	return ipks
}

// EndTime set endTime
func (ipks *IndexPriceKlinesService) EndTime(endTime int64) *IndexPriceKlinesService {
	ipks.endTime = &endTime
	return ipks
}

// Do send request
func (ipks *IndexPriceKlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/indexPriceKlines",
	}
	r.setParam("pair", ipks.pair)
	r.setParam("interval", ipks.interval)
	if ipks.limit != nil {
		r.setParam("limit", *ipks.limit)
	}
	if ipks.startTime != nil {
		r.setParam("startTime", *ipks.startTime)
	}
	if ipks.endTime != nil {
		r.setParam("endTime", *ipks.endTime)
	}
	data, _, err := ipks.c.callAPI(ctx, r, opts...)
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
