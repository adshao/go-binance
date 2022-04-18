package binance

import (
	"context"
	"encoding/json"
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

// TransferToSubAccountResponse define transfer to subaccount response
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

// TransferToSubAccountResponse define transfer to subaccount response
type SubaccountSpotSummaryResponse struct {
	TotalCount                int64                       `json:"totalCount"`
	MasterAccountTotalAsset   string                      `json:"masterAccountTotalAsset"`
	SpotSubUserAssetBtcVoList []SpotSubUserAssetBtcVoList `json:"spotSubUserAssetBtcVoList"`
}

type SpotSubUserAssetBtcVoList struct {
	Email      string `json:"email"`
	TotalAsset string `json:"totalAsset"`
}
