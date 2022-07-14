package binance

import (
	"context"
	"fmt"
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
	r.setParam("symbol", s.symbol)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*CommissionRate{}, err
	}
	j, err := newJSON(data)
	if err != nil {
		return []*CommissionRate{}, err
	}
	num := len(j.MustArray())
	res = make([]*CommissionRate, num)
	for i := 0; i < num; i++ {
		item := j.GetIndex(i)
		if len(item.MustArray()) < 11 {
			err = fmt.Errorf("invalid Commission Rate response")
			return []*CommissionRate{}, err
		}
		res[i] = &CommissionRate{
			Symbol:              item.GetIndex(0).MustString(),
			MakerCommissionRate: item.GetIndex(1).MustString(),
			TakerCommissionRate: item.GetIndex(2).MustString(),
		}
	}
	return res, nil
}

// Commission Rate
type CommissionRate struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}
