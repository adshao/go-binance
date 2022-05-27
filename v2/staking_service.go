package binance

import (
	"context"
	"net/http"
)

// StakingProductPositionService fetches the staking product positions
type StakingProductPositionService struct {
	c         *Client
	product   StakingProduct
	productId *string
	asset     *string
	current   *int32
	size      *int32
}

// Product sets the product parameter.
func (s *StakingProductPositionService) Product(product StakingProduct) *StakingProductPositionService {
	s.product = product
	return s
}

// ProductId sets the productId parameter.
func (s *StakingProductPositionService) ProductId(productId string) *StakingProductPositionService {
	s.productId = &productId
	return s
}

// Asset sets the asset parameter.
func (s *StakingProductPositionService) Asset(asset string) *StakingProductPositionService {
	s.asset = &asset
	return s
}

// Current sets the current parameter.
func (s *StakingProductPositionService) Current(current int32) *StakingProductPositionService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *StakingProductPositionService) Size(size int32) *StakingProductPositionService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *StakingProductPositionService) Do(ctx context.Context) (*StakingProductPositions, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/position",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	if s.productId != nil {
		r.setParam("productId", *s.productId)
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(StakingProductPositions)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// StakingProductPositions represents a list of staking product positions.
type StakingProductPositions []StakingProductPosition

type StakingProductPosition struct {
	PositionId                 int64  `json:"positionId"`
	ProductId                  string `json:"productId"`
	Asset                      string `json:"asset"`
	Amount                     string `json:"amount"`
	PurchaseTime               int64  `json:"purchaseTime"`
	Duration                   int64  `json:"duration"`
	AccrualDays                int64  `json:"accrualDays"`
	RewardAsset                string `json:"rewardAsset"`
	APY                        string `json:"apy"`
	RewardAmount               string `json:"rewardAmt"`
	ExtraRewardAsset           string `json:"extraRewardAsset"`
	ExtraRewardAPY             string `json:"extraRewardAPY"`
	EstimatedExtraRewardAmount string `json:"estExtraRewardAmt"`
	NextInterestPay            string `json:"nextInterestPay"`
	NextInterestPayDate        string `json:"nextInterestPayDate"`
	PayInterestPeriod          int64  `json:"payInterestPeriod"`
	RedeemAmountEarly          string `json:"redeemAmountEarly"`
	InterestEndDate            int64  `json:"interestEndDate"`
	DeliverDate                int64  `json:"deliverDate"`
	RedeemPeriod               int64  `json:"redeemPeriod"`
	RedeemingAmount            string `json:"redeemingAmt"`
	PartialAmountDeliverDate   int64  `json:"partialAmtDeliverDate"`
	CanRedeemEarly             bool   `json:"canRedeemEarly"`
	Renewable                  bool   `json:"renewable"`
	Type                       string `json:"type"`
	Status                     string `json:"status"`
}
