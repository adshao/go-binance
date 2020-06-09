package binance

import (
	"context"
	"encoding/json"
)

// ListDepositsService fetches deposit history.
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-history-user_data
type ListDepositsService struct {
	c         *Client
	asset     *string
	status    *int
	startTime *int64
	endTime   *int64
}

// Asset sets the asset parameter.
func (s *ListDepositsService) Asset(asset string) *ListDepositsService {
	s.asset = &asset
	return s
}

// Status sets the status parameter.
func (s *ListDepositsService) Status(status int) *ListDepositsService {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDepositsService) StartTime(startTime int64) *ListDepositsService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDepositsService) EndTime(endTime int64) *ListDepositsService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *ListDepositsService) Do(ctx context.Context) (deposits []*Deposit, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/depositHistory.html",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, err := s.c.callAPI(ctx, r)
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

// DepositHistoryResponse represents a response from ListDepositsService.
type DepositHistoryResponse struct {
	Success  bool       `json:"success"`
	Deposits []*Deposit `json:"depositList"`
}

// Deposit represents a single deposit entry.
type Deposit struct {
	InsertTime int64   `json:"insertTime"`
	Amount     float64 `json:"amount"`
	Asset      string  `json:"asset"`
	Address    string  `json:"address"`
	AddressTag string  `json:"addressTag"`
	TxID       string  `json:"txId"`
	Status     int     `json:"status"`
}

// GetDepositsAddressService retrieves the details of a deposit address.
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-address-supporting-network-user_data
type GetDepositsAddressService struct {
	c      *Client
	asset  string
	status *bool
}

// Asset sets the asset parameter (MANDATORY).
func (s *GetDepositsAddressService) Asset(v string) *GetDepositsAddressService {
	s.asset = v
	return s
}

// Status sets the status parameter.
func (s *GetDepositsAddressService) Status(v bool) *GetDepositsAddressService {
	s.status = &v
	return s
}

// Do sends the request.
func (s *GetDepositsAddressService) Do(ctx context.Context) (*GetDepositAddressResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/depositAddress.html",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if v := s.status; v != nil {
		r.setParam("status", *v)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &GetDepositAddressResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetDepositAddressResponse represents a response from GetDepositsAddressService.
type GetDepositAddressResponse struct {
	Success    bool   `json:"success"`
	Address    string `json:"address"`
	AddressTag string `json:"addressTag"`
	Asset      string `json:"asset"`
	URL        string `json:"url"`
}
