package binance

import (
	"context"
	"encoding/json"
)

type AssetTransferType string

const (
	//目前支持的type划转类型
	MAIN_UMFUTURE AssetTransferType = "MAIN_UMFUTURE"
	MAIN_CMFUTURE AssetTransferType = "MAIN_CMFUTURE"
	UMFUTURE_MAIN AssetTransferType = "UMFUTURE_MAIN"
	CMFUTURE_MAIN AssetTransferType = "CMFUTURE_MAIN"
)


type AssetTransferService struct {
	c *Client

	assetType AssetTransferType
	asset string
	amount string
	recvWindow int64
}

func (s *AssetTransferService) AssetType(a AssetTransferType) {
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



	if s.asset != "" {
		r.setParam("asset", s.asset)
	}
	if s.assetType != "" {
		r.setParam("type", s.assetType)
	}
	if s.amount != "" {
		r.setParam("amount", s.amount)
	}

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

