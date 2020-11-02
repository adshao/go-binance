/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 19 June 2020 - 18:51:45
** @Filename:				0_dust_service.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package binance

import (
	"context"
	"encoding/json"
)

// ListDustLogService fetch small amounts of assets exchanged versus BNB
// See https://binance-docs.github.io/apidocs/spot/en/#dustlog-user_data
type ListDustLogService struct {
	c *Client
}

// Do sends the request.
func (s *ListDustLogService) Do(ctx context.Context) (withdraws *DustLogResponseWrapper, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/userAssetDribbletLog.html",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(DustLogResponseWrapper)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}

// DustLogResponseWrapper represents a response from ListDustLogService.
type DustLogResponseWrapper struct {
	Success bool        `json:"success"`
	Results *DustResult `json:"results"`
}

// DustResult represents the result of a DustLog API Call.
type DustResult struct {
	Total uint8     `json:"total"` //Total counts of exchange
	Rows  []DustRow `json:"rows"`
}

// DustRow represents one dust log row
type DustRow struct {
	TransferedTotal    string    `json:"transfered_total"`     //Total transfered BNB amount for this exchange.
	ServiceChargeTotal string    `json:"service_charge_total"` //Total service charge amount for this exchange.
	TranID             int       `json:"tran_id"`
	Logs               []DustLog `json:"logs"` //Details of this exchange.
	OperateTime        string    `json:"operate_time"`
}

// DustLog represents one dust log informations
type DustLog struct {
	TranID              int    `json:"tranId"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	UID                 string `json:"uid"`
	Amount              string `json:"amount"`
	OperateTime         string `json:"operateTime"` //The time of this exchange.
	TransferedAmount    string `json:"transferedAmount"`
	FromAsset           string `json:"fromAsset"`
}

// DustTransferService convert dust assets to BNB.
// See https://binance-docs.github.io/apidocs/spot/en/#dust-transfer-user_data
type DustTransferService struct {
	c     *Client
	asset []string
}

// Asset set asset.
func (s *DustTransferService) Asset(asset []string) *DustTransferService {
	s.asset = asset
	return s
}

// Do sends the request.
func (s *DustTransferService) Do(ctx context.Context) (withdraws *DustTransferResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/asset/dust",
		secType:  secTypeSigned,
	}
	for _, a := range s.asset {
		r.addParam("asset", a)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(DustTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}

// DustTransferResponse represents the response from DustTransferService.
type DustTransferResponse struct {
	TotalServiceCharge string                `json:"totalServiceCharge"`
	TotalTransfered    string                `json:"totalTransfered"`
	TransferResult     []*DustTransferResult `json:"transferResult"`
}

// DustTransferResult represents the result of a dust transfer.
type DustTransferResult struct {
	Amount              string `json:"amount"`
	FromAsset           string `json:"fromAsset"`
	OperateTime         int64  `json:"operateTime"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	TranID              int64  `json:"tranId"`
	TransferedAmount    string `json:"transferedAmount"`
}
