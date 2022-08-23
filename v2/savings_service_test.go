package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type savingsServiceTestSuite struct {
	baseTestSuite
}

func TestSavingsService(t *testing.T) {
	suite.Run(t, new(savingsServiceTestSuite))
}

func (s *savingsServiceTestSuite) TestListSavingsFlexibleProducts() {
	data := []byte(`[
    {
        "asset": "BTC",
        "avgAnnualInterestRate": "0.00250025",
        "canPurchase": true,
        "canRedeem": true,
        "dailyInterestPerThousand": "0.00685000",
        "featured": true,
        "minPurchaseAmount": "0.01000000",
        "productId": "BTC001",
        "purchasedAmount": "16.32467016",
        "status": "PURCHASING",
        "upLimit": "200.00000000",
        "upLimitPerUser": "5.00000000"
    },
    {
        "asset": "BUSD",
        "avgAnnualInterestRate": "0.01228590",
        "canPurchase": true,
        "canRedeem": true,
        "dailyInterestPerThousand": "0.03836000",
        "featured": true,
        "minPurchaseAmount": "0.10000000",
        "productId": "BUSD001",
        "purchasedAmount": "10.38932339",
        "status": "PURCHASING",
        "upLimit": "100000.00000000",
        "upLimitPerUser": "50000.00000000"
    }
]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"status":   "ALL",
			"featured": "ALL",
			"current":  1,
			"size":     50,
		})
		s.assertRequestEqual(e, r)
	})

	flexibleProductList, err := s.client.NewListSavingsFlexibleProductsService().
		Status("ALL").
		Featured("ALL").
		Current(1).
		Size(50).
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(flexibleProductList, 2)
	s.assertSavingsFlexibleProductEqual(&SavingsFlexibleProduct{
		Asset:                    "BTC",
		AvgAnnualInterestRate:    "0.00250025",
		CanPurchase:              true,
		CanRedeem:                true,
		DailyInterestPerThousand: "0.00685000",
		Featured:                 true,
		MinPurchaseAmount:        "0.01000000",
		ProductId:                "BTC001",
		PurchasedAmount:          "16.32467016",
		Status:                   "PURCHASING",
		UpLimit:                  "200.00000000",
		UpLimitPerUser:           "5.00000000",
	}, flexibleProductList[0])
	s.assertSavingsFlexibleProductEqual(&SavingsFlexibleProduct{
		Asset:                    "BUSD",
		AvgAnnualInterestRate:    "0.01228590",
		CanPurchase:              true,
		CanRedeem:                true,
		DailyInterestPerThousand: "0.03836000",
		Featured:                 true,
		MinPurchaseAmount:        "0.10000000",
		ProductId:                "BUSD001",
		PurchasedAmount:          "10.38932339",
		Status:                   "PURCHASING",
		UpLimit:                  "100000.00000000",
		UpLimitPerUser:           "50000.00000000",
	}, flexibleProductList[1])
}

func (s *savingsServiceTestSuite) assertSavingsFlexibleProductEqual(e, a *SavingsFlexibleProduct) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.AvgAnnualInterestRate, a.AvgAnnualInterestRate, "AvgAnnualInterestRate")
	r.Equal(e.CanPurchase, a.CanPurchase, "CanPurchase")
	r.Equal(e.CanRedeem, a.CanRedeem, "CanRedeem")
	r.Equal(e.DailyInterestPerThousand, a.DailyInterestPerThousand, "DailyInterestPerThousand")
	r.Equal(e.Featured, a.Featured, "Featured")
	r.Equal(e.MinPurchaseAmount, a.MinPurchaseAmount, "MinPurchaseAmount")
	r.Equal(e.ProductId, a.ProductId, "ProductId")
	r.Equal(e.PurchasedAmount, a.PurchasedAmount, "PurchasedAmount")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.UpLimit, a.UpLimit, "UpLimit")
	r.Equal(e.UpLimitPerUser, a.UpLimitPerUser, "UpLimitPerUser")
}

func (s *savingsServiceTestSuite) TestPurchaseSavingsFlexibleProduct() {
	data := []byte(`{ "purchaseId": 40607 }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"productId": "BTC001",
			"amount":    0.52,
		})
		s.assertRequestEqual(e, r)
	})

	purchaseId, err := s.client.NewPurchaseSavingsFlexibleProductService().
		ProductId("BTC001").
		Amount(0.52).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(purchaseId, uint64(40607), "Purchase Id")
}

