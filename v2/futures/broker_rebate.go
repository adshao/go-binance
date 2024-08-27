package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	BrokerRebateStatusPending uint8 = 0
	BrokerRebateStatusSuccess uint8 = 1
	BrokerRebateStatusFailed  uint8 = 2

	BrokerRebateFuturesTypeUSDT int64 = 1
	BrokerRebateFuturesTypeCoin int64 = 2

	BrokerRebateMaxLimit int = 100
)

// BrokerRebateService queries broker futures commission rebate recent records
type BrokerRebateService struct {
	c           *Client
	futuresType int64
	startTime   int64
	endTime     int64
	page        *int
	size        *int
	recvWindow  *int64
}

// FuturesType sets futuresType (1: USDT Futures, 2: Coin Futures)
func (s *BrokerRebateService) FuturesType(futuresType int64) *BrokerRebateService {
	s.futuresType = futuresType
	return s
}

// StartTime sets startTime
func (s *BrokerRebateService) StartTime(startTime int64) *BrokerRebateService {
	s.startTime = startTime
	return s
}

// EndTime sets endTime
func (s *BrokerRebateService) EndTime(endTime int64) *BrokerRebateService {
	s.endTime = endTime
	return s
}

// Page sets page
func (s *BrokerRebateService) Page(page int) *BrokerRebateService {
	s.page = &page
	return s
}

// Size sets size
func (s *BrokerRebateService) Size(size int) *BrokerRebateService {
	s.size = &size
	return s
}

// RecvWindow sets recvWindow
func (s *BrokerRebateService) RecvWindow(recvWindow int64) *BrokerRebateService {
	s.recvWindow = &recvWindow
	return s
}

// Do sends the request
func (s *BrokerRebateService) Do(ctx context.Context, opts ...RequestOption) (res []*RebateRecord, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/rebate/futures/recentRecord",
		secType:  secTypeSigned,
	}
	r.setParam("futuresType", s.futuresType)
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*RebateRecord{}, err
	}
	res = make([]*RebateRecord, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*RebateRecord{}, err
	}
	return res, nil
}

// RebateRecord defines the structure of the futures rebate record
type RebateRecord struct {
	SubAccountID string `json:"subaccountId"`
	Income       string `json:"income"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	TradeID      int64  `json:"tradeId"`
	Time         int64  `json:"time"`
	Status       int    `json:"status"`
}
