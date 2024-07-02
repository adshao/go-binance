package binance

import (
	"context"
	"net/http"
)

// TransferToSubAccountService transfer to subaccount
type TransferToSubAccountService struct {
	c       *Client
	toEmail string
	asset   string
	amount  string
}

// ToEmail set toEmail
func (s *TransferToSubAccountService) ToEmail(toEmail string) *TransferToSubAccountService {
	s.toEmail = toEmail
	return s
}

// Asset set asset
func (s *TransferToSubAccountService) Asset(asset string) *TransferToSubAccountService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *TransferToSubAccountService) Amount(amount string) *TransferToSubAccountService {
	s.amount = amount
	return s
}

func (s *TransferToSubAccountService) transferToSubaccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"toEmail": s.toEmail,
		"asset":   s.asset,
		"amount":  s.amount,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *TransferToSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *TransferToSubAccountResponse, err error) {
	data, err := s.transferToSubaccount(ctx, "/sapi/v1/sub-account/transfer/subToSub", opts...)
	if err != nil {
		return nil, err
	}
	res = &TransferToSubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TransferToSubAccountResponse define transfer to subaccount response
type TransferToSubAccountResponse struct {
	TxnID int64 `json:"txnId"`
}

type SubaccountDepositAddressService struct {
	c       *Client
	email   string
	coin    string
	network string
}

// Email set email
func (s *SubaccountDepositAddressService) Email(email string) *SubaccountDepositAddressService {
	s.email = email
	return s
}

// Coin set coin
func (s *SubaccountDepositAddressService) Coin(coin string) *SubaccountDepositAddressService {
	s.coin = coin
	return s
}

// Network set network
func (s *SubaccountDepositAddressService) Network(network string) *SubaccountDepositAddressService {
	s.network = network
	return s
}

func (s *SubaccountDepositAddressService) subaccountDepositAddress(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email":   s.email,
		"coin":    s.coin,
		"network": s.network,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountDepositAddressResponse, err error) {
	data, err := s.subaccountDepositAddress(ctx, "/sapi/v1/capital/deposit/subAddress", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountDepositAddressResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubaccountDepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	URL     string `json:"url"`
}

type SubaccountAssetsService struct {
	c     *Client
	email string
}

// Email set email
func (s *SubaccountAssetsService) Email(email string) *SubaccountAssetsService {
	s.email = email
	return s
}

func (s *SubaccountAssetsService) subaccountAssets(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email": s.email,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountAssetsService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountAssetsResponse, err error) {
	data, err := s.subaccountAssets(ctx, "/sapi/v3/sub-account/assets", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountAssetsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubaccountAssetsResponse Query Sub-account Assets response
type SubaccountAssetsResponse struct {
	Balances []AssetBalance `json:"balances"`
}

type AssetBalance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free"`
	Locked float64 `json:"locked"`
}

type SubaccountSpotSummaryService struct {
	c     *Client
	email *string
	page  *int32
	size  *int32
}

// Email set email
func (s *SubaccountSpotSummaryService) Email(email string) *SubaccountSpotSummaryService {
	s.email = &email
	return s
}

func (s *SubaccountSpotSummaryService) Page(page int32) *SubaccountSpotSummaryService {
	s.page = &page
	return s
}

func (s *SubaccountSpotSummaryService) Size(size int32) *SubaccountSpotSummaryService {
	s.size = &size
	return s
}

func (s *SubaccountSpotSummaryService) subaccountSpotSummary(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	if s.size != nil {
		r.setParam("size", *s.size)
	}

	if s.page != nil {
		r.setParam("page", *s.page)
	}

	if s.email != nil {
		r.setParam("email", *s.email)
	}
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountSpotSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountSpotSummaryResponse, err error) {
	data, err := s.subaccountSpotSummary(ctx, "/sapi/v1/sub-account/spotSummary", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountSpotSummaryResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubaccountSpotSummaryResponse Query Sub-account Spot Assets Summary response
type SubaccountSpotSummaryResponse struct {
	TotalCount                int64                       `json:"totalCount"`
	MasterAccountTotalAsset   string                      `json:"masterAccountTotalAsset"`
	SpotSubUserAssetBtcVoList []SpotSubUserAssetBtcVoList `json:"spotSubUserAssetBtcVoList"`
}

type SpotSubUserAssetBtcVoList struct {
	Email      string `json:"email"`
	TotalAsset string `json:"totalAsset"`
}

// SubAccountListService Query Sub-account List (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#query-sub-account-list-for-master-account
type SubAccountListService struct {
	c           *Client
	email       *string
	isFreeze    bool
	page, limit int
}

func (s *SubAccountListService) Email(v string) *SubAccountListService {
	s.email = &v
	return s
}

func (s *SubAccountListService) IsFreeze(v bool) *SubAccountListService {
	s.isFreeze = v
	return s
}

func (s *SubAccountListService) Page(v int) *SubAccountListService {
	s.page = v
	return s
}

func (s *SubAccountListService) Limit(v int) *SubAccountListService {
	s.limit = v
	return s
}

func (s *SubAccountListService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountList, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/list",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	if s.isFreeze {
		r.setParam("isFreeze", "true")
	} else {
		r.setParam("isFreeze", "false")
	}
	if s.page > 0 {
		r.setParam("page", s.page)
	}
	if s.limit > 200 {
		r.setParam("limit", 200)
	} else if s.limit <= 0 {
		r.setParam("limit", 10)
	} else {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountList)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountList struct {
	SubAccounts []SubAccount `json:"subAccounts"`
}

type SubAccount struct {
	Email                       string `json:"email"`
	IsFreeze                    bool   `json:"isFreeze"`
	CreateTime                  uint64 `json:"createTime"`
	IsManagedSubAccount         bool   `json:"isManagedSubAccount"`
	IsAssetManagementSubAccount bool   `json:"isAssetManagementSubAccount"`
}

// ManagedSubAccountDepositService
// Deposit Assets Into The Managed Sub-account（For Investor Master Account）
// https://binance-docs.github.io/apidocs/spot/en/#deposit-assets-into-the-managed-sub-account-for-investor-master-account
type ManagedSubAccountDepositService struct {
	c       *Client
	toEmail string
	asset   string
	amount  float64
}

func (s *ManagedSubAccountDepositService) ToEmail(email string) *ManagedSubAccountDepositService {
	s.toEmail = email
	return s
}

func (s *ManagedSubAccountDepositService) Asset(asset string) *ManagedSubAccountDepositService {
	s.asset = asset
	return s
}

func (s *ManagedSubAccountDepositService) Amount(amount float64) *ManagedSubAccountDepositService {
	s.amount = amount
	return s
}

type ManagedSubAccountDepositResponse struct {
	ID int64 `json:"tranId"`
}

// Do send request
func (s *ManagedSubAccountDepositService) Do(ctx context.Context, opts ...RequestOption) (*ManagedSubAccountDepositResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/managed-subaccount/deposit",
		secType:  secTypeSigned,
	}

	r.setParam("toEmail", s.toEmail)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &ManagedSubAccountDepositResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ManagedSubAccountWithdrawalService
// Withdrawal Assets From The Managed Sub-account（For Investor Master Account）
// https://binance-docs.github.io/apidocs/spot/en/#withdrawl-assets-from-the-managed-sub-account-for-investor-master-account
type ManagedSubAccountWithdrawalService struct {
	c            *Client
	fromEmail    string
	asset        string
	amount       float64
	transferDate int64 // Withdrawals is automatically occur on the transfer date(UTC0). If a date is not selected, the withdrawal occurs right now
}

func (s *ManagedSubAccountWithdrawalService) FromEmail(email string) *ManagedSubAccountWithdrawalService {
	s.fromEmail = email
	return s
}

func (s *ManagedSubAccountWithdrawalService) Asset(asset string) *ManagedSubAccountWithdrawalService {
	s.asset = asset
	return s
}

func (s *ManagedSubAccountWithdrawalService) Amount(amount float64) *ManagedSubAccountWithdrawalService {
	s.amount = amount
	return s
}

func (s *ManagedSubAccountWithdrawalService) TransferDate(val int64) *ManagedSubAccountWithdrawalService {
	s.transferDate = val
	return s
}

type ManagedSubAccountWithdrawalResponse struct {
	ID int64 `json:"tranId"`
}

// Do send request
func (s *ManagedSubAccountWithdrawalService) Do(ctx context.Context, opts ...RequestOption) (*ManagedSubAccountWithdrawalResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/managed-subaccount/withdraw",
		secType:  secTypeSigned,
	}

	r.setParam("fromEmail", s.fromEmail)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	if s.transferDate > 0 {
		r.setParam("transferDate", s.transferDate)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &ManagedSubAccountWithdrawalResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ManagedSubAccountAssetsService
// Query Managed Sub-account Asset Details（For Investor Master Account）
// https://binance-docs.github.io/apidocs/spot/en/#query-managed-sub-account-asset-details-for-investor-master-account
type ManagedSubAccountAssetsService struct {
	c     *Client
	email string
}

func (s *ManagedSubAccountAssetsService) Email(email string) *ManagedSubAccountAssetsService {
	s.email = email
	return s
}

type ManagedSubAccountAsset struct {
	Coin             string `json:"coin"`
	Name             string `json:"name"`
	TotalBalance     string `json:"totalBalance"`
	AvailableBalance string `json:"availableBalance"`
	InOrder          string `json:"inOrder"`
	BtcValue         string `json:"btcValue"`
}

func (s *ManagedSubAccountAssetsService) Do(ctx context.Context, opts ...RequestOption) ([]*ManagedSubAccountAsset, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/managed-subaccount/asset",
		secType:  secTypeSigned,
	}

	r.setParam("email", s.email)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*ManagedSubAccountAsset, 0)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// SubAccountFuturesAccountService Get Detail on Sub-account's Futures Account (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#get-detail-on-sub-account-39-s-futures-account-for-master-account
type SubAccountFuturesAccountService struct {
	c     *Client
	email *string
}

func (s *SubAccountFuturesAccountService) Email(v string) *SubAccountFuturesAccountService {
	s.email = &v
	return s
}

func (s *SubAccountFuturesAccountService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesAccount, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/futures/account",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountFuturesAccount struct {
	Email                       string                          `json:"email"`
	Asset                       string                          `json:"asset"`
	Assets                      []SubAccountFuturesAccountAsset `json:"assets"`
	CanDeposit                  bool                            `json:"canDeposit"`
	CanTrade                    bool                            `json:"canTrade"`
	CanWithdraw                 bool                            `json:"canWithdraw"`
	FeeTier                     int                             `json:"feeTier"`
	MaxWithdrawAmount           string                          `json:"maxWithdrawAmount"`
	TotalInitialMargin          string                          `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string                          `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string                          `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string                          `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string                          `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string                          `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string                          `json:"totalWalletBalance"`
	UpdateTime                  int64                           `json:"updateTime"`
}

type SubAccountFuturesAccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}

// SubaccountFuturesSummaryV1Service Get Summary of Sub-account's Futures Account (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#get-summary-of-sub-account-39-s-futures-account-for-master-account
type SubAccountFuturesSummaryV1Service struct {
	c *Client
}

func (s *SubAccountFuturesSummaryV1Service) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesSummaryV1, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/futures/accountSummary",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesSummaryV1)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountFuturesSummaryCommon struct {
	Asset                       string `json:"asset"`
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
}

type SubAccountFuturesSummaryV1 struct {
	SubAccountFuturesSummaryCommon
	SubAccountList []SubAccountFuturesSummaryV1SubAccountList `json:"subAccountList"`
}

type SubAccountFuturesSummaryV1SubAccountList struct {
	Email string `json:"email"`
	SubAccountFuturesSummaryCommon
}

// SubAccountFuturesTransferV1Service Futures Transfer for Sub-account (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#futures-transfer-for-sub-account-for-master-account
type SubAccountFuturesTransferV1Service struct {
	c      *Client
	email  string
	asset  string
	amount float64
	/*
		1: transfer from subaccount's spot account to its USDT-margined futures account
		2: transfer from subaccount's USDT-margined futures account to its spot account
		3: transfer from subaccount's spot account to its COIN-margined futures account
		4:transfer from subaccount's COIN-margined futures account to its spot account
	*/
	transferType int
}

func (s *SubAccountFuturesTransferV1Service) Email(v string) *SubAccountFuturesTransferV1Service {
	s.email = v
	return s
}

func (s *SubAccountFuturesTransferV1Service) Asset(v string) *SubAccountFuturesTransferV1Service {
	s.asset = v
	return s
}

func (s *SubAccountFuturesTransferV1Service) Amount(v float64) *SubAccountFuturesTransferV1Service {
	s.amount = v
	return s
}

func (s *SubAccountFuturesTransferV1Service) TransferType(v int) *SubAccountFuturesTransferV1Service {
	s.transferType = v
	return s
}

func (s *SubAccountFuturesTransferV1Service) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesTransferResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/futures/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"email":  s.email,
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountFuturesTransferResponse struct {
	// seems api doc bug, return `tranId` as int64 actually in production environment
	TranID int64 `json:"tranId"`
}

// Sub-account Transfer History (For Sub-account)
// https://binance-docs.github.io/apidocs/spot/en/#sub-account-transfer-history-for-sub-account
type SubAccountTransferHistoryService struct {
	c                 *Client
	asset             *string
	transferType      *SubAccountTransferType
	startTime         *int64
	endTime           *int64
	limit             *int
	returnFailHistory *bool
}

func (s *SubAccountTransferHistoryService) Asset(v string) *SubAccountTransferHistoryService {
	s.asset = &v
	return s
}

func (s *SubAccountTransferHistoryService) TransferType(v SubAccountTransferType) *SubAccountTransferHistoryService {
	s.transferType = &v
	return s
}

func (s *SubAccountTransferHistoryService) StartTime(v int64) *SubAccountTransferHistoryService {
	s.startTime = &v
	return s
}

func (s *SubAccountTransferHistoryService) EndTime(v int64) *SubAccountTransferHistoryService {
	s.endTime = &v
	return s
}

func (s *SubAccountTransferHistoryService) Limit(v int) *SubAccountTransferHistoryService {
	s.limit = &v
	return s
}

func (s *SubAccountTransferHistoryService) ReturnFailHistory(v bool) *SubAccountTransferHistoryService {
	s.returnFailHistory = &v
	return s
}

func (s *SubAccountTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccountTransferHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/transfer/subUserHistory",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.transferType != nil {
		r.setParam("type", *s.transferType)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.returnFailHistory != nil {
		r.setParam("returnFailHistory", *s.returnFailHistory)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccountTransferHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountTransferHistory struct {
	CounterParty    string                 `json:"counterParty"`
	Email           string                 `json:"email"`
	Type            SubAccountTransferType `json:"type"`
	Asset           string                 `json:"asset"`
	Qty             string                 `json:"qty"`
	FromAccountType AccountType            `json:"fromAccountType"`
	ToAccountType   AccountType            `json:"toAccountType"`
	Status          string                 `json:"status"`
	TranID          int64                  `json:"tranId"`
	Time            int64                  `json:"time"`
}

// Create virtual sub-account
type CrtVirtSubAccService struct {
	c                *Client
	subAccountString string
	recvWindow       *int64
}

type CrtVirtSubAccRsp struct {
	Email string `json:"email"`
}

func (s *CrtVirtSubAccService) SubAccountString(subAccountString string) *CrtVirtSubAccService {
	s.subAccountString = subAccountString
	return s
}

func (s *CrtVirtSubAccService) RecvWindow(recvWindow int64) *CrtVirtSubAccService {
	s.recvWindow = &recvWindow
	return s
}

func (s *CrtVirtSubAccService) Do(ctx context.Context, opts ...RequestOption) (res *CrtVirtSubAccRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/virtualSubAccount",
		secType:  secTypeSigned,
	}
	r.setFormParam("subAccountString", s.subAccountString)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CrtVirtSubAccRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Sub-Account spot transfer history
type SubAccSpotTrfHisService struct {
	c          *Client
	fromEmail  *string // Sub-account email
	toEmail    *string // true or false
	startTime  *uint64
	endTime    *uint64
	page       *int32 // default 1
	limit      *int32 // default 1, max 200
	recvWindow *int64
}

type SubAccSpotTrf struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Asset  string `json:"asset"`
	Qty    string `json:"qty"`
	Status string `json:"status"`
	TranId uint64 `json:"tranId"`
	Time   uint64 `json:"time"`
}

func (s *SubAccSpotTrfHisService) FromEmail(fromEmail string) *SubAccSpotTrfHisService {
	s.fromEmail = &fromEmail
	return s
}

func (s *SubAccSpotTrfHisService) ToEmail(toEmail string) *SubAccSpotTrfHisService {
	s.toEmail = &toEmail
	return s
}

func (s *SubAccSpotTrfHisService) StartTime(startTime uint64) *SubAccSpotTrfHisService {
	s.startTime = &startTime
	return s
}

func (s *SubAccSpotTrfHisService) EndTime(endTime uint64) *SubAccSpotTrfHisService {
	s.endTime = &endTime
	return s
}

func (s *SubAccSpotTrfHisService) Page(page int32) *SubAccSpotTrfHisService {
	s.page = &page
	return s
}

func (s *SubAccSpotTrfHisService) Limit(limit int32) *SubAccSpotTrfHisService {
	s.limit = &limit
	return s
}

func (s *SubAccSpotTrfHisService) RecvWindow(recvWindow int64) *SubAccSpotTrfHisService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccSpotTrfHisService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccSpotTrf, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/sub/transfer/history",
		secType:  secTypeSigned,
	}
	if s.fromEmail != nil {
		r.setParam("fromEmail", *s.fromEmail)
	}
	if s.toEmail != nil {
		r.setParam("toEmail", *s.toEmail)
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
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccSpotTrf, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Sub-Account futures transfer history
type SubAccFutTrfHisService struct {
	c           *Client
	email       string // Sub-account email
	futuresType int64  // 1: usdt 2: coin
	startTime   *int64
	endTime     *int64
	page        *int32 // default 1
	limit       *int32 // default 50, max 500
	recvWindow  *int64
}

type SubAccFutTrfHisRsp struct {
	Success     bool            `json:"success"`
	FuturesType int32           `json:"futuresType"`
	Transfers   []*SubAccFutTrf `json:"transfers"`
}

type SubAccFutTrf struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Asset  string `json:"asset"`
	Qty    string `json:"qty"`
	Status string `json:"status"`
	TranId uint64 `json:"tranId"`
	Time   uint64 `json:"time"`
}

func (s *SubAccFutTrfHisService) Email(email string) *SubAccFutTrfHisService {
	s.email = email
	return s
}

func (s *SubAccFutTrfHisService) FuturesType(futuresType int64) *SubAccFutTrfHisService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccFutTrfHisService) StartTime(startTime int64) *SubAccFutTrfHisService {
	s.startTime = &startTime
	return s
}

func (s *SubAccFutTrfHisService) EndTime(endTime int64) *SubAccFutTrfHisService {
	s.endTime = &endTime
	return s
}

func (s *SubAccFutTrfHisService) Page(page int32) *SubAccFutTrfHisService {
	s.page = &page
	return s
}

func (s *SubAccFutTrfHisService) Limit(limit int32) *SubAccFutTrfHisService {
	s.limit = &limit
	return s
}

func (s *SubAccFutTrfHisService) RecvWindow(recvWindow int64) *SubAccFutTrfHisService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccFutTrfHisService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccFutTrfHisRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/futures/internalTransfer",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("futuresType", s.futuresType)
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
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccFutTrfHisRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Execute sub account futures balance transfer
type SubAccFutTrfService struct {
	c           *Client
	fromEmail   string // sender email
	toEmail     string // receiver email
	futuresType int64  // 1: usdt 2: coin
	asset       string
	amount      string // decimal format
	recvWindow  *int64
}

type SubAccFutTrfRsp struct {
	Success bool   `json:"success"`
	TxnId   string `json:"txnId"`
}

func (s *SubAccFutTrfService) FromEmail(fromEmail string) *SubAccFutTrfService {
	s.fromEmail = fromEmail
	return s
}

func (s *SubAccFutTrfService) ToEmail(toEmail string) *SubAccFutTrfService {
	s.toEmail = toEmail
	return s
}

func (s *SubAccFutTrfService) FuturesType(futuresType int64) *SubAccFutTrfService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccFutTrfService) Asset(asset string) *SubAccFutTrfService {
	s.asset = asset
	return s
}

func (s *SubAccFutTrfService) Amount(amount string) *SubAccFutTrfService {
	s.amount = amount
	return s
}

func (s *SubAccFutTrfService) RecvWindow(recvWindow int64) *SubAccFutTrfService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccFutTrfService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccFutTrfRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/futures/internalTransfer",
		secType:  secTypeSigned,
	}
	r.setFormParam("fromEmail", s.fromEmail)
	r.setFormParam("toEmail", s.toEmail)
	r.setFormParam("futuresType", s.futuresType)
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccFutTrfRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get sub account deposit record
type SubAccDepRecService struct {
	c          *Client
	email      string // sub-account email
	coin       *string
	status     *int32 // 0(0:pending,6: credited but cannot withdraw,7:Wrong Deposit,8:Waiting User confirm,1:success)
	startTime  *int64
	endTime    *int64
	limit      *int // sub-account email
	offset     *int // default 0
	txId       *string
	recvWindow *int64
}

