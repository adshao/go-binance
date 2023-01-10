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
		method:   http.MethodGet,
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
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/subAccount/margin",
		secType:  secTypeSigned,
	}
	r.setParam("subaccountId", s.subaccountId)
	r.setParam("margin", true)
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
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/subAccount/futures",
		secType:  secTypeSigned,
	}
	r.setParam("subaccountId", s.subaccountId)
	r.setParam("futures", true)
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
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", s.subAccountId)
	r.setParam("canTrade", s.canTrade)
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
	SecretKey    int64  `json:"secretKey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}
