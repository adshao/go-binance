package binance

import (
	"context"
	"net/http"
)

// GetSwapPoolService get swap pool market data
type GetSwapPoolService struct {
	c *Client
}

// Do send request
func (s *GetSwapPoolService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapPool, error) {
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
