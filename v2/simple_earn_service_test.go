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
