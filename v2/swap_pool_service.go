package binance

import (
	"context"
	"net/http"
)

// GetAllSwapPoolService get swap pool market data
type GetAllSwapPoolService struct {
	c *Client
}

// Do send request
func (s *GetAllSwapPoolService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapPool, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/pools",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []*SwapPool{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type SwapPool struct {
	PoolId   int64    `json:"poolId"`
	PoolName string   `json:"poolName"`
	Assets   []string `json:"assets"`
}

// GetSwapPoolDetailService get swap pool detail by pool id
type GetSwapPoolDetailService struct {
	c      *Client
	poolId *int64
}

// PoolId set poolId
func (s *GetSwapPoolDetailService) PoolId(poolId int64) *GetSwapPoolDetailService {
	s.poolId = &poolId
	return s
}

type SwapPoolDetail struct {
	PoolId     int64                 `json:"poolId"`
	PoolName   string                `json:"poolName"`
	UpdateTime int64                 `json:"updateTime"`
	Liquidity  map[string]string     `json:"liquidity"`
	Share      *PoolShareInformation `json:"share"`
}

type PoolShareInformation struct {
	ShareAmount     float64           `json:"shareAmount"`
	SharePercentage float64           `json:"sharePercentage"`
	Assets          map[string]string `json:"asset"`
}

// Do sends the request.
func (s *GetSwapPoolDetailService) Do(ctx context.Context) ([]*SwapPoolDetail, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/liquidity",
		secType:  secTypeSigned,
	}
	if s.poolId != nil {
		r.setParam("poolId", *s.poolId)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := []*SwapPoolDetail{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
