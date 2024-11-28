package delivery

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TradeServiceTestSuite struct {
	baseTestSuite
}

func TestTradeService(t *testing.T) {
	suite.Run(t, new(TradeServiceTestSuite))
}

func (s *TradeServiceTestSuite) TestUserTrades() {
	res, err := s.client.NewListAccountTradeService().Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 2)
}
