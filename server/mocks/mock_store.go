// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Brightscout/mattermost-plugin-servicenow-virtual-agent/server/plugin (interfaces: Store)

// Package mock_plugin is a generated GoMock package.
package mock_plugin

import (
	gomock "github.com/golang/mock/gomock"
	serializer "github.com/mattermost/mattermost-plugin-servicenow-virtual-agent/server/serializer"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteUser mocks base method
func (m *MockStore) DeleteUser(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockStoreMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0)
}

// LoadUser mocks base method
func (m *MockStore) LoadUser(arg0 string) (*serializer.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUser", arg0)
	ret0, _ := ret[0].(*serializer.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUser indicates an expected call of LoadUser
func (mr *MockStoreMockRecorder) LoadUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUser", reflect.TypeOf((*MockStore)(nil).LoadUser), arg0)
}

// LoadUserWithSysID mocks base method
func (m *MockStore) LoadUserWithSysID(arg0 string) (*serializer.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUserWithSysID", arg0)
	ret0, _ := ret[0].(*serializer.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUserWithSysID indicates an expected call of LoadUserWithSysID
func (mr *MockStoreMockRecorder) LoadUserWithSysID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUserWithSysID", reflect.TypeOf((*MockStore)(nil).LoadUserWithSysID), arg0)
}

// StoreOAuth2State mocks base method
func (m *MockStore) StoreOAuth2State(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreOAuth2State", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreOAuth2State indicates an expected call of StoreOAuth2State
func (mr *MockStoreMockRecorder) StoreOAuth2State(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreOAuth2State", reflect.TypeOf((*MockStore)(nil).StoreOAuth2State), arg0)
}

// StoreUser mocks base method
func (m *MockStore) StoreUser(arg0 *serializer.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreUser indicates an expected call of StoreUser
func (mr *MockStoreMockRecorder) StoreUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreUser", reflect.TypeOf((*MockStore)(nil).StoreUser), arg0)
}

// VerifyOAuth2State mocks base method
func (m *MockStore) VerifyOAuth2State(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOAuth2State", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyOAuth2State indicates an expected call of VerifyOAuth2State
func (mr *MockStoreMockRecorder) VerifyOAuth2State(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOAuth2State", reflect.TypeOf((*MockStore)(nil).VerifyOAuth2State), arg0)
}
