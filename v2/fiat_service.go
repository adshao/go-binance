package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// FiatDepositWithdrawHistoryService retrieve the fiat deposit/withdraw history
type FiatDepositWithdrawHistoryService struct {
	c               *Client
	transactionType TransactionType
	beginTime       *int64
	endTime         *int64
	page            *int32
	rows            *int32
}

// TransactionType set transactionType
func (s *FiatDepositWithdrawHistoryService) TransactionType(transactionType TransactionType) *FiatDepositWithdrawHistoryService {
	s.transactionType = transactionType
	return s
}

// BeginTime set beginTime
func (s *FiatDepositWithdrawHistoryService) BeginTime(beginTime int64) *FiatDepositWithdrawHistoryService {
	s.beginTime = &beginTime
	return s
}

// EndTime set endTime
func (s *FiatDepositWithdrawHistoryService) EndTime(endTime int64) *FiatDepositWithdrawHistoryService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *FiatDepositWithdrawHistoryService) Page(page int32) *FiatDepositWithdrawHistoryService {
	s.page = &page
	return s
}

// Rows set rows
func (s *FiatDepositWithdrawHistoryService) Rows(rows int32) *FiatDepositWithdrawHistoryService {
	s.rows = &rows
	return s
}

// Do send request
func (s *FiatDepositWithdrawHistoryService) Do(ctx context.Context, opts ...RequestOption) (*FiatDepositWithdrawHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/fiat/orders",
		secType:  secTypeSigned,
	}
	r.setParam("transactionType", s.transactionType)
	if s.beginTime != nil {
		r.setParam("beginTime", *s.beginTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.rows != nil {
		r.setParam("rows", *s.rows)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := FiatDepositWithdrawHistory{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// FiatDepositWithdrawHistory define the fiat deposit/withdraw history
type FiatDepositWithdrawHistory struct {
	Code    string                           `json:"code"`
	Message string                           `json:"message"`
	Data    []FiatDepositWithdrawHistoryItem `json:"data"`
	Total   int32                            `json:"total"`
	Success bool                             `json:"success"`
}

// FiatDepositWithdrawHistoryItem define a fiat deposit/withdraw history item
type FiatDepositWithdrawHistoryItem struct {
	OrderNo         string `json:"orderNo"`
	FiatCurrency    string `json:"fiatCurrency"`
	IndicatedAmount string `json:"indicatedAmount"`
	Amount          string `json:"amount"`
	TotalFee        string `json:"totalFee"`
	Method          string `json:"method"`
	Status          string `json:"status"`
	CreateTime      int64  `json:"createTime"`
	UpdateTime      int64  `json:"updateTime"`
}