type SubAccDepRec struct {
	Id            string `json:"id"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int32  `json:"status"`
	Address       string `json:"address"`
	TxId          string `json:"txId"`
	AddressTag    string `json:"addressTag"`
	InsertTime    int64  `json:"insertTime"`
	TransferType  int32  `json:"transferType"`
	ConfirmTimes  string `json:"confirmTimes"`
	UnlockConfirm int32  `json:"unlockConfirm"`
	WalletType    int32  `json:"walletType"`
}

func (s *SubAccDepRecService) Email(email string) *SubAccDepRecService {
	s.email = email
	return s
}

func (s *SubAccDepRecService) Coin(coin string) *SubAccDepRecService {
	s.coin = &coin
	return s
}

func (s *SubAccDepRecService) Status(status int32) *SubAccDepRecService {
	s.status = &status
	return s
}

func (s *SubAccDepRecService) StartTime(startTime int64) *SubAccDepRecService {
	s.startTime = &startTime
	return s
}

func (s *SubAccDepRecService) EndTime(endTime int64) *SubAccDepRecService {
	s.endTime = &endTime
	return s
}

func (s *SubAccDepRecService) Limit(limit int) *SubAccDepRecService {
	s.limit = &limit
	return s
}

func (s *SubAccDepRecService) Offset(offset int) *SubAccDepRecService {
	s.offset = &offset
	return s
}

func (s *SubAccDepRecService) TxId(txId string) *SubAccDepRecService {
	s.txId = &txId
	return s
}

func (s *SubAccDepRecService) RecvWindow(recvWindow int64) *SubAccDepRecService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccDepRecService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccDepRec, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/subHisrec",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.txId != nil {
		r.setParam("txId", *s.txId)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccDepRec, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get sub account margin futures status
type SubAccMFStatusService struct {
	c          *Client
	email      *string // sub-account email
	recvWindow *int64
}

type SubAccMFStatus struct {
	Email            string `json:"email"`
	IsSubUserEnabled bool   `json:"isSubUserEnabled"`
	IsUserActive     bool   `json:"isUserActive"`
	InsertTime       int64  `json:"insertTime"`
	IsMarginEnabled  bool   `json:"isMarginEnabled"`
	IsFutureEnabled  bool   `json:"isFutureEnabled"`
	Mobile           int64  `json:"mobile"`
}

func (s *SubAccMFStatusService) Email(email string) *SubAccMFStatusService {
	s.email = &email
	return s
}

func (s *SubAccMFStatusService) RecvWindow(recvWindow int64) *SubAccMFStatusService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccMFStatusService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccMFStatus, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/status",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccMFStatus, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// sub account margin enable
type SubAccMarginEnableService struct {
	c          *Client
	email      string // sub-account email
	recvWindow *int64
}

type SubAccMarginEnableRsp struct {
	Email           string `json:"email"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
}

