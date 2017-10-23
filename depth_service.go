package binance

import (
	"context"
)

type DepthService struct {
	c      *Client
	symbol string
	limit  *int
}

func (s *DepthService) Symbol(symbol string) *DepthService {
	s.symbol = symbol
	return s
}

func (s *DepthService) Limit(limit int) *DepthService {
	s.limit = &limit
	return s
}

func (s *DepthService) Do(ctx context.Context, opts ...RequestOption) (res *DepthResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/depth",
	}
	r.SetParam("symbol", s.symbol)
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
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
	res.Bids = make([]Bid, bidsLen)
	for i := 0; i < bidsLen; i++ {
		item := j.Get("bids").GetIndex(i)
		res.Bids[i] = Bid{
			Price:    item.GetIndex(0).MustString(),
			Quantity: item.GetIndex(1).MustString(),
		}
	}
	asksLen := len(j.Get("asks").MustArray())
	res.Asks = make([]Ask, bidsLen)
	for i := 0; i < asksLen; i++ {
		item := j.Get("asks").GetIndex(i)
		res.Asks[i] = Ask{
			Price:    item.GetIndex(0).MustString(),
			Quantity: item.GetIndex(1).MustString(),
		}
	}
	return
}

type DepthResponse struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

type Bid struct {
	Price    string
	Quantity string
}

type Ask struct {
	Price    string
	Quantity string
}
