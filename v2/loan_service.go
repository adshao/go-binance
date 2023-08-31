package binance

import (
	"context"
	"net/http"
)

type CryptoLoanAsset struct {
	LoanCoin               string `json:"loanCoin"`
	D7HourlyInterestRate   string `json:"_7dHourlyInterestRate"`
	D7DailyInterestRate    string `json:"_7dDailyInterestRate"`
	D14HourlyInterestRate  string `json:"_14dHourlyInterestRate"`
	D14DailyInterestRate   string `json:"_14dDailyInterestRate"`
	D30HourlyInterestRate  string `json:"_30dHourlyInterestRate"`
	D30DailyInterestRate   string `json:"_30dDailyInterestRate"`
	D90HourlyInterestRate  string `json:"_90dHourlyInterestRate"`
	D90DailyInterestRate   string `json:"_90dDailyInterestRate"`
	D180HourlyInterestRate string `json:"_180dHourlyInterestRate"`
	D180DailyInterestRate  string `json:"_180dDailyInterestRate"`
	MinLimit               string `json:"minLimit"`
	MaxLimit               string `json:"maxLimit"`
	VipLevel               int64  `json:"vipLevel"`
}

type CryptoLoanableAssets struct {
	Rows  []CryptoLoanAsset `json:"rows"`
	Total int64             `json:"total"`
}

type CryptoLoanService struct {
	c *Client
}

// Do send request
func (s *CryptoLoanService) Do(ctx context.Context, opts ...RequestOption) (res *CryptoLoanableAssets, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/loanable/data",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(CryptoLoanableAssets)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VIPLoanAsset struct {
	LoanCoin              string `json:"loanCoin"`
	D30DailyInterestRate  string `json:"_30dDailyInterestRate"`
	D30YearlyInterestRate string `json:"_30dYearlyInterestRate"`
	D60DailyInterestRate  string `json:"_60dDailyInterestRate"`
	D60YearlyInterestRate string `json:"_60dYearlyInterestRate"`
	MinLimit              string `json:"minLimit"`
	MaxLimit              string `json:"maxLimit"`
	VipLevel              int64  `json:"vipLevel"`
}

type VipLoanableAssets struct {
	Rows  []CryptoLoanAsset `json:"rows"`
	Total int64             `json:"total"`
}

type VipLoanService struct {
	c *Client
}

// Do send request
func (s *VipLoanService) Do(ctx context.Context, opts ...RequestOption) (res *VipLoanableAssets, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/loanable/data",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(VipLoanableAssets)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