func (s *SubAccMarginEnableService) Email(email string) *SubAccMarginEnableService {
	s.email = email
	return s
}

func (s *SubAccMarginEnableService) RecvWindow(recvWindow int64) *SubAccMarginEnableService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccMarginEnableService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccMarginEnableRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/margin/enable",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccMarginEnableRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get sub-account margin account detail
type SubAccMarginAccService struct {
	c          *Client
	email      string // sub-account email
	recvWindow *int64
}

type SubAccMarginAcc struct {
	Email                 string               `json:"email"`
	MarginLevel           string               `json:"marginLevel"`
	TotalAssetOfBtc       string               `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc   string               `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc    string               `json:"totalNetAssetOfBtc"`
	MarginTradeCoeffVo    *MarginTradeCoeffVo  `json:"marginTradeCoeffVo"`
	MarginUserAssetVoList []*MarginUserAssetVo `json:"marginUserAssetVoList"`
}

type MarginTradeCoeffVo struct {
	ForceLiquidationBar string `json:"forceLiquidationBar"`
	MarginCallBar       string `json:"marginCallBar"`
	NormalBar           string `json:"normalBar"`
}

type MarginUserAssetVo struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

func (s *SubAccMarginAccService) Email(email string) *SubAccMarginAccService {
	s.email = email
	return s
}

