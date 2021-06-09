package binance

import (
	"context"
	"encoding/json"
)

const (
	//目前支持的type划转类型
	AssetTransferType_MAIN_UMFUTURE = "MAIN_UMFUTURE"
	AssetTransferType_MAIN_CMFUTURE = "MAIN_CMFUTURE"
	AssetTransferType_UMFUTURE_MAIN = "UMFUTURE_MAIN"
	AssetTransferType_CMFUTURE_MAIN = "CMFUTURE_MAIN"
)


type AssetTransferService struct {
	c *Client

	assetType string
	asset string
	amount string
	recvWindow int64
}

func (s *AssetTransferService) AssetType(a string) {
	s.assetType = a
}

func (s *AssetTransferService) Asset(a string) {
	s.asset = a
}

func (s *AssetTransferService) Amount(a string) {
	s.amount = a
}



// Do send request
func (s *AssetTransferService) Do(ctx context.Context, opts ...RequestOption) (res *AssetTransferResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/asset/transfer",
		secType:  secTypeSigned,
	}

	r.setParams(params{
		"type": s.assetType,
		"asset":   s.asset,
		"amount": s.amount,
		"recvWindow": s.recvWindow,
	})


	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(AssetTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}


type AssetTransferResponse struct {
	TranId  string     `json:"tranId"`
}