func (s *savingsServiceTestSuite) TestReedemSavingsFlexibleProduct() {
	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"productId": "BTC001",
			"amount":    0.52,
			"type":      "FAST",
		})
		s.assertRequestEqual(e, r)
	})

	err := s.client.NewRedeemSavingsFlexibleProductService().
		ProductId("BTC001").
		Amount(0.52).
		Type("FAST").
		Do(newContext())

	r := s.r()
	r.NoError(err)
}

func (s *savingsServiceTestSuite) TestListSavingsFixedAndActivityProducts() {
	data := []byte(`[
    {
        "asset": "USDT",
        "displayPriority": 1,
        "duration": 90,
        "interestPerLot": "1.35810000",
        "interestRate": "0.05510000",
        "lotSize": "100.00000000",
        "lotsLowLimit": 1,
        "lotsPurchased": 74155,
        "lotsUpLimit": 80000,
        "maxLotsPerUser": 2000,
        "needKyc": false,
        "projectId": "CUSDT90DAYSS001",
        "projectName": "USDT",
        "status": "PURCHASING",
        "type": "CUSTOMIZED_FIXED",
        "withAreaLimitation": false
    }
]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":     "USDT",
			"type":      "ACTIVITY",
			"status":    "ALL",
			"isSortAsc": false,
			"sortBy":    "INTEREST_RATE",
			"current":   5,
			"size":      15,
		})
		s.assertRequestEqual(e, r)
	})

	flexibleProductList, err := s.client.NewListSavingsFixedAndActivityProductsService().
		Asset("USDT").
		Type("ACTIVITY").
		Status("ALL").
		IsSortAsc(false).
		SortBy("INTEREST_RATE").
		Current(5).
		Size(15).
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(flexibleProductList, 1)
	s.assertSavingsFixedProductEqual(&SavingsFixedProduct{
		Asset:              "USDT",
		DisplayPriority:    1,
		Duration:           90,
		InterestPerLot:     "1.35810000",
		InterestRate:       "0.05510000",
		LotSize:            "100.00000000",
		LotsLowLimit:       1,
		LotsPurchased:      74155,
		LotsUpLimit:        80000,
		MaxLotsPerUser:     2000,
		NeedKyc:            false,
		ProjectId:          "CUSDT90DAYSS001",
		ProjectName:        "USDT",
		Status:             "PURCHASING",
		Type:               "CUSTOMIZED_FIXED",
		WithAreaLimitation: false,
	}, flexibleProductList[0])
}

func (s *savingsServiceTestSuite) assertSavingsFixedProductEqual(e, a *SavingsFixedProduct) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.DisplayPriority, a.DisplayPriority, "DisplayPriority")
	r.Equal(e.Duration, a.Duration, "Duration")
	r.Equal(e.InterestPerLot, a.InterestPerLot, "InterestPerLot")
	r.Equal(e.InterestRate, a.InterestRate, "InterestRate")
	r.Equal(e.LotSize, a.LotSize, "LotSize")
	r.Equal(e.LotsLowLimit, a.LotsLowLimit, "LotsLowLimit")
	r.Equal(e.LotsPurchased, a.LotsPurchased, "LotsPurchased")
	r.Equal(e.LotsUpLimit, a.LotsUpLimit, "LotsUpLimit")
	r.Equal(e.MaxLotsPerUser, a.MaxLotsPerUser, "MaxLotsPerUser")
	r.Equal(e.NeedKyc, a.NeedKyc, "NeedKyc")
	r.Equal(e.ProjectId, a.ProjectId, "ProjectId")
	r.Equal(e.ProjectName, a.ProjectName, "ProjectName")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.WithAreaLimitation, a.WithAreaLimitation, "WithAreaLimitation")
}

func (s *savingsServiceTestSuite) TestSavingFlexibleProductPositionsService() {
	data := []byte(`[
		{
			"asset": "BUSD",
			"productId": "BUSD001",
			"productName": "BUSD",
			"avgAnnualInterestRate": "0.09998802",
			"annualInterestRate": "0.1",
			"dailyInterestRate": "0.00017529",
			"totalInterest": "12.95020362",
			"totalAmount": "1234.56789",
			"todayPurchasedAmount": "0",
			"redeemingAmount": "0",
			"freeAmount": "1234.56789",
			"freezeAmount": "0",
			"lockedAmount": "0",
			"canRedeem": true
	}
]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{})
		s.assertRequestEqual(e, r)
	})

	flexibleProductList, err := s.client.NewSavingFlexibleProductPositionsService().
		Asset("").
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(flexibleProductList, 1)
	s.assertSavingFlexibleProductPosition(&SavingFlexibleProductPosition{
		Asset:                 "BUSD",
		ProductId:             "BUSD001",
		ProductName:           "BUSD",
		AvgAnnualInterestRate: "0.09998802",
		AnnualInterestRate:    "0.1",
		DailyInterestRate:     "0.00017529",
		TotalInterest:         "12.95020362",
		TotalAmount:           "1234.56789",
		TotalPurchasedAmount:  "0",
		RedeemingAmount:       "0",
		FreeAmount:            "1234.56789",
		FreezeAmount:          "0",
		LockedAmount:          "0",
		CanRedeem:             true,
	}, flexibleProductList[0])
}

