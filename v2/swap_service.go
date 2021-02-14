package binance

import (
	"context"
	"encoding/json"
)

// LiquidityRemovalType represents liquidity removal type
type LiquidityRemovalType string

const (
	// LiquidityRemovalTypeSingle represents liquidity removal of single asset
	LiquidityRemovalTypeSingle = "SINGLE"
	// LiquidityRemovalTypeCombination represents liquidity removal of combination of assets
	LiquidityRemovalTypeCombination = "COMBINATION"
)

// ListAllSwapPoolsService gets metadata about all swap pools.
type ListAllSwapPoolsService struct {
	c *Client
}

// Do sends the request.
func (s *ListAllSwapPoolsService) Do(ctx context.Context) (ListAllSwapPoolsResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/bswap/pools",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := ListAllSwapPoolsResponse{}
	if err := json.Unmarshal(data, &res); err != nil {
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

// ListLiquidityService gets liquidity information and user share of a pool.
type ListLiquidityService struct {
	c      *Client
	poolID int64
}

// Do sends the request.
func (s *ListLiquidityService) Do(ctx context.Context) (ListLiquidityResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/bswap/liquidity",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := ListLiquidityResponse{}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListLiquidityResponse represents a response from ListLiquidityService.
type ListLiquidityResponse []struct {
	PoolID     int               `json:"poolId"`
	PoolName   string            `json:"poolName"`
	UpdateTime int64             `json:"updateTime"`
	Liquidity  map[string]string `json:"liquidity"`
	Share      struct {
		ShareAmount     string            `json:"shareAmount"`
		SharePercentage string            `json:"sharePercentage"`
		Asset           map[string]string `json:"asset"`
	} `json:"share"`
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

// RemoveLiquidityService removes liquidity from swap pool.
type RemoveLiquidityService struct {
	c           *Client
	poolID      int64
	removalType LiquidityRemovalType
	asset       string
	shareAmount float64
}

// PoolID sets the poolID parameter (MANDATORY).
func (s *RemoveLiquidityService) PoolID(v int64) *RemoveLiquidityService {
	s.poolID = v
	return s
}

// Type sets the type parameter (MANDATORY).
func (s *RemoveLiquidityService) Type(v LiquidityRemovalType) *RemoveLiquidityService {
	s.removalType = v
	return s
}

// Asset sets the asset parameter (MANDATORY).
func (s *RemoveLiquidityService) Asset(v string) *RemoveLiquidityService {
	s.asset = v
	return s
}

// ShareAmount sets the shareAmount parameter (MANDATORY).
func (s *RemoveLiquidityService) ShareAmount(v float64) *RemoveLiquidityService {
	s.shareAmount = v
	return s
}

// Do sends the request.
func (s *RemoveLiquidityService) Do(ctx context.Context) (*RemoveLiquidityResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/bswap/liquidityRemove",
		secType:  secTypeSigned,
	}
	r.setParam("poolId", s.poolID)
	r.setParam("type", s.removalType)
	r.setParam("asset", s.asset)
	r.setParam("shareAmount", s.shareAmount)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(RemoveLiquidityResponse)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// RemoveLiquidityResponse represents a response from RemoveLiquidityService.
type RemoveLiquidityResponse struct {
	OperationID int `json:"operationId"`
}
