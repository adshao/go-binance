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

	if len(data) == 0 || string(data) == "{}" {
		return &DeleteBrokerSubAccountApiKeyResponse{}, nil
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
func (s *BrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res []BrokerSubAccountsResponse, err error) {
	data, err := s.brokerSubAccount(ctx, "/sapi/v1/broker/subAccount", opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// BrokerSubAccountsResponse Query Sub Account response
type BrokerSubAccountsResponse struct {
	SubAccountID          string `json:"subaccountId"`
	Email                 string `json:"email"`
	Tag                   string `json:"tag"`
	MakerCommission       string `json:"makerCommission"`
	TakerCommission       string `json:"takerCommission"`
	MarginMakerCommission string `json:"marginMakerCommission"`
	MarginTakerCommission string `json:"marginTakerCommission"`
	CreateTime            int64  `json:"createTime"`
}

// UpdateIPBrokerSubAccountService Update IP Restriction for Sub-Account API key (For Master Account)
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#update-ip-restriction-for-sub-account-api-key-for-master-account
type UpdateIPBrokerSubAccountService struct {
	c                *Client
	subAccountId     string
	subAccountApiKey string
	status           string
	ipAddress        *string
}

// SubAccountID set subAccountID required
func (s *UpdateIPBrokerSubAccountService) SubAccountID(subAccountID string) *UpdateIPBrokerSubAccountService {
	s.subAccountId = subAccountID
	return s
}

// SubAccountApiKey set api key required
func (s *UpdateIPBrokerSubAccountService) SubAccountApiKey(subAccountApiKey string) *UpdateIPBrokerSubAccountService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

// Status set status required
// IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
func (s *UpdateIPBrokerSubAccountService) Status(status string) *UpdateIPBrokerSubAccountService {
	s.status = status
	return s
}

// IPAddress set ipAddress Insert static IP in batch, separated by commas.
func (s *UpdateIPBrokerSubAccountService) IPAddress(ipAddress string) *UpdateIPBrokerSubAccountService {
	s.ipAddress = &ipAddress
	return s
}

func (s *UpdateIPBrokerSubAccountService) updateIPBrokerSubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
		"status":           s.status,
	}
	r.setParams(m)

	if s.ipAddress != nil {
		r.setParam("ipAddress", *s.ipAddress)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *UpdateIPBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *UpdateIPBrokerSubAccountsResponse, err error) {
	data, err := s.updateIPBrokerSubAccount(ctx, "/sapi/v2/broker/subAccountApi/ipRestriction", opts...)
	if err != nil {
		return nil, err
	}
	res = &UpdateIPBrokerSubAccountsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateIPBrokerSubAccountsResponse Update IP Restriction for Sub-Account API key (For Master Account) response
type UpdateIPBrokerSubAccountsResponse struct {
	ApiKey     string   `json:"apiKey"`
	Status     string   `json:"status"`
	IPList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
}

// IPBrokerSubAccountService Update IP Restriction for Sub-Account API key (For Master Account)
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#get-ip-restriction-for-sub-account-api-key
type IPBrokerSubAccountService struct {
	c                *Client
	subAccountId     string
	subAccountApiKey string
}

// SubAccountID set subAccountID required
func (s *IPBrokerSubAccountService) SubAccountID(subAccountID string) *IPBrokerSubAccountService {
	s.subAccountId = subAccountID
	return s
}

// SubAccountApiKey set api key required
func (s *IPBrokerSubAccountService) SubAccountApiKey(subAccountApiKey string) *IPBrokerSubAccountService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *IPBrokerSubAccountService) ipBrokerSubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodGet,
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
func (s *IPBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *IPBrokerSubAccountResponse, err error) {
	data, err := s.ipBrokerSubAccount(ctx, "/sapi/v1/broker/subAccountApi/ipRestriction", opts...)
	if err != nil {
		return nil, err
	}
	res = &IPBrokerSubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IPBrokerSubAccountResponse Get IP Restriction for Sub Account Api Key response
type IPBrokerSubAccountResponse struct {
	SubAccountID string   `json:"subaccountId"`
	IPRestrict   bool     `json:"ipRestrict"`
	ApiKey       string   `json:"apikey"`
	IPList       []string `json:"ipList"`
	UpdateTime   int64    `json:"updateTime"`
}

// DeleteIPBrokerSubAccountService Update IP Restriction for Sub-Account API key (For Master Account)
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#get-ip-restriction-for-sub-account-api-key
type DeleteIPBrokerSubAccountService struct {
	c                *Client
	subAccountId     string
	subAccountApiKey string
	ipAddress        *string
}

// SubAccountID set subAccountID required
func (s *DeleteIPBrokerSubAccountService) SubAccountID(subAccountID string) *DeleteIPBrokerSubAccountService {
	s.subAccountId = subAccountID
	return s
}

// SubAccountApiKey set api key required
func (s *DeleteIPBrokerSubAccountService) SubAccountApiKey(subAccountApiKey string) *DeleteIPBrokerSubAccountService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

// IPAddress set ipAddress Insert static IP in batch, separated by commas.
func (s *DeleteIPBrokerSubAccountService) IPAddress(ipAddress string) *DeleteIPBrokerSubAccountService {
	s.ipAddress = &ipAddress
	return s
}

func (s *DeleteIPBrokerSubAccountService) deleteIPBrokerSubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
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

	if s.ipAddress != nil {
		r.setParam("ipAddress", *s.ipAddress)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *DeleteIPBrokerSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *DeleteIPBrokerSubAccountResponse, err error) {
	data, err := s.deleteIPBrokerSubAccount(ctx, "/sapi/v1/broker/subAccountApi/ipRestriction/ipList", opts...)
	if err != nil {
		return nil, err
	}

	res = &DeleteIPBrokerSubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteIPBrokerSubAccountResponse Get IP Restriction for Sub Account Api Key response
type DeleteIPBrokerSubAccountResponse struct {
	SubAccountID string   `json:"subaccountId"`
	ApiKey       string   `json:"apikey"`
	IPList       []string `json:"ipList"`
	UpdateTime   int64    `json:"updateTime"`
}

// EnableUniversalTransferPermissionService Enable Universal Transfer Permission For Sub Account Api Key
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#enable-universal-transfer-permission-for-sub-account-api-key
type EnableUniversalTransferPermissionService struct {
	c                    *Client
	subAccountId         string
	subAccountApiKey     string
	canUniversalTransfer bool
}

// SubAccountID set subAccountID required
func (s *EnableUniversalTransferPermissionService) SubAccountID(subAccountID string) *EnableUniversalTransferPermissionService {
	s.subAccountId = subAccountID
	return s
}

// SubAccountApiKey set api key required
func (s *EnableUniversalTransferPermissionService) SubAccountApiKey(subAccountApiKey string) *EnableUniversalTransferPermissionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

// CanUniversalTransfer set canUniversalTransfer required
func (s *EnableUniversalTransferPermissionService) CanUniversalTransfer(canUniversalTransfer bool) *EnableUniversalTransferPermissionService {
	s.canUniversalTransfer = canUniversalTransfer
	return s
}

func (s *EnableUniversalTransferPermissionService) enableUniversalTransferPermission(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":         s.subAccountId,
		"subAccountApiKey":     s.subAccountApiKey,
		"canUniversalTransfer": s.canUniversalTransfer,
	}
	r.setParams(m)

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *EnableUniversalTransferPermissionService) Do(ctx context.Context, opts ...RequestOption) (res *EnableUniversalTransferPermissionResponse, err error) {
	data, err := s.enableUniversalTransferPermission(ctx, "/sapi/v1/broker/subAccountApi/permission/universalTransfer", opts...)
	if err != nil {
		return nil, err
	}
	res = &EnableUniversalTransferPermissionResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// EnableUniversalTransferPermissionResponse Enable Universal Transfer Permission For Sub Account Api Key response
type EnableUniversalTransferPermissionResponse struct {
	SubAccountID         string `json:"subaccountId"`
	ApiKey               string `json:"apikey"`
	CanUniversalTransfer bool   `json:"canUniversalTransfer"`
}

// UniversalTransferService Universal Transfer
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#enable-universal-transfer-permission-for-sub-account-api-key
type UniversalTransferService struct {
	c      *Client
	fromId string
	toId   string
	// SPOT, USDT_FUTURE, COIN_FUTURE
	fromAccountType string
	// SPOT,USDT_FUTURE,COIN_FUTURE
	toAccountType string
	// Client transfer id, must be unique. The max length is 32 characters
	clientTranId string
	asset        string
	amount       float64
}

// FromID set fromID
func (s *UniversalTransferService) FromID(fromID string) *UniversalTransferService {
	s.fromId = fromID
	return s
}

// ToID set toID
func (s *UniversalTransferService) ToID(toID string) *UniversalTransferService {
	s.toId = toID
	return s
}

// FromAccountType set fromAccountType required
func (s *UniversalTransferService) FromAccountType(fromAccountType string) *UniversalTransferService {
	s.fromAccountType = fromAccountType
	return s
}

// ToAccountType set toAccountType required
func (s *UniversalTransferService) ToAccountType(toAccountType string) *UniversalTransferService {
	s.toAccountType = toAccountType
	return s
}

// ClientTranID set clientTranID
func (s *UniversalTransferService) ClientTranID(clientTranID string) *UniversalTransferService {
	s.clientTranId = clientTranID
	return s
}

// Asset set asset required
func (s *UniversalTransferService) Asset(asset string) *UniversalTransferService {
	s.asset = asset
	return s
}

// Amount set amount required
func (s *UniversalTransferService) Amount(amount float64) *UniversalTransferService {
	s.amount = amount
	return s
}

func (s *UniversalTransferService) universalTransfer(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"fromId":          s.fromId,
		"toId":            s.toId,
		"fromAccountType": s.fromAccountType,
		"toAccountType":   s.toAccountType,
		"clientTranId":    s.clientTranId,
		"asset":           s.asset,
		"amount":          s.amount,
	}
	r.setParams(m)

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *UniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (res *UniversalTransferResponse, err error) {
	data, err := s.universalTransfer(ctx, "/sapi/v1/broker/universalTransfer", opts...)
	if err != nil {
		return nil, err
	}
	res = &UniversalTransferResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UniversalTransferResponse Enable Universal Transfer Permission For Sub Account Api Key response
type UniversalTransferResponse struct {
	TxID         int64  `json:"txnId"`
	ClientTranID string `json:"clientTranId"`
}

// UniversalTransferHistoryService Query Universal Transfer History
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#query-universal-transfer-history
type UniversalTransferHistoryService struct {
	c *Client
	// Either fromId or toId must be sent.
	fromId *string
	// Either fromId or toId must be sent.
	toId *string
	// Client transfer id, must be unique. The max length is 32 characters
	clientTranId string
	startTime    *int64
	endTime      *int64
	page         int64
	//default 500, max 500
	limit         int64
	showAllStatus bool
}

// FromID set fromID
func (s *UniversalTransferHistoryService) FromID(fromID string) *UniversalTransferHistoryService {
	s.fromId = &fromID
	return s
}

// ToID set toID
func (s *UniversalTransferHistoryService) ToID(toID string) *UniversalTransferHistoryService {
	s.toId = &toID
	return s
}

// ClientTranID set clientTranID
func (s *UniversalTransferHistoryService) ClientTranID(clientTranID string) *UniversalTransferHistoryService {
	s.clientTranId = clientTranID
	return s
}

// StartTime set startTime
func (s *UniversalTransferHistoryService) StartTime(startTime int64) *UniversalTransferHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UniversalTransferHistoryService) EndTime(endTime int64) *UniversalTransferHistoryService {
	s.endTime = &endTime
	return s
}

// Page set page required
func (s *UniversalTransferHistoryService) Page(page int64) *UniversalTransferHistoryService {
	s.page = page
	return s
}

// Limit set limit required
func (s *UniversalTransferHistoryService) Limit(limit int64) *UniversalTransferHistoryService {
	s.limit = limit
	return s
}

// ShowAllStatus set showAllStatus required
func (s *UniversalTransferHistoryService) ShowAllStatus(showAllStatus bool) *UniversalTransferHistoryService {
	s.showAllStatus = showAllStatus
	return s
}

func (s *UniversalTransferHistoryService) universalTransferHistory(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	m := params{
		"clientTranId":  s.clientTranId,
		"page":          s.page,
		"limit":         s.limit,
		"showAllStatus": s.showAllStatus,
	}
	r.setParams(m)

	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.toId != nil {
		r.setParam("toId", *s.toId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.startTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *UniversalTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*UniversalTransferHistoryResponse, err error) {
	data, err := s.universalTransferHistory(ctx, "/sapi/v1/broker/universalTransfer", opts...)
	if err != nil {
		return nil, err
	}
	fmt.Println("UniversalTransferHistoryService: ", string(data))
	res = make([]*UniversalTransferHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UniversalTransferHistoryResponse Query Universal Transfer History response
type UniversalTransferHistoryResponse struct {
	ToID            string `json:"toId"`
	Asset           string `json:"asset"`
	Qty             string `json:"qty"`
	Time            int64  `json:"time"`
	Status          string `json:"status"`
	TxID            int64  `json:"txnId"`
	ClientTranID    string `json:"clientTranId"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
}

// SubAccountDepositHistoryService Get Sub Account Deposit History
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#get-sub-account-deposit-history
type SubAccountDepositHistoryService struct {
	c            *Client
	subAccountId string
	coin         *string
	status       *int64
	startTime    *int64
	endTime      *int64
	//default 500, max 500
	limit  *int64
	offset *int64
}

// SubAccountID set subAccountId
func (s *SubAccountDepositHistoryService) SubAccountID(subAccountId string) *SubAccountDepositHistoryService {
	s.subAccountId = subAccountId
	return s
}

// Coin set cion
func (s *SubAccountDepositHistoryService) Coin(coin string) *SubAccountDepositHistoryService {
	s.coin = &coin
	return s
}

// Status set status
func (s *SubAccountDepositHistoryService) Status(status int64) *SubAccountDepositHistoryService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *SubAccountDepositHistoryService) StartTime(startTime int64) *SubAccountDepositHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *SubAccountDepositHistoryService) EndTime(endTime int64) *SubAccountDepositHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *SubAccountDepositHistoryService) Limit(limit int64) *SubAccountDepositHistoryService {
	s.limit = &limit
	return s
}

// Offset set offset
func (s *SubAccountDepositHistoryService) Offset(offset int64) *SubAccountDepositHistoryService {
	s.offset = &offset
	return s
}

func (s *SubAccountDepositHistoryService) subAccountDepositHistory(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	m := params{
		"subAccountId": s.subAccountId,
	}
	r.setParams(m)

	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.startTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubAccountDepositHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccountDepositHistoryResponse, err error) {
	data, err := s.subAccountDepositHistory(ctx, "/sapi/v1/broker/subAccount/depositHist", opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccountDepositHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubAccountDepositHistoryResponse Get Sub Account Deposit History
type SubAccountDepositHistoryResponse struct {
	DepositID        int64  `json:"depositId"`
	SubAccountID     string `json:"subAccountId"`
	Address          string `json:"address"`
	AddressTag       string `json:"addressTag"`
	Amount           string `json:"amount"`
	Coin             string `json:"coin"`
	InsertTime       int64  `json:"insertTime"`
	TransferType     int64  `json:"transferType"`
	Network          string `json:"network"`
	Status           int64  `json:"status"`
	TxID             string `json:"txId"`
	SourceAddress    string `json:"sourceAddress"`
	ConfirmTimes     string `json:"confirmTimes"`
	SelfReturnStatus int64  `json:"selfReturnStatus"`
}

// SubAccountTransferFuturesService Sub Account Transfer（FUTURES）
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#sub-account-transferfutures
type SubAccountTransferFuturesService struct {
	c *Client
	// Either fromId or toId must be sent.
	fromId *string
	// Either fromId or toId must be sent.
	toId         *string
	futuresType  int64
	asset        string
	amount       float64
	clientTranId string
}

// FromID set fromId
func (s *SubAccountTransferFuturesService) FromID(fromId string) *SubAccountTransferFuturesService {
	s.fromId = &fromId
	return s
}

// ToID set toId
func (s *SubAccountTransferFuturesService) ToID(toId string) *SubAccountTransferFuturesService {
	s.toId = &toId
	return s
}

// FuturesType set futuresType
func (s *SubAccountTransferFuturesService) FuturesType(futuresType int64) *SubAccountTransferFuturesService {
	s.futuresType = futuresType
	return s
}

// Asset set startTime
func (s *SubAccountTransferFuturesService) Asset(asset string) *SubAccountTransferFuturesService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *SubAccountTransferFuturesService) Amount(amount float64) *SubAccountTransferFuturesService {
	s.amount = amount
	return s
}

// ClientTranId set clientTranId
func (s *SubAccountTransferFuturesService) ClientTranId(clientTranId string) *SubAccountTransferFuturesService {
	s.clientTranId = clientTranId
	return s
}

func (s *SubAccountTransferFuturesService) subAccountTransferFutures(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	m := params{
		"futuresType":  s.futuresType,
		"asset":        s.asset,
		"amount":       s.amount,
		"clientTranId": s.clientTranId,
	}
	r.setParams(m)

	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.toId != nil {
		r.setParam("toId", *s.toId)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubAccountTransferFuturesService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountTransferFuturesResponse, err error) {
	data, err := s.subAccountTransferFutures(ctx, "/sapi/v1/broker/transfer/futures", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubAccountTransferFuturesResponse{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubAccountTransferFuturesResponse Sub Account Transfer（FUTURES)
type SubAccountTransferFuturesResponse struct {
	Success      bool   `json:"success"`
	TxnID        int64  `json:"txnId"`
	ClientTranID string `json:"clientTranId"`
}

// SubAccountTransferFuturesHistoryService Query Sub Account Transfer History（FUTURES）
// https://binance-docs.github.io/Brokerage-API/Brokerage_Operation_Endpoints/#query-sub-account-transfer-historyfutures
type SubAccountTransferFuturesHistoryService struct {
	c            *Client
	subAccountId string
	futuresType  int64
	clientTranId *string
	startTime    *int64
	endTime      *int64
	//default 50, max 500
	limit *int64
	page  *int64
}

// SubAccountID set subAccountId required
func (s *SubAccountTransferFuturesHistoryService) SubAccountID(subAccountId string) *SubAccountTransferFuturesHistoryService {
	s.subAccountId = subAccountId
	return s
}

// FuturesType set futuresType required
func (s *SubAccountTransferFuturesHistoryService) FuturesType(futuresType int64) *SubAccountTransferFuturesHistoryService {
	s.futuresType = futuresType
	return s
}

// ClientTranId set clientTranId
func (s *SubAccountTransferFuturesHistoryService) ClientTranId(clientTranId string) *SubAccountTransferFuturesHistoryService {
	s.clientTranId = &clientTranId
	return s
}

// StartTime set startTime
func (s *SubAccountTransferFuturesHistoryService) StartTime(startTime int64) *SubAccountTransferFuturesHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *SubAccountTransferFuturesHistoryService) EndTime(endTime int64) *SubAccountTransferFuturesHistoryService {
	s.endTime = &endTime
	return s
}

// Page set page required
func (s *SubAccountTransferFuturesHistoryService) Page(page int64) *SubAccountTransferFuturesHistoryService {
	s.page = &page
	return s
}

// Limit set limit required
func (s *SubAccountTransferFuturesHistoryService) Limit(limit int64) *SubAccountTransferFuturesHistoryService {
	s.limit = &limit
	return s
}

func (s *SubAccountTransferFuturesHistoryService) subAccountTransferFuturesHistory(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	m := params{
		"subAccountId": s.subAccountId,
		"futuresType":  s.futuresType,
	}
	r.setParams(m)

	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
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
	if s.page != nil {
		r.setParam("page", *s.page)
	}

	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubAccountTransferFuturesHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountTransferFuturesHistoryResponse, err error) {
	data, err := s.subAccountTransferFuturesHistory(ctx, "/sapi/v1/broker/transfer/futures", opts...)
	if err != nil {
		return nil, err
	}

	res = &SubAccountTransferFuturesHistoryResponse{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubAccountTransferFuturesHistoryResponse Query Sub Account Transfer History（FUTURES）
type SubAccountTransferFuturesHistoryResponse struct {
	Success     bool        `json:"success"`
	FuturesType int64       `json:"futuresType"`
	Transfers   []Transfers `json:"transfers"`
}

type Transfers struct {
	FromSubAccountID string `json:"from,omitempty"`
	ToSubAccountID   string `json:"to,omitempty"`
	Asset            string `json:"asset"`
	Quantity         string `json:"qty"`
	TxID             string `json:"tranId"`
	ClientTranID     string `json:"clientTranId"`
	Time             int64  `json:"time"`
}
