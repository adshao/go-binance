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
	ShareAmount     string            `json:"shareAmount"`
	SharePercentage string            `json:"sharePercentage"`
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

// AddLiquidityPreviewService to preview the quote/base qty needed when adding assets to a liquidity pool with an estimated share after adding the liquidity
type AddLiquidityPreviewService struct {
	c             *Client
	poolId        *int64
	operationType *LiquidityOperationType
	quoteAsset    *string
	quoteQty      *float64
}

type AddLiquidityPreviewResponse struct {
	QuoteAsset string `json:"quoteAsset"`
	BaseAsset  string `json:"baseAsset"` // only existed when type is COMBINATION
	QuoteAmt   string `json:"quoteAmt"`
	BaseAmt    string `json:"baseAmt"` // only existed when type is COMBINATION
	Price      string `json:"price"`
	Share      string `json:"share"`
	Slippage   string `json:"slippage"`
	Fee        string `json:"fee"`
}

// Do sends the request.
func (s *AddLiquidityPreviewService) Do(ctx context.Context) (*AddLiquidityPreviewResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/addLiquidityPreview",
		secType:  secTypeSigned,
	}

	r.setParam("poolId", *s.poolId)
	r.setParam("type", *s.operationType)
	r.setParam("quoteAsset", *s.quoteAsset)
	r.setParam("quoteQty", *s.quoteQty)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &AddLiquidityPreviewResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
