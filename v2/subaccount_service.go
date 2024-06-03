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
