package binance

import (
	"context"
	"encoding/json"
)

// ListLendingPurchaseService fetches the saving purchases
type ListLendingPurchaseService struct {
	c           *Client
	lendingType *string
	asset       *string
	startTime   *int64
	endTime     *int64
	current     *int
	size        *int
}

// Asset sets the asset parameter.
func (s *ListLendingPurchaseService) Asset(asset string) *ListLendingPurchaseService {
	s.asset = &asset
	return s
}

// Current set the currently querying page. Start from 1. Default:1
func (s *ListLendingPurchaseService) Current(current int) *ListLendingPurchaseService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *ListLendingPurchaseService) Size(size int) *ListLendingPurchaseService {
	s.size = &size
	return s
}

// LendingType sets the lendingType parameter.
func (s *ListLendingPurchaseService) LendingType(lendingType string) *ListLendingPurchaseService {
	s.lendingType = &lendingType
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListLendingPurchaseService) StartTime(startTime int64) *ListLendingPurchaseService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListLendingPurchaseService) EndTime(endTime int64) *ListLendingPurchaseService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *ListLendingPurchaseService) Do(ctx context.Context) (*[]LendingPurchaseResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/lending/union/purchaseRecord",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	} else {
		r.setParam("size", 100)
	}
	if s.lendingType != nil {
		r.setParam("lendingType", *s.lendingType)
	} else {
		r.setParam("lendingType", `DAILY`)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	r.setParam("recvWindow", 60000)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new([]LendingPurchaseResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LendingPurchaseResponse represents a response from ListLendingPurchaseService.
type LendingPurchaseResponse struct {
	Amount      string `json:"amount"`
	Asset       string `json:"asset"`
	CreateTime  int64  `json:"createTime"`
	LendingType string `json:"lendingType"`
	ProductName string `json:"productName"`
	PurchaseID  int    `json:"purchaseId"`
	Lot         int    `json:"lot"` //When lendingType == REGULAR
	Status      string `json:"status"`
}
