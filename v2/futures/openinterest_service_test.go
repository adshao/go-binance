package futures

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
		"openInterest": "10659.509", 
		"symbol": "BTCUSDT",
		"time": 1589437530011
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
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
		Time:         1589437530011,
		OpenInterest: "10659.509",
	}

	s.r().Equal(e.Symbol, res.Symbol, "Symbol")
	s.r().Equal(e.OpenInterest, res.OpenInterest, "OpenInterest")
	s.r().Equal(e.Time, res.Time, "Time")
}
