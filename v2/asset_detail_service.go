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
		endpoint: "/wapi/v3/assetDetail.html",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(AssetDetailResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res.AssetDetails, nil
}

// AssetDetail represents the detail of an asset
type AssetDetail struct {
	MinWithdrawAmount string  `json:"minWithdrawAmount"`
	DepositStatus     bool    `json:"depositStatus"`
	WithdrawFee       float64 `json:"withdrawFee"`
	WithdrawStatus    bool    `json:"withdrawStatus"`
	DepositTip        string  `json:"depositTip"`
}

// AssetDetailResponse represents a response from AssetDetailService.
type AssetDetailResponse struct {
	Success      bool                   `json:"success"`
	AssetDetails map[string]AssetDetail `json:"assetDetail"`
}
