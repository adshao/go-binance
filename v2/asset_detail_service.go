package binance

import (
	"context"
	"encoding/json"
)

// GetAssetDetailService fetches all asset detail.
//
// See https://binance-docs.github.io/apidocs/spot/en/#asset-detail-user_data
type GetAssetDetailService struct {
	c *Client
}

// Do sends the request.
func (s *GetAssetDetailService) Do(ctx context.Context) (assetDetails map[string]AssetDetail, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/asset/assetDetail",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}

	res := make(map[string]AssetDetail, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// AssetDetail represents the detail of an asset
type AssetDetail struct {
	MinWithdrawAmount string `json:"minWithdrawAmount"`
	DepositStatus     bool   `json:"depositStatus"`
	WithdrawFee       string `json:"withdrawFee"`
	WithdrawStatus    bool   `json:"withdrawStatus"`
	DepositTip        string `json:"depositTip"`
}

// AssetDetailResponse represents a response from AssetDetailService.
type AssetDetailResponse struct {
	AssetDetails map[string]AssetDetail `json:"assetDetail"`
}
