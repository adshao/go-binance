package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type SymbolConfService struct {
	c          *Client
	symbol     string
	recvWindow *int64
	timestamp  *int64
}

// Symbol set symbol
func (s *SymbolConfService) Symbol(symbol string) *SymbolConfService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *SymbolConfService) RecvWindow(recvWindow int64) *SymbolConfService {
	s.recvWindow = &recvWindow
	return s
}

// Timestamp set timestamp
func (s *SymbolConfService) Timestamp(timestamp int64) *SymbolConfService {
	s.timestamp = &timestamp
	return s

}

// Do send request
func (s *SymbolConfService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolConf, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/symbolConfig",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	if s.timestamp != nil {
		r.setParam("timestamp", *s.timestamp)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SymbolConf, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type SymbolConf struct {
	Symbol           string `json:"symbol"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  bool   `json:"isAutoAddMargin"`
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
}
