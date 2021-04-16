package binance

import (
	"context"
	"encoding/json"
)

// https://binance-docs.github.io/apidocs/spot/en/#list-all-swap-pools-market_data
type ListBSwapPoolsService struct {
	c *Client
}

type BSwapPool struct {
	PoolID   int64    `json:"poolId"`
	PoolName string   `json:"poolName"`
	Assets   []string `json:"assets"`
}

// Do sends the request.
func (bs *ListBSwapPoolsService) Do(ctx context.Context, opts ...RequestOption) ([]*BSwapPool, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/bswap/pools",
		secType:  secTypeSigned,
	}
	data, err := bs.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := make([]*BSwapPool, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// https://binance-docs.github.io/apidocs/spot/en/#get-liquidity-information-of-a-pool-user_data
type GetBSwapPoolLiquidityInfoService struct {
	c      *Client
	poolID *int64
}

type BSwapPoolLiquidityInfo struct {
	PoolID     int64             `json:"poolId"`
	PoolName   string            `json:"poolName"`
	UpdateTime int64             `json:"updateTime"`
	Liquidity  map[string]string `json:"liquidity"`
	Share      *BSwapShare       `json:"share"`
}

type BSwapShare struct {
	ShareAmount     string            `json:"shareAmount"`
	SharePercentage string            `json:"sharePercentage"`
	Asset           map[string]string `json:"asset"`
}

// PoolID set poolID
func (bs *GetBSwapPoolLiquidityInfoService) PoolID(poolID int64) *GetBSwapPoolLiquidityInfoService {
	bs.poolID = &poolID
	return bs
}

// Do send request
func (bs *GetBSwapPoolLiquidityInfoService) Do(ctx context.Context, opts ...RequestOption) ([]*BSwapPoolLiquidityInfo, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/bswap/liquidity",
		secType:  secTypeSigned,
	}
	if bs.poolID != nil {
		r.setParam("poolId", *bs.poolID)
	}
	data, err := bs.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := make([]*BSwapPoolLiquidityInfo, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// https://binance-docs.github.io/apidocs/spot/en/#add-liquidity-trade
type AddBSwapLiquidityService struct {
	c        *Client
	poolID   *int64
	asset    *string
	quantity *string
}

// PoolID set poolID (MANDATORY)
func (bs *AddBSwapLiquidityService) PoolID(poolID int64) *AddBSwapLiquidityService {
	bs.poolID = &poolID
	return bs
}

// Asset set asset (MANDATORY)
func (bs *AddBSwapLiquidityService) Asset(asset string) *AddBSwapLiquidityService {
	bs.asset = &asset
	return bs
}

// Quantity set quantity (MANDATORY)
func (bs *AddBSwapLiquidityService) Quantity(quantity string) *AddBSwapLiquidityService {
	bs.quantity = &quantity
	return bs
}

func (bs *AddBSwapLiquidityService) Do(ctx context.Context, opts ...RequestOption) (*BSwapLiquidityTradeResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/bswap/liquidityAdd",
		secType:  secTypeSigned,
	}
	if bs.poolID != nil {
		r.setParam("poolId", *bs.poolID)
	}
	if bs.asset != nil {
		r.setParam("asset", *bs.asset)
	}
	if bs.quantity != nil {
		r.setParam("quantity", *bs.quantity)
	}
	data, err := bs.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(BSwapLiquidityTradeResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// https://binance-docs.github.io/apidocs/spot/en/#remove-liquidity-trade
type RemoveBSwapLiquidityService struct {
	c           *Client
	poolID      *int64
	removalType *BSwapRemovalType
	asset       *string
	shareAmount *string
}

// PoolID set poolID (MANDATORY)
func (bs *RemoveBSwapLiquidityService) PoolID(poolID int64) *RemoveBSwapLiquidityService {
	bs.poolID = &poolID
	return bs
}

// Type set removalType (BSwapRemovalTypeSingle, BSwapRemovalTypeCombination) (MANDATORY)
func (bs *RemoveBSwapLiquidityService) Type(removalType BSwapRemovalType) *RemoveBSwapLiquidityService {
	bs.removalType = &removalType
	return bs
}

// Asset set asset (MANDATORY)
func (bs *RemoveBSwapLiquidityService) Asset(asset string) *RemoveBSwapLiquidityService {
	bs.asset = &asset
	return bs
}

// ShareAmount set shareAmount (MANDATORY)
func (bs *RemoveBSwapLiquidityService) ShareAmount(shareAmount string) *RemoveBSwapLiquidityService {
	bs.shareAmount = &shareAmount
	return bs
}

func (bs *RemoveBSwapLiquidityService) Do(ctx context.Context, opts ...RequestOption) (*BSwapLiquidityTradeResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/bswap/liquidityRemove",
		secType:  secTypeSigned,
	}
	if bs.poolID != nil {
		r.setParam("poolId", *bs.poolID)
	}
	if bs.removalType != nil {
		r.setParam("type", *bs.removalType)
	}
	if bs.asset != nil {
		r.setParam("asset", *bs.asset)
	}
	if bs.shareAmount != nil {
		r.setParam("shareAmount", *bs.shareAmount)
	}
	data, err := bs.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(BSwapLiquidityTradeResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type BSwapLiquidityTradeResponse struct {
	OperationId int64 `json:"operationId"`
}

// https://binance-docs.github.io/apidocs/spot/en/#get-liquidity-operation-record-user_data
type GetBSwapLiquidityOperationRecordService struct {
	c           *Client
	operationID *int64
	poolID      *int64
	operation   *BSwapLiquidityOperationType
	startTime   *int64
	endTime     *int64
	limit       *int64
}

type BSwapLiquidityOperationRecord struct {
	OperationId int64                       `json:"operationId"`
	PoolID      int64                       `json:"poolId"`
	PoolName    string                      `json:"poolName"`
	Operation   BSwapLiquidityOperationType `json:"operation"`
	Status      BSwapStatusType             `json:"status"`
	UpdateTime  int64                       `json:"updateTime"`
	ShareAmount string                      `json:"shareAmount"`
}

// OperationID set operationID
func (bs *GetBSwapLiquidityOperationRecordService) OperationID(operationID int64) *GetBSwapLiquidityOperationRecordService {
	bs.operationID = &operationID
	return bs
}

// PoolID set poolID
func (bs *GetBSwapLiquidityOperationRecordService) PoolID(poolID int64) *GetBSwapLiquidityOperationRecordService {
	bs.poolID = &poolID
	return bs
}

// Operation set operation (BSwapLiquidityOperationTypeAdd, BSwapLiquidityOperationTypeRemove)
func (bs *GetBSwapLiquidityOperationRecordService) Operation(operation BSwapLiquidityOperationType) *GetBSwapLiquidityOperationRecordService {
	bs.operation = &operation
	return bs
}

// StartTime set startTime
func (bs *GetBSwapLiquidityOperationRecordService) StartTime(startTime int64) *GetBSwapLiquidityOperationRecordService {
	bs.startTime = &startTime
	return bs
}

// EndTime set endTime
func (bs *GetBSwapLiquidityOperationRecordService) EndTime(endTime int64) *GetBSwapLiquidityOperationRecordService {
	bs.endTime = &endTime
	return bs
}

// Limit set limit
func (bs *GetBSwapLiquidityOperationRecordService) Limit(limit int64) *GetBSwapLiquidityOperationRecordService {
	bs.limit = &limit
	return bs
}

func (bs *GetBSwapLiquidityOperationRecordService) Do(ctx context.Context, opts ...RequestOption) ([]*BSwapLiquidityOperationRecord, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/bswap/liquidityOps",
		secType:  secTypeSigned,
	}
	if bs.operationID != nil {
		r.setParam("operationId", *bs.operationID)
	}
	if bs.poolID != nil {
		r.setParam("poolId", *bs.poolID)
	}
	if bs.operation != nil {
		r.setParam("operation", *bs.operation)
	}
	if bs.startTime != nil {
		r.setParam("startTime", *bs.startTime)
	}
	if bs.endTime != nil {
		r.setParam("endTime", *bs.endTime)
	}
	if bs.limit != nil {
		r.setParam("limit", *bs.limit)
	}
	data, err := bs.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := make([]*BSwapLiquidityOperationRecord, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// https://binance-docs.github.io/apidocs/spot/en/#request-quote-user_data
type RequestBSwapQuoteService struct {
	c          *Client
	quoteAsset *string
	baseAsset  *string
	quoteQty   *string
}

type BSwapQuoteResponse struct {
	QuoteAsset string  `json:"quoteAsset"`
	BaseAsset  string  `json:"baseAsset"`
	QuoteQty   string `json:"quoteQty"`
	BaseQty    string `json:"baseQty"`
	Price      string `json:"price"`
	Slippage   string `json:"slippage"`
	Fee        string `json:"fee"`
}

// QuoteAsset set quoteAsset (MANDATORY)
func (bs *RequestBSwapQuoteService) QuoteAsset(quoteAsset string) *RequestBSwapQuoteService {
	bs.quoteAsset = &quoteAsset
	return bs
}

// BaseAsset set baseAsset (MANDATORY)
func (bs *RequestBSwapQuoteService) BaseAsset(baseAsset string) *RequestBSwapQuoteService {
	bs.baseAsset = &baseAsset
	return bs
}

// QuoteQty set quoteQty (MANDATORY)
func (bs *RequestBSwapQuoteService) QuoteQty(quoteQty string) *RequestBSwapQuoteService {
	bs.quoteQty = &quoteQty
	return bs
}

func (bs *RequestBSwapQuoteService) Do(ctx context.Context, opts ...RequestOption) (*BSwapQuoteResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/bswap/quote",
		secType:  secTypeSigned,
	}
	if bs.quoteAsset != nil {
		r.setParam("quoteAsset", *bs.quoteAsset)
	}
	if bs.baseAsset != nil {
		r.setParam("baseAsset", *bs.baseAsset)
	}
	if bs.quoteQty != nil {
		r.setParam("quoteQty", *bs.quoteQty)
	}
	data, err := bs.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(BSwapQuoteResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// https://binance-docs.github.io/apidocs/spot/en/#swap-trade
type SwapBSwapService struct {
	c          *Client
	quoteAsset *string
	baseAsset  *string
	quoteQty   *string
}

// QuoteAsset set quoteAsset
func (bs *SwapBSwapService) QuoteAsset(quoteAsset string) *SwapBSwapService {
	bs.quoteAsset = &quoteAsset
	return bs
}

// BaseAsset set baseAsset
func (bs *SwapBSwapService) BaseAsset(baseAsset string) *SwapBSwapService {
	bs.baseAsset = &baseAsset
	return bs
}

// QuoteQty set quoteQty
func (bs *SwapBSwapService) QuoteQty(quoteQty string) *SwapBSwapService {
	bs.quoteQty = &quoteQty
	return bs
}

type SwapBSwapResponse struct {
	SwapId int64 `json:"swapId"`
}

func (bs *SwapBSwapService) Do(ctx context.Context, opts ...RequestOption) (*SwapBSwapResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/bswap/swap",
		secType:  secTypeSigned,
	}
	if bs.quoteAsset != nil {
		r.setParam("quoteAsset", *bs.quoteAsset)
	}
	if bs.baseAsset != nil {
		r.setParam("baseAsset", *bs.baseAsset)
	}
	if bs.quoteQty != nil {
		r.setParam("quoteQty", *bs.quoteQty)
	}
	data, err := bs.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(SwapBSwapResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// https://binance-docs.github.io/apidocs/spot/en/#get-swap-history-user_data
type GetBSwapSwapHistoryService struct {
	c          *Client
	swapId     *int64
	startTime  *int64
	endTime    *int64
	status     *BSwapStatusType
	quoteAsset *string
	baseAsset  *string
	limit      *int64
}

type BSwapSwapHistory struct {
	SwapId     int64           `json:"swapId"`
	SwapTime   int64         `json:"swapTime"`
	Status     BSwapStatusType `json:"status"`
	BaseAsset  string          `json:"baseAsset"`
	QuoteAsset string          `json:"quoteAsset"`
	QuoteQty   string         `json:"quoteQty"`
	BaseQty    string         `json:"baseQty"`
	Price      string         `json:"price"`
	Fee        string         `json:"fee"`
}

// SwapId set swapId
func (bs *GetBSwapSwapHistoryService) SwapId(swapId int64) *GetBSwapSwapHistoryService {
	bs.swapId = &swapId
	return bs
}

// StartTime set startTime
func (bs *GetBSwapSwapHistoryService) StartTime(startTime int64) *GetBSwapSwapHistoryService {
	bs.startTime = &startTime
	return bs
}

// EndTime set endTime
func (bs *GetBSwapSwapHistoryService) EndTime(endTime int64) *GetBSwapSwapHistoryService {
	bs.endTime = &endTime
	return bs
}

// Status set status
func (bs *GetBSwapSwapHistoryService) Status(status BSwapStatusType) *GetBSwapSwapHistoryService {
	bs.status = &status
	return bs
}

// QuoteAsset set quoteAsset
func (bs *GetBSwapSwapHistoryService) QuoteAsset(quoteAsset string) *GetBSwapSwapHistoryService {
	bs.quoteAsset = &quoteAsset
	return bs
}

// BaseAsset set baseAsset
func (bs *GetBSwapSwapHistoryService) BaseAsset(baseAsset string) *GetBSwapSwapHistoryService {
	bs.baseAsset = &baseAsset
	return bs
}

// Limit set limit
func (bs *GetBSwapSwapHistoryService) Limit(limit int64) *GetBSwapSwapHistoryService {
	bs.limit = &limit
	return bs
}

func (bs *GetBSwapSwapHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]*BSwapSwapHistory, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/bswap/swap",
		secType:  secTypeSigned,
	}
	if bs.swapId != nil {
		r.setParam("swapId", *bs.swapId)
	}
	if bs.startTime != nil {
		r.setParam("startTime", *bs.startTime)
	}
	if bs.endTime != nil {
		r.setParam("endTime", *bs.endTime)
	}
	if bs.status != nil {
		r.setParam("status", *bs.status)
	}
	if bs.quoteAsset != nil {
		r.setParam("quoteAsset", *bs.quoteAsset)
	}
	if bs.baseAsset != nil {
		r.setParam("baseAsset", *bs.baseAsset)
	}
	if bs.limit != nil {
		r.setParam("limit", *bs.limit)
	}
	data, err := bs.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := make([]*BSwapSwapHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
