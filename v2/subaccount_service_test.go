package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type subAccountServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountService(t *testing.T) {
	suite.Run(t, new(subAccountServiceTestSuite))
}

func (s *subAccountServiceTestSuite) TestSubaccountDepositAddressService() {
	data := []byte(`
	{
		"address":"TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV",
		"coin":"USDT",
		"tag":"a_tag",
		"url":"https://tronscan.org/#/address/TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	coin := "a_coin"
	network := "a_network"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"email":   email,
			"coin":    coin,
			"network": network,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewSubaccountDepositAddressService().
		Email(email).
		Coin(coin).
		Network(network).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV", res.Address, "Address")
	r.Equal("USDT", res.Coin, "Coin")
	r.Equal("a_tag", res.Tag, "Tag")
	r.Equal("https://tronscan.org/#/address/TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV", res.URL, "URL")
}

func (s *subAccountServiceTestSuite) TestSubAccountListService() {
	data := []byte(`
	{
		"subAccounts":[
			{
				"email":"testsub@gmail.com",
				"isFreeze":false,
				"createTime":1544433328000,
				"isManagedSubAccount": false,
				"isAssetManagementSubAccount": false
			}
		]
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	isFreeze := false
	page := 1
	limit := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"email":    email,
			"isFreeze": isFreeze,
			"page":     1,
			"limit":    10,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewSubAccountListService().
		Email(email).
		IsFreeze(false).
		Page(page).
		Limit(limit).
		Do(newContext())

	r := s.r()
	r.NoError(err)

	s.assertSubAccountEqual(SubAccount{
		Email:      email,
		CreateTime: 1544433328000,
	}, res.SubAccounts[0])
}

func (s *subAccountServiceTestSuite) assertSubAccountEqual(e, a SubAccount) {
	r := s.r()
	r.Equal(e.Email, a.Email, "Email")
	r.Equal(e.IsFreeze, a.IsFreeze, "IsFreeze")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.IsManagedSubAccount, a.IsManagedSubAccount, "IsManagedSubAccount")
	r.Equal(e.IsAssetManagementSubAccount, a.IsAssetManagementSubAccount, "IsAssetManagementSubAccount")
}

func (s *subAccountServiceTestSuite) TestSubManagedSubAccountDepositService() {
	data := []byte(` { "tranId": 12345678 } `)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	asset := "USDT"
	amount := 1.0
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"toEmail": email,
			"asset":   asset,
			"amount":  amount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewManagedSubAccountDepositService().
		ToEmail(email).
		Asset(asset).
		Amount(amount).
		Do(newContext())

	r := s.r()
	r.NoError(err)

	r.Equal(int64(12345678), res.ID)
}

func (s *subAccountServiceTestSuite) TestSubManagedSubAccountWithdrawalService() {
	data := []byte(` { "tranId": 12345678 } `)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	asset := "USDT"
	amount := 1.0
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"fromEmail": email,
			"asset":     asset,
			"amount":    amount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewManagedSubAccountWithdrawalService().
		FromEmail(email).
		Asset(asset).
		Amount(amount).
		Do(newContext())

	r := s.r()
	r.NoError(err)

	r.Equal(int64(12345678), res.ID)
}

func (s *subAccountServiceTestSuite) TestSubManagedSubAccountAssetsService() {
	data := []byte(`
		[
			{
			"coin": "INJ",                
			"name": "Injective Protocol", 
			"totalBalance": "0",          
			"availableBalance": "0",      
			"inOrder": "0",                
			"btcValue": "0"               
			}
		]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{"email": email})
		s.assertRequestEqual(e, r)
	})

	assets, err := s.client.NewManagedSubAccountAssetsService().Email(email).Do(newContext())

	r := s.r()
	r.NoError(err)

	s.assertAssetsEqual(&ManagedSubAccountAsset{
		Coin:             "INJ",
		Name:             "Injective Protocol",
		TotalBalance:     "0",
		AvailableBalance: "0",
		InOrder:          "0",
		BtcValue:         "0",
	}, assets[0])
}

func (s *subAccountServiceTestSuite) assertAssetsEqual(e, a *ManagedSubAccountAsset) {
	r := s.r()
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.Name, a.Name, "Name")
	r.Equal(e.TotalBalance, a.TotalBalance, "TotalBalance")
	r.Equal(e.AvailableBalance, a.AvailableBalance, "AvailableBalance")
	r.Equal(e.InOrder, a.InOrder, "InOrder")
	r.Equal(e.BtcValue, a.BtcValue, "BtcValue")
}

func (s *subAccountServiceTestSuite) TestSubAccountFuturesService() {
	data := []byte(`
		{
		    "email": "abc@test.com",
		    "asset": "USDT",
		    "assets":[
		        {
		            "asset": "USDT",
		            "initialMargin": "0.00000000",
		            "maintenanceMargin": "0.00000000",
		            "marginBalance": "0.88308000",
		            "maxWithdrawAmount": "0.88308000",
		            "openOrderInitialMargin": "0.00000000",
		            "positionInitialMargin": "0.00000000",
		            "unrealizedProfit": "0.00000000",
		            "walletBalance": "0.88308000"
		         }
		    ],
		    "canDeposit": true,
		    "canTrade": true,
		    "canWithdraw": true,
		    "feeTier": 2,
		    "maxWithdrawAmount": "0.88308000",
		    "totalInitialMargin": "0.00000000",
		    "totalMaintenanceMargin": "0.00000000",
		    "totalMarginBalance": "0.88308000",
		    "totalOpenOrderInitialMargin": "0.00000000",
		    "totalPositionInitialMargin": "0.00000000",
		    "totalUnrealizedProfit": "0.00000000",
		    "totalWalletBalance": "0.88308000",
		    "updateTime": 1576756674610
		 }
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "abc@test.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{"email": email})
		s.assertRequestEqual(e, r)
	})

	account, err := s.client.NewSubAccountFuturesAccountService().Email(email).Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("abc@test.com", account.Email, "Email")
	r.Equal("USDT", account.Asset, "Asset")
	r.Equal(true, account.CanDeposit, "CanDeposit")
	r.Equal(true, account.CanTrade, "CanTrade")
	r.Equal(true, account.CanWithdraw, "CanWithdraw")
	r.Equal(2, account.FeeTier, "FeeTier")
	r.Equal("0.88308000", account.MaxWithdrawAmount, "MaxWithdrawAmount")
	r.Equal("0.00000000", account.TotalInitialMargin, "TotalInitialMargin")
	r.Equal("0.00000000", account.TotalMaintenanceMargin, "TotalMaintenanceMargin")
	r.Equal("0.88308000", account.TotalMarginBalance, "TotalMarginBalance")
	r.Equal("0.00000000", account.TotalOpenOrderInitialMargin, "TotalOpenOrderInitialMargin")
	r.Equal("0.00000000", account.TotalPositionInitialMargin, "TotalPositionInitialMargin")
	r.Equal("0.00000000", account.TotalUnrealizedProfit, "TotalUnrealizedProfit")
	r.Equal("0.88308000", account.TotalWalletBalance, "TotalWalletBalance")
	r.Equal(int64(1576756674610), account.UpdateTime, "UpdateTime")

	s.assertAccountFuturesAssetsEqual(&SubAccountFuturesAccountAsset{
		Asset:                  "USDT",
		InitialMargin:          "0.00000000",
		MaintenanceMargin:      "0.00000000",
		MarginBalance:          "0.88308000",
		MaxWithdrawAmount:      "0.88308000",
		OpenOrderInitialMargin: "0.00000000",
		PositionInitialMargin:  "0.00000000",
		UnrealizedProfit:       "0.00000000",
		WalletBalance:          "0.88308000",
	}, &account.Assets[0])
}