func (s *SubAccMarginAccService) RecvWindow(recvWindow int64) *SubAccMarginAccService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccMarginAccService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccMarginAcc, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/margin/account",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccMarginAcc)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get sub-account margin account summary
type SubAccMarginAccSummService struct {
	c          *Client
	recvWindow *int64
}

type SubAccMarginAccSumm struct {
	TotalAssetOfBtc     string         `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string         `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string         `json:"totalNetAssetOfBtc"`
	SubAccountList      []*MSubAccount `json:"subAccountList"`
}

type MSubAccount struct {
	Email               string `json:"email"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
}

func (s *SubAccMarginAccSummService) RecvWindow(recvWindow int64) *SubAccMarginAccSummService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccMarginAccSummService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccMarginAccSumm, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/margin/accountSummary",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccMarginAccSumm)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type SubAccFuturesEnableService struct {
	c          *Client
	email      string // sub-account email
	recvWindow *int64
}

type SubAccFuturesEnableRsp struct {
	Email            string `json:"email"`
	IsFuturesEnabled bool   `json:"isFuturesEnabled"`
}

func (s *SubAccFuturesEnableService) Email(email string) *SubAccFuturesEnableService {
	s.email = email
	return s
}

func (s *SubAccFuturesEnableService) RecvWindow(recvWindow int64) *SubAccFuturesEnableService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccFuturesEnableService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccFuturesEnableRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/futures/enable",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccFuturesEnableRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get the target sub-account futures account detail, v2 interface.
type SubAccFuturesAccService struct {
	c           *Client
	email       string // sub-account email
	futuresType int32  // 1:USDT Margined Futures, 2:COIN Margined Futures
	recvWindow  *int64
}

type SubAccFuturesAccSvcRsp struct {
	FutureAccountResp   *SubAccFuturesAcc  `json:"futureAccountResp"`   // set while futuresType=1(USDT margined)
	DeliveryAccountResp *SubAccDeliveryAcc `json:"deliveryAccountResp"` // set while futuresType=2(COIN margined)
}

type SubAccFuturesAcc struct {
	Email                       string    `json:"email"`
	Asset                       string    `json:"asset"`
	Assets                      []*FAsset `json:"assets"`
	CanDeposit                  bool      `json:"canDeposit"`
	CanTrade                    bool      `json:"canTrade"`
	CanWithdraw                 bool      `json:"canWithdraw"`
	FeeTier                     int32     `json:"feeTier"`
	MaxWithdrawAmount           string    `json:"maxWithdrawAmount"`
	TotalInitialMargin          string    `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string    `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string    `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string    `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string    `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string    `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string    `json:"totalWalletBalance"`
	UpdateTime                  int64     `json:"updateTime"`
}

type SubAccDeliveryAcc struct {
	Email       string    `json:"email"`
	Assets      []*FAsset `json:"assets"`
	CanDeposit  bool      `json:"canDeposit"`
	CanTrade    bool      `json:"canTrade"`
	CanWithdraw bool      `json:"canWithdraw"`
	FeeTier     int32     `json:"feeTier"`
	UpdateTime  int64     `json:"updateTime"`
}

type FAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}

func (s *SubAccFuturesAccService) Email(email string) *SubAccFuturesAccService {
	s.email = email
	return s
}

func (s *SubAccFuturesAccService) FuturesType(futuresType int32) *SubAccFuturesAccService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccFuturesAccService) RecvWindow(recvWindow int64) *SubAccFuturesAccService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccFuturesAccService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccFuturesAccSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v2/sub-account/futures/account",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("futuresType", s.futuresType)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccFuturesAccSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get sub-account futures account summary, include U-M and C-M, v2 interface
type SubAccFuturesAccSummService struct {
	c           *Client
	futuresType int32  // 1:USDT Margined Futures, 2:COIN Margined Futures
	page        *int32 // default 1
	limit       *int32 // default 10, max 20
	recvWindow  *int64
}

type SubAccFuturesAccSummSvcRsp struct {
	FutureAccountSummaryResp   *SubAccFuturesAccSumm  `json:"futureAccountSummaryResp"`   // set while futuresType=1
	DeliveryAccountSummaryResp *SubAccDeliveryAccSumm `json:"deliveryAccountSummaryResp"` // set while futuresType=2
}

type SubAccFuturesAccSumm struct {
	TotalInitialMargin          string         `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string         `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string         `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string         `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string         `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string         `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string         `json:"totalWalletBalance"`
	Asset                       string         `json:"asset"`
	SubAccountList              []*FSubAccount `json:"subAccountList"`
}

type FSubAccount struct {
	Email                       string `json:"email"`
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
	Asset                       string `json:"asset"`
}

type SubAccDeliveryAccSumm struct {
	TotalMarginBalanceOfBTC    string         `json:"totalMarginBalanceOfBTC"`
	TotalUnrealizedProfitOfBTC string         `json:"totalUnrealizedProfitOfBTC"`
	TotalWalletBalanceOfBTC    string         `json:"totalWalletBalanceOfBTC"`
	Asset                      string         `json:"asset"`
	SubAccountList             []*DSubAccount `json:"subAccountList"`
}

type DSubAccount struct {
	Email                 string `json:"email"`
	TotalMarginBalance    string `json:"totalMarginBalance"`
	TotalUnrealizedProfit string `json:"totalUnrealizedProfit"`
	TotalWalletBalance    string `json:"totalWalletBalance"`
	Asset                 string `json:"asset"`
}

func (s *SubAccFuturesAccSummService) FuturesType(futuresType int32) *SubAccFuturesAccSummService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccFuturesAccSummService) Page(page int32) *SubAccFuturesAccSummService {
	s.page = &page
	return s
}

func (s *SubAccFuturesAccSummService) Limit(limit int32) *SubAccFuturesAccSummService {
	s.limit = &limit
	return s
}

func (s *SubAccFuturesAccSummService) RecvWindow(recvWindow int64) *SubAccFuturesAccSummService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccFuturesAccSummService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccFuturesAccSummSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v2/sub-account/futures/accountSummary",
		secType:  secTypeSigned,
	}
	r.setParam("futuresType", s.futuresType)
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccFuturesAccSummSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get target sub-account futures position information, include U-M and C-M, v2 interface.
type SubAccFuturesPositionsService struct {
	c           *Client
	email       string
	futuresType int32 // 1:USDT Margined Futures, 2:COIN Margined Futures
	recvWindow  *int64
}

type SubAccFuturesPosSvcRsp struct {
	FuturePositionRiskVos   []*SubAccFuturesPosition  `json:"futurePositionRiskVos"`   // set while futuresType=1
	DeliveryPositionRiskVos []*SubAccDeliveryPosition `json:"deliveryPositionRiskVos"` // set while futuresType=2
}

type SubAccFuturesPosition struct {
	EntryPrice       string `json:"entryPrice"`
	Leverage         string `json:"leverage"`
	MaxNotional      string `json:"maxNotional"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmount   string `json:"positionAmount"`
	Symbol           string `json:"symbol"`
	UnrealizedProfit string `json:"unrealizedProfit"`
}

