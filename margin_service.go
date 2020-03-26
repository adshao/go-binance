package binance

import (
	"context"
	"encoding/json"
)

// MarginTransferService transfer between spot account and margin account
type MarginTransferService struct {
	c            *Client
	asset        string
	amount       string
	transferType int
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginTransferService) Asset(asset string) *MarginTransferService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginTransferService) Amount(amount string) *MarginTransferService {
	s.amount = amount
	return s
}

// Type 1: transfer from main account to margin account 2: transfer from margin account to main account
func (s *MarginTransferService) Type(transferType MarginTransferType) *MarginTransferService {
	s.transferType = int(transferType)
	return s
}

// Do send request
func (s *MarginTransferService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/margin/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TransactionResponse define transaction response
type TransactionResponse struct {
	TranID int64 `json:"tranId"`
}

// MarginLoanService apply for a loan
type MarginLoanService struct {
	c      *Client
	asset  string
	amount string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginLoanService) Asset(asset string) *MarginLoanService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginLoanService) Amount(amount string) *MarginLoanService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginLoanService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/margin/loan",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginRepayService repay loan for margin account
type MarginRepayService struct {
	c      *Client
	asset  string
	amount string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginRepayService) Asset(asset string) *MarginRepayService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginRepayService) Amount(amount string) *MarginRepayService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginRepayService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/margin/repay",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListMarginLoansService list loan record
type ListMarginLoansService struct {
	c         *Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListMarginLoansService) Asset(asset string) *ListMarginLoansService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *ListMarginLoansService) TxID(txID int64) *ListMarginLoansService {
	s.txID = &txID
	return s
}

// StartTime set start time
func (s *ListMarginLoansService) StartTime(startTime int64) *ListMarginLoansService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *ListMarginLoansService) EndTime(endTime int64) *ListMarginLoansService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginLoansService) Current(current int64) *ListMarginLoansService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginLoansService) Size(size int64) *ListMarginLoansService {
	s.size = &size
	return s
}

// Do send request
func (s *ListMarginLoansService) Do(ctx context.Context, opts ...RequestOption) (res *MarginLoanResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/loan",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.txID != nil {
		r.setParam("txId", *s.txID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginLoanResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginLoanResponse define margin loan response
type MarginLoanResponse struct {
	Rows  []MarginLoan `json:"rows"`
	Total int64        `json:"total"`
}

// MarginLoan define margin loan
type MarginLoan struct {
	Asset     string               `json:"asset"`
	Principal string               `json:"principal"`
	Timestamp int64                `json:"timestamp"`
	Status    MarginLoanStatusType `json:"status"`
}

// ListMarginRepaysService list repay record
type ListMarginRepaysService struct {
	c         *Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListMarginRepaysService) Asset(asset string) *ListMarginRepaysService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *ListMarginRepaysService) TxID(txID int64) *ListMarginRepaysService {
	s.txID = &txID
	return s
}

// StartTime set start time
func (s *ListMarginRepaysService) StartTime(startTime int64) *ListMarginRepaysService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *ListMarginRepaysService) EndTime(endTime int64) *ListMarginRepaysService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginRepaysService) Current(current int64) *ListMarginRepaysService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginRepaysService) Size(size int64) *ListMarginRepaysService {
	s.size = &size
	return s
}

