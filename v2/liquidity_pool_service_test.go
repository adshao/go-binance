package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type liquidityPoolServiceTestSuite struct {
	baseTestSuite
}

func TestLiquidityPoolService(t *testing.T) {
	suite.Run(t, new(liquidityPoolServiceTestSuite))
}

func (s *liquidityPoolServiceTestSuite) TestSwapService() {
	data := []byte(`{
		"swapId": 2314
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"quoteAsset": "USDT",
			"baseAsset":  "BUSD",
			"quoteQty":   1000,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewSwapService().
		QuoteAsset("USDT").
		BaseAsset("BUSD").
		QuoteQty(1000).
		Do(newContext())
	s.r().NoError(err)

	e := &SwapResponse{
		SwapId: 2314,
	}
	s.assertSwapEqual(e, res)
}

func (s *liquidityPoolServiceTestSuite) assertSwapEqual(e *SwapResponse, a *SwapResponse) {
	r := s.r()

	r.Equal(e.SwapId, a.SwapId, "SwapId")
}

func (s *liquidityPoolServiceTestSuite) TestGetSwapQuoteService() {
	data := []byte(`{
		"quoteAsset": "USDT",
		"baseAsset": "BUSD",
		"quoteQty": "300000",
		"baseQty": "299975",
		"price": "1.00008334",
		"slippage": "0.00007245",
		"fee": "120"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"quoteAsset": "USDT",
			"baseAsset":  "BUSD",
			"quoteQty":   1000,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetSwapQuoteService().
		QuoteAsset("USDT").
		BaseAsset("BUSD").
		QuoteQty(1000).
		Do(newContext())
	s.r().NoError(err)

	e := &GetSwapQuoteResponse{
		QuoteAsset: "USDT",
		BaseAsset:  "BUSD",
		QuoteQty:   "300000",
		BaseQty:    "299975",
		Price:      "1.00008334",
		Slippage:   "0.00007245",
		Fee:        "120",
	}
	s.assertGetSwapQuoteEqual(e, res)
}

func (s *liquidityPoolServiceTestSuite) assertGetSwapQuoteEqual(e *GetSwapQuoteResponse, a *GetSwapQuoteResponse) {
	r := s.r()

	r.Equal(e.BaseQty, a.BaseQty, "BaseQty")
	r.Equal(e.BaseAsset, a.BaseAsset, "BaseAsset")
	r.Equal(e.QuoteQty, a.QuoteQty, "QuoteQty")
	r.Equal(e.Fee, a.Fee, "Fee")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.Slippage, a.Slippage, "Slippage")
}

func (s *liquidityPoolServiceTestSuite) TestAddLiquidityPreviewService() {
	data := []byte(`{
		"quoteAsset": "USDT",
		"baseAsset": "BUSD",
		"quoteAmt": "300000",
		"baseAmt": "299975",
		"price": "1.00008334",
		"share": "1.23",
		"slippage": "0.00007245",
		"fee": "120"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"poolId":     2,
			"type":       LiquidityOperationTypeCombination,
			"quoteAsset": "USDT",
			"quoteQty":   1000,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewAddLiquidityPreviewService().
		PoolId(2).
		OperationType(LiquidityOperationTypeCombination).
		QuoteAsset("USDT").
		QuoteQty(1000).
		Do(newContext())
	s.r().NoError(err)

	e := &AddLiquidityPreviewResponse{
		QuoteAsset: "USDT",
		BaseAsset:  "BUSD",
		QuoteAmt:   "300000",
		BaseAmt:    "299975",
		Price:      "1.00008334",
		Share:      "1.23",
		Slippage:   "0.00007245",
		Fee:        "120",
	}
	s.assertAddLiquidityPreviewEqual(e, res)
}

func (s *liquidityPoolServiceTestSuite) assertAddLiquidityPreviewEqual(e *AddLiquidityPreviewResponse, a *AddLiquidityPreviewResponse) {
	r := s.r()

	r.Equal(e.BaseAmt, a.BaseAmt, "BaseAmt")
	r.Equal(e.BaseAsset, a.BaseAsset, "BaseAsset")
	r.Equal(e.QuoteAmt, a.QuoteAmt, "QuoteAmt")
	r.Equal(e.Fee, a.Fee, "Fee")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.Slippage, a.Slippage, "Slippage")
	r.Equal(e.Share, a.Share, "Share")
}

func (s *liquidityPoolServiceTestSuite) TestGetLiquidityPoolDetail() {
	data := []byte(`[
		{
			"poolId": 2,
			"poolName": "BUSD/USDT",
			"updateTime": 1565769342148,
			"liquidity": {
				"BUSD": "100000315.79",
				"USDT": "99999245.54"
			},
			"share": {
				"shareAmount": "12415",
				"sharePercentage": "0.00006207",
				"asset": {
					"BUSD": "6207.02",
					"USDT": "6206.95"
				}
			}
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"poolId": 2,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetLiquidityPoolDetailService().PoolId(2).Do(newContext())
	s.r().NoError(err)
	e := []*LiquidityPoolDetail{
		{
			PoolId:     2,
			PoolName:   "BUSD/USDT",
			UpdateTime: 1565769342148,
			Liquidity: map[string]string{
				"BUSD": "100000315.79",
				"USDT": "99999245.54",
			},
			Share: &PoolShareInformation{
				ShareAmount:     "12415",
				SharePercentage: "0.00006207",
				Assets: map[string]string{
					"BUSD": "6207.02",
					"USDT": "6206.95",
				},
			},
		},
	}
	s.assertLiquidityPoolDetailsEqual(e, 1, res)
}

func (s *liquidityPoolServiceTestSuite) assertLiquidityPoolDetailsEqual(e []*LiquidityPoolDetail, expectLen int, a []*LiquidityPoolDetail) {
	r := s.r()

	r.Len(a, expectLen)
	for i := 0; i < len(a); i++ {
		s.assertLiquidityPoolDetailEqual(e[i], a[i])
	}
}

func (s *liquidityPoolServiceTestSuite) assertLiquidityPoolDetailEqual(e *LiquidityPoolDetail, a *LiquidityPoolDetail) {
	r := s.r()

	r.Equal(e.PoolId, a.PoolId, "PoolId")
	r.Equal(e.PoolName, a.PoolName, "PoolName")
	r.Equal(e.Share.ShareAmount, a.Share.ShareAmount, "Share.ShareAmount")
	r.Equal(e.Share.SharePercentage, a.Share.SharePercentage, "Share.SharePercentage")

	r.Len(a.Liquidity, len(e.Liquidity), "Liquidity.Len")
	for k, v := range e.Liquidity {
		r.Equal(v, a.Liquidity[k], "Liquidity")
	}

	r.Len(a.Share.Assets, len(e.Share.Assets), "Share.Assets.Len")
	for k, v := range e.Share.Assets {
		r.Equal(v, a.Share.Assets[k], "Share.Assets")
	}
}

func (s *liquidityPoolServiceTestSuite) TestGetLiquidityPoolList() {
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

	res, err := s.client.NewGetAllLiquidityPoolService().Do(newContext())
	s.r().NoError(err)
	e := []*LiquidityPool{
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
	s.assertLiquidityPoolSliceEqual(e, 3, res)
}

func (s *liquidityPoolServiceTestSuite) assertLiquidityPoolSliceEqual(e []*LiquidityPool, expectLen int, a []*LiquidityPool) {
	r := s.r()

	r.Len(a, expectLen)
	for i := 0; i < len(a); i++ {
		s.assertLiquidityPoolEqual(e[i], a[i])
	}
}

func (s *liquidityPoolServiceTestSuite) assertLiquidityPoolEqual(e *LiquidityPool, a *LiquidityPool) {
	r := s.r()

	r.Equal(e.PoolId, a.PoolId, "PoolId")
	r.Equal(e.PoolName, a.PoolName, "PoolName")
	r.ElementsMatch(e.Assets, a.Assets, "Assets")
}
