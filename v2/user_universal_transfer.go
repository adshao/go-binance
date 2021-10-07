/*******************************************************************************
** @Author:					Alessandro Maioli <Nicemine>
** @Email:					alessandro.maioli@gmail.com
** @Date:					Wednesday 29 September 2021 - 18:51:45
** @Filename:				user_universal_transfer.go
**
** @Last modified by:
*******************************************************************************/

package binance

import (
	"context"
	"encoding/json"
)

// CreateUserUniversalTransferService submits a transfer request.
//
// See https://binance-docs.github.io/apidocs/spot/en/#user-universal-transfer-user_data
type CreateUserUniversalTransferService struct {
	c          *Client
	types      string
	asset      string
	amount     float64
	fromSymbol *string
	toSymbol   *string
}

// Coin sets the coin parameter (MANDATORY).
func (s *CreateUserUniversalTransferService) Type(v string) *CreateUserUniversalTransferService {
	s.types = v
	return s
}

// Asset sets the Asset parameter (MANDATORY).
func (s *CreateUserUniversalTransferService) Asset(v string) *CreateUserUniversalTransferService {
	s.asset = v
	return s
}

// Amount sets the Amount parameter (MANDATORY).
func (s *CreateUserUniversalTransferService) Amount(v float64) *CreateUserUniversalTransferService {
	s.amount = v
	return s
}

// fromSymbol sets the fromSymbol parameter
func (s *CreateUserUniversalTransferService) FromSymbol(v string) *CreateUserUniversalTransferService {
	s.fromSymbol = &v
	return s
}

// toSymbol sets the toSymbol parameter
func (s *CreateUserUniversalTransferService) ToSymbol(v string) *CreateUserUniversalTransferService {
	s.toSymbol = &v
	return s
}

// Do sends the request.
func (s *CreateUserUniversalTransferService) Do(ctx context.Context) (*CreateUserUniversalTransferResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/asset/transfer",
		secType:  secTypeSigned,
	}
	r.setParam("types", s.types)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	if v := s.fromSymbol; v != nil {
		r.setParam("fromSymbol", *v)
	}
	if v := s.toSymbol; v != nil {
		r.setParam("toSymbol", *v)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &CreateUserUniversalTransferResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateUserUniversalTransferResponse represents a response from CreateUserUniversalTransferResponse.
type CreateUserUniversalTransferResponse struct {
	ID int64 `json:"tranId"`
}

// ListUserUniversalTransfer fetches transfer history.
//
// See https://binance-docs.github.io/apidocs/spot/en/#query-user-universal-transfer-history-user_data
// type ListUserUniversalTransfer struct {
// 	c          *Client
// 	types      string
// 	startTime  *int64
// 	endTime    *int64
// 	current    *int
// 	size       *int
// 	fromSymbol *string
// 	toSymbol   *string
// }

// // Type sets the type parameter.
// func (s *ListUserUniversalTransfer) Type(v string) *ListUserUniversalTransfer {
// 	s.types = v
// 	return s
// }

// // StartTime sets the startTime parameter.
// func (s *ListUserUniversalTransfer) StartTime(v int64) *ListUserUniversalTransfer {
// 	s.startTime = &v
// 	return s
// }

// // EndTime sets the startTime parameter.
// func (s *ListUserUniversalTransfer) EndTime(v int64) *ListUserUniversalTransfer {
// 	s.startTime = &v
// 	return s
// }

// // Current sets the current parameter.
// func (s *ListUserUniversalTransfer) Current(v int) *ListUserUniversalTransfer {
// 	s.current = &v
// 	return s
// }

// // Size sets the size parameter.
// func (s *ListUserUniversalTransfer) Size(v int) *ListUserUniversalTransfer {
// 	s.current = &v
// 	return s
// }

// // FromSymbol set fromSymbol
// func (s *ListUserUniversalTransfer) FromSymbol(v string) *ListUserUniversalTransfer {
// 	s.fromSymbol = &v
// 	return s
// }

// // ToSymbol set toSymbol
// func (s *ListUserUniversalTransfer) ToSymbol(v string) *ListUserUniversalTransfer {
// 	s.toSymbol = &v
// 	return s
// }

// // Do sends the request.
// func (s *ListUserUniversalTransfer) Do(ctx context.Context) (res []*TransferResult, err error) {
// 	r := &request{
// 		method:   "GET",
// 		endpoint: "/sapi/v1/asset/transfer",
// 		secType:  secTypeSigned,
// 	}
// 	r.setParam("types", s.types)
// 	if s.startTime != nil {
// 		r.setParam("startTime", *s.startTime)
// 	}
// 	if s.endTime != nil {
// 		r.setParam("endTime", *s.endTime)
// 	}
// 	if s.current != nil {
// 		r.setParam("current", *s.current)
// 	}
// 	if s.size != nil {
// 		r.setParam("size", *s.size)
// 	}
// 	if s.fromSymbol != nil {
// 		r.setParam("fromSymbol", *s.fromSymbol)
// 	}
// 	if s.toSymbol != nil {
// 		r.setParam("toSymbol", *s.toSymbol)
// 	}
// 	data, err := s.c.callAPI(ctx, r)
// 	if err != nil {
// 		return
// 	}
// 	res = make([]*TransferResult, 0)
// 	err = json.Unmarshal(data, &res)
// 	if err != nil {
// 		return
// 	}
// 	return res, nil
// }

// // Withdraw represents a single withdraw entry.
// type TransferResult struct {
// 	Total    uint8      `json:"total"`
// 	Transfer []Transfer `json:"rows"`
// }

// type Transfer struct {
// 	Asset     string `json:"asset"`
// 	Amount    string `json:"amount"`
// 	Type      string `json:"type"`
// 	Status    string `json:"status"`
// 	TranId    string `json:"tranId"`
// 	Timestamp string `json:"timestamp"`
// }
