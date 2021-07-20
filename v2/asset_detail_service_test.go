package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type assetDetailServiceTestSuite struct {
	baseTestSuite
}

func TestAssetDetailService(t *testing.T) {
	suite.Run(t, new(assetDetailServiceTestSuite))
}

func (s *withdrawServiceTestSuite) TestGetAssetDetail() {
	data := []byte(`
	{
		"CTR": {
			"minWithdrawAmount": 70.00000000,
			"depositStatus": false,
			"withdrawFee": 35,
			"withdrawStatus": true,
			"depositTip": "Delisted, Deposit Suspended"
		},
		"SKY": {
			"minWithdrawAmount": 0.02000000,
			"depositStatus": true,
			"withdrawFee": 0.01,
			"withdrawStatus": true
		}
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAssetDetailService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal(res["CTR"].DepositStatus, false, "depositStatus")
	s.r().Equal(res["SKY"].DepositStatus, true, "depositStatus")
}