type SubAccDeliveryPosition struct {
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	Leverage         string `json:"leverage"`
	Isolated         string `json:"isolated"`
	IsolatedWallet   string `json:"isolatedWallet"`
	IsolatedMargin   string `json:"isolatedMargin"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	PositionSide     string `json:"positionSide"`
	PositionAmount   string `json:"positionAmount"`
	Symbol           string `json:"symbol"`
	UnrealizedProfit string `json:"unrealizedProfit"`
}

func (s *SubAccFuturesPositionsService) Email(email string) *SubAccFuturesPositionsService {
	s.email = email
	return s
}

func (s *SubAccFuturesPositionsService) FuturesType(futuresType int32) *SubAccFuturesPositionsService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccFuturesPositionsService) RecvWindow(recvWindow int64) *SubAccFuturesPositionsService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccFuturesPositionsService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccFuturesPosSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v2/sub-account/futures/positionRisk",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("futuresType", s.futuresType)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccFuturesPosSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// execute sub-account margin account transfer
type SubAccMarginTrfService struct {
	c            *Client
	email        string
	asset        string
	amount       string // decimal format
	transferType int32  // 1: Transfer from the spot account of the sub account to its leverage account; 2: Transfer from the leverage account of the sub account to its spot account
	recvWindow   *int64
}

type SubAccMarginTrfRsp struct {
	TxnId string `json:"txnId"`
}

func (s *SubAccMarginTrfService) Email(email string) *SubAccMarginTrfService {
	s.email = email
	return s
}

func (s *SubAccMarginTrfService) Asset(asset string) *SubAccMarginTrfService {
	s.asset = asset
	return s
}

func (s *SubAccMarginTrfService) Amount(amount string) *SubAccMarginTrfService {
	s.amount = amount
	return s
}

func (s *SubAccMarginTrfService) TransferType(transferType int32) *SubAccMarginTrfService {
	s.transferType = transferType
	return s
}

func (s *SubAccMarginTrfService) RecvWindow(recvWindow int64) *SubAccMarginTrfService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccMarginTrfService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccMarginTrfRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/margin/transfer",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	r.setFormParam("type", s.transferType)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccMarginTrfRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// sub-account transfer balance to master-account
type SubAccTrfSubToMasterService struct {
	c          *Client
	asset      string
	amount     string // decimal format
	recvWindow *int64
}

type SubAccTrfSubToMasterRsp struct {
	TxnId string `json:"txnId"`
}

func (s *SubAccTrfSubToMasterService) Asset(asset string) *SubAccTrfSubToMasterService {
	s.asset = asset
	return s
}

func (s *SubAccTrfSubToMasterService) Amount(amount string) *SubAccTrfSubToMasterService {
	s.amount = amount
	return s
}

func (s *SubAccTrfSubToMasterService) RecvWindow(recvWindow int64) *SubAccTrfSubToMasterService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccTrfSubToMasterService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccTrfSubToMasterRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/transfer/subToMaster",
		secType:  secTypeSigned,
	}
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccTrfSubToMasterRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Universal transfer of master and sub accounts
type SubAccUnivTrfService struct {
	c               *Client
	fromEmail       *string
	toEmail         *string
	fromAccountType string  // "SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	toAccountType   string  // "SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	clientTranId    *string // Non repeatable
	symbol          *string // Only used under ISOLATEd_MARGIN type
	asset           string
	amount          string // decimal format
	recvWindow      *int64
}

type SubAccUnivTrfRsp struct {
	TranId       int64  `json:"tranId"`
	ClientTranId string `json:"clientTranId"`
}

func (s *SubAccUnivTrfService) FromEmail(fromEmail string) *SubAccUnivTrfService {
	s.fromEmail = &fromEmail
	return s
}

func (s *SubAccUnivTrfService) ToEmail(toEmail string) *SubAccUnivTrfService {
	s.toEmail = &toEmail
	return s
}

func (s *SubAccUnivTrfService) FromAccountType(fromAccountType string) *SubAccUnivTrfService {
	s.fromAccountType = fromAccountType
	return s
}

func (s *SubAccUnivTrfService) ToAccountType(toAccountType string) *SubAccUnivTrfService {
	s.toAccountType = toAccountType
	return s
}

func (s *SubAccUnivTrfService) ClientTranId(clientTranId string) *SubAccUnivTrfService {
	s.clientTranId = &clientTranId
	return s
}

func (s *SubAccUnivTrfService) Symbol(symbol string) *SubAccUnivTrfService {
	s.symbol = &symbol
	return s
}

func (s *SubAccUnivTrfService) Asset(asset string) *SubAccUnivTrfService {
	s.asset = asset
	return s
}

func (s *SubAccUnivTrfService) Amount(amount string) *SubAccUnivTrfService {
	s.amount = amount
	return s
}

func (s *SubAccUnivTrfService) RecvWindow(recvWindow int64) *SubAccUnivTrfService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccUnivTrfService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccUnivTrfRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}
	if s.fromEmail != nil {
		r.setFormParam("fromEmail", *s.fromEmail)
	}
	if s.toEmail != nil {
		r.setFormParam("toEmail", *s.toEmail)
	}
	r.setFormParam("fromAccountType", s.fromAccountType)
	r.setFormParam("toAccountType", s.toAccountType)
	if s.clientTranId != nil {
		r.setFormParam("clientTranId", *s.clientTranId)
	}
	if s.symbol != nil {
		r.setFormParam("symbol", *s.symbol)
	}
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccUnivTrfRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query the universal transfer history of sub and master accounts
type SubAccUnivTrfHisService struct {
	c            *Client
	fromEmail    *string
	toEmail      *string
	clientTranId *string // Non repeatable
	startTime    *int64
	endTime      *int64
	page         *int32 // default 1
	limit        *int32 // default 500, max 500
	recvWindow   *int64
}

type SubAccUnivTrfHisServiceRsp struct {
	Result     []*SubAccUnivTrfRec `json:"result"`
	TotalCount int64               `json:"totalCount"`
}

type SubAccUnivTrfRec struct {
	TranId          int64  `json:"tranId"`
	FromEmail       string `json:"fromEmail"`
	ToEmail         string `json:"toEmail"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	CreateTimeStamp int64  `json:"createTimeStamp"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	ClientTranId    string `json:"clientTranId"`
}

func (s *SubAccUnivTrfHisService) FromEmail(fromEmail string) *SubAccUnivTrfHisService {
	s.fromEmail = &fromEmail
	return s
}

func (s *SubAccUnivTrfHisService) ToEmail(toEmail string) *SubAccUnivTrfHisService {
	s.toEmail = &toEmail
	return s
}

func (s *SubAccUnivTrfHisService) ClientTranId(clientTranId string) *SubAccUnivTrfHisService {
	s.clientTranId = &clientTranId
	return s
}

func (s *SubAccUnivTrfHisService) StartTime(startTime int64) *SubAccUnivTrfHisService {
	s.startTime = &startTime
	return s
}

func (s *SubAccUnivTrfHisService) EndTime(endTime int64) *SubAccUnivTrfHisService {
	s.endTime = &endTime
	return s
}

func (s *SubAccUnivTrfHisService) Page(page int32) *SubAccUnivTrfHisService {
	s.page = &page
	return s
}

func (s *SubAccUnivTrfHisService) Limit(limit int32) *SubAccUnivTrfHisService {
	s.limit = &limit
	return s
}

func (s *SubAccUnivTrfHisService) RecvWindow(recvWindow int64) *SubAccUnivTrfHisService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccUnivTrfHisService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccUnivTrfHisServiceRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}
	if s.fromEmail != nil {
		r.setParam("fromEmail", *s.fromEmail)
	}
	if s.toEmail != nil {
		r.setParam("toEmail", *s.toEmail)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
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
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccUnivTrfHisServiceRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Binance Leveraged Tokens enable
type SubAccBlvtEnableService struct {
	c          *Client
	email      string // sub-account email
	enableBlvt bool   // Only true for now
	recvWindow *int64
}

type SubAccBlvtEnableSvcRsp struct {
	Email      string `json:"email"`
	EnableBlvt bool   `json:"enableBlvt"`
}

func (s *SubAccBlvtEnableService) Email(email string) *SubAccBlvtEnableService {
	s.email = email
	return s
}

func (s *SubAccBlvtEnableService) EnableBlvt(enableBlvt bool) *SubAccBlvtEnableService {
	s.enableBlvt = enableBlvt
	return s
}

func (s *SubAccBlvtEnableService) RecvWindow(recvWindow int64) *SubAccBlvtEnableService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccBlvtEnableService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccBlvtEnableSvcRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/blvt/enable",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	r.setFormParam("enableBlvt", s.enableBlvt)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccBlvtEnableSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// query sub-account api ip restriction
type SubAccApiIpRestrictionService struct {
	c                *Client
	email            string // sub-account email
	subAccountApiKey string
	recvWindow       *int64
}

type SubAccApiIpRestrictSvcRsp struct {
	IpRestrict string   `json:"ipRestrict"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}

