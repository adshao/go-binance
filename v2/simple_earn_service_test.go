package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type simpleEarnServiceTestSuite struct {
	baseTestSuite
}

func TestSimpleEarnService(t *testing.T) {
	suite.Run(t, new(simpleEarnServiceTestSuite))
}

func (s *simpleEarnServiceTestSuite) TestGetAccount() {
	data := []byte(`{
  "totalAmountInBTC": "0.01067982",
  "totalAmountInUSDT": "77.13289230",
  "totalFlexibleAmountInBTC": "0.00000000",
  "totalFlexibleAmountInUSDT": "0.00000000",
  "totalLockedInBTC": "0.01067982",
  "totalLockedInUSDT": "77.13289230"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	account, err := s.client.NewGetSimpleEarnAccountService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal("0.01067982", account.TotalAmountInBTC)
	s.r().Equal("77.13289230", account.TotalAmountInUSDT)
	s.r().Equal("0.00000000", account.TotalFlexibleAmountInBTC)
	s.r().Equal("0.00000000", account.TotalFlexibleAmountInUSDT)
	s.r().Equal("0.01067982", account.TotalLockedInBTC)
	s.r().Equal("77.13289230", account.TotalLockedInUSDT)
}

func (s *simpleEarnServiceTestSuite) TestListFlexibleProduct() {
	data := []byte(`{
  "rows": [
    {
      "asset": "BTC",
      "latestAnnualPercentageRate": "0.05000000",  
      "tierAnnualPercentageRate": {               
        "0-5BTC": "0.05",
        "5-10BTC": "0.03"
      },
      "airDropPercentageRate": "0.05000000",      
      "canPurchase": true,
      "canRedeem": true,
      "isSoldOut": true,
      "hot": true,                                
      "minPurchaseAmount": "0.01000000",
      "productId": "BTC001",
      "subscriptionStartTime": 1646182276000,
      "status": "PURCHASING"                      
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": "BTC",
		})
		s.assertRequestEqual(e, r)
	})

	product, err := s.client.NewListSimpleEarnFlexibleProductService().Asset("BTC").Do(newContext())
	s.r().NoError(err)
	s.r().Equal("BTC001", product.Rows[0].ProductId)
	s.r().Equal("BTC", product.Rows[0].Asset)
	s.r().Equal("0.05000000", product.Rows[0].LatestAnnualPercentageRate)
	s.r().Equal(map[string]string{"0-5BTC": "0.05", "5-10BTC": "0.03"}, product.Rows[0].TierAnnualPercentageRate)
	s.r().Equal("0.05000000", product.Rows[0].AirDropPercentageRate)
	s.r().Equal(true, product.Rows[0].CanPurchase)
	s.r().Equal(true, product.Rows[0].CanRedeem)
	s.r().Equal(true, product.Rows[0].IsSoldOut)
	s.r().Equal(true, product.Rows[0].Hot)
	s.r().Equal("0.01000000", product.Rows[0].MinPurchaseAmount)
	s.r().Equal("1646182276000", product.Rows[0].SubscriptionStartTime)
	s.r().Equal("PURCHASING", product.Rows[0].Status)
	s.r().Equal(1, product.Total)
}

