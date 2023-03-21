package binance

import (
	"context"
	"net/http"
)

type BrokerSubAccountService struct {
	c *Client
}

// Do sends the request.
func (s *BrokerSubAccountService) Do(ctx context.Context) (*BrokerSubAccount, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(BrokerSubAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type BrokerSubAccount struct {
	SubAccountID string `json:"subaccountId"`
	Email        string `json:"email"`
	Tag          string `json:"tag"`
}

type BrokerMarginService struct {
	c            *Client
	subaccountId *string
	margin       *bool
}

// SubaccountID set subaccountId
func (s *BrokerMarginService) SubaccountID(subaccountId string) *BrokerMarginService {
	s.subaccountId = &subaccountId
	return s
}

// Margin set margin
func (s *BrokerMarginService) Margin(margin bool) *BrokerMarginService {
	s.margin = &margin
	return s
}

// Do sends the request.
func (s *BrokerMarginService) Do(ctx context.Context) (*BrokerMargin, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccount/margin",
		secType:  secTypeSigned,
	}
	r.setParam("subaccountId", *s.subaccountId)
	r.setParam("margin", *s.margin)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(BrokerMargin)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type BrokerMargin struct {
	SubAccountID string `json:"subaccountId"`
	EnableMargin bool   `json:"enableMargin"`
	UpdateTime   int64  `json:"updateTime"`
}

type BrokerFutureService struct {
	c            *Client
	subaccountId *string
	futures      *bool
}

// SubaccountID set subaccountId
func (s *BrokerFutureService) SubaccountID(subaccountId string) *BrokerFutureService {
	s.subaccountId = &subaccountId
	return s
}

// Margin set margin
func (s *BrokerFutureService) Futures(futures bool) *BrokerFutureService {
	s.futures = &futures
	return s
}

// Do sends the request.
func (s *BrokerFutureService) Do(ctx context.Context) (*BrokerFuture, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccount/futures",
		secType:  secTypeSigned,
	}
	r.setParam("subaccountId", *s.subaccountId)
	r.setParam("futures", *s.futures)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(BrokerFuture)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type BrokerFuture struct {
	SubAccountID  string `json:"subaccountId"`
	EnableFutures bool   `json:"enableFutures"`
	UpdateTime    int64  `json:"updateTime"`
}

type BrokerAPIKeyService struct {
	c            *Client
	subAccountId *string
	canTrade     *bool
	marginTrade  *bool
	futuresTrade *bool
}

// SubaccountID set subaccountId
func (s *BrokerAPIKeyService) SubaccountID(subAccountId string) *BrokerAPIKeyService {
	s.subAccountId = &subAccountId
	return s
}

// CanTrade set canTrade
func (s *BrokerAPIKeyService) CanTrade(canTrade bool) *BrokerAPIKeyService {
	s.canTrade = &canTrade
	return s
}

// MarginTrade set marginTrade
func (s *BrokerAPIKeyService) MarginTrade(marginTrade bool) *BrokerAPIKeyService {
	s.marginTrade = &marginTrade
	return s
}

// FuturesTrade set futuresTrade
func (s *BrokerAPIKeyService) FuturesTrade(futuresTrade bool) *BrokerAPIKeyService {
	s.futuresTrade = &futuresTrade
	return s
}

// Do sends the request.
func (s *BrokerAPIKeyService) Do(ctx context.Context) (*BrokerAPIKey, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	r.setParam("canTrade", *s.canTrade)
	if s.marginTrade != nil {
		r.setParam("marginTrade", *s.marginTrade)
	}
	if s.futuresTrade != nil {
		r.setParam("futuresTrade", *s.futuresTrade)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(BrokerAPIKey)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type BrokerAPIKey struct {
	SubAccountID string `json:"subaccountId"`
	APIKey       string `json:"apiKey"`
	SecretKey    string `json:"secretKey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

type BrokerSpotSubaccountTransferService struct {
	c            *Client
	fromId       *string
	toId         *string
	clientTranId *string
	asset        *string
	amount       *float64
}

// BrokerSpotSubaccountTransferService set fromId
func (s *BrokerSpotSubaccountTransferService) FromID(fromId string) *BrokerSpotSubaccountTransferService {
	s.fromId = &fromId
	return s
}

// BrokerSpotSubaccountTransferService set toId
func (s *BrokerSpotSubaccountTransferService) ToId(toId string) *BrokerSpotSubaccountTransferService {
	s.toId = &toId
	return s
}

// BrokerSpotSubaccountTransferService set clientTranId
func (s *BrokerSpotSubaccountTransferService) ClientTraniD(clientTranId string) *BrokerSpotSubaccountTransferService {
	s.clientTranId = &clientTranId
	return s
}

// BrokerSpotSubaccountTransferService set asset
func (s *BrokerSpotSubaccountTransferService) Asset(asset string) *BrokerSpotSubaccountTransferService {
	s.asset = &asset
	return s
}

// BrokerSpotSubaccountTransferService set asset
func (s *BrokerSpotSubaccountTransferService) Amount(amount float64) *BrokerSpotSubaccountTransferService {
	s.amount = &amount
	return s
}

// Do sends the request.
func (s *BrokerSpotSubaccountTransferService) Do(ctx context.Context) (*BrokerSpotSubaccountTransfer, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/transfer",
		secType:  secTypeSigned,
	}
	r.setParam("asset", *s.asset)
	r.setParam("amount", *s.amount)
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.toId != nil {
		r.setParam("toId", *s.toId)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(BrokerSpotSubaccountTransfer)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type BrokerSpotSubaccountTransfer struct {
	TxnID        int64  `json:"txnId"`
	ClientTranId string `json:"clientTranId"`
}

type BrokerSpotTransferHistoryService struct {
	c             *Client
	fromId        *string
	toId          *string
	clientTranId  *string
	showAllStatus *bool
	startTime     *int64
	endTime       *int64
	page          *int
	limit         *int
}

// BrokerSpotTransferHistoryService set fromId
func (s *BrokerSpotTransferHistoryService) FromID(fromId string) *BrokerSpotTransferHistoryService {
	s.fromId = &fromId
	return s
}

// BrokerSpotTransferHistoryService set toId
func (s *BrokerSpotTransferHistoryService) ToId(toId string) *BrokerSpotTransferHistoryService {
	s.toId = &toId
	return s
}

// BrokerSpotTransferHistoryService set clientTranId
func (s *BrokerSpotTransferHistoryService) ClientTraniD(clientTranId string) *BrokerSpotTransferHistoryService {
	s.clientTranId = &clientTranId
	return s
}

// BrokerSpotTransferHistoryService set showAllStatus
func (s *BrokerSpotTransferHistoryService) ShowAllStatus(showAllStatus bool) *BrokerSpotTransferHistoryService {
	s.showAllStatus = &showAllStatus
	return s
}

// BrokerSpotTransferHistoryService set startTime
func (s *BrokerSpotTransferHistoryService) StartTime(startTime int64) *BrokerSpotTransferHistoryService {
	s.startTime = &startTime
	return s
}

// BrokerSpotTransferHistoryService set page
func (s *BrokerSpotTransferHistoryService) Page(page int) *BrokerSpotTransferHistoryService {
	s.page = &page
	return s
}

// BrokerSpotTransferHistoryService set limit
func (s *BrokerSpotTransferHistoryService) Limit(limit int) *BrokerSpotTransferHistoryService {
	s.limit = &limit
	return s
}

// BrokerSpotTransferHistoryService set endTime
func (s *BrokerSpotTransferHistoryService) EndTime(endTime int64) *BrokerSpotTransferHistoryService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *BrokerSpotTransferHistoryService) Do(ctx context.Context) (*[]BrokerSpotTransferHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/transfer",
		secType:  secTypeSigned,
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.toId != nil {
		r.setParam("toId", *s.toId)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.showAllStatus != nil {
		r.setParam("showAllStatus", *s.showAllStatus)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new([]BrokerSpotTransferHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type BrokerSpotTransferHistory struct {
	FromId       string `json:"fromId"`
	ToId         string `json:"toId"`
	Asset        string `json:"asset"`
	Qty          string `json:"qty"`
	Time         int64  `json:"time"`
	TxnId        string `json:"txnId"`
	ClientTranId string `json:"clientTranId"`
	Status       string `json:"status"`
}
