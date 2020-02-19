package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseIncomeHistoryTestSuite struct {
	baseTestSuite
}

func TestIncomeHistoryTestService(t *testing.T) {
	suite.Run(t, new(incomeHistoryServiceTestSuite))
}

type incomeHistoryServiceTestSuite struct {
	baseIncomeHistoryTestSuite
}

func (s *incomeHistoryServiceTestSuite) TestIncomeHistory() {
	data := []byte(`[
		{
			"symbol": "BTCUSDT",
			"incomeType": "COMMISSION", 
			"income": "-0.01000000",
			"asset": "USDT",
			"info":"",  
			"time": 1570636800000
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	orders, err := s.client.NewGetIncomeHistoryService().Symbol(symbol).
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &IncomeHistory{
		Income:     "-0.01000000",
		Info:       "",
		Asset:      "USDT",
		Symbol:     "BTCUSDT",
		Time:       1578047897183,
		IncomeType: "COMMISSION",
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *incomeHistoryServiceTestSuite) assertOrderEqual(e, a *IncomeHistory) {
	r := s.r()
	r.Equal(e.Income, a.Income, "Income")
	r.Equal(e.Info, a.Info, "Info")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Time, e.Time, "Time")
	r.Equal(e.IncomeType, a.IncomeType, "IncomeType")
}