func (s *simpleEarnServiceTestSuite) TestListLockedProduct() {
	data := []byte(`{
  "rows": [
    {
      "projectId": "Axs*90",
      "detail": {
        "asset": "AXS",                
        "rewardAsset": "AXS",          
        "duration": 90,                
        "renewable": true,             
        "isSoldOut": true,
        "apr": "1.2069",
        "status": "CREATED",           
        "subscriptionStartTime": 1646182276000,
        "extraRewardAsset": "BNB",
        "extraRewardAPR": "0.23"
      },
      "quota": {
        "totalPersonalQuota": "2",     
        "minimum": "0.001"              
      }
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

	product, err := s.client.NewListSimpleEarnLockedProductService().Asset("AXS").Do(newContext())
	s.r().NoError(err)
	s.r().Equal("Axs*90", product.Rows[0].ProjectId)
	s.r().Equal("AXS", product.Rows[0].Detail.Asset)
	s.r().Equal("AXS", product.Rows[0].Detail.RewardAsset)
	s.r().Equal(90, product.Rows[0].Detail.Duration)
	s.r().Equal(true, product.Rows[0].Detail.Renewable)
	s.r().Equal(true, product.Rows[0].Detail.IsSoldOut)
	s.r().Equal("1.2069", product.Rows[0].Detail.Apr)
	s.r().Equal("CREATED", product.Rows[0].Detail.Status)
	s.r().EqualValues(1646182276000, product.Rows[0].Detail.SubscriptionStartTime)
	s.r().Equal("BNB", product.Rows[0].Detail.ExtraRewardAsset)
	s.r().Equal("0.23", product.Rows[0].Detail.ExtraRewardAPR)
	s.r().Equal("2", product.Rows[0].Quota.TotalPersonalQuota)
	s.r().Equal("0.001", product.Rows[0].Quota.Minimum)
	s.r().Equal(1, product.Total)
}

func (s *simpleEarnServiceTestSuite) TestGetSimpleEarnFlexiblePositionService() {
	data := []byte(`{
  "rows": [
    {
      "totalAmount": "75.46000000",
      "tierAnnualPercentageRate": {
        "0-5BTC": "0.05",
        "5-10BTC": "0.03"
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
		e := newSignedRequest().setParams(params{
			"asset": "USDT",
		})
		s.assertRequestEqual(e, r)
	})

	position, err := s.client.NewGetSimpleEarnFlexiblePositionService().Asset("USDT").Do(newContext())
	s.r().NoError(err)
	s.r().Equal("75.46000000", position.Rows[0].TotalAmount)
	s.r().Equal(map[string]string{"0-5BTC": "0.05", "5-10BTC": "0.03"}, position.Rows[0].TierAnnualPercentageRate)
	s.r().Equal("0.02599895", position.Rows[0].LatestAnnualPercentageRate)
	s.r().Equal("0.02599895", position.Rows[0].YesterdayAirdropPercentageRate)
	s.r().Equal("USDT", position.Rows[0].Asset)
	s.r().Equal("BETH", position.Rows[0].AirDropAsset)
	s.r().Equal(true, position.Rows[0].CanRedeem)
	s.r().Equal("232.23123213", position.Rows[0].CollateralAmount)
	s.r().Equal("USDT001", position.Rows[0].ProductId)
	s.r().Equal("0.10293829", position.Rows[0].YesterdayRealTimeRewards)
	s.r().Equal("0.22759183", position.Rows[0].CumulativeBonusRewards)
	s.r().Equal("0.22759183", position.Rows[0].CumulativeRealTimeRewards)
	s.r().Equal("0.45459183", position.Rows[0].CumulativeTotalRewards)
	s.r().Equal(true, position.Rows[0].AutoSubscribe)
	s.r().Equal(1, position.Total)
}

func (s *simpleEarnServiceTestSuite) TestGetSimpleEarnLockedPositionService() {
	data := []byte(`{
  "rows": [
    {
      "positionId": 123123,
      "parentPositionId": 123122,
      "projectId": "Axs*90",
      "asset": "AXS",
      "amount": "122.09202928",
      "purchaseTime": 1646182276000,
      "duration": "60",
      "accrualDays": "4",
      "rewardAsset": "AXS",
      "APY": "0.2032",
      "rewardAmt": "5.17181528",
      "extraRewardAsset": "BNB",
      "extraRewardAPR": "0.0203",
      "estExtraRewardAmt": "5.17181528",
      "nextPay": "1.29295383",
      "nextPayDate": 1646697600000,
      "payPeriod": "1",
      "redeemAmountEarly": "2802.24068892",
      "rewardsEndDate": 1651449600000,
      "deliverDate": 1651536000000,
      "redeemPeriod": "1",
      "redeemingAmt": "232.2323",
      "redeemTo": "FLEXIBLE",
      "partialAmtDeliverDate": 1651536000000,
      "canRedeemEarly": true,
      "canFastRedemption": true,
      "autoSubscribe": true,
      "type": "AUTO",
      "status": "HOLDING",
      "canReStake": true
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":      "AXS",
			"positionId": 123123,
			"projectId":  "Axs*90",
			"current":    1,
			"size":       10,
		})
		s.assertRequestEqual(e, r)
	})

	position, err := s.client.NewGetSimpleEarnLockedPositionService().
		Asset("AXS").
		PositionId(123123).
		ProjectId("Axs*90").
		Current(1).
		Size(10).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(123123, position.Rows[0].PositionId)
	s.r().Equal(123122, position.Rows[0].ParentPositionId)
	s.r().Equal("Axs*90", position.Rows[0].ProjectId)
	s.r().Equal("AXS", position.Rows[0].Asset)
	s.r().Equal("122.09202928", position.Rows[0].Amount)
	s.r().EqualValues(1646182276000, position.Rows[0].PurchaseTime)
	s.r().Equal("60", position.Rows[0].Duration)
	s.r().Equal("4", position.Rows[0].AccrualDays)
	s.r().Equal("AXS", position.Rows[0].RewardAsset)
	s.r().Equal("0.2032", position.Rows[0].APY)
	s.r().Equal("5.17181528", position.Rows[0].RewardAmt)
	s.r().Equal("BNB", position.Rows[0].ExtraRewardAsset)
	s.r().Equal("0.0203", position.Rows[0].ExtraRewardAPR)
	s.r().Equal("5.17181528", position.Rows[0].EstExtraRewardAmt)
	s.r().Equal("1.29295383", position.Rows[0].NextPay)
	s.r().EqualValues(1646697600000, position.Rows[0].NextPayDate)
	s.r().Equal("1", position.Rows[0].PayPeriod)
	s.r().Equal("2802.24068892", position.Rows[0].RedeemAmountEarly)
	s.r().EqualValues(1651449600000, position.Rows[0].RewardsEndDate)
	s.r().EqualValues(1651536000000, position.Rows[0].DeliverDate)
	s.r().Equal("1", position.Rows[0].RedeemPeriod)
	s.r().Equal("232.2323", position.Rows[0].RedeemingAmt)
	s.r().Equal("FLEXIBLE", position.Rows[0].RedeemTo)
	s.r().EqualValues(1651536000000, position.Rows[0].PartialAmtDeliverDate)
	s.r().Equal(true, position.Rows[0].CanRedeemEarly)
	s.r().Equal(true, position.Rows[0].CanFastRedemption)
	s.r().Equal(true, position.Rows[0].AutoSubscribe)
	s.r().Equal("AUTO", position.Rows[0].Type)
	s.r().Equal("HOLDING", position.Rows[0].Status)
	s.r().Equal(true, position.Rows[0].CanReStake)
	s.r().Equal(1, position.Total)
}
