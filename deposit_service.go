package binance

import (
	"context"
	"encoding/json"
)

// ListDepositsService list deposits
type ListDepositsService struct {
	c         *Client
	asset     *string
	status    *int
	startTime *int64
	endTime   *int64
}

// Asset set asset
func (s *ListDepositsService) Asset(asset string) *ListDepositsService {
	s.asset = &asset
	return s
}

// Status set status
func (s *ListDepositsService) Status(status int) *ListDepositsService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *ListDepositsService) StartTime(startTime int64) *ListDepositsService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListDepositsService) EndTime(endTime int64) *ListDepositsService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ListDepositsService) Do(ctx context.Context, opts ...RequestOption) (deposits []*Deposit, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/depositHistory.html",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.asset != nil {
		m["asset"] = *s.asset
	}
	if s.status != nil {
		m["status"] = *s.status
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	r.setParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res := new(DepositHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res.Deposits, nil
}

// DepositHistoryResponse define deposit history
type DepositHistoryResponse struct {
	Success  bool       `json:"success"`
	Deposits []*Deposit `json:"depositList"`
}

// Deposit define deposit info
type Deposit struct {
	InsertTime int64   `json:"insertTime"`
	Amount     float64 `json:"amount"`
	Asset      string  `json:"asset"`
	Address    string  `json:"address"`
	AddressTag string  `json:"addressTag"`
	TxID       string  `json:"txId"`
	Status     int     `json:"status"`
}

// GetDepositAddressService get deposit address
type GetDepositAddressService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetDepositAddressService) Asset(asset string) *GetDepositAddressService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (res *DepositAddressResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/depositAddress.html",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(DepositAddressResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DepositAddressResponse deposit address
type DepositAddressResponse struct {
	Address    string `json:"address"`
	Success    bool   `json:"success"`
	AddressTag string `json:"addressTag"`
	Asset      string `json:"asset"`
}
