package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type payServiceTestSuite struct {
	baseTestSuite
}

func TestPayService(t *testing.T) {
	suite.Run(t, new(payServiceTestSuite))
}

func (s *payServiceTestSuite) TestPayTradeHistory() {
	data := []byte(`{
	"code": "000000",
   	"message": "success",
   	"data": [
   		{
       		"orderType": "C2C", 
       		"transactionId": "M_P_71505104267788288",  
       		"transactionTime": 1610090460133, 
       		"amount": "23.72469206", 
       		"currency": "BNB", 
       		"fundsDetail": [ 
               {
                "currency": "USDT", 
                "amount": "1.2" 
                },
                {
                 "currency": "ETH",
                 "amount": "0.0001"
                }
          	]
     	}
   	],
  	"success": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewPayTradeHistoryService().Do(newContext())
	s.r().NoError(err)
	e := &PayTradeHistory{
		Code:    "000000",
		Message: "success",
		Data: []PayTradeItem{
			{
				OrderType:       "C2C",
				TransactionID:   "M_P_71505104267788288",
				TransactionTime: 1610090460133,
				Amount:          "23.72469206",
				Currency:        "BNB",
				FundsDetail: []FundsDetail{
					{
						Currency: "USDT",
						Amount:   "1.2",
					},
					{
						Currency: "ETH",
						Amount:   "0.0001",
					},
				},
			},
		},
		Success: true,
	}
	s.assertPayTradeHistoryEqual(e, res)
}

func (s *payServiceTestSuite) assertPayTradeHistoryEqual(e, a *PayTradeHistory) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Message, a.Message, "Message")
	r.Equal(e.Success, a.Success, "Success")

	r.Len(a.Data, len(e.Data))
	for i := 0; i < len(a.Data); i++ {
		s.assertPayTradeItemEqual(&e.Data[i], &a.Data[i])
	}
}

func (s *payServiceTestSuite) assertPayTradeItemEqual(e, a *PayTradeItem) {
	r := s.r()
	r.Equal(e.OrderType, a.OrderType, "OrderType")
	r.Equal(e.TransactionID, a.TransactionID, "TransactionID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Currency, a.Currency, "Currency")
	r.Equal(e.FundsDetail, a.FundsDetail, "FundsDetail")
}
