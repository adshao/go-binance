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

func (s *savingsServiceTestSuite) TestGetFlexibleProductPositionService() {
	data := []byte(`{
		"rows":[
		  {
			"totalAmount": "75.46000000",
			"tierAnnualPercentageRate": {
			  "0-5BTC": 0.05,
			  "5-10BTC": 0.03
			},
			"latestAnnualPercentageRate": "0.02599895",
			"yesterdayAirdropPercentageRate": "0.02599895",
			"asset": "USDT",
			"airDropAsset": "BETH",
			"canRedeem": true,
			"collateralAmount": "232.23123213",
			"productId": "USDT001",
			"yesterdayRealTimeRewards": "0.10293829",
			"cumulativeBonusRewards": "0.22759183",
			"cumulativeRealTimeRewards": "0.22759183",
			"cumulativeTotalRewards": "0.45459183",
			"autoSubscribe": true
		  }
		],
		"total": 1
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{})
		s.assertRequestEqual(e, r)
	})

	positionsList, err := s.client.NewGetFlexibleProductPosition().
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(positionsList.Rows, 1)
	s.assertFlexibleProductPosition(&FlexibleProductPosition{
		TotalAmount: "75.46000000",
		TierAnnualPercentageRate: map[string]float64{
			"0-5BTC":  0.05,
			"5-10BTC": 0.03,
		},
		LatestAnnualPercentageRate:     "0.02599895",
		YesterdayAirdropPercentageRate: "0.02599895",
		Asset:                          "USDT",
		AirDropAsset:                   "BETH",
		CanRedeem:                      true,
		CollateralAmount:               "232.23123213",
		ProductID:                      "USDT001",
		YesterdayRealTimeRewards:       "0.10293829",
		CumulativeBonusRewards:         "0.22759183",
		CumulativeRealTimeRewards:      "0.22759183",
		CumulativeTotalRewards:         "0.45459183",
		AutoSubscribe:                  true,
	}, &positionsList.Rows[0])
}

func (s *savingsServiceTestSuite) assertFlexibleProductPosition(e, a *FlexibleProductPosition) {
	r := s.r()
	r.Equal(e.TotalAmount, a.TotalAmount, "TotalAmount")
	r.Equal(e.TierAnnualPercentageRate, a.TierAnnualPercentageRate, "TierAnnualPercentageRate")
	r.Equal(e.LatestAnnualPercentageRate, a.LatestAnnualPercentageRate, "LatestAnnualPercentageRate")
	r.Equal(e.YesterdayAirdropPercentageRate, a.YesterdayAirdropPercentageRate, "YesterdayAirdropPercentageRate")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.AirDropAsset, a.AirDropAsset, "AirDropAsset")
	r.Equal(e.CanRedeem, a.CanRedeem, "CanRedeem")
	r.Equal(e.CollateralAmount, a.CollateralAmount, "CollateralAmount")
	r.Equal(e.ProductID, a.ProductID, "ProductID")
	r.Equal(e.YesterdayRealTimeRewards, a.YesterdayRealTimeRewards, "YesterdayRealTimeRewards")
	r.Equal(e.CumulativeBonusRewards, a.CumulativeBonusRewards, "CumulativeBonusRewards")
	r.Equal(e.CumulativeRealTimeRewards, a.CumulativeRealTimeRewards, "CumulativeRealTimeRewards")
	r.Equal(e.CumulativeTotalRewards, a.CumulativeTotalRewards, "CumulativeTotalRewards")
	r.Equal(e.AutoSubscribe, a.AutoSubscribe, "AutoSubscribe")
}

func (s *savingsServiceTestSuite) TestGetLockedProductPositionService() {
	data := []byte(`{
		"rows":[
		  {
			"positionId": "123123",
			"projectId": "Axs*90",
			"asset": "AXS",
			"amount": "122.09202928",
			"purchaseTime": "1646182276000",
			"duration": "60",
			"accrualDays": "4",
			"rewardAsset": "AXS",
			"APY": "0.23",
			"isRenewable": true,
			"isAutoRenew": true,
			"redeemDate": "1732182276000"
		  }
		],
		"total": 1
	  }`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": "AXS",
		})
		s.assertRequestEqual(e, r)
	})

	positionsList, err := s.client.NewGetLockedProductPosition().
		Asset("AXS").
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(positionsList.Rows, 1)
	s.assertLockedProductPosition(&LockedProductPosition{
		PositionID:   "123123",
		ProjectID:    "Axs*90",
		Asset:        "AXS",
		Amount:       "122.09202928",
		PurchaseTime: "1646182276000",
		Duration:     "60",
		AccrualDays:  "4",
		RewardAsset:  "AXS",
		Apy:          "0.23",
		IsRenewable:  true,
		IsAutoRenew:  true,
		RedeemDate:   "1732182276000",
	}, &positionsList.Rows[0])
}

func (s *savingsServiceTestSuite) assertLockedProductPosition(e, a *LockedProductPosition) {
	r := s.r()
	r.Equal(e.PositionID, a.PositionID, "PositionID")
	r.Equal(e.ProjectID, a.ProjectID, "ProjectID")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.PurchaseTime, a.PurchaseTime, "PurchaseTime")
	r.Equal(e.Duration, a.Duration, "Duration")
	r.Equal(e.AccrualDays, a.AccrualDays, "AccrualDays")
	r.Equal(e.RewardAsset, a.RewardAsset, "RewardAsset")
	r.Equal(e.Apy, a.Apy, "Apy")
	r.Equal(e.IsRenewable, a.IsRenewable, "IsRenewable")
	r.Equal(e.IsAutoRenew, a.IsAutoRenew, "IsAutoRenew")
	r.Equal(e.RedeemDate, a.RedeemDate, "RedeemDate")
}

func (s *savingsServiceTestSuite) TestGetFlexibleRewardHistoryService() {
	data := []byte(`{
		"rows": [
		  {
			"asset": "BUSD",
			"rewards": "0.00006408",
			"projectId": "USDT001",
			"type": "BONUS",
			"time": 1577233578000
		  },
		  {
			"asset": "USDT",
			"rewards": "0.00687654",
			"projectId": "USDT001",
			"type": "REALTIME",
			"time": 1577233562000
		  }
		],
		"total": 2
	  }`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{})
		s.assertRequestEqual(e, r)
	})

	rewardHistoryRes, err := s.client.NewGetFlexibleRewardHistory().
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(rewardHistoryRes.Rows, 2)
	s.assertFlexibleRewardHistory(&FlexibleRewardHistory{
		Asset:     "BUSD",
		Rewards:   "0.00006408",
		ProjectID: "USDT001",
		Typ:       "BONUS",
		Time:      1577233578000,
	}, &rewardHistoryRes.Rows[0])
	s.assertFlexibleRewardHistory(&FlexibleRewardHistory{
		Asset:     "USDT",
		Rewards:   "0.00687654",
		ProjectID: "USDT001",
		Typ:       "REALTIME",
		Time:      1577233562000,
	}, &rewardHistoryRes.Rows[1])
}

func (s *savingsServiceTestSuite) assertFlexibleRewardHistory(e, a *FlexibleRewardHistory) {
	r := s.r()
	r.Equal(e.ProjectID, a.ProjectID, "ProjectID")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Rewards, a.Rewards, "Rewards")
	r.Equal(e.Typ, a.Typ, "Type")
	r.Equal(e.Time, a.Time, "Time")
}
