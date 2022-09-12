package binance

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type subAccountServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountService(t *testing.T) {
	suite.Run(t, new(subAccountServiceTestSuite))
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
