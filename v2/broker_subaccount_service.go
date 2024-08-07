package binance

import (
	"context"
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
