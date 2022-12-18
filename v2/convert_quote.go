package binance

import (
	"context"
	"net/http"
)

type ConvertGetQuoteService struct {
	c          *Client
	fromAsset  string
	toAsset    string
	fromAmount float64
	toAmount   float64
	validTime  string
}

// FromAsset set fromAsset
func (s *ConvertGetQuoteService) FromAsset(fromAsset string) *ConvertGetQuoteService {
	s.fromAsset = fromAsset
	return s
}

// ToAsset set toAsset
func (s *ConvertGetQuoteService) ToAsset(toAsset string) *ConvertGetQuoteService {
	s.toAsset = toAsset
	return s
}

// FromAmount set fromAmount
func (s *ConvertGetQuoteService) FromAmount(fromAmount float64) *ConvertGetQuoteService {
	s.fromAmount = fromAmount
	return s
}

// ToAmount set toAmount
func (s *ConvertGetQuoteService) ToAmount(toAmount float64) *ConvertGetQuoteService {
	s.toAmount = toAmount
	return s
}

// ValidTime set validTime
func (s *ConvertGetQuoteService) ValidTime(validTime string) *ConvertGetQuoteService {
	s.validTime = validTime
	return s
}

// Do send request
func (s *ConvertGetQuoteService) Do(ctx context.Context, opts ...RequestOption) (*ConvertGetQuote, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/getQuote",
		secType:  secTypeSigned,
	}
	r.setParam("fromAsset", s.fromAsset)
	r.setParam("toAsset", s.toAsset)
	if s.fromAmount > 0 {
		r.setParam("fromAmount", s.fromAmount)
	} else if s.toAmount > 0 {
		r.setParam("toAmount", s.toAmount)
	}
	if s.validTime != "" {
		r.setParam("validTime", s.validTime)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := ConvertGetQuote{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ConvertGetQuote define the convert accept quote
type ConvertGetQuote struct {
	QuoteID        string `json:"quoteId"`
	Ratio          string `json:"ratio"`
	InverseRatio   string `json:"inverseRatio"`
	ValidTimestamp int    `json:"validTimestamp"`
	ToAmount       string `json:"toAmount"`
	FromAmount     string `json:"fromAmount"`
}

type ConvertAcceptQuoteService struct {
	c       *Client
	quoteID string
}

// QuoteID set quoteID
func (s *ConvertAcceptQuoteService) QuoteID(quoteID string) *ConvertAcceptQuoteService {
	s.quoteID = quoteID
	return s
}

// Do send request
func (s *ConvertAcceptQuoteService) Do(ctx context.Context, opts ...RequestOption) (*ConvertAcceptQuote, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/acceptQuote",
		secType:  secTypeSigned,
	}
	r.setParam("quoteId", s.quoteID)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := ConvertAcceptQuote{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ConvertAcceptQuote define the convert accept quote
type ConvertAcceptQuote struct {
	OrderId     string `json:"orderId"`
	CreateTime  int    `json:"createTime"`
	OrderStatus string `json:"orderStatus"`
}

type ConvertOrderStatusService struct {
	c       *Client
	orderID string
}

// OrderID set quoteID
func (s *ConvertOrderStatusService) OrderID(orderID string) *ConvertOrderStatusService {
	s.orderID = orderID
	return s
}

// Do send request
func (s *ConvertOrderStatusService) Do(ctx context.Context, opts ...RequestOption) (*ConvertOrderStatus, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/orderStatus",
		secType:  secTypeSigned,
	}
	r.setParam("orderId", s.orderID)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := ConvertOrderStatus{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ConvertAcceptQuote define the convert accept quote
type ConvertOrderStatus struct {
	OrderId      int    `json:"orderId"`
	OrderStatus  string `json:"orderStatus"`
	FromAsset    string `json:"fromAsset"`
	FromAmount   string `json:"fromAmount"`
	ToAsset      string `json:"toAsset"`
	ToAmount     string `json:"toAmount"`
	Ratio        string `json:"ratio"`
	InverseRatio string `json:"inverseRatio"`
	CreateTime   int    `json:"createTime"`
}
