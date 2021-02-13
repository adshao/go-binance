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

// AddLiquidityService adds liquidity into swap pool.
type AddLiquidityService struct {
	c        *Client
	poolID   int64
	asset    string
	quantity float64
}

// PoolID sets the poolID parameter (MANDATORY).
func (s *AddLiquidityService) PoolID(v int64) *AddLiquidityService {
	s.poolID = v
	return s
}

// Asset sets the asset parameter (MANDATORY).
func (s *AddLiquidityService) Asset(v string) *AddLiquidityService {
	s.asset = v
	return s
}

// Quantity sets the quantity parameter (MANDATORY).
func (s *AddLiquidityService) Quantity(v float64) *AddLiquidityService {
	s.quantity = v
	return s
}

// Do sends the request.
func (s *AddLiquidityService) Do(ctx context.Context) (*AddLiquidityResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/bswap/liquidityAdd",
		secType:  secTypeSigned,
	}
	r.setParam("poolId", s.poolID)
	r.setParam("asset", s.asset)
	r.setParam("quantity", s.quantity)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(AddLiquidityResponse)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// AddLiquidityResponse represents a response from AddLiquidityService.
type AddLiquidityResponse struct {
	OperationID int `json:"operationId"`
}
