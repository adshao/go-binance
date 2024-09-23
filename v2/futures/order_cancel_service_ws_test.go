package futures

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/adshao/go-binance/v2/futures/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func (s *orderCancelServiceWsTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
	s.signedKey = "HMAC"
	s.timeOffset = 0

	s.requestID = "e2a85d9f-07a5-4f94-8d5f-789dc3deb098"

	s.ctrl = gomock.NewController(s.T())
	s.client = mock.NewMockwsClient(s.ctrl)

	s.orderCancel = &OrderCancelWsService{
		c: s.client,
	}

	s.orderCancelRequest = NewOrderCancelRequest().OrigClientOrderID(s.requestID)
}

func (s *orderCancelServiceWsTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

type orderCancelServiceWsTestSuite struct {
	suite.Suite
	apiKey     string
	secretKey  string
	signedKey  string
	timeOffset int64

	ctrl   *gomock.Controller
	client *mock.MockwsClient

	requestID string

	orderCancel        *OrderCancelWsService
	orderCancelRequest *OrderCancelRequest
}

func TestOrderCancelServiceWs(t *testing.T) {
	suite.Run(t, new(orderCancelServiceWsTestSuite))
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancel() {
	s.expectCalls(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).AnyTimes()

	err := s.orderCancel.Do(s.requestID, s.orderCancelRequest)
	s.NoError(err)
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancel_EmptyRequestID() {
	s.expectCalls(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	err := s.orderCancel.Do("", s.orderCancelRequest)
	s.ErrorIs(err, ErrorRequestIDNotSet)
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancel_EmptyApiKey() {
	s.expectCalls("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderCancel.Do(s.requestID, s.orderCancelRequest)
	s.ErrorIs(err, ErrorApiKeyIsNotSet)
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancel_EmptySecretKey() {
	s.expectCalls(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderCancel.Do(s.requestID, s.orderCancelRequest)
	s.ErrorIs(err, ErrorSecretKeyIsNotSet)
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancel_EmptySignKeyType() {
	s.expectCalls(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderCancel.Do(s.requestID, s.orderCancelRequest)
	s.Error(err)
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancelSync() {
	s.expectCalls(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	orderCancelResponse := OrderCancelWsResponse{
		Id:     s.requestID,
		Status: 200,
		Result: CancelOrderResult{
			CancelOrderResponse{
				ClientOrderID: s.requestID,
			},
		},
	}

	rawResponseData, err := json.Marshal(orderCancelResponse)
	s.NoError(err)

	s.client.EXPECT().WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(rawResponseData, nil).Times(1)

	req := s.orderCancelRequest
	response, err := s.orderCancel.SyncDo(s.requestID, req)
	s.Require().NoError(err)
	s.Equal(s.requestID, response.Result.ClientOrderID)
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancelSync_EmptyRequestID() {
	s.expectCalls(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	req := s.orderCancelRequest
	response, err := s.orderCancel.SyncDo("", req)
	s.Nil(response)
	s.ErrorIs(err, ErrorRequestIDNotSet)
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancelSync_EmptyApiKey() {
	s.expectCalls("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderCancel.SyncDo(s.requestID, s.orderCancelRequest)
	s.Nil(response)
	s.ErrorIs(err, ErrorApiKeyIsNotSet)
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancelSync_EmptySecretKey() {
	s.expectCalls(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderCancel.SyncDo(s.requestID, s.orderCancelRequest)
	s.Nil(response)
	s.ErrorIs(err, ErrorSecretKeyIsNotSet)
}

func (s *orderCancelServiceWsTestSuite) TestOrderCancelSync_EmptySignKeyType() {
	s.expectCalls(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderCancel.SyncDo(s.requestID, s.orderCancelRequest)
	s.Nil(response)
	s.Error(err)
}

func (s *orderCancelServiceWsTestSuite) expectCalls(apiKey, secretKey, signKeyType string, timeOffset int64) {
	s.client.EXPECT().GetApiKey().Return(apiKey).AnyTimes()
	s.client.EXPECT().GetSecretKey().Return(secretKey).AnyTimes()
	s.client.EXPECT().GetKeyType().Return(signKeyType).AnyTimes()
	s.client.EXPECT().GetTimeOffset().Return(timeOffset).AnyTimes()
}
