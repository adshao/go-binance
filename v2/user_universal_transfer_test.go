package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type userUniversalTransferTestSuite struct {
	baseTestSuite
}

func TestUserUniversalTransferService(t *testing.T) {
	suite.Run(t, new(userUniversalTransferTestSuite))
}

func (s *userUniversalTransferTestSuite) TestUserUniversalTransfer() {
	data := []byte(`
	{
		"tranId":13526853623
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	types := "MAIN_C2C"
	asset := "USDT"
	amount := 0.1
	fromSymbol := "USDT"
	toSymbol := "USDT"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"type":       types,
			"asset":      asset,
			"amount":     amount,
			"fromSymbol": fromSymbol,
			"toSymbol":   toSymbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUserUniversalTransferService().
		Type(types).
		Asset(asset).
		Amount(amount).
		FromSymbol(fromSymbol).
		ToSymbol(toSymbol).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(int64(13526853623), res.ID)
}
