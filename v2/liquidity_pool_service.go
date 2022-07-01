package binance

import (
	"context"
	"net/http"
)

// GetAllSwapPoolService get swap pool market data
type GetAllLiquidityPoolService struct {
	c *Client
}

// Do send request
func (s *GetAllLiquidityPoolService) Do(ctx context.Context, opts ...RequestOption) ([]*LiquidityPool, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/pools",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []*LiquidityPool{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type LiquidityPool struct {
	PoolId   int64    `json:"poolId"`
	PoolName string   `json:"poolName"`
	Assets   []string `json:"assets"`
}

// GetLiquidityPoolDetailService get swap pool detail by pool id
type GetLiquidityPoolDetailService struct {
	c      *Client
	poolId *int64
}

// PoolId set poolId
func (s *GetLiquidityPoolDetailService) PoolId(poolId int64) *GetLiquidityPoolDetailService {
	s.poolId = &poolId
	return s
}

type LiquidityPoolDetail struct {
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
func (s *GetLiquidityPoolDetailService) Do(ctx context.Context) ([]*LiquidityPoolDetail, error) {
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
	res := []*LiquidityPoolDetail{}
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

// PoolId set poolId
func (s *AddLiquidityPreviewService) PoolId(poolId int64) *AddLiquidityPreviewService {
	s.poolId = &poolId
	return s
}

// QuoteAsset set quoteAsset
func (s *AddLiquidityPreviewService) QuoteAsset(quoteAsset string) *AddLiquidityPreviewService {
	s.quoteAsset = &quoteAsset
	return s
}

// QuoteQty set quoteQty
func (s *AddLiquidityPreviewService) QuoteQty(quoteQty float64) *AddLiquidityPreviewService {
	s.quoteQty = &quoteQty
	return s
}

// OperationType set operationType
func (s *AddLiquidityPreviewService) OperationType(operationType LiquidityOperationType) *AddLiquidityPreviewService {
	s.operationType = &operationType
	return s
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

// GetSwapQuoteService get the quote
type GetSwapQuoteService struct {
	c          *Client
	quoteAsset *string
	baseAsset  *string
	quoteQty   *float64
}

// QuoteAsset set quoteAsset
func (s *GetSwapQuoteService) QuoteAsset(quoteAsset string) *GetSwapQuoteService {
	s.quoteAsset = &quoteAsset
	return s
}

// QuoteQty set quoteQty
func (s *GetSwapQuoteService) QuoteQty(quoteQty float64) *GetSwapQuoteService {
	s.quoteQty = &quoteQty
	return s
}

// BaseAsset set baseAsset
func (s *GetSwapQuoteService) BaseAsset(baseAsset string) *GetSwapQuoteService {
	s.baseAsset = &baseAsset
	return s
}

type GetSwapQuoteResponse struct {
	QuoteAsset string `json:"quoteAsset"`
	BaseAsset  string `json:"baseAsset"`
	QuoteQty   string `json:"quoteQty"`
	BaseQty    string `json:"baseQty"`
	Price      string `json:"price"`
	Slippage   string `json:"slippage"`
	Fee        string `json:"fee"`
}

// Do sends the request.
func (s *GetSwapQuoteService) Do(ctx context.Context) (*GetSwapQuoteResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/quote",
		secType:  secTypeSigned,
	}

	r.setParam("quoteAsset", *s.quoteAsset)
	r.setParam("baseAsset", *s.baseAsset)
	r.setParam("quoteQty", *s.quoteQty)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &GetSwapQuoteResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SwapService swap tokens in liquidity pool
type SwapService struct {
	c          *Client
	quoteAsset *string
	baseAsset  *string
	quoteQty   *float64
}

// QuoteAsset set quoteAsset
func (s *SwapService) QuoteAsset(quoteAsset string) *SwapService {
	s.quoteAsset = &quoteAsset
	return s
}

// QuoteQty set quoteQty
func (s *SwapService) QuoteQty(quoteQty float64) *SwapService {
	s.quoteQty = &quoteQty
	return s
}

// BaseAsset set baseAsset
func (s *SwapService) BaseAsset(baseAsset string) *SwapService {
	s.baseAsset = &baseAsset
	return s
}

type SwapResponse struct {
	SwapId int64 `json:"swapId"`
}

// Do sends the request.
func (s *SwapService) Do(ctx context.Context) (*SwapResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/bswap/swap",
		secType:  secTypeSigned,
	}

	r.setParam("quoteAsset", *s.quoteAsset)
	r.setParam("baseAsset", *s.baseAsset)
	r.setParam("quoteQty", *s.quoteQty)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &SwapResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AddLiquidityService to add liquidity
type AddLiquidityService struct {
	c             *Client
	poolId        *int64
	operationType *LiquidityOperationType
	quoteAsset    *string
	quoteQty      *float64
}

// PoolId set poolId
func (s *AddLiquidityService) PoolId(poolId int64) *AddLiquidityService {
	s.poolId = &poolId
	return s
}

// QuoteAsset set quoteAsset
func (s *AddLiquidityService) QuoteAsset(quoteAsset string) *AddLiquidityService {
	s.quoteAsset = &quoteAsset
	return s
}

// QuoteQty set quoteQty
func (s *AddLiquidityService) QuoteQty(quoteQty float64) *AddLiquidityService {
	s.quoteQty = &quoteQty
	return s
}

// OperationType set operationType
func (s *AddLiquidityService) OperationType(operationType LiquidityOperationType) *AddLiquidityService {
	s.operationType = &operationType
	return s
}

type AddLiquidityResponse struct {
	OperationId int64 `json:"operationId"`
}

// Do sends the request.
func (s *AddLiquidityService) Do(ctx context.Context) (*AddLiquidityResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/liquidityAdd",
		secType:  secTypeSigned,
	}

	r.setParam("poolId", *s.poolId)
	r.setParam("type", *s.operationType)
	r.setParam("asset", *s.quoteAsset)
	r.setParam("quantity", *s.quoteQty)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &AddLiquidityResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
