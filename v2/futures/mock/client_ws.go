// Code generated by MockGen. DO NOT EDIT.
// Source: client_ws.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockwsClient is a mock of wsClient interface.
type MockwsClient struct {
	ctrl     *gomock.Controller
	recorder *MockwsClientMockRecorder
}

// MockwsClientMockRecorder is the mock recorder for MockwsClient.
type MockwsClientMockRecorder struct {
	mock *MockwsClient
}

// NewMockwsClient creates a new mock instance.
func NewMockwsClient(ctrl *gomock.Controller) *MockwsClient {
	mock := &MockwsClient{ctrl: ctrl}
	mock.recorder = &MockwsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockwsClient) EXPECT() *MockwsClientMockRecorder {
	return m.recorder
}

// GetApiKey mocks base method.
func (m *MockwsClient) GetApiKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApiKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetApiKey indicates an expected call of GetApiKey.
func (mr *MockwsClientMockRecorder) GetApiKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiKey", reflect.TypeOf((*MockwsClient)(nil).GetApiKey))
}

// GetKeyType mocks base method.
func (m *MockwsClient) GetKeyType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetKeyType indicates an expected call of GetKeyType.
func (mr *MockwsClientMockRecorder) GetKeyType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyType", reflect.TypeOf((*MockwsClient)(nil).GetKeyType))
}

// GetReadChannel mocks base method.
func (m *MockwsClient) GetReadChannel() <-chan []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadChannel")
	ret0, _ := ret[0].(<-chan []byte)
	return ret0
}

// GetReadChannel indicates an expected call of GetReadChannel.
func (mr *MockwsClientMockRecorder) GetReadChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadChannel", reflect.TypeOf((*MockwsClient)(nil).GetReadChannel))
}

// GetReadErrorChannel mocks base method.
func (m *MockwsClient) GetReadErrorChannel() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadErrorChannel")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// GetReadErrorChannel indicates an expected call of GetReadErrorChannel.
func (mr *MockwsClientMockRecorder) GetReadErrorChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadErrorChannel", reflect.TypeOf((*MockwsClient)(nil).GetReadErrorChannel))
}

// GetReconnectCount mocks base method.
func (m *MockwsClient) GetReconnectCount() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReconnectCount")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetReconnectCount indicates an expected call of GetReconnectCount.
func (mr *MockwsClientMockRecorder) GetReconnectCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReconnectCount", reflect.TypeOf((*MockwsClient)(nil).GetReconnectCount))
}

// GetSecretKey mocks base method.
func (m *MockwsClient) GetSecretKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSecretKey indicates an expected call of GetSecretKey.
func (mr *MockwsClientMockRecorder) GetSecretKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretKey", reflect.TypeOf((*MockwsClient)(nil).GetSecretKey))
}

// GetTimeOffset mocks base method.
func (m *MockwsClient) GetTimeOffset() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeOffset")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetTimeOffset indicates an expected call of GetTimeOffset.
func (mr *MockwsClientMockRecorder) GetTimeOffset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeOffset", reflect.TypeOf((*MockwsClient)(nil).GetTimeOffset))
}

// Wait mocks base method.
func (m *MockwsClient) Wait(timeout time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Wait", timeout)
}

// Wait indicates an expected call of Wait.
func (mr *MockwsClientMockRecorder) Wait(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockwsClient)(nil).Wait), timeout)
}

// Write mocks base method.
func (m *MockwsClient) Write(id string, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", id, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockwsClientMockRecorder) Write(id, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockwsClient)(nil).Write), id, data)
}

// WriteSync mocks base method.
func (m *MockwsClient) WriteSync(id string, data []byte, timeout time.Duration) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteSync", id, data, timeout)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteSync indicates an expected call of WriteSync.
func (mr *MockwsClientMockRecorder) WriteSync(id, data, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteSync", reflect.TypeOf((*MockwsClient)(nil).WriteSync), id, data, timeout)
}

// MockwsConnection is a mock of wsConnection interface.
type MockwsConnection struct {
	ctrl     *gomock.Controller
	recorder *MockwsConnectionMockRecorder
}

// MockwsConnectionMockRecorder is the mock recorder for MockwsConnection.
type MockwsConnectionMockRecorder struct {
	mock *MockwsConnection
}

// NewMockwsConnection creates a new mock instance.
func NewMockwsConnection(ctrl *gomock.Controller) *MockwsConnection {
	mock := &MockwsConnection{ctrl: ctrl}
	mock.recorder = &MockwsConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockwsConnection) EXPECT() *MockwsConnectionMockRecorder {
	return m.recorder
}

// ReadMessage mocks base method.
func (m *MockwsConnection) ReadMessage() (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadMessage indicates an expected call of ReadMessage.
func (mr *MockwsConnectionMockRecorder) ReadMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockwsConnection)(nil).ReadMessage))
}

// WriteMessage mocks base method.
func (m *MockwsConnection) WriteMessage(messageType int, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", messageType, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage.
func (mr *MockwsConnectionMockRecorder) WriteMessage(messageType, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockwsConnection)(nil).WriteMessage), messageType, data)
}
