package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type accountServiceTestSuite struct {
	baseTestSuite
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountServiceTestSuite))
}

func (s *accountServiceTestSuite) TestGetAccount() {
	data := []byte(`{
        "makerCommission": 15,
        "takerCommission": 15,
        "buyerCommission": 0,
        "sellerCommission": 0,
        "canTrade": true,
        "canWithdraw": true,
        "canDeposit": true,
        "balances": [
            {
                "asset": "BTC",
                "free": "4723846.89208129",
                "locked": "0.00000000"
            },
            {
                "asset": "LTC",
                "free": "4763368.68006011",
                "locked": "0.00000000"
            }
        ]
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAccountService().Do(newContext())
	s.r().NoError(err)
	e := &Account{
		MakerCommission:  15,
		TakerCommission:  15,
		BuyerCommission:  0,
		SellerCommission: 0,
		CanTrade:         true,
		CanWithdraw:      true,
		CanDeposit:       true,
		Balances: []Balance{
			{
				Asset:  "BTC",
				Free:   "4723846.89208129",
				Locked: "0.00000000",
			},
			{
				Asset:  "LTC",
				Free:   "4763368.68006011",
				Locked: "0.00000000",
			},
		},
	}
	s.assertAccountEqual(e, res)
}

func (s *accountServiceTestSuite) assertAccountEqual(e, a *Account) {
	r := s.r()
	r.Equal(e.MakerCommission, a.MakerCommission, "MakerCommission")
	r.Equal(e.TakerCommission, a.TakerCommission, "TakerCommission")
	r.Equal(e.BuyerCommission, a.BuyerCommission, "BuyerCommission")
	r.Equal(e.SellerCommission, a.SellerCommission, "SellerCommission")
	r.Equal(e.CanTrade, a.CanTrade, "CanTrade")
	r.Equal(e.CanWithdraw, a.CanWithdraw, "CanWithdraw")
	r.Equal(e.CanDeposit, a.CanDeposit, "CanDeposit")
	r.Len(a.Balances, len(e.Balances))
	for i := 0; i < len(a.Balances); i++ {
		r.Equal(e.Balances[i].Asset, a.Balances[i].Asset, "Asset")
		r.Equal(e.Balances[i].Free, a.Balances[i].Free, "Free")
		r.Equal(e.Balances[i].Locked, a.Balances[i].Locked, "Locked")
	}
}
