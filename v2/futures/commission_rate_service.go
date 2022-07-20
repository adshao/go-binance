package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type CommissionRateService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (service *CommissionRateService) Symbol(symbol string) *CommissionRateService {
	service.symbol = symbol
	return service
}

// Do send request
func (s *CommissionRateService) Do(ctx context.Context, opts ...RequestOption) (res []*CommissionRate, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/commissionRate",
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*CommissionRate{}, err
	}
	res = make([]*CommissionRate, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CommissionRate{}, err
	}
	return res, nil
}

// Commission Rate
type CommissionRate struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}
