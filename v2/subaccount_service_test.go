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
