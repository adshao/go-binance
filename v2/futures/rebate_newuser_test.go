package futures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseRebateNewUserTestSuite struct {
	baseTestSuite
}

func TestRebateNewUserTestService(t *testing.T) {
	suite.Run(t, new(baseRebateNewUserServiceTestSuite))
}

type baseRebateNewUserServiceTestSuite struct {
	baseRebateNewUserTestSuite
}

func (s *baseRebateNewUserServiceTestSuite) TestRebateNewUser() {
	data := []byte(`
		{
			"brokerId": "123456",
			"rebateWorking": true, 
			"ifNewUser": true
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	brokerageID := "123456"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"brokerId":   brokerageID,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	check, err := s.client.NewGetRebateNewUserService().BrokerageID(brokerageID).
		Do(newContext(), WithRecvWindow(recvWindow))
	fmt.Println(check)
	r := s.r()
	r.NoError(err)
	e := &RebateNewUser{
		BrokerId:      "123456",
		RebateWorking: true,
		IfNewUser:     true,
	}
	s.assertOrderEqual(e, check)
}

func (s *baseRebateNewUserServiceTestSuite) assertOrderEqual(e, a *RebateNewUser) {
	r := s.r()
	r.Equal(e.BrokerId, a.BrokerId, "BrokerageId")
	r.Equal(e.IfNewUser, a.IfNewUser, "New User check")
	r.Equal(e.RebateWorking, a.RebateWorking, "Rebate Working check")
}
