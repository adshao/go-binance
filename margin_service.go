package binance

import (
	"context"
	"encoding/json"
)

// MarginTransferService transfer between spot account and margin account
type MarginTransferService struct {
	c            *Client
	asset        string
	amount       string
	transferType int
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginTransferService) Asset(asset string) *MarginTransferService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginTransferService) Amount(amount string) *MarginTransferService {
	s.amount = amount
	return s
}

// Type 1: transfer from main account to margin account 2: transfer from margin account to main account
func (s *MarginTransferService) Type(transferType MarginTransferType) *MarginTransferService {
	s.transferType = int(transferType)
	return s
}

// Do send request
func (s *MarginTransferService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/margin/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TransactionResponse define transaction response
type TransactionResponse struct {
	TranID int64 `json:"tranId"`
}

// MarginLoanService apply for a loan
type MarginLoanService struct {
	c      *Client
	asset  string
	amount string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginLoanService) Asset(asset string) *MarginLoanService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginLoanService) Amount(amount string) *MarginLoanService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginLoanService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/margin/loan",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginRepayService repay loan for margin account
type MarginRepayService struct {
	c      *Client
	asset  string
	amount string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginRepayService) Asset(asset string) *MarginRepayService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginRepayService) Amount(amount string) *MarginRepayService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginRepayService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/margin/repay",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListMarginLoansService list loan record
type ListMarginLoansService struct {
	c         *Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListMarginLoansService) Asset(asset string) *ListMarginLoansService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *ListMarginLoansService) TxID(txID int64) *ListMarginLoansService {
	s.txID = &txID
	return s
}

// StartTime set start time
func (s *ListMarginLoansService) StartTime(startTime int64) *ListMarginLoansService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *ListMarginLoansService) EndTime(endTime int64) *ListMarginLoansService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginLoansService) Current(current int64) *ListMarginLoansService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginLoansService) Size(size int64) *ListMarginLoansService {
	s.size = &size
	return s
}

// Do send request
func (s *ListMarginLoansService) Do(ctx context.Context, opts ...RequestOption) (res *MarginLoanResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/loan",
	}
	r.setParam("asset", s.asset)
	if s.txID != nil {
		r.setParam("txId", *s.txID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginLoanResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginLoanResponse define margin loan response
type MarginLoanResponse struct {
	Rows  []MarginLoan `json:"rows"`
	Total int64        `json:"total"`
}

// MarginLoan define margin loan
type MarginLoan struct {
	Asset     string               `json:"asset"`
	Principal string               `json:"principal"`
	Timestamp int64                `json:"timestamp"`
	Status    MarginLoanStatusType `json:"status"`
}
