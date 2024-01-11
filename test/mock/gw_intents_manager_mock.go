// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bloXroute-Labs/gateway/v2/nodes (interfaces: IntentsManager)
//
// Generated by this command:
//
//	mockgen -destination ../test/mock/gw_intents_manager_mock.go -package mock . IntentsManager
//
// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	bxmessage "github.com/bloXroute-Labs/gateway/v2/bxmessage"
	gateway "github.com/bloXroute-Labs/gateway/v2/protobuf"
	gomock "go.uber.org/mock/gomock"
)

// MockIntentsManager is a mock of IntentsManager interface.
type MockIntentsManager struct {
	ctrl     *gomock.Controller
	recorder *MockIntentsManagerMockRecorder
}

// MockIntentsManagerMockRecorder is the mock recorder for MockIntentsManager.
type MockIntentsManagerMockRecorder struct {
	mock *MockIntentsManager
}

// NewMockIntentsManager creates a new mock instance.
func NewMockIntentsManager(ctrl *gomock.Controller) *MockIntentsManager {
	mock := &MockIntentsManager{ctrl: ctrl}
	mock.recorder = &MockIntentsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntentsManager) EXPECT() *MockIntentsManagerMockRecorder {
	return m.recorder
}

// AddIntentsSubscription mocks base method.
func (m *MockIntentsManager) AddIntentsSubscription(arg0 *gateway.IntentsRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddIntentsSubscription", arg0)
}

// AddIntentsSubscription indicates an expected call of AddIntentsSubscription.
func (mr *MockIntentsManagerMockRecorder) AddIntentsSubscription(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIntentsSubscription", reflect.TypeOf((*MockIntentsManager)(nil).AddIntentsSubscription), arg0)
}

// AddSolutionsSubscription mocks base method.
func (m *MockIntentsManager) AddSolutionsSubscription(arg0 *gateway.IntentSolutionsRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSolutionsSubscription", arg0)
}

// AddSolutionsSubscription indicates an expected call of AddSolutionsSubscription.
func (mr *MockIntentsManagerMockRecorder) AddSolutionsSubscription(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSolutionsSubscription", reflect.TypeOf((*MockIntentsManager)(nil).AddSolutionsSubscription), arg0)
}

// IntentsSubscriptionExists mocks base method.
func (m *MockIntentsManager) IntentsSubscriptionExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntentsSubscriptionExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IntentsSubscriptionExists indicates an expected call of IntentsSubscriptionExists.
func (mr *MockIntentsManagerMockRecorder) IntentsSubscriptionExists(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntentsSubscriptionExists", reflect.TypeOf((*MockIntentsManager)(nil).IntentsSubscriptionExists), arg0)
}

// RmIntentsSubscription mocks base method.
func (m *MockIntentsManager) RmIntentsSubscription(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RmIntentsSubscription", arg0)
}

// RmIntentsSubscription indicates an expected call of RmIntentsSubscription.
func (mr *MockIntentsManagerMockRecorder) RmIntentsSubscription(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RmIntentsSubscription", reflect.TypeOf((*MockIntentsManager)(nil).RmIntentsSubscription), arg0)
}

// RmSolutionsSubscription mocks base method.
func (m *MockIntentsManager) RmSolutionsSubscription(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RmSolutionsSubscription", arg0)
}

// RmSolutionsSubscription indicates an expected call of RmSolutionsSubscription.
func (mr *MockIntentsManagerMockRecorder) RmSolutionsSubscription(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RmSolutionsSubscription", reflect.TypeOf((*MockIntentsManager)(nil).RmSolutionsSubscription), arg0)
}

// SolutionsSubscriptionExists mocks base method.
func (m *MockIntentsManager) SolutionsSubscriptionExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SolutionsSubscriptionExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SolutionsSubscriptionExists indicates an expected call of SolutionsSubscriptionExists.
func (mr *MockIntentsManagerMockRecorder) SolutionsSubscriptionExists(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SolutionsSubscriptionExists", reflect.TypeOf((*MockIntentsManager)(nil).SolutionsSubscriptionExists), arg0)
}

// SubscriptionMessages mocks base method.
func (m *MockIntentsManager) SubscriptionMessages() []bxmessage.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionMessages")
	ret0, _ := ret[0].([]bxmessage.Message)
	return ret0
}

// SubscriptionMessages indicates an expected call of SubscriptionMessages.
func (mr *MockIntentsManagerMockRecorder) SubscriptionMessages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionMessages", reflect.TypeOf((*MockIntentsManager)(nil).SubscriptionMessages))
}
