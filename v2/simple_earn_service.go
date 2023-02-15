package binance

import (
	"context"
	"fmt"
	"net/http"
)

// https://binance-docs.github.io/apidocs/spot/en/#lending-account-user_data
type SimpleEarnService struct {
	c           *Client
	lendingType *string
	asset       *string
	startTime   *int64
	endTime     *int64
	current     *int64
	size        *int64
}

func (s *SimpleEarnService) LendingType(lendingType string) *SimpleEarnService {
	s.lendingType = &lendingType
	return s
}

func (s *SimpleEarnService) Asset(asset string) *SimpleEarnService {
	s.asset = &asset
	return s
}

func (s *SimpleEarnService) StartTime(startTime int64) *SimpleEarnService {
	s.startTime = &startTime
	return s
}

func (s *SimpleEarnService) EndTime(endTime int64) *SimpleEarnService {
	s.endTime = &endTime
	return s
}

func (s *SimpleEarnService) Current(current int64) *SimpleEarnService {
	s.current = &current
	return s
}

func (s *SimpleEarnService) Size(size int64) *SimpleEarnService {
	s.size = &size
	return s
}

type DataType int

const (
	Subscription DataType = iota + 1
	Redemption
	Interest
)

type EarnHistory struct {
	Amount      float64 `json:",string"`
	Asset       string
	CreateTime  int64
	LendingType string
	ProductName string
	PurchaseID  int64
	Status      string
	Type        string
	Time        int64
	Interest    float64 `json:",string"`
}

func (s *SimpleEarnService) Do(ctx context.Context, dataType DataType) (res []*EarnHistory, err error) {
	var endpoint string
	switch dataType {
	case Subscription:
		endpoint = "/sapi/v1/lending/union/purchaseRecord"
	case Redemption:
		endpoint = "/sapi/v1/lending/union/redemptionRecord"
	case Interest:
		endpoint = "/sapi/v1/lending/union/interestHistory"
	default:
		return nil, fmt.Errorf("type error")
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	if s.lendingType != nil {
		r.setParam("lendingType", *s.lendingType)
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.current)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make([]*EarnHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}
