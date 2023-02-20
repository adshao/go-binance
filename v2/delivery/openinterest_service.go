package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetOpenInterestService get present open interest of a specific symbol.
type GetOpenInterestService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetOpenInterestService) Symbol(symbol string) *GetOpenInterestService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetOpenInterestService) Do(ctx context.Context, opts ...RequestOption) (res *OpenInterest, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/openInterest",
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(OpenInterest)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type OpenInterest struct {
	Symbol       string `json:"symbol"`
	Pair         string `json:"pair"`
	OpenInterest string `json:"openInterest"`
	ContractType string `json:"contractType"`
	Time         int64  `json:"time"`
}
