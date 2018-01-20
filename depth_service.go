package binance

import (
	"context"
	"strconv"
)

// DepthService show depth info
type DepthService struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *DepthService) Symbol(symbol string) *DepthService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *DepthService) Limit(limit int) *DepthService {
	s.limit = &limit
	return s
}

// Do send request
func (s *DepthService) Do(ctx context.Context, opts ...RequestOption) (res *DepthResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/depth",
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	j, err := newJSON(data)
	if err != nil {
		return
	}
	res = new(DepthResponse)
	res.LastUpdateID = j.Get("lastUpdateId").MustInt64()
	bidsLen := len(j.Get("bids").MustArray())
	res.Bids = make([]BookEntry, bidsLen)
	for i := 0; i < bidsLen; i++ {
		item := j.Get("bids").GetIndex(i)
		res.Bids[i] = BookEntry{
			item.GetIndex(0).MustString(),
			item.GetIndex(1).MustString(),
		}
	}
	asksLen := len(j.Get("asks").MustArray())
	res.Asks = make([]BookEntry, asksLen)
	for i := 0; i < asksLen; i++ {
		item := j.Get("asks").GetIndex(i)
		res.Asks[i] = BookEntry{
			item.GetIndex(0).MustString(),
			item.GetIndex(1).MustString(),
		}
	}
	return
}

// DepthResponse define depth info with bids and asks
type DepthResponse struct {
	LastUpdateID int64       `json:"lastUpdateId"`
	Bids         []BookEntry `json:"bids"`
	Asks         []BookEntry `json:"asks"`
}

// BookEntry = bid or ask info with price and quantity
type BookEntry []string

// Price = bid or ask price
func (be *BookEntry) Price() float64 {
	out, _ := strconv.ParseFloat((*be)[0], 64)
	return out
}

// Quantity = bid or ask size
func (be *BookEntry) Quantity() float64 {
	out, _ := strconv.ParseFloat((*be)[1], 64)
	return out
}
