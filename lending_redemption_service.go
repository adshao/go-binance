package binance

import (
	"context"
	"encoding/json"
)

// ListLendingRedemptionService fetches the saving purchases
type ListLendingRedemptionService struct {
	c           *Client
	lendingType *string
	asset       *string
	startTime   *int64
	endTime     *int64
	size        *int
	current     *int
}

// Asset sets the asset parameter.
func (s *ListLendingRedemptionService) Asset(asset string) *ListLendingRedemptionService {
	s.asset = &asset
	return s
}

// Current sets the current parameter.
func (s *ListLendingRedemptionService) Current(current int) *ListLendingRedemptionService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *ListLendingRedemptionService) Size(size int) *ListLendingRedemptionService {
	s.size = &size
	return s
}

// LendingType sets the lendingType parameter.
func (s *ListLendingRedemptionService) LendingType(lendingType string) *ListLendingRedemptionService {
	s.lendingType = &lendingType
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListLendingRedemptionService) StartTime(startTime int64) *ListLendingRedemptionService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListLendingRedemptionService) EndTime(endTime int64) *ListLendingRedemptionService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *ListLendingRedemptionService) Do(ctx context.Context) (*[]LendingRedemptionResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/lending/union/redemptionRecord",
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
	res := new([]LendingRedemptionResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LendingRedemptionResponse represents a response from ListLendingRedemptionService.
type LendingRedemptionResponse struct {
	CreateTime  int64  `json:"createTime"`
	Amount      string `json:"amount"`
	Asset       string `json:"asset"`
	Interest    string `json:"interest"` //lendingType == REGULAR
	Status      string `json:"status"`
	Type        string `json:"type"`
	ProjectName string `json:"projectName"`
	ProjectID   string `json:"projectId"`
	Principal   string `json:"principal"`
}
