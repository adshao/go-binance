package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type basePositionMarginHistoryTestSuite struct {
	baseTestSuite
}

func TestPositionMarginHistoryTestService(t *testing.T) {
	suite.Run(t, new(positionMarginHistoryServiceTestSuite))
}

type positionMarginHistoryServiceTestSuite struct {
	basePositionMarginHistoryTestSuite
}

func (s *positionMarginHistoryServiceTestSuite) TestPositionMarginHistory() {
	data := []byte(`[
		{
			"amount": "23.36332311",
			"asset": "USDT",
			"symbol": "BTCUSDT",
			"time": 1578047897183,
			"type": 1
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
	orders, err := s.client.NewGetPositionMarginHistoryService().Symbol(symbol).
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &PositionMarginHistory{
		Amount: "23.36332311",
		Asset:  "USDT",
		Symbol: "BTCUSDT",
		Time:   1578047897183,
		Type:   1,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *positionMarginHistoryServiceTestSuite) assertOrderEqual(e, a *PositionMarginHistory) {
	r := s.r()
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Time, e.Time, "Time")
	r.Equal(e.Type, a.Type, "Type")
}
