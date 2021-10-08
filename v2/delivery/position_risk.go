package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetPositionRiskService get account balance
type GetPositionRiskService struct {
	c           *Client
	pair        *string
	marginAsset *string
}

// MarginAsset set margin asset
func (s *GetPositionRiskService) MarginAsset(marginAsset string) *GetPositionRiskService {
	s.marginAsset = &marginAsset
	return s
}

// Pair set pair
func (s *GetPositionRiskService) Pair(pair string) *GetPositionRiskService {
	s.pair = &pair
	return s
}

// Do send request
func (s *GetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/positionRisk",
		secType:  secTypeSigned,
	}
	if s.marginAsset != nil {
		r.setParam("marginAsset", *s.marginAsset)
	}
	if s.pair != nil {
		r.setParam("pair", *s.pair)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionRisk{}, err
	}
	res = make([]*PositionRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionRisk{}, err
	}
	return res, nil
}

// PositionRisk define position risk info
type PositionRisk struct {
	Symbol           string `json:"symbol"`
	PositionAmt      string `json:"positionAmt"`
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	LiquidationPrice string `json:"liquidationPrice"`
	Leverage         string `json:"leverage"`
	MaxQuantity      string `json:"maxQty"`
	MarginType       string `json:"marginType"`
	IsolatedMargin   string `json:"isolatedMargin"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	PositionSide     string `json:"positionSide"`
}
