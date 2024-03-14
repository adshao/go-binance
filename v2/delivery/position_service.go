package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChangeLeverageService change user's initial leverage of specific symbol market
type ChangeLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeLeverageService) Symbol(symbol string) *ChangeLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeLeverageService) Leverage(leverage int) *ChangeLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *SymbolLeverage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/dapi/v1/leverage",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SymbolLeverage)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SymbolLeverage define leverage info of symbol
type SymbolLeverage struct {
	Leverage    int    `json:"leverage"`
	MaxQuantity string `json:"maxQty"`
	Symbol      string `json:"symbol"`
}

// ChangeMarginTypeService change user's margin type of specific symbol market
type ChangeMarginTypeService struct {
	c          *Client
	symbol     string
	marginType MarginType
}

// Symbol set symbol
func (s *ChangeMarginTypeService) Symbol(symbol string) *ChangeMarginTypeService {
	s.symbol = symbol
	return s
}

// MarginType set margin type
func (s *ChangeMarginTypeService) MarginType(marginType MarginType) *ChangeMarginTypeService {
	s.marginType = marginType
	return s
}

// Do send request
func (s *ChangeMarginTypeService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/dapi/v1/marginType",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":     s.symbol,
		"marginType": s.marginType,
	})
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePositionMarginService update isolated position margin
type UpdatePositionMarginService struct {
	c            *Client
	symbol       string
	positionSide *PositionSideType
	amount       string
	actionType   int
}

// Symbol set symbol
func (s *UpdatePositionMarginService) Symbol(symbol string) *UpdatePositionMarginService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *UpdatePositionMarginService) PositionSide(positionSide PositionSideType) *UpdatePositionMarginService {
	s.positionSide = &positionSide
	return s
}

// Amount set position margin amount
func (s *UpdatePositionMarginService) Amount(amount string) *UpdatePositionMarginService {
	s.amount = amount
	return s
}

// Type set action type: 1: Add position margin，2: Reduce position margin
func (s *UpdatePositionMarginService) Type(actionType int) *UpdatePositionMarginService {
	s.actionType = actionType
	return s
}

// Do send request
func (s *UpdatePositionMarginService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/dapi/v1/positionMargin",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
		"amount": s.amount,
		"type":   s.actionType,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	r.setFormParams(m)

	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ChangePositionModeService change user's position mode
type ChangePositionModeService struct {
	c        *Client
	dualSide string
}

// Change user's position mode: true - Hedge Mode, false - One-way Mode
func (s *ChangePositionModeService) DualSide(dualSide bool) *ChangePositionModeService {
	if dualSide {
		s.dualSide = "true"
	} else {
		s.dualSide = "false"
	}
	return s
}

// Do send request
func (s *ChangePositionModeService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/dapi/v1/positionSide/dual",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"dualSidePosition": s.dualSide,
	})
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetPositionModeService get user's position mode
type GetPositionModeService struct {
	c *Client
}

// Response of user's position mode
type PositionMode struct {
	DualSidePosition bool `json:"dualSidePosition"`
}

// Do send request
func (s *GetPositionModeService) Do(ctx context.Context, opts ...RequestOption) (res *PositionMode, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/positionSide/dual",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{})
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = &PositionMode{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
