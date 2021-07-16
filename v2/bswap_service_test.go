package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type bswapServiceTestSuite struct {
	baseTestSuite
}

func TestBSwapService(t *testing.T) {
	suite.Run(t, new(bswapServiceTestSuite))
}

func (s *bswapServiceTestSuite) TestListBSwapPoolsService() {
	data := []byte(`
        [
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
			}
		]
    `)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewListBSwapPoolsService().Do(newContext())
	s.r().NoError(err)
	e := []*BSwapPool{
		{
			PoolID:   2,
			PoolName: "BUSD/USDT",
			Assets:   []string{"BUSD", "USDT"},
		},
		{
			PoolID:   3,
			PoolName: "BUSD/DAI",
			Assets:   []string{"BUSD", "DAI"},
		},
	}
	s.assertBSwapPoolEqual(e, res)
}

func (s *bswapServiceTestSuite) assertBSwapPoolEqual(e, a []*BSwapPool) {
	r := s.r()
	r.Equal(len(e), len(a), "equal length")
	for i := 0; i < len(a); i++ {
		r.Equal(e[i].PoolID, a[i].PoolID, "PoolID")
		r.Equal(e[i].PoolName, a[i].PoolName, "PoolName")
		r.Equal(e[i].Assets, a[i].Assets, "Assets")
	}
}

