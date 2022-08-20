package binance

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type getBNBBurnServiceTestSuite struct {
	baseTestSuite
}

func TestGetBNBBurnService(t *testing.T) {
	suite.Run(t, new(getBNBBurnServiceTestSuite))
}

func (s *getBNBBurnServiceTestSuite) TestGetBNBBurn() {
	data := []byte(`
	{
   		"spotBNBBurn":true,
   		"interestBNBBurn": false   
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	bnbBurn, err := s.client.NewGetBNBBurnService().
		Do(context.Background())
	r := s.r()
	r.NoError(err)

	assertBnbBurn(r, bnbBurn)
}

func assertBnbBurn(r *require.Assertions, bnbBurn *BNBBurn) {
	r.Equal(true, bnbBurn.SpotBNBBurn)
	r.Equal(false, bnbBurn.InterestBNBBurn)
}

type toggleBNBBurnServiceTestSuite struct {
	baseTestSuite
}

func TestToggleBNBBurnService(t *testing.T) {
	suite.Run(t, new(toggleBNBBurnServiceTestSuite))
}

func (s *toggleBNBBurnServiceTestSuite) TestToggleBNBBurn() {
	data := []byte(`
	{
   		"spotBNBBurn":true,
   		"interestBNBBurn": false   
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			`spotBNBBurn`:     true,
			`interestBNBBurn`: false,
		})
		s.assertRequestEqual(e, r)
	})

	bnbBurn, err := s.client.NewToggleBNBBurnService().
		SpotBNBBurn(true).
		InterestBNBBurn(false).
		Do(context.Background())
	r := s.r()
	r.NoError(err)

	assertBnbBurn(r, bnbBurn)
}