func (s *savingsServiceTestSuite) assertSavingFlexibleProductPosition(e, a *SavingFlexibleProductPosition) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.ProductId, a.ProductId, "ProductId")
	r.Equal(e.ProductName, a.ProductName, "ProductName")
	r.Equal(e.AvgAnnualInterestRate, a.AvgAnnualInterestRate, "AvgAnnualInterestRate")
	r.Equal(e.AnnualInterestRate, a.AnnualInterestRate, "AnnualInterestRate")
	r.Equal(e.DailyInterestRate, a.DailyInterestRate, "DailyInterestRate")
	r.Equal(e.TotalInterest, a.TotalInterest, "TotalInterest")
	r.Equal(e.TotalAmount, a.TotalAmount, "TotalAmount")
	r.Equal(e.TotalPurchasedAmount, a.TotalPurchasedAmount, "TotalPurchasedAmount")
	r.Equal(e.RedeemingAmount, a.RedeemingAmount, "RedeemingAmount")
	r.Equal(e.FreeAmount, a.FreeAmount, "FreeAmount")
	r.Equal(e.FreezeAmount, a.FreezeAmount, "FreezeAmount")
	r.Equal(e.LockedAmount, a.LockedAmount, "LockedAmount")
	r.Equal(e.CanRedeem, a.CanRedeem, "CanRedeem")
}

func (s *savingsServiceTestSuite) TestSavingFixedProjectPositionsService() {
	data := []byte(`[
		{
			"asset": "USDT",
			"canTransfer": true,
			"createTimestamp": 1587010770000,
			"duration": 14,
			"startTime": 1587081600000,
			"endTime": 1588291200000,
			"purchaseTime": 1587010771000,
			"interest": "0.19950000",
			"interestRate": "0.05201250",
			"lot": 1,
			"positionId": 51724,
			"principal": "100.00000000",
			"projectId": "CUSDT14DAYSS001",
			"projectName": "USDT",
			"redeemDate": "2020-05-01",
			"status": "HOLDING",
			"type": "CUSTOMIZED_FIXED"
	}
]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"status": "HOLDING",
		})
		s.assertRequestEqual(e, r)
	})

	positionsList, err := s.client.NewSavingFixedProjectPositionsService().
		Asset("").
		Status("HOLDING").
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(positionsList, 1)
	s.assertSavingFixedProjectPositionsService(&SavingFixedProjectPosition{
		Asset:           "USDT",
		CanTransfer:     true,
		CreateTimestamp: 1587010770000,
		Duration:        14,
		StartTime:       1587081600000,
		EndTime:         1588291200000,
		PurchaseTime:    1587010771000,
		Interest:        "0.19950000",
		InterestRate:    "0.05201250",
		Lot:             1,
		PositionId:      51724,
		Principal:       "100.00000000",
		ProjectId:       "CUSDT14DAYSS001",
		ProjectName:     "USDT",
		RedeemDate:      "2020-05-01",
		Status:          "HOLDING",
		ProjectType:     "CUSTOMIZED_FIXED",
	}, positionsList[0])
}

func (s *savingsServiceTestSuite) assertSavingFixedProjectPositionsService(e, a *SavingFixedProjectPosition) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.CanTransfer, a.CanTransfer, "CanTransfer")
	r.Equal(e.CreateTimestamp, a.CreateTimestamp, "CreateTimestamp")
	r.Equal(e.Duration, a.Duration, "Duration")
	r.Equal(e.StartTime, a.StartTime, "StartTime")
	r.Equal(e.EndTime, a.EndTime, "EndTime")
	r.Equal(e.PurchaseTime, a.PurchaseTime, "PurchaseTime")
	r.Equal(e.Interest, a.Interest, "Interest")
	r.Equal(e.InterestRate, a.InterestRate, "InterestRate")
	r.Equal(e.Lot, a.Lot, "Lot")
	r.Equal(e.PositionId, a.PositionId, "PositionId")
	r.Equal(e.Principal, a.Principal, "Principal")
	r.Equal(e.ProjectId, a.ProjectId, "ProjectId")
	r.Equal(e.ProjectName, a.ProjectName, "ProjectName")
	r.Equal(e.RedeemDate, a.RedeemDate, "RedeemDate")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.ProjectType, a.ProjectType, "ProjectType")
}