func (s *bswapServiceTestSuite) TestGetBSwapPoolLiquidityInfoService() {
	data := []byte(`
        [
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
		]
    `)
	s.mockDo(data, nil)
	defer s.assertDo()

	poolID := int64(2)
	s.assertReq(func(r *request) {
		e := newSignedRequest().addParam("poolId", poolID)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetBSwapPoolLiquidityInfoService().PoolID(poolID).Do(newContext())
	s.r().NoError(err)
	e := []*BSwapPoolLiquidityInfo{
		{
			PoolID:     2,
			PoolName:   "BUSD/USDT",
			UpdateTime: 1565769342148,
			Liquidity: map[string]string{
				"BUSD": "100000315.79",
				"USDT": "99999245.54",
			},
			Share: &BSwapShare{
				ShareAmount:     "12415",
				SharePercentage: "0.00006207",
				Asset: map[string]string{
					"BUSD": "6207.02",
					"USDT": "6206.95",
				},
			},
		},
	}
	s.assertBSwapPoolLiquidityInfoEqual(e, res)
}

func (s *bswapServiceTestSuite) assertBSwapPoolLiquidityInfoEqual(e, a []*BSwapPoolLiquidityInfo) {
	r := s.r()
	r.Equal(len(e), len(a), "equal length")
	for i := 0; i < len(a); i++ {
		r.Equal(e[i].PoolID, a[i].PoolID, "PoolID")
		r.Equal(e[i].PoolName, a[i].PoolName, "PoolName")
		r.Equal(e[i].UpdateTime, a[i].UpdateTime, "UpdateTime")
		r.Equal(e[i].Liquidity, a[i].Liquidity, "Liquidity")
		r.Equal(e[i].Share.ShareAmount, a[i].Share.ShareAmount, "Share.ShareAmount")
		r.Equal(e[i].Share.SharePercentage, a[i].Share.SharePercentage, "Share.SharePercentage")
		r.Equal(e[i].Share.Asset, a[i].Share.Asset, "Share.Asset")
	}
}

func (s *bswapServiceTestSuite) TestAddBSwapLiquidityService() {
	data := []byte(`
		{    
			"operationId": 12341
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	poolID := int64(1234)
	asset := "USDT"
	quantity := 12.3
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"poolId":   poolID,
			"asset":    asset,
			"quantity": quantity,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewAddBSwapLiquidityService().PoolID(poolID).Asset(asset).Quantity(quantity).
		Do(newContext())
	s.r().NoError(err)
	e := &BSwapLiquidityTradeResponse{OperationId: 12341}
	s.assertBSwapLiquidityTradeResponseEqual(e, res)
}

func (s *bswapServiceTestSuite) assertBSwapLiquidityTradeResponseEqual(e, a *BSwapLiquidityTradeResponse) {
	r := s.r()
	r.Equal(e.OperationId, a.OperationId, "OperationId")
}

func (s *bswapServiceTestSuite) TestRemoveBSwapLiquidityService() {
	data := []byte(`
		{    
			"operationId": 12341
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	poolID := int64(1234)
	removalType := BSwapRemovalTypeSingle
	asset := "USDT"
	shareAmount := 12.3
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"poolId":      poolID,
			"type":        removalType,
			"asset":       asset,
			"shareAmount": shareAmount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewRemoveBSwapLiquidityService().PoolID(poolID).Type(removalType).Asset(asset).ShareAmount(shareAmount).
		Do(newContext())
	s.r().NoError(err)
	e := &BSwapLiquidityTradeResponse{OperationId: 12341}
	s.assertBSwapLiquidityTradeResponseEqual(e, res)
}

func (s *bswapServiceTestSuite) TestGetBSwapLiquidityOperationRecordService() {
	data := []byte(`
        [
			{
				"operationId": 12342,
				"poolId": 2,
				"poolName": "BUSD/USDT",
				"operation": "REMOVE",
				"status": 1,
				"updateTime": 1565769342159,
				"shareAmount": "10.1"
			}
		]
    `)
	s.mockDo(data, nil)
	defer s.assertDo()

	poolID := int64(2)
	operationID := int64(12342)
	operation := BSwapLiquidityOperationTypeAdd
	startTime := int64(1565769342148)
	endTime := int64(1565769342150)
	limit := int64(20)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"poolId":      poolID,
			"operationId": operationID,
			"operation":   operation,
			"startTime":   startTime,
			"endTime":     endTime,
			"limit":       limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetBSwapLiquidityOperationRecordService().
		OperationID(operationID).PoolID(poolID).Operation(operation).StartTime(startTime).EndTime(endTime).Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	e := []*BSwapLiquidityOperationRecord{
		{
			OperationId: 12342,
			PoolID:      2,
			PoolName:    "BUSD/USDT",
			Operation:   BSwapLiquidityOperationTypeRemove,
			Status:      BSwapStatusTypeSuccess,
			UpdateTime:  1565769342159,
			ShareAmount: "10.1",
		},
	}
	s.assertBSwapLiquidityOperationRecordEqual(e, res)
}

func (s *bswapServiceTestSuite) assertBSwapLiquidityOperationRecordEqual(e, a []*BSwapLiquidityOperationRecord) {
	r := s.r()
	r.Equal(len(e), len(a), "equal length")
	for i := 0; i < len(a); i++ {
		r.Equal(e[i].OperationId, a[i].OperationId, "OperationId")
		r.Equal(e[i].PoolID, a[i].PoolID, "PoolID")
		r.Equal(e[i].PoolName, a[i].PoolName, "PoolName")
		r.Equal(e[i].Operation, a[i].Operation, "Operation")
		r.Equal(e[i].Status, a[i].Status, "Status")
		r.Equal(e[i].UpdateTime, a[i].UpdateTime, "UpdateTime")
		r.Equal(e[i].ShareAmount, a[i].ShareAmount, "ShareAmount")
	}
}

func (s *bswapServiceTestSuite) TestRequestBSwapQuoteService() {
	data := []byte(`
		{
			"quoteAsset": "USDT",
			"baseAsset": "BUSD",
			"quoteQty": "300000",
			"baseQty": "299975",
			"price": "1.00008334",
			"slippage": "0.00007245",
			"fee": "120"
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	quoteAsset := "USDT"
	baseAsset := "BUSD"
	quoteQty := float64(300000)

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"quoteAsset": quoteAsset,
			"baseAsset":  baseAsset,
			"quoteQty":   quoteQty,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewRequestBSwapQuoteService().QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).
		Do(newContext())
	s.r().NoError(err)
	e := &BSwapQuoteResponse{
		QuoteAsset: "USDT",
		BaseAsset:  "BUSD",
		QuoteQty:   "300000",
		BaseQty:    "299975",
		Price:      "1.00008334",
		Slippage:   "0.00007245",
		Fee:        "120",
	}
	s.assertBSwapQuoteResponseEqual(e, res)
}

func (s *bswapServiceTestSuite) assertBSwapQuoteResponseEqual(e, a *BSwapQuoteResponse) {
	r := s.r()
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.BaseAsset, a.BaseAsset, "BaseAsset")
	r.Equal(e.QuoteQty, a.QuoteQty, "QuoteQty")
	r.Equal(e.BaseQty, a.BaseQty, "BaseQty")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Slippage, a.Slippage, "Slippage")
	r.Equal(e.Fee, a.Fee, "Fee")
}

func (s *bswapServiceTestSuite) TestSwapBSwapService() {
	data := []byte(`
		{
			"swapId": 2314
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	quoteAsset := "USDT"
	baseAsset := "BUSD"
	quoteQty := float64(300000)

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"quoteAsset": quoteAsset,
			"baseAsset":  baseAsset,
			"quoteQty":   quoteQty,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewSwapBSwapService().QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).
		Do(newContext())
	s.r().NoError(err)
	e := &SwapBSwapResponse{SwapId: 2314}
	s.assertSwapBSwapResponseEqual(e, res)
}

func (s *bswapServiceTestSuite) assertSwapBSwapResponseEqual(e, a *SwapBSwapResponse) {
	r := s.r()
	r.Equal(e.SwapId, a.SwapId, "SwapId")
}

func (s *bswapServiceTestSuite) TestGetBSwapSwapHistoryService() {
	data := []byte(`
        [
			{
				"swapId": 2314,
				"swapTime": 1565770342148,
				"status": 0,
				"quoteAsset": "USDT",
				"baseAsset": "BUSD",
				"quoteQty": "300000",
				"baseQty": "299975",
				"price": "1.00008334",
				"fee": "120"
			}
		]
    `)
	s.mockDo(data, nil)
	defer s.assertDo()

	swapID := int64(2314)
	startTime := int64(1565770342146)
	endTime := int64(1565770342149)
	status := BSwapStatusTypePending
	quoteAsset := "USDT"
	baseAsset := "BUSD"
	limit := int64(20)

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"swapId":     swapID,
			"startTime":  startTime,
			"endTime":    endTime,
			"status":     status,
			"quoteAsset": quoteAsset,
			"baseAsset":  baseAsset,
			"limit":      limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetBSwapSwapHistoryService().
		SwapId(swapID).StartTime(startTime).EndTime(endTime).Status(status).
		QuoteAsset(quoteAsset).BaseAsset(baseAsset).Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	e := []*BSwapSwapHistory{
		{
			SwapId:     2314,
			SwapTime:   1565770342148,
			Status:     BSwapStatusTypePending,
			QuoteAsset: "USDT",
			BaseAsset:  "BUSD",
			QuoteQty:   "300000",
			BaseQty:    "299975",
			Price:      "1.00008334",
			Fee:        "120",
		},
	}
	s.assertSwapHistoryEqual(e, res)
}

func (s *bswapServiceTestSuite) assertSwapHistoryEqual(e, a []*BSwapSwapHistory) {
	r := s.r()
	r.Equal(len(e), len(a), "equal length")
	for i := 0; i < len(a); i++ {
		r.Equal(e[i].SwapId, a[i].SwapId, "SwapId")
		r.Equal(e[i].SwapTime, a[i].SwapTime, "SwapTime")
		r.Equal(e[i].Status, a[i].Status, "Status")
		r.Equal(e[i].QuoteAsset, a[i].QuoteAsset, "QuoteAsset")
		r.Equal(e[i].BaseAsset, a[i].BaseAsset, "BaseAsset")
		r.Equal(e[i].QuoteQty, a[i].QuoteQty, "QuoteQty")
		r.Equal(e[i].BaseQty, a[i].BaseQty, "BaseQty")
		r.Equal(e[i].Price, a[i].Price, "Price")
		r.Equal(e[i].Fee, a[i].Fee, "Fee")
	}
}