func (s *SubAccApiIpRestrictionService) Email(email string) *SubAccApiIpRestrictionService {
	s.email = email
	return s
}

func (s *SubAccApiIpRestrictionService) SubAccountApiKey(subAccountApiKey string) *SubAccApiIpRestrictionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *SubAccApiIpRestrictionService) RecvWindow(recvWindow int64) *SubAccApiIpRestrictionService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccApiIpRestrictionService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccApiIpRestrictSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/subAccountApi/ipRestriction",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("subAccountApiKey", s.subAccountApiKey)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccApiIpRestrictSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// delete sub-account ip restriction
type SubAccApiDelIpRestrictionService struct {
	c                *Client
	email            string // sub-account email
	subAccountApiKey string
	ipAddress        *string // Can be deleted in batches, separated by commas
	recvWindow       *int64
}

type SubAccApiDelIpRestrictSvcRsp struct {
	IpRestrict string   `json:"ipRestrict"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}

func (s *SubAccApiDelIpRestrictionService) Email(email string) *SubAccApiDelIpRestrictionService {
	s.email = email
	return s
}

func (s *SubAccApiDelIpRestrictionService) SubAccountApiKey(subAccountApiKey string) *SubAccApiDelIpRestrictionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *SubAccApiDelIpRestrictionService) IpAddress(ipAddress string) *SubAccApiDelIpRestrictionService {
	s.ipAddress = &ipAddress
	return s
}

func (s *SubAccApiDelIpRestrictionService) RecvWindow(recvWindow int64) *SubAccApiDelIpRestrictionService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccApiDelIpRestrictionService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccApiDelIpRestrictSvcRsp, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/sub-account/subAccountApi/ipRestriction/ipList",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	r.setFormParam("subAccountApiKey", s.subAccountApiKey)
	if s.ipAddress != nil {
		r.setFormParam("ipAddress", *s.ipAddress)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccApiDelIpRestrictSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// add sub-account ip restriction
type SubAccApiAddIpRestrictionService struct {
	c                *Client
	email            string // sub-account email
	subAccountApiKey string
	status           string  // IP restriction status. 1 or not filled in (null)=IP not restricted. 2=Access is limited to trusted IPs only.
	ipAddress        *string // IP can be filled in bulk, separated by commas
	recvWindow       *int64
}

type SubAccApiAddIpRestrictSvcRsp struct {
	IpRestrict string   `json:"ipRestrict"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}

func (s *SubAccApiAddIpRestrictionService) Email(email string) *SubAccApiAddIpRestrictionService {
	s.email = email
	return s
}

func (s *SubAccApiAddIpRestrictionService) SubAccountApiKey(subAccountApiKey string) *SubAccApiAddIpRestrictionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *SubAccApiAddIpRestrictionService) Status(status string) *SubAccApiAddIpRestrictionService {
	s.status = status
	return s
}

func (s *SubAccApiAddIpRestrictionService) IpAddress(ipAddress string) *SubAccApiAddIpRestrictionService {
	s.ipAddress = &ipAddress
	return s
}

