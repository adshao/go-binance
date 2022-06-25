package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type swapPoolServiceTestSuite struct {
	baseTestSuite
}

func TestSwapPoolService(t *testing.T) {
	suite.Run(t, new(swapPoolServiceTestSuite))
}

func (s *swapPoolServiceTestSuite) TestGetSwapPoolList() {
	data := []byte(`[
		{
			"poolId": 2,
			"poolName": "BUSD/USDT",
			"assets": [
				"BUSD",
				"USDT"
			]
		},
		{
			"poolId": 3,
			"poolName": "BUSD/DAI",
			"assets": [
				"BUSD",
				"DAI"
			]
		},
		{
			"poolId": 4,
			"poolName": "USDT/DAI",
			"assets": [
				"USDT",
				"DAI"
			]
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetSwapPoolService().Do(newContext())
	s.r().NoError(err)
	e := []*SwapPool{
		{
			PoolId:   2,
			PoolName: "BUSD/USDT",
			Assets:   []string{"BUSD", "USDT"},
		},
		{
			PoolId:   3,
			PoolName: "BUSD/DAI",
			Assets:   []string{"BUSD", "DAI"},
		},
		{
			PoolId:   4,
			PoolName: "USDT/DAI",
			Assets:   []string{"USDT", "DAI"},
		},
	}
	s.assertSwapPoolSliceEqual(e, 3, res)
}

func (s *swapPoolServiceTestSuite) assertSwapPoolSliceEqual(e []*SwapPool, expectLen int, a []*SwapPool) {
	r := s.r()

	r.Len(a, expectLen)
	for i := 0; i < len(a); i++ {
		s.assertSwapPoolEqual(e[i], a[i])
	}
}

func (s *swapPoolServiceTestSuite) assertSwapPoolEqual(e *SwapPool, a *SwapPool) {
	r := s.r()

	r.Equal(e.PoolId, a.PoolId, "PoolId")
	r.Equal(e.PoolName, a.PoolName, "PoolName")
	r.ElementsMatch(e.Assets, a.Assets, "Assets")
}
