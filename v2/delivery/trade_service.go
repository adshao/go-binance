package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListAccountTradeService define account trade list service
type ListAccountTradeService struct {
	c         *Client
	symbol    string
	orderId   *int64
	startTime *int64
	endTime   *int64
	fromID    *int64
	limit     *int
}

// Symbol set symbol
func (s *ListAccountTradeService) Symbol(symbol string) *ListAccountTradeService {
	s.symbol = symbol
	return s
}

// OrderID set orderId
func (s *ListAccountTradeService) OrderID(orderID int64) *ListAccountTradeService {
	s.orderId = &orderID
	return s
}

// StartTime set startTime
func (s *ListAccountTradeService) StartTime(startTime int64) *ListAccountTradeService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListAccountTradeService) EndTime(endTime int64) *ListAccountTradeService {
	s.endTime = &endTime
	return s
}

// FromID set fromID
func (s *ListAccountTradeService) FromID(fromID int64) *ListAccountTradeService {
	s.fromID = &fromID
	return s
}

// Limit set limit
func (s *ListAccountTradeService) Limit(limit int) *ListAccountTradeService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListAccountTradeService) Do(ctx context.Context, opts ...RequestOption) (res []*AccountTrade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/userTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromID", *s.fromID)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*AccountTrade{}, err
	}
	res = make([]*AccountTrade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*AccountTrade{}, err
	}
	return res, nil
}

// AccountTrade define account trade
type AccountTrade struct {
	Symbol          string           `json:"symbol"`
	ID              int64            `json:"id"`
	OrderID         int64            `json:"orderId"`
	Pair            string           `json:"pair"`
	Side            SideType         `json:"side"`
	Price           string           `json:"price"`
	Quantity        string           `json:"qty"`
	RealizedPnl     string           `json:"realizedPnl"`
	MarginAsset     string           `json:"marginAsset"`
	BaseQuantity    string           `json:"baseQty"`
	Commission      string           `json:"commission"`
	CommissionAsset string           `json:"commissionAsset"`
	Time            int64            `json:"time"`
	PositionSide    PositionSideType `json:"positionSide"`
	Buyer           bool             `json:"buyer"`
	Maker           bool             `json:"maker"`
}