func (s *SubAccApiAddIpRestrictionService) RecvWindow(recvWindow int64) *SubAccApiAddIpRestrictionService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccApiAddIpRestrictionService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccApiAddIpRestrictSvcRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/subAccountApi/ipRestriction",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	r.setFormParam("subAccountApiKey", s.subAccountApiKey)
	r.setFormParam("status", s.status)
	if s.ipAddress != nil {
		r.setFormParam("ipAddress", *s.ipAddress)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccApiAddIpRestrictSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MngSubAccWithdrawService struct {
	c            *Client
	fromEmail    string
	asset        string
	amount       string // decimal format
	transferDate *int64 // The withdrawal will automatically occur on the selected date (UTC0). If no date is selected, the withdrawal will take effect immediately
	recvWindow   *int64
}

type MngSubAccWithdrawSvcRsp struct {
	TranId int64 `json:"tranId"`
}

func (s *MngSubAccWithdrawService) FromEmail(fromEmail string) *MngSubAccWithdrawService {
	s.fromEmail = fromEmail
	return s
}

func (s *MngSubAccWithdrawService) Asset(asset string) *MngSubAccWithdrawService {
	s.asset = asset
	return s
}

func (s *MngSubAccWithdrawService) Amount(amount string) *MngSubAccWithdrawService {
	s.amount = amount
	return s
}

func (s *MngSubAccWithdrawService) TransferDate(transferDate int64) *MngSubAccWithdrawService {
	s.transferDate = &transferDate
	return s
}

func (s *MngSubAccWithdrawService) RecvWindow(recvWindow int64) *MngSubAccWithdrawService {
	s.recvWindow = &recvWindow
	return s
}

func (s *MngSubAccWithdrawService) Do(ctx context.Context, opts ...RequestOption) (res *MngSubAccWithdrawSvcRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/managed-subaccount/withdraw",
		secType:  secTypeSigned,
	}
	r.setFormParam("fromEmail", s.fromEmail)
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	if s.transferDate != nil {
		r.setFormParam("transferDate", *s.transferDate)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MngSubAccWithdrawSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query asset snapshot of managed-sub account
type MngSubAccSnapshotService struct {
	c          *Client
	email      string
	accType    string // SPOT (spot), MARGIN (full position), FUTURES (U-based contract)
	startTime  *int64
	endTime    *int64
	limit      *int32 // min 7, max 30, default 7
	recvWindow *int64
}

type MngSubAccSnapshotSvcRsp struct {
	Code        int64         `json:"code"`
	Msg         string        `json:"msg"`
	SnapshotVos []*SnapshotVo `json:"snapshotVos"`
}

type SnapshotVo struct {
	Data       SnapshotVoData `json:"data"`
	Type       string         `json:"type"`
	UpdateTime int64          `json:"updateTime"`
}

type SnapshotVoData struct {
	Balances            []*SnapShotSpotBalance `json:"balances"`            // set while SnapshotVo.type=spot
	TotalAssetOfBtc     string                 `json:"totalAssetOfBtc"`     // set while SnapshotVo.type is one of spot and margin
	MarginLevel         string                 `json:"marginLevel"`         // set while SnapshotVo.type=margin
	TotalLiabilityOfBtc string                 `json:"totalLiabilityOfBtc"` // set while SnapshotVo.type=margin
	TotalNetAssetOfBtc  string                 `json:"totalNetAssetOfBtc"`  // set while SnapshotVo.type=margin
	UserAssets          []*MarginUserAsset     `json:"userAssets"`          // set while SnapshotVo.type=margin
	Assets              []*FuturesAsset        `json:"assets"`              // set while SnapshotVo.type=futures
	Position            []*FuturesPosition     `json:"position"`            // set while SnapshotVo.type=futures
}

type SnapShotSpotBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type MarginUserAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

type FuturesAsset struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	WalletBalance string `json:"walletBalance"`
}

type FuturesPosition struct {
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
}

func (s *MngSubAccSnapshotService) Email(email string) *MngSubAccSnapshotService {
	s.email = email
	return s
}

func (s *MngSubAccSnapshotService) AccType(accType string) *MngSubAccSnapshotService {
	s.accType = accType
	return s
}

func (s *MngSubAccSnapshotService) StartTime(startTime int64) *MngSubAccSnapshotService {
	s.startTime = &startTime
	return s
}

func (s *MngSubAccSnapshotService) EndTime(endTime int64) *MngSubAccSnapshotService {
	s.endTime = &endTime
	return s
}

func (s *MngSubAccSnapshotService) Limit(limit int32) *MngSubAccSnapshotService {
	s.limit = &limit
	return s
}

func (s *MngSubAccSnapshotService) RecvWindow(recvWindow int64) *MngSubAccSnapshotService {
	s.recvWindow = &recvWindow
	return s
}

func (s *MngSubAccSnapshotService) Do(ctx context.Context, opts ...RequestOption) (res *MngSubAccSnapshotSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/accountSnapshot",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("type", s.accType)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MngSubAccSnapshotSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// managed-sub account query transfer log, this interface is for investor
type MngSubAccQryTrfLogForInvestorService struct {
	c                           *Client
	email                       string
	startTime                   int64
	endTime                     int64 // The start and end time intervals cannot exceed six months
	page                        int32
	limit                       int32   // max 500
	transfers                   *string // from/to
	transferFunctionAccountType *string // SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE
}

type MngSubAccQryTrfLogForInvestorSvcRsp struct {
	ManagerSubTransferHistoryVos []*MgnSubTrfHisVo `json:"managerSubTransferHistoryVos"`
	Count                        int32             `json:"count"`
}

type MgnSubTrfHisVo struct {
	FromEmail       string `json:"fromEmail"`
	FromAccountType string `json:"fromAccountType"`
	ToEmail         string `json:"toEmail"`
	ToAccountType   string `json:"toAccountType"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	ScheduledData   int64  `json:"scheduledData"`
	CreateTime      int64  `json:"createTime"`
	Status          string `json:"status"`
	TranId          int64  `json:"tranId"`
}

func (s *MngSubAccQryTrfLogForInvestorService) Email(email string) *MngSubAccQryTrfLogForInvestorService {
	s.email = email
	return s
}

func (s *MngSubAccQryTrfLogForInvestorService) StartTime(startTime int64) *MngSubAccQryTrfLogForInvestorService {
	s.startTime = startTime
	return s
}

func (s *MngSubAccQryTrfLogForInvestorService) EndTime(endTime int64) *MngSubAccQryTrfLogForInvestorService {
	s.endTime = endTime
	return s
}

func (s *MngSubAccQryTrfLogForInvestorService) Page(page int32) *MngSubAccQryTrfLogForInvestorService {
	s.page = page
	return s
}

func (s *MngSubAccQryTrfLogForInvestorService) Limit(limit int32) *MngSubAccQryTrfLogForInvestorService {
	s.limit = limit
	return s
}

func (s *MngSubAccQryTrfLogForInvestorService) Transfers(transfers string) *MngSubAccQryTrfLogForInvestorService {
	s.transfers = &transfers
	return s
}

func (s *MngSubAccQryTrfLogForInvestorService) TransferFunctionAccountType(transferFunctionAccountType string) *MngSubAccQryTrfLogForInvestorService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *MngSubAccQryTrfLogForInvestorService) Do(ctx context.Context, opts ...RequestOption) (res *MngSubAccQryTrfLogForInvestorSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/queryTransLogForInvestor",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	r.setParam("page", s.page)
	r.setParam("limit", s.limit)
	if s.transfers != nil {
		r.setParam("transfers", *s.transfers)
	}
	if s.transferFunctionAccountType != nil {
		r.setParam("transferFunctionAccountType", *s.transferFunctionAccountType)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MngSubAccQryTrfLogForInvestorSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MngSubAccQryTrfLogForTradeParentService struct {
	c                           *Client
	email                       string
	startTime                   int64
	endTime                     int64 // The start and end time intervals cannot exceed six months
	page                        int32
	limit                       int32   // max 500
	transfers                   *string // from/to
	transferFunctionAccountType *string // SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE
}

type MngSubAccQryTrfLogForTradeParentSvcRsp struct {
	ManagerSubTransferHistoryVos []*MgnSubTrfHisVo `json:"managerSubTransferHistoryVos"`
	Count                        int32             `json:"count"`
}

func (s *MngSubAccQryTrfLogForTradeParentService) Email(email string) *MngSubAccQryTrfLogForTradeParentService {
	s.email = email
	return s
}

func (s *MngSubAccQryTrfLogForTradeParentService) StartTime(startTime int64) *MngSubAccQryTrfLogForTradeParentService {
	s.startTime = startTime
	return s
}

func (s *MngSubAccQryTrfLogForTradeParentService) EndTime(endTime int64) *MngSubAccQryTrfLogForTradeParentService {
	s.endTime = endTime
	return s
}

func (s *MngSubAccQryTrfLogForTradeParentService) Page(page int32) *MngSubAccQryTrfLogForTradeParentService {
	s.page = page
	return s
}

func (s *MngSubAccQryTrfLogForTradeParentService) Limit(limit int32) *MngSubAccQryTrfLogForTradeParentService {
	s.limit = limit
	return s
}

func (s *MngSubAccQryTrfLogForTradeParentService) Transfers(transfers string) *MngSubAccQryTrfLogForTradeParentService {
	s.transfers = &transfers
	return s
}

func (s *MngSubAccQryTrfLogForTradeParentService) TransferFunctionAccountType(transferFunctionAccountType string) *MngSubAccQryTrfLogForTradeParentService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *MngSubAccQryTrfLogForTradeParentService) Do(ctx context.Context, opts ...RequestOption) (res *MngSubAccQryTrfLogForTradeParentSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/queryTransLogForTradeParent",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	r.setParam("page", s.page)
	r.setParam("limit", s.limit)
	if s.transfers != nil {
		r.setParam("transfers", *s.transfers)
	}
	if s.transferFunctionAccountType != nil {
		r.setParam("transferFunctionAccountType", *s.transferFunctionAccountType)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MngSubAccQryTrfLogForTradeParentSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Investor account inquiry custody account futures assets
type MngSubAccQryFuturesAssetService struct {
	c     *Client
	email string
}

type MngSubAccQryFuturesAssetSvcRsp struct {
	Code        int32                 `json:"code"`
	Message     string                `json:"message"`
	SnapshotVos []*MgnSubFutAccSnapVo `json:"snapshotVos"`
}

type MgnSubFutAccSnapVo struct {
	Type       string                  `json:"type"`
	UpdateTime int64                   `json:"updateTime"`
	Data       *MgnSubFutAccSnapVoData `json:"data"`
}

type MgnSubFutAccSnapVoData struct {
	Assets   []*MgnSubFutAccSnapVoDataAsset `json:"assets"`
	Position []*MgnSubFutAccSnapVoDataPos   `json:"position"`
}

type MgnSubFutAccSnapVoDataAsset struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	WalletBalance string `json:"walletBalance"`
}

type MgnSubFutAccSnapVoDataPos struct {
	Symbol      string `json:"symbol"`
	EntryPrice  string `json:"entryPrice"`
	MarkPrice   string `json:"markPrice"`
	PositionAmt string `json:"positionAmt"`
}

func (s *MngSubAccQryFuturesAssetService) Email(email string) *MngSubAccQryFuturesAssetService {
	s.email = email
	return s
}

func (s *MngSubAccQryFuturesAssetService) Do(ctx context.Context, opts ...RequestOption) (res *MngSubAccQryFuturesAssetSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/fetch-future-asset",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MngSubAccQryFuturesAssetSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Investor account inquiry for leveraged assets in custodial accounts
type MngSubAccQryMarginAssetService struct {
	c     *Client
	email string
}

type MngSubAccQryMgnAssetSvcRsp struct {
	MarginLevel         string               `json:"marginLevel"`
	TotalAssetOfBtc     string               `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string               `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string               `json:"totalNetAssetOfBtc"`
	UserAssets          []*MngSubAccMgnAsset `json:"userAssets"`
}

type MngSubAccMgnAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

func (s *MngSubAccQryMarginAssetService) Email(email string) *MngSubAccQryMarginAssetService {
	s.email = email
	return s
}

func (s *MngSubAccQryMarginAssetService) Do(ctx context.Context, opts ...RequestOption) (res *MngSubAccQryMgnAssetSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/marginAsset",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MngSubAccQryMgnAssetSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query sub account assets, v4 interface.
type SubAccAssetService struct {
	c          *Client
	email      string
	recvWindow *int64
}

type SubAccAssetSvcRsp struct {
	Balances []*SubAccAssetBalance `json:"balances"`
}

type SubAccAssetBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

func (s *SubAccAssetService) Email(email string) *SubAccAssetService {
	s.email = email
	return s
}

func (s *SubAccAssetService) RecvWindow(recvWindow int64) *SubAccAssetService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccAssetService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccAssetSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v4/sub-account/assets",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccAssetSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query the list of managed-accounts
type MgnSubAccInfoService struct {
	c          *Client
	email      string
	page       *int32 // default 1
	limit      *int32 // default 20, max 20
	recvWindow *int64
}

type MgnSubAccInfoSvcRsp struct {
	Total                    int32                  `json:"total"`
	ManagerSubUserInfoVoList []*MgnSubAccUserInfoVo `json:"managerSubUserInfoVoList"`
}

type MgnSubAccUserInfoVo struct {
	RootUserId               int64  `json:"rootUserId"`
	ManagersubUserId         int64  `json:"managersubUserId"`
	BindParentUserId         int64  `json:"bindParentUserId"`
	Email                    string `json:"email"`
	InsertTimeStamp          int64  `json:"insertTimeStamp"`
	BindParentEmail          string `json:"bindParentEmail"`
	IsSubUserEnabled         bool   `json:"isSubUserEnabled"`
	IsUserActive             bool   `json:"isUserActive"`
	IsMarginEnabled          bool   `json:"isMarginEnabled"`
	IsFutureEnabled          bool   `json:"isFutureEnabled"`
	IsSignedLVTRiskAgreement bool   `json:"isSignedLVTRiskAgreement"`
}

func (s *MgnSubAccInfoService) Email(email string) *MgnSubAccInfoService {
	s.email = email
	return s
}

func (s *MgnSubAccInfoService) Page(page int32) *MgnSubAccInfoService {
	s.page = &page
	return s
}

func (s *MgnSubAccInfoService) Limit(limit int32) *MgnSubAccInfoService {
	s.limit = &limit
	return s
}

func (s *MgnSubAccInfoService) RecvWindow(recvWindow int64) *MgnSubAccInfoService {
	s.recvWindow = &recvWindow
	return s
}

func (s *MgnSubAccInfoService) Do(ctx context.Context, opts ...RequestOption) (res *MgnSubAccInfoSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/info",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MgnSubAccInfoSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query sub account transaction volume statistics list
type SubAccTxnStatsService struct {
	c          *Client
	email      string
	recvWindow *int64
}

type SubAccTxnStatsSvcRsp struct {
	Recent30BtcTotal         string         `json:"recent30BtcTotal"`
	Recent30BtcFuturesTotal  string         `json:"recent30BtcFuturesTotal"`
	Recent30BtcMarginTotal   string         `json:"recent30BtcMarginTotal"`
	Recent30BusdTotal        string         `json:"recent30BusdTotal"`
	Recent30BusdFuturesTotal string         `json:"recent30BusdFuturesTotal"`
	Recent30BusdMarginTotal  string         `json:"recent30BusdMarginTotal"`
	TradeInfoVos             []*TradeInfoVo `json:"tradeInfoVos"`
}

type TradeInfoVo struct {
	UserId      int64   `json:"userId"`
	Btc         float64 `json:"btc"`
	BtcFutures  float64 `json:"btcFutures"`
	BtcMargin   float64 `json:"btcMargin"`
	Busd        float64 `json:"busd"`
	BusdFutures float64 `json:"busdFutures"`
	BusdMargin  float64 `json:"busdMargin"`
	Date        int64   `json:"date"`
}

func (s *SubAccTxnStatsService) Email(email string) *SubAccTxnStatsService {
	s.email = email
	return s
}

func (s *SubAccTxnStatsService) RecvWindow(recvWindow int64) *SubAccTxnStatsService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccTxnStatsService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccTxnStatsSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/transaction-statistics",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccTxnStatsSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Obtain the recharge address for the custody account
type MgnSubAccDepositAddrService struct {
	c          *Client
	email      string
	coin       string
	network    *string // The network can be obtained from GET /sapi/v1/capital/deposit/address. When the network is not transmitted, return the default network of the coin
	recvWindow *int64
}

type MgnSubAccDepositAddrSvcRsp struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

func (s *MgnSubAccDepositAddrService) Email(email string) *MgnSubAccDepositAddrService {
	s.email = email
	return s
}

func (s *MgnSubAccDepositAddrService) Coin(coin string) *MgnSubAccDepositAddrService {
	s.coin = coin
	return s
}

func (s *MgnSubAccDepositAddrService) Network(network string) *MgnSubAccDepositAddrService {
	s.network = &network
	return s
}

func (s *MgnSubAccDepositAddrService) RecvWindow(recvWindow int64) *MgnSubAccDepositAddrService {
	s.recvWindow = &recvWindow
	return s
}

func (s *MgnSubAccDepositAddrService) Do(ctx context.Context, opts ...RequestOption) (res *MgnSubAccDepositAddrSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/deposit/address",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("coin", s.coin)
	if s.network != nil {
		r.setParam("network", *s.network)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MgnSubAccDepositAddrSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type SubAccOptionsEnableService struct {
	c          *Client
	email      string
	recvWindow *int64
}

type SubAccOptionsEnableSvcRsp struct {
	Email             string `json:"email"`
	IsEOptionsEnabled bool   `json:"isEOptionsEnabled"`
}

func (s *SubAccOptionsEnableService) Email(email string) *SubAccOptionsEnableService {
	s.email = email
	return s
}

func (s *SubAccOptionsEnableService) RecvWindow(recvWindow int64) *SubAccOptionsEnableService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccOptionsEnableService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccOptionsEnableSvcRsp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/eoptions/enable",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccOptionsEnableSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query transfer records of managed-sub accounts
type MgnSubAccQryTrfLogService struct {
	c                           *Client
	startTime                   int64
	endTime                     int64 // End time (start time end time interval cannot exceed six months)
	page                        int32
	limit                       int32   // max 500
	transfers                   *string // from/to
	transferFunctionAccountType *string // SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE
	recvWindow                  *int64
}

type MgnSubAccQryTrfLogSvcRsp struct {
	ManagerSubTransferHistoryVos []*MgnSubTrfHisVo `json:"managerSubTransferHistoryVos"`
	Count                        int32             `json:"count"`
}

func (s *MgnSubAccQryTrfLogService) StartTime(startTime int64) *MgnSubAccQryTrfLogService {
	s.startTime = startTime
	return s
}

func (s *MgnSubAccQryTrfLogService) EndTime(endTime int64) *MgnSubAccQryTrfLogService {
	s.endTime = endTime
	return s
}

func (s *MgnSubAccQryTrfLogService) Page(page int32) *MgnSubAccQryTrfLogService {
	s.page = page
	return s
}

func (s *MgnSubAccQryTrfLogService) Limit(limit int32) *MgnSubAccQryTrfLogService {
	s.limit = limit
	return s
}

func (s *MgnSubAccQryTrfLogService) Transfers(transfers string) *MgnSubAccQryTrfLogService {
	s.transfers = &transfers
	return s
}

func (s *MgnSubAccQryTrfLogService) TransferFunctionAccountType(transferFunctionAccountType string) *MgnSubAccQryTrfLogService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *MgnSubAccQryTrfLogService) RecvWindow(recvWindow int64) *MgnSubAccQryTrfLogService {
	s.recvWindow = &recvWindow
	return s
}

func (s *MgnSubAccQryTrfLogService) Do(ctx context.Context, opts ...RequestOption) (res *MgnSubAccQryTrfLogSvcRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/query-trans-log",
		secType:  secTypeSigned,
	}
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	r.setParam("page", s.page)
	r.setParam("limit", s.limit)
	if s.transfers != nil {
		r.setParam("transfers", *s.transfers)
	}
	if s.transferFunctionAccountType != nil {
		r.setParam("transferFunctionAccountType", *s.transferFunctionAccountType)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MgnSubAccQryTrfLogSvcRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
