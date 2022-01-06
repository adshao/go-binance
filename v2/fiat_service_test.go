package binance

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type fiatServiceTestSuite struct {
	baseTestSuite
}

func TestFiatService(t *testing.T) {
	suite.Run(t, new(fiatServiceTestSuite))
}

func (s *fiatServiceTestSuite) TestFiatDepositWithdrawHistory() {
	data := []byte(`{
	   "code": "000000",
	   "message": "success",
	   "data": [
	   {
		  "orderNo":"7d76d611-0568-4f43-afb6-24cac7767365",
		  "fiatCurrency": "BRL",
		  "indicatedAmount": "10.00",
		  "amount": "10.00",
		  "totalFee": "0.00",
		  "method": "BankAccount",
		  "status": "Expired",
		  "createTime": 1626144956000,
		  "updateTime": 1626400907000
	   }
	   ],
	   "total": 1,
	   "success": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	transactionType := TransactionTypeDeposit
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"transactionType": transactionType,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewFiatDepositWithdrawHistoryService().
		TransactionType(transactionType).
		Do(newContext())
	s.r().NoError(err)
	e := &FiatDepositWithdrawHistory{
		Code:    "000000",
		Message: "success",
		Data: []FiatDepositWithdrawHistoryItem{
			{
				OrderNo:         "7d76d611-0568-4f43-afb6-24cac7767365",
				FiatCurrency:    "BRL",
				IndicatedAmount: "10.00",
				Amount:          "10.00",
				TotalFee:        "0.00",
				Method:          "BankAccount",
				Status:          "Expired",
				CreateTime:      1626144956000,
				UpdateTime:      1626400907000,
			},
		},
		Total:   1,
		Success: true,
	}
	s.assertFiatDepositWithdrawHistoryEqual(e, res)
}

func (s *fiatServiceTestSuite) assertFiatDepositWithdrawHistoryEqual(e, a *FiatDepositWithdrawHistory) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Message, a.Message, "Message")
	r.Equal(e.Total, a.Total, "Total")
	r.Equal(e.Success, a.Success, "Success")

	r.Len(a.Data, len(e.Data))
	for i := 0; i < len(a.Data); i++ {
		s.assertFiatDepositWithdrawHistoryItemEqual(&e.Data[i], &a.Data[i])
	}
}

func (s *fiatServiceTestSuite) assertFiatDepositWithdrawHistoryItemEqual(e, a *FiatDepositWithdrawHistoryItem) {
	r := s.r()
	r.Equal(e.OrderNo, a.OrderNo, "OrderNo")
	r.Equal(e.FiatCurrency, a.FiatCurrency, "FiatCurrency")
	r.Equal(e.IndicatedAmount, a.IndicatedAmount, "IndicatedAmount")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.TotalFee, a.TotalFee, "TotalFee")
	r.Equal(e.Method, a.Method, "Method")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
}
