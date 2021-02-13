package binance

import (
	"context"
	"encoding/json"
)

// ListAllSwapPoolsService gets metadata about all swap pools.
type ListAllSwapPoolsService struct {
	c *Client
}

// Do sends the request.
func (s *ListAllSwapPoolsService) Do(ctx context.Context) (*ListAllSwapPoolsResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/bswap/pools",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(ListAllSwapPoolsResponse)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListAllSwapPoolsResponse represents a response from ListAllSwapPoolsService.
type ListAllSwapPoolsResponse []struct {
	PoolID   int      `json:"poolId"`
	PoolName string   `json:"poolName"`
	Assets   []string `json:"assets"`
}
