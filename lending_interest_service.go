package binance

import (
	"context"
	"encoding/json"
)

// ListLendingInterestService fetches the saving purchases
type ListLendingInterestService struct {
	c           *Client
	lendingType *string
	asset       *string
	startTime   *int64
	endTime     *int64
	size        *int
	current     *int
}

// Asset sets the asset parameter.
func (s *ListLendingInterestService) Asset(asset string) *ListLendingInterestService {
	s.asset = &asset
	return s
}

// Size sets the size parameter.
func (s *ListLendingInterestService) Size(size int) *ListLendingInterestService {
	s.size = &size
	return s
}

// Current set the currently querying page. Start from 1. Default:1
func (s *ListLendingInterestService) Current(current int) *ListLendingInterestService {
	s.current = &current
	return s
}

// LendingType sets the lendingType parameter.
func (s *ListLendingInterestService) LendingType(lendingType string) *ListLendingInterestService {
	s.lendingType = &lendingType
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListLendingInterestService) StartTime(startTime int64) *ListLendingInterestService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListLendingInterestService) EndTime(endTime int64) *ListLendingInterestService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *ListLendingInterestService) Do(ctx context.Context) (*[]LendingInterestResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/lending/union/interestHistory",
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
	res := new([]LendingInterestResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LendingInterestResponse represents a response from ListLendingInterestService.
type LendingInterestResponse struct {
	Time        int64  `json:"time"`
	Asset       string `json:"asset"`
	ProductName string `json:"productName"`
	Interest    string `json:"interest"`
	LendingType string `json:"lendingType"`
}
