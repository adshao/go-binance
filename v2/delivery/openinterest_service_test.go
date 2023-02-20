package delivery

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type openInterestServiceTestSuite struct {
	baseTestSuite
}

func TestGetOpenInterestService(t *testing.T) {
	suite.Run(t, new(openInterestServiceTestSuite))
}

func (s *openInterestServiceTestSuite) TestGetOpenInterest() {
	data := []byte(`{
		"symbol": "BTCUSD_200626",
		"pair": "BTCUSD",
		"openInterest": "15004",
		"contractType": "CURRENT_QUARTER",
		"time": 1591261042378
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200626"
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetOpenInterestService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)

	e := &OpenInterest{
		Symbol:       symbol,
		Pair:         "BTCUSD",
		OpenInterest: "15004",
		ContractType: "CURRENT_QUARTER",
		Time:         1591261042378,
	}

	s.r().Equal(e.Symbol, res.Symbol, "Symbol")
	s.r().Equal(e.OpenInterest, res.OpenInterest, "OpenInterest")
	s.r().Equal(e.Time, res.Time, "Time")
}