// Do send request
func (s *ListMarginRepaysService) Do(ctx context.Context, opts ...RequestOption) (res *MarginRepayResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/repay",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.txID != nil {
		r.setParam("txId", *s.txID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginRepayResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginRepayResponse define margin repay response
type MarginRepayResponse struct {
	Rows  []MarginRepay `json:"rows"`
	Total int64         `json:"total"`
}

// MarginRepay define margin repay
type MarginRepay struct {
	Asset     string                `json:"asset"`
	Amount    string                `json:"amount"`
	Interest  string                `json:"interest"`
	Principal string                `json:"principal"`
	Timestamp int64                 `json:"timestamp"`
	Status    MarginRepayStatusType `json:"status"`
	TxID      int64                 `json:"txId"`
}

// GetMarginAccountService get margin account info
type GetMarginAccountService struct {
	c *Client
}

// Do send request
func (s *GetMarginAccountService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccount, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginAccount define margin account info
type MarginAccount struct {
	BorrowEnabled       bool        `json:"borrowEnabled"`
	MarginLevel         string      `json:"marginLevel"`
	TotalAssetOfBTC     string      `json:"totalAssetOfBtc"`
	TotalLiabilityOfBTC string      `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBTC  string      `json:"totalNetAssetOfBtc"`
	TradeEnabled        bool        `json:"tradeEnabled"`
	TransferEnabled     bool        `json:"transferEnabled"`
	UserAssets          []UserAsset `json:"userAssets"`
}

// UserAsset define user assets of margin account
type UserAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// GetMarginAssetService get margin asset info
type GetMarginAssetService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMarginAssetService) Asset(asset string) *GetMarginAssetService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMarginAssetService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAsset, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/asset",
		secType:  secTypeAPIKey,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginAsset)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginAsset define margin asset info
type MarginAsset struct {
	FullName      string `json:"assetFullName"`
	Name          string `json:"assetName"`
	Borrowable    bool   `json:"isBorrowable"`
	Mortgageable  bool   `json:"isMortgageable"`
	UserMinBorrow string `json:"userMinBorrow"`
	UserMinRepay  string `json:"userMinRepay"`
}

// GetMarginPairService get margin pair info
type GetMarginPairService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetMarginPairService) Symbol(symbol string) *GetMarginPairService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetMarginPairService) Do(ctx context.Context, opts ...RequestOption) (res *MarginPair, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/pair",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginPair)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginPair define margin pair info
type MarginPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// GetMarginAllPairsService get margin pair info
type GetMarginAllPairsService struct {
	c *Client
}

// Do send request
func (s *GetMarginAllPairsService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAllPair, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/allPairs",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAllPair{}, err
	}
	res = make([]*MarginAllPair, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAllPair{}, err
	}
	return res, nil
}

// MarginAllPair define margin pair info
type MarginAllPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// GetMarginPriceIndexService get margin price index
type GetMarginPriceIndexService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetMarginPriceIndexService) Symbol(symbol string) *GetMarginPriceIndexService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetMarginPriceIndexService) Do(ctx context.Context, opts ...RequestOption) (res *MarginPriceIndex, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/priceIndex",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginPriceIndex)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginPriceIndex define margin price index
type MarginPriceIndex struct {
	CalcTime int64  `json:"calcTime"`
	Price    string `json:"price"`
	Symbol   string `json:"symbol"`
}

// ListMarginTradesService list trades
type ListMarginTradesService struct {
	c         *Client
	symbol    string
	startTime *int64
	endTime   *int64
	limit     *int
	fromID    *int64
}

// Symbol set symbol
func (s *ListMarginTradesService) Symbol(symbol string) *ListMarginTradesService {
	s.symbol = symbol
	return s
}

// StartTime set starttime
func (s *ListMarginTradesService) StartTime(startTime int64) *ListMarginTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListMarginTradesService) EndTime(endTime int64) *ListMarginTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListMarginTradesService) Limit(limit int) *ListMarginTradesService {
	s.limit = &limit
	return s
}

// FromID set fromID
func (s *ListMarginTradesService) FromID(fromID int64) *ListMarginTradesService {
	s.fromID = &fromID
	return s
}

// Do send request
func (s *ListMarginTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*TradeV3, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/myTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*TradeV3{}, err
	}
	res = make([]*TradeV3, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*TradeV3{}, err
	}
	return res, nil
}

// GetMaxBorrowableService get max borrowable of asset
type GetMaxBorrowableService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMaxBorrowableService) Asset(asset string) *GetMaxBorrowableService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMaxBorrowableService) Do(ctx context.Context, opts ...RequestOption) (res *MaxBorrowable, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/maxBorrowable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxBorrowable)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MaxBorrowable define max borrowable response
type MaxBorrowable struct {
	Amount string `json:"amount"`
}

// GetMaxTransferableService get max transferable of asset
type GetMaxTransferableService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMaxTransferableService) Asset(asset string) *GetMaxTransferableService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMaxTransferableService) Do(ctx context.Context, opts ...RequestOption) (res *MaxTransferable, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/maxTransferable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxTransferable)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MaxTransferable define max transferable response
type MaxTransferable struct {
	Amount string `json:"amount"`
}

// StartMarginUserStreamService create listen key for margin user stream service
type StartMarginUserStreamService struct {
	c *Client
}

// Do send request
func (s *StartMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveMarginUserStreamService update listen key
type KeepaliveMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveMarginUserStreamService) ListenKey(listenKey string) *KeepaliveMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "PUT",
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseMarginUserStreamService delete listen key
type CloseMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseMarginUserStreamService) ListenKey(listenKey string) *CloseMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}
