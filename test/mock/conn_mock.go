// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bloXroute-Labs/gateway/v2/connections (interfaces: Conn)
//
// Generated by this command:
//
//	mockgen -destination ../test/mock/conn_mock.go -package mock . Conn
//
// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	bxmessage "github.com/bloXroute-Labs/gateway/v2/bxmessage"
	connections "github.com/bloXroute-Labs/gateway/v2/connections"
	logger "github.com/bloXroute-Labs/gateway/v2/logger"
	types "github.com/bloXroute-Labs/gateway/v2/types"
	utils "github.com/bloXroute-Labs/gateway/v2/utils"
	gomock "go.uber.org/mock/gomock"
)

// MockConn is a mock of Conn interface.
type MockConn struct {
	ctrl     *gomock.Controller
	recorder *MockConnMockRecorder
}

// MockConnMockRecorder is the mock recorder for MockConn.
type MockConnMockRecorder struct {
	mock *MockConn
}

// NewMockConn creates a new mock instance.
func NewMockConn(ctrl *gomock.Controller) *MockConn {
	mock := &MockConn{ctrl: ctrl}
	mock.recorder = &MockConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConn) EXPECT() *MockConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConn) Close(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConn)(nil).Close), arg0)
}

// Connect mocks base method.
func (m *MockConn) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockConnMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockConn)(nil).Connect))
}

// Disable mocks base method.
func (m *MockConn) Disable(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disable", arg0)
}

// Disable indicates an expected call of Disable.
func (mr *MockConnMockRecorder) Disable(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockConn)(nil).Disable), arg0)
}

// GetAccountID mocks base method.
func (m *MockConn) GetAccountID() types.AccountID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountID")
	ret0, _ := ret[0].(types.AccountID)
	return ret0
}

// GetAccountID indicates an expected call of GetAccountID.
func (mr *MockConnMockRecorder) GetAccountID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountID", reflect.TypeOf((*MockConn)(nil).GetAccountID))
}

// GetCapabilities mocks base method.
func (m *MockConn) GetCapabilities() types.CapabilityFlags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapabilities")
	ret0, _ := ret[0].(types.CapabilityFlags)
	return ret0
}

// GetCapabilities indicates an expected call of GetCapabilities.
func (mr *MockConnMockRecorder) GetCapabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapabilities", reflect.TypeOf((*MockConn)(nil).GetCapabilities))
}

// GetConnectedAt mocks base method.
func (m *MockConn) GetConnectedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetConnectedAt indicates an expected call of GetConnectedAt.
func (mr *MockConnMockRecorder) GetConnectedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectedAt", reflect.TypeOf((*MockConn)(nil).GetConnectedAt))
}

// GetConnectionState mocks base method.
func (m *MockConn) GetConnectionState() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectionState")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetConnectionState indicates an expected call of GetConnectionState.
func (mr *MockConnMockRecorder) GetConnectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectionState", reflect.TypeOf((*MockConn)(nil).GetConnectionState))
}

// GetConnectionType mocks base method.
func (m *MockConn) GetConnectionType() utils.NodeType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectionType")
	ret0, _ := ret[0].(utils.NodeType)
	return ret0
}

// GetConnectionType indicates an expected call of GetConnectionType.
func (mr *MockConnMockRecorder) GetConnectionType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectionType", reflect.TypeOf((*MockConn)(nil).GetConnectionType))
}

// GetLocalPort mocks base method.
func (m *MockConn) GetLocalPort() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalPort")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetLocalPort indicates an expected call of GetLocalPort.
func (mr *MockConnMockRecorder) GetLocalPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalPort", reflect.TypeOf((*MockConn)(nil).GetLocalPort))
}

// GetNetworkNum mocks base method.
func (m *MockConn) GetNetworkNum() types.NetworkNum {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkNum")
	ret0, _ := ret[0].(types.NetworkNum)
	return ret0
}

// GetNetworkNum indicates an expected call of GetNetworkNum.
func (mr *MockConnMockRecorder) GetNetworkNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkNum", reflect.TypeOf((*MockConn)(nil).GetNetworkNum))
}

// GetNodeID mocks base method.
func (m *MockConn) GetNodeID() types.NodeID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeID")
	ret0, _ := ret[0].(types.NodeID)
	return ret0
}

// GetNodeID indicates an expected call of GetNodeID.
func (mr *MockConnMockRecorder) GetNodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeID", reflect.TypeOf((*MockConn)(nil).GetNodeID))
}

// GetPeerEnode mocks base method.
func (m *MockConn) GetPeerEnode() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerEnode")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPeerEnode indicates an expected call of GetPeerEnode.
func (mr *MockConnMockRecorder) GetPeerEnode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerEnode", reflect.TypeOf((*MockConn)(nil).GetPeerEnode))
}

