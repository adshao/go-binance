package futures

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/adshao/go-binance/v2/futures/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func (s *orderPlaceServiceWsTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
	s.signedKey = "HMAC"
	s.timeOffset = 0

	s.requestID = "e2a85d9f-07a5-4f94-8d5f-789dc3deb098"

	s.symbol = "BTCUSDT"
	s.side = SideTypeSell
	s.orderType = OrderTypeLimit
	s.timeInForce = TimeInForceTypeGTC
	s.quantity = "0.1001"
	s.price = "50000"
	s.newClientOrderID = "testOrder"

	s.ctrl = gomock.NewController(s.T())
	s.client = mock.NewMockwsClient(s.ctrl)

	s.orderPlace = &OrderPlaceWsService{
		c: s.client,
	}

	s.orderPlaceRequest = NewOrderPlaceWsRequest().
		Symbol(s.symbol).
		Side(s.side).
		Type(s.orderType).
		TimeInForce(s.timeInForce).
		Quantity(s.quantity).
		Price(s.price).
		NewClientOrderID(s.newClientOrderID)
}

func (s *orderPlaceServiceWsTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

type orderPlaceServiceWsTestSuite struct {
	suite.Suite
	apiKey     string
	secretKey  string
	signedKey  string
	timeOffset int64

	ctrl   *gomock.Controller
	client *mock.MockwsClient

	requestID        string
	symbol           string
	side             SideType
	orderType        OrderType
	timeInForce      TimeInForceType
	quantity         string
	price            string
	newClientOrderID string

	orderPlace        *OrderPlaceWsService
	orderPlaceRequest *OrderPlaceWsRequest
}

func TestOrderPlaceServiceWsPlace(t *testing.T) {
	suite.Run(t, new(orderPlaceServiceWsTestSuite))
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace() {
	s.expectCalls(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).AnyTimes()

	err := s.orderPlace.Do(s.requestID, s.orderPlaceRequest)
	s.NoError(err)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace_EmptyRequestID() {
	s.expectCalls(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	err := s.orderPlace.Do("", s.orderPlaceRequest)
	s.ErrorIs(err, ErrorRequestIDNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace_EmptyApiKey() {
	s.expectCalls("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderPlace.Do(s.requestID, s.orderPlaceRequest)
	s.ErrorIs(err, ErrorApiKeyIsNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace_EmptySecretKey() {
	s.expectCalls(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderPlace.Do(s.requestID, s.orderPlaceRequest)
	s.ErrorIs(err, ErrorSecretKeyIsNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace_EmptySignKeyType() {
	s.expectCalls(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderPlace.Do(s.requestID, s.orderPlaceRequest)
	s.Error(err)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync() {
	s.expectCalls(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	orderPlaceResponse := CreateOrderWsResponse{
		Id:     s.requestID,
		Status: 200,
		Result: CreateOrderResult{
			CreateOrderResponse{
				Symbol:        s.symbol,
				OrderID:       0,
				ClientOrderID: s.newClientOrderID,
				Price:         s.price,
				TimeInForce:   s.timeInForce,
				Type:          s.orderType,
				Side:          s.side,
			},
		},
	}

	rawResponseData, err := json.Marshal(orderPlaceResponse)
	s.NoError(err)

	s.client.EXPECT().WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(rawResponseData, nil).Times(1)

	req := s.orderPlaceRequest
	response, err := s.orderPlace.SyncDo(s.requestID, req)
	s.Require().NoError(err)
	s.Equal(*req.newClientOrderID, response.Result.ClientOrderID)
	s.Equal(req.symbol, response.Result.Symbol)
	s.Equal(req.orderType, response.Result.Type)
	s.Equal(*req.price, response.Result.Price)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync_EmptyRequestID() {
	s.expectCalls(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	req := s.orderPlaceRequest
	response, err := s.orderPlace.SyncDo("", req)
	s.Nil(response)
	s.ErrorIs(err, ErrorRequestIDNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync_EmptyApiKey() {
	s.expectCalls("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderPlace.SyncDo(s.requestID, s.orderPlaceRequest)
	s.Nil(response)
	s.ErrorIs(err, ErrorApiKeyIsNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync_EmptySecretKey() {
	s.expectCalls(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderPlace.SyncDo(s.requestID, s.orderPlaceRequest)
	s.Nil(response)
	s.ErrorIs(err, ErrorSecretKeyIsNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync_EmptySignKeyType() {
	s.expectCalls(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderPlace.SyncDo(s.requestID, s.orderPlaceRequest)
	s.Nil(response)
	s.Error(err)
}

func (s *orderPlaceServiceWsTestSuite) expectCalls(apiKey, secretKey, signKeyType string, timeOffset int64) {
	s.client.EXPECT().GetApiKey().Return(apiKey).AnyTimes()
	s.client.EXPECT().GetSecretKey().Return(secretKey).AnyTimes()
	s.client.EXPECT().GetKeyType().Return(signKeyType).AnyTimes()
	s.client.EXPECT().GetTimeOffset().Return(timeOffset).AnyTimes()
}
