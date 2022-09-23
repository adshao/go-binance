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
