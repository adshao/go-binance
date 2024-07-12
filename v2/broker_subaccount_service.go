package binance

import (
	"context"
	"fmt"
	"net/http"
)

// CreateBrokerSubAccountService Create a Sub Account
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#create-a-sub-account
type CreateBrokerSubAccountService struct {
	c   *Client
	tag string
}

// Tag set random string required
func (s *CreateBrokerSubAccountService) Tag(tag string) *CreateBrokerSubAccountService {
	s.tag = tag
	return s
}

func (s *CreateBrokerSubAccountService) createBrokerSubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"tag": s.tag,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *CreateBrokerSubAccountResponse, err error) {
	data, err := s.createBrokerSubAccount(ctx, "/sapi/v1/broker/subAccount", opts...)
	if err != nil {
		return nil, err
	}
	res = &CreateBrokerSubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateBrokerSubAccountResponse Create a Virtual Sub-account response
type CreateBrokerSubAccountResponse struct {
	SubAccountID string `json:"subaccountId"`
	Email        string `json:"email"`
	Tag          string `json:"tag"`
}

// EnableFuturesBrokerSubAccountService Enable Futures for Sub Account
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#enable-futures-for-sub-account
type EnableFuturesBrokerSubAccountService struct {
	c            *Client
	subAccountId string
	futures      bool
}

// SubAccountID set subAccountID required
func (s *EnableFuturesBrokerSubAccountService) SubAccountID(subAccountID string) *EnableFuturesBrokerSubAccountService {
	s.subAccountId = subAccountID
	return s
}

// Futures set true only, required
func (s *EnableFuturesBrokerSubAccountService) Futures(futures bool) *EnableFuturesBrokerSubAccountService {
	s.futures = futures
	return s
}

func (s *EnableFuturesBrokerSubAccountService) enableFuturesBrokerSubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"futures":      s.futures,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *EnableFuturesBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *EnableFuturesBrokerSubAccountResponse, err error) {
	data, err := s.enableFuturesBrokerSubAccount(ctx, "/sapi/v1/broker/subAccount/futures", opts...)
	if err != nil {
		return nil, err
	}
	res = &EnableFuturesBrokerSubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// EnableFuturesBrokerSubAccountResponse Enable Futures for Sub Account response
type EnableFuturesBrokerSubAccountResponse struct {
	SubAccountID  string `json:"subaccountId"`
	EnableFutures bool   `json:"enableFutures"`
	UpdateTime    int64  `json:"updateTime"`
}

// CreateBrokerApiKeySubAccountService Create Api Key for Sub Account
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#create-api-key-for-sub-account
// You need to enable "trade" option for the api key which requests this endpoint
// Sub account should be enable margin before its api-key's marginTrade being enabled
// Sub account should be enable futures before its api-key's futuresTrade being enabled
type CreateBrokerApiKeySubAccountService struct {
	c            *Client
	subAccountId string
	canTrade     bool
	marginTrade  *bool
	futuresTrade *bool
}

// SubAccountID set subAccountID required
func (s *CreateBrokerApiKeySubAccountService) SubAccountID(subAccountID string) *CreateBrokerApiKeySubAccountService {
	s.subAccountId = subAccountID
	return s
}

// CanTrade required
func (s *CreateBrokerApiKeySubAccountService) CanTrade(canTrade bool) *CreateBrokerApiKeySubAccountService {
	s.canTrade = canTrade
	return s
}

// MarginTrade Sub account should be enable margin before its api-key's marginTrade being enabled
func (s *CreateBrokerApiKeySubAccountService) MarginTrade(marginTrade bool) *CreateBrokerApiKeySubAccountService {
	s.marginTrade = &marginTrade
	return s
}

// FuturesTrade Sub account should be enable futures before its api-key's futuresTrade being enabled
func (s *CreateBrokerApiKeySubAccountService) FuturesTrade(futuresTrade bool) *CreateBrokerApiKeySubAccountService {
	s.futuresTrade = &futuresTrade
	return s
}

func (s *CreateBrokerApiKeySubAccountService) createBrokerApiKeySubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"canTrade":     s.canTrade,
	}
	r.setParams(m)

	if s.marginTrade != nil {
		r.setParam("marginTrade", *s.marginTrade)
	}

	if s.futuresTrade != nil {
		r.setParam("futuresTrade", *s.futuresTrade)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateBrokerApiKeySubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *CreateBrokerApiKeySubAccountResponse, err error) {
	data, err := s.createBrokerApiKeySubAccount(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	if err != nil {
		return nil, err
	}
	res = &CreateBrokerApiKeySubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateBrokerApiKeySubAccountResponse Create Api Key for Sub Account response
type CreateBrokerApiKeySubAccountResponse struct {
	SubAccountID string `json:"subaccountId"`
	Apikey       string `json:"apikey"`
	SecretKey    string `json:"secretKey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

// DeleteBrokerSubAccountApiKeyService Delete Sub Account Api Key
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#delete-sub-account-api-key
type DeleteBrokerSubAccountApiKeyService struct {
	c                *Client
	subAccountId     string
	subAccountApiKey string
}

// SubAccountID set subAccountID required
func (s *DeleteBrokerSubAccountApiKeyService) SubAccountID(subAccountID string) *DeleteBrokerSubAccountApiKeyService {
	s.subAccountId = subAccountID
	return s
}

// SubAccountApiKey required
func (s *DeleteBrokerSubAccountApiKeyService) SubAccountApiKey(subAccountApiKey string) *DeleteBrokerSubAccountApiKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *DeleteBrokerSubAccountApiKeyService) deleteBrokerApiKeySubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
	}
	r.setParams(m)

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *DeleteBrokerSubAccountApiKeyService) Do(ctx context.Context, opts ...RequestOption) (res *DeleteBrokerSubAccountApiKeyResponse, err error) {
	data, err := s.deleteBrokerApiKeySubAccount(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	if err != nil {
		return nil, err
	}
	res = &DeleteBrokerSubAccountApiKeyResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteBrokerSubAccountApiKeyResponse Delete Sub Account Api Key response
type DeleteBrokerSubAccountApiKeyResponse struct{}

// BrokerSubAccountApiKeysService Query Sub Account Api Key
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#query-sub-account-api-key
type BrokerSubAccountApiKeysService struct {
	c                *Client
	subAccountId     string
	subAccountApiKey *string
	page             int64
	size             int64
}

// SubAccountID set subAccountID required
func (s *BrokerSubAccountApiKeysService) SubAccountID(subAccountID string) *BrokerSubAccountApiKeysService {
	s.subAccountId = subAccountID
	return s
}

// SubAccountApiKey set SubAccountApiKey
func (s *BrokerSubAccountApiKeysService) SubAccountApiKey(subAccountApiKey string) *BrokerSubAccountApiKeysService {
	s.subAccountApiKey = &subAccountApiKey
	return s
}

// Page set page
func (s *BrokerSubAccountApiKeysService) Page(page int64) *BrokerSubAccountApiKeysService {
	if page == 0 {
		page = 1
	}
	s.page = page
	return s
}

// Size set count objects per page
func (s *BrokerSubAccountApiKeysService) Size(size int64) *BrokerSubAccountApiKeysService {
	if size > 500 {
		size = 500
	}
	if size == 0 {
		size = 1
	}
	s.size = size
	return s
}

func (s *BrokerSubAccountApiKeysService) brokerSubAccountApiKey(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"page":         s.page,
		"size":         s.size,
	}
	r.setParams(m)

	if s.subAccountApiKey != nil {
		r.setParam("subAccountApiKey", *s.subAccountApiKey)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *BrokerSubAccountApiKeysService) Do(ctx context.Context, opts ...RequestOption) (res *BrokerSubAccountApiKeysResponse, err error) {
	data, err := s.brokerSubAccountApiKey(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	if err != nil {
		return nil, err
	}
	res = &BrokerSubAccountApiKeysResponse{}
	fmt.Println("res: ", string(data))
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// BrokerSubAccountApiKeysResponse Query Sub Account Api Key response
type BrokerSubAccountApiKeysResponse struct {
	SubAccountsApiKeys []BrokerSubAccountApiKeys
}

type BrokerSubAccountApiKeys struct {
	SubAccountID string `json:"subaccountId"`
	ApiKey       string `json:"apikey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

// ChangeBrokerSubAccountApiPermissionService Change Sub Account Api Permission
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#change-sub-account-api-permission
type ChangeBrokerSubAccountApiPermissionService struct {
	c                *Client
	subAccountId     string
	subAccountApiKey string
	canTrade         bool
	marginTrade      bool
	futuresTrade     bool
}

// SubAccountID set subAccountID required
func (s *ChangeBrokerSubAccountApiPermissionService) SubAccountID(subAccountID string) *ChangeBrokerSubAccountApiPermissionService {
	s.subAccountId = subAccountID
	return s
}

// SubAccountApiKey set subAccountApiKey required
func (s *ChangeBrokerSubAccountApiPermissionService) SubAccountApiKey(subAccountApiKey string) *ChangeBrokerSubAccountApiPermissionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

// CanTrade set canTrade required
func (s *ChangeBrokerSubAccountApiPermissionService) CanTrade(canTrade bool) *ChangeBrokerSubAccountApiPermissionService {
	s.canTrade = canTrade
	return s
}

// MarginTrade set marginTrade required
// Sub account should be enable margin before its api-key's marginTrade being enabled.
func (s *ChangeBrokerSubAccountApiPermissionService) MarginTrade(marginTrade bool) *ChangeBrokerSubAccountApiPermissionService {
	s.marginTrade = marginTrade
	return s
}

// FuturesTrade set futuresTrade required
// Sub account should be enable futures before its api-key's futuresTrade being enabled.
func (s *ChangeBrokerSubAccountApiPermissionService) FuturesTrade(futuresTrade bool) *ChangeBrokerSubAccountApiPermissionService {
	s.futuresTrade = futuresTrade
	return s
}

func (s *ChangeBrokerSubAccountApiPermissionService) changeBrokerSubAccountApiPermission(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
		"canTrade":         s.canTrade,
		"marginTrade":      s.marginTrade,
		"futuresTrade":     s.futuresTrade,
	}
	r.setParams(m)

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *ChangeBrokerSubAccountApiPermissionService) Do(ctx context.Context, opts ...RequestOption) (res *ChangeBrokerSubAccountApiPermissionResponse, err error) {
	data, err := s.changeBrokerSubAccountApiPermission(ctx, "/sapi/v1/broker/subAccountApi/permission", opts...)
	if err != nil {
		return nil, err
	}
	res = &ChangeBrokerSubAccountApiPermissionResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ChangeBrokerSubAccountApiPermissionResponse Change Sub Account Api Permission response
type ChangeBrokerSubAccountApiPermissionResponse struct {
	SubAccountID string `json:"subaccountId"`
	ApiKey       string `json:"apikey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

// BrokerSubAccountService Query Sub Account
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#query-sub-account
type BrokerSubAccountService struct {
	c            *Client
	subAccountId *string
	page         int64
	size         int64
}

// SubAccountID set subAccountID
func (s *BrokerSubAccountService) SubAccountID(subAccountID string) *BrokerSubAccountService {
	s.subAccountId = &subAccountID
	return s
}

// Page set page
func (s *BrokerSubAccountService) Page(page int64) *BrokerSubAccountService {
	if page == 0 {
		page = 1
	}
	s.page = page
	return s
}

// Size set count objects per page
func (s *BrokerSubAccountService) Size(size int64) *BrokerSubAccountService {
	if size > 500 {
		size = 500
	}
	if size == 0 {
		size = 1
	}
	s.size = size
	return s
}

func (s *BrokerSubAccountService) brokerSubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"page": s.page,
		"size": s.size,
	}
	r.setParams(m)

	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *BrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *BrokerSubAccountsResponse, err error) {
	data, err := s.brokerSubAccount(ctx, "/sapi/v1/broker/subAccount", opts...)
	if err != nil {
		return nil, err
	}
	res = &BrokerSubAccountsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// BrokerSubAccountsResponse Query Sub Account response
type BrokerSubAccountsResponse struct {
	SubAccountsApiKeys []BrokerSubAccount
}

type BrokerSubAccount struct {
	SubAccountID          string  `json:"subaccountId"`
	Email                 string  `json:"email"`
	Tag                   string  `json:"tag"`
	MakerCommission       float64 `json:"makerCommission"`
	TakerCommission       float64 `json:"takerCommission"`
	MarginMakerCommission int64   `json:"marginMakerCommission"`
	MarginTakerCommission int64   `json:"marginTakerCommission"`
	CreateTime            int64   `json:"createTime"`
}
