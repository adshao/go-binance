package binance

import (
	"context"
	"encoding/json"
)

//查询手续费服务
type AssetTradeFeeService struct {
	c *Client

	symbol     string
	recvWindow int64
}

func (s *AssetTradeFeeService) Symbol(a string) *AssetTradeFeeService {
	s.symbol = a
	return s
}

func (s *AssetTradeFeeService) RecvWindow(a int64) *AssetTradeFeeService {
	s.recvWindow = a
	return s
}

// Do send request
func (s *AssetTradeFeeService) Do(ctx context.Context, opts ...RequestOption) (res []AssetTradeFeeResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/asset/tradeFee",
		secType:  secTypeSigned,
	}

	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	if s.recvWindow != 0 {
		r.setParam("recvWindow", s.recvWindow)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]AssetTradeFeeResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type AssetTradeFeeResponse struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}