func (s *subAccountServiceTestSuite) assertAccountFuturesAssetsEqual(e, a *SubAccountFuturesAccountAsset) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.InitialMargin, a.InitialMargin, "InitialMargin")
	r.Equal(e.MaintenanceMargin, a.MaintenanceMargin, "MaintenanceMargin")
	r.Equal(e.MarginBalance, a.MarginBalance, "MarginBalance")
	r.Equal(e.MaxWithdrawAmount, a.MaxWithdrawAmount, "MaxWithdrawAmount")
	r.Equal(e.OpenOrderInitialMargin, a.OpenOrderInitialMargin, "OpenOrderInitialMargin")
	r.Equal(e.PositionInitialMargin, a.PositionInitialMargin, "PositionInitialMargin")
	r.Equal(e.UnrealizedProfit, a.UnrealizedProfit, "UnrealizedProfit")
	r.Equal(e.WalletBalance, a.WalletBalance, "WalletBalance")
}

func (s *subAccountServiceTestSuite) TestSubAccountFuturesTransferService() {
	data := []byte(`
		{
		    "tranId": 123456789
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "abc@test.com"
	asset := "USDT"
	amount := 1.0
	tType := 1
	s.assertReq(func(r *request) {
		e := newSignedRequest().
			setParams(params{
				"email":  email,
				"asset":  asset,
				"amount": amount,
				"type":   tType,
			})
		s.assertRequestEqual(e, r)
	})

	response, err := s.client.NewSubAccountFuturesTransferV1Service().Email(email).Asset(asset).Amount(amount).TransferType(tType).Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(int64(123456789), response.TranID, "TranID")

}
