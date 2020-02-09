package futures

import (
	"context"
	"encoding/json"
)

// GetBalanceService get account balance
type GetBalanceService struct {
	c *Client
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*Balance, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/balance",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Balance{}, err
	}
	res = make([]*Balance, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Balance{}, err
	}
	return res, nil
}

// Balance define user balance of your account
type Balance struct {
	AccountAlias      string `json:"accountAlias"`
	Asset             string `json:"asset"`
	Balance           string `json:"balance"`
	WithdrawAvailable string `json:"withdrawAvailable"`
}

// GetAccountService get account info
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Account define account info
type Account struct {
	Assets                      []*AccountAsset    `json:"assets"`
	CanDeposit                  bool               `json:"canDeposit"`
	CanTrade                    bool               `json:"canTrade"`
	CanWithdraw                 bool               `json:"canWithdraw"`
	FeeTier                     int                `json:"feeTier"`
	MaxWithdrawAmount           string             `json:"maxWithdrawAmount"`
	Positions                   []*AccountPosition `json:"positions"`
	TotalInitialMargin          string             `json:"totalInitialMargin"`
	TotalMaintMargin            string             `json:"totalMaintMargin"`
	TotalMarginBalance          string             `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string             `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string             `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string             `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string             `json:"totalWalletBalance"`
	UpdateTime                  int64              `json:"updateTime"`
}

// AccountAsset define account asset
type AccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}

// AccountPosition define account position
type AccountPosition struct {
	Isolated               bool   `json:"isolated"`
	Leverage               string `json:"leverage"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	Symbol                 string `json:"symbol"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
}