// GetPeerIP mocks base method.
func (m *MockConn) GetPeerIP() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerIP")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPeerIP indicates an expected call of GetPeerIP.
func (mr *MockConnMockRecorder) GetPeerIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerIP", reflect.TypeOf((*MockConn)(nil).GetPeerIP))
}

// GetPeerPort mocks base method.
func (m *MockConn) GetPeerPort() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerPort")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetPeerPort indicates an expected call of GetPeerPort.
func (mr *MockConnMockRecorder) GetPeerPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerPort", reflect.TypeOf((*MockConn)(nil).GetPeerPort))
}

// GetVersion mocks base method.
func (m *MockConn) GetVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockConnMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockConn)(nil).GetVersion))
}

// ID mocks base method.
func (m *MockConn) ID() connections.Socket {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(connections.Socket)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockConnMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockConn)(nil).ID))
}

// IsDisabled mocks base method.
func (m *MockConn) IsDisabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDisabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDisabled indicates an expected call of IsDisabled.
func (mr *MockConnMockRecorder) IsDisabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDisabled", reflect.TypeOf((*MockConn)(nil).IsDisabled))
}

// IsInitiator mocks base method.
func (m *MockConn) IsInitiator() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInitiator")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInitiator indicates an expected call of IsInitiator.
func (mr *MockConnMockRecorder) IsInitiator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInitiator", reflect.TypeOf((*MockConn)(nil).IsInitiator))
}

// IsLocalGEO mocks base method.
func (m *MockConn) IsLocalGEO() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLocalGEO")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLocalGEO indicates an expected call of IsLocalGEO.
func (mr *MockConnMockRecorder) IsLocalGEO() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLocalGEO", reflect.TypeOf((*MockConn)(nil).IsLocalGEO))
}

// IsOpen mocks base method.
func (m *MockConn) IsOpen() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOpen")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOpen indicates an expected call of IsOpen.
func (mr *MockConnMockRecorder) IsOpen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOpen", reflect.TypeOf((*MockConn)(nil).IsOpen))
}

// IsPrivateNetwork mocks base method.
func (m *MockConn) IsPrivateNetwork() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrivateNetwork")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPrivateNetwork indicates an expected call of IsPrivateNetwork.
func (mr *MockConnMockRecorder) IsPrivateNetwork() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrivateNetwork", reflect.TypeOf((*MockConn)(nil).IsPrivateNetwork))
}

// IsSameRegion mocks base method.
func (m *MockConn) IsSameRegion() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSameRegion")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSameRegion indicates an expected call of IsSameRegion.
func (mr *MockConnMockRecorder) IsSameRegion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSameRegion", reflect.TypeOf((*MockConn)(nil).IsSameRegion))
}

// Log mocks base method.
func (m *MockConn) Log() *logger.Entry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log")
	ret0, _ := ret[0].(*logger.Entry)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockConnMockRecorder) Log() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockConn)(nil).Log))
}

// Protocol mocks base method.
func (m *MockConn) Protocol() bxmessage.Protocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(bxmessage.Protocol)
	return ret0
}

// Protocol indicates an expected call of Protocol.
func (mr *MockConnMockRecorder) Protocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockConn)(nil).Protocol))
}

// ReadMessages mocks base method.
func (m *MockConn) ReadMessages(arg0 func(bxmessage.MessageBytes), arg1 time.Duration, arg2 int, arg3 func([]byte) int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessages", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMessages indicates an expected call of ReadMessages.
func (mr *MockConnMockRecorder) ReadMessages(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessages", reflect.TypeOf((*MockConn)(nil).ReadMessages), arg0, arg1, arg2, arg3)
}

// Send mocks base method.
func (m *MockConn) Send(arg0 bxmessage.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockConnMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockConn)(nil).Send), arg0)
}

// SendWithDelay mocks base method.
func (m *MockConn) SendWithDelay(arg0 bxmessage.Message, arg1 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendWithDelay", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendWithDelay indicates an expected call of SendWithDelay.
func (mr *MockConnMockRecorder) SendWithDelay(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendWithDelay", reflect.TypeOf((*MockConn)(nil).SendWithDelay), arg0, arg1)
}

// SetProtocol mocks base method.
func (m *MockConn) SetProtocol(arg0 bxmessage.Protocol) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetProtocol", arg0)
}

// SetProtocol indicates an expected call of SetProtocol.
func (mr *MockConnMockRecorder) SetProtocol(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProtocol", reflect.TypeOf((*MockConn)(nil).SetProtocol), arg0)
}
