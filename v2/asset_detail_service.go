package binance

import (
	"context"
	"encoding/json"
)

// GetAssetDetailService fetches all asset detail.
//
// See https://binance-docs.github.io/apidocs/spot/en/#asset-detail-user_data
type GetAssetDetailService struct {
	c     *Client
	asset *string
}

// Asset sets the asset parameter.
func (s *GetAssetDetailService) Asset(asset string) *GetAssetDetailService {
	s.asset = &asset
	return s
}

// Do sends the request.
func (s *GetAssetDetailService) Do(ctx context.Context) (res map[string]AssetDetail, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/asset/assetDetail",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make(map[string]AssetDetail)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// AssetDetail represents the detail of an asset
type AssetDetail struct {
	MinWithdrawAmount float64 `json:"minWithdrawAmount"`
	DepositStatus     bool    `json:"depositStatus"`
	WithdrawFee       float64 `json:"withdrawFee"`
	WithdrawStatus    bool    `json:"withdrawStatus"`
	DepositTip        string  `json:"depositTip"`
}
