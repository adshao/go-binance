package futures

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

func (s *accountServiceTestSuite) TestGetBalance() {
	data := []byte(`[
		{
			"accountAlias": "SgsR",
			"asset": "USDT",
			"balance": "122607.35137903",
			"crossWalletBalance": "23.72469206",
			"crossUnPnl": "0.00000000",
			"availableBalance": "23.72469206",
			"maxWithdrawAmount": "23.72469206"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetBalanceService().Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 1)
	e := &Balance{
		AccountAlias:       "SgsR",
		Asset:              "USDT",
		Balance:            "122607.35137903",
		CrossWalletBalance: "23.72469206",
		CrossUnPnl:         "0.00000000",
		AvailableBalance:   "23.72469206",
		MaxWithdrawAmount:  "23.72469206",
	}
	s.assertBalanceEqual(e, res[0])
}

func (s *accountServiceTestSuite) assertBalanceEqual(e, a *Balance) {
	r := s.r()
	r.Equal(e.AccountAlias, a.AccountAlias, "AccountAlias")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Balance, a.Balance, "Balance")
	r.Equal(e.CrossWalletBalance, a.CrossWalletBalance, "CrossWalletBalance")
	r.Equal(e.CrossUnPnl, a.CrossUnPnl, "CrossUnPnl")
	r.Equal(e.AvailableBalance, a.AvailableBalance, "AvailableBalance")
	r.Equal(e.MaxWithdrawAmount, a.MaxWithdrawAmount, "MaxWithdrawAmount")
}

func (s *accountServiceTestSuite) TestGetAccount() {
	data := []byte(`{
		"assets": [
			{
				"asset": "USDT",
				"initialMargin": "0.33683000",
				"maintMargin": "0.02695000",
				"marginBalance": "8.74947592",
				"maxWithdrawAmount": "8.41264592",
				"openOrderInitialMargin": "0.00000000",
				"positionInitialMargin": "0.33683000",
				"unrealizedProfit": "-0.44537584",
				"walletBalance": "9.19485176"
			}
		 ],
		 "canDeposit": true,
		 "canTrade": true,
		 "canWithdraw": true,
		 "feeTier": 2,
		 "maxWithdrawAmount": "8.41264592",
		 "positions": [
			 {
				"isolated": false, 
				"leverage": "20",
				"initialMargin": "0.33683",
				"maintMargin": "0.02695",
				"openOrderInitialMargin": "0.00000",
				"positionInitialMargin": "0.33683",
				"symbol": "BTCUSDT",
				"unrealizedProfit": "-0.44537584",
				"entryPrice": "8950.5",
				"maxNotional": "250000",
				"positionSide": "BOTH",
				"positionAmt": "0.436",
				"notional":"0.1234",
				"isolatedWallet":"0.5678",
				"updateTime":1618646402359
			 }
		 ],
		 "totalInitialMargin": "0.33683000",
		 "totalMaintMargin": "0.02695000",
		 "totalMarginBalance": "8.74947592",
		 "totalOpenOrderInitialMargin": "0.00000000",
		 "totalPositionInitialMargin": "0.33683000",
		 "totalUnrealizedProfit": "-0.44537584",
		 "totalWalletBalance": "9.19485176",
		 "updateTime": 0
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
		Assets: []*AccountAsset{
			{
				Asset:                  "USDT",
				InitialMargin:          "0.33683000",
				MaintMargin:            "0.02695000",
				MarginBalance:          "8.74947592",
				MaxWithdrawAmount:      "8.41264592",
				OpenOrderInitialMargin: "0.00000000",
				PositionInitialMargin:  "0.33683000",
				UnrealizedProfit:       "-0.44537584",
				WalletBalance:          "9.19485176",
			},
		},
		CanTrade:          true,
		CanWithdraw:       true,
		CanDeposit:        true,
		FeeTier:           2,
		MaxWithdrawAmount: "8.41264592",
		Positions: []*AccountPosition{
			{
				Isolated:               false,
				Leverage:               "20",
				InitialMargin:          "0.33683",
				MaintMargin:            "0.02695",
				OpenOrderInitialMargin: "0.00000",
				PositionInitialMargin:  "0.33683",
				Symbol:                 "BTCUSDT",
				UnrealizedProfit:       "-0.44537584",
				EntryPrice:             "8950.5",
				MaxNotional:            "250000",
				PositionSide:           "BOTH",
				PositionAmt:            "0.436",
				Notional:               "0.1234",
				IsolatedWallet:         "0.5678",
				UpdateTime:             1618646402359,
			},
		},
		TotalInitialMargin:          "0.33683000",
		TotalMaintMargin:            "0.02695000",
		TotalMarginBalance:          "8.74947592",
		TotalOpenOrderInitialMargin: "0.00000000",
		TotalPositionInitialMargin:  "0.33683000",
		TotalUnrealizedProfit:       "-0.44537584",
		TotalWalletBalance:          "9.19485176",
		UpdateTime:                  0,
	}
	s.assertAccountEqual(e, res)
}

func (s *accountServiceTestSuite) assertAccountEqual(e, a *Account) {
	r := s.r()
	r.Equal(e.CanDeposit, a.CanDeposit, "CanDeposit")
	r.Equal(e.CanTrade, a.CanTrade, "CanTrade")
	r.Equal(e.CanWithdraw, a.CanWithdraw, "CanWithdraw")
	r.Equal(e.FeeTier, a.FeeTier, "FeeTier")
	r.Equal(e.MaxWithdrawAmount, a.MaxWithdrawAmount, "MaxWithdrawAmount")
	r.Equal(e.TotalInitialMargin, a.TotalInitialMargin, "TotalInitialMargin")
	r.Equal(e.TotalMaintMargin, a.TotalMaintMargin, "TotalMaintMargin")
	r.Equal(e.TotalMarginBalance, a.TotalMarginBalance, "TotalMarginBalance")
	r.Equal(e.TotalOpenOrderInitialMargin, a.TotalOpenOrderInitialMargin, "TotalOpenOrderInitialMargin")
	r.Equal(e.TotalPositionInitialMargin, a.TotalPositionInitialMargin, "TotalPositionInitialMargin")
	r.Equal(e.TotalUnrealizedProfit, a.TotalUnrealizedProfit, "TotalUnrealizedProfit")
	r.Equal(e.TotalWalletBalance, a.TotalWalletBalance, "TotalWalletBalance")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")

	r.Len(a.Assets, len(e.Assets))
	for i := 0; i < len(a.Assets); i++ {
		r.Equal(e.Assets[i].Asset, a.Assets[i].Asset, "Asset")
		r.Equal(e.Assets[i].InitialMargin, a.Assets[i].InitialMargin, "InitialMargin")
		r.Equal(e.Assets[i].MaintMargin, a.Assets[i].MaintMargin, "MaintMargin")
		r.Equal(e.Assets[i].MarginBalance, a.Assets[i].MarginBalance, "MarginBalance")
		r.Equal(e.Assets[i].MaxWithdrawAmount, a.Assets[i].MaxWithdrawAmount, "MaxWithdrawAmount")
		r.Equal(e.Assets[i].OpenOrderInitialMargin, a.Assets[i].OpenOrderInitialMargin, "OpenOrderInitialMargin")
		r.Equal(e.Assets[i].PositionInitialMargin, a.Assets[i].PositionInitialMargin, "PositionInitialMargin")
		r.Equal(e.Assets[i].UnrealizedProfit, a.Assets[i].UnrealizedProfit, "UnrealizedProfit")
		r.Equal(e.Assets[i].WalletBalance, a.Assets[i].WalletBalance, "WalletBalance")
	}

	r.Len(a.Positions, len(e.Positions))
	for i := 0; i < len(a.Positions); i++ {
		r.Equal(e.Positions[i].Isolated, a.Positions[i].Isolated, "Isolated")
		r.Equal(e.Positions[i].Leverage, a.Positions[i].Leverage, "Leverage")
		r.Equal(e.Positions[i].InitialMargin, a.Positions[i].InitialMargin, "InitialMargin")
		r.Equal(e.Positions[i].MaintMargin, a.Positions[i].MaintMargin, "MaintMargin")
		r.Equal(e.Positions[i].OpenOrderInitialMargin, a.Positions[i].OpenOrderInitialMargin, "OpenOrderInitialMargin")
		r.Equal(e.Positions[i].PositionInitialMargin, a.Positions[i].PositionInitialMargin, "PositionInitialMargin")
		r.Equal(e.Positions[i].Symbol, a.Positions[i].Symbol, "Symbol")
		r.Equal(e.Positions[i].UnrealizedProfit, a.Positions[i].UnrealizedProfit, "UnrealizedProfit")
		r.Equal(e.Positions[i].EntryPrice, a.Positions[i].EntryPrice, "EntryPrice")
		r.Equal(e.Positions[i].MaxNotional, a.Positions[i].MaxNotional, "MaxNotional")
		r.Equal(e.Positions[i].PositionSide, a.Positions[i].PositionSide, "PositionSide")
		r.Equal(e.Positions[i].PositionAmt, a.Positions[i].PositionAmt, "PositionAmt")
		r.Equal(e.Positions[i].Notional, a.Positions[i].Notional, "Notional")
		r.Equal(e.Positions[i].IsolatedWallet, a.Positions[i].IsolatedWallet, "IsolatedWallet")
		r.Equal(e.Positions[i].UpdateTime, a.Positions[i].UpdateTime, "UpdateTime")
	}
}
