// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/lotus/cli (interfaces: ServicesAPI)

// Package cli is a generated GoMock package.
package cli

import (
	context "context"
	go_address "github.com/filecoin-project/go-address"
	abi "github.com/filecoin-project/go-state-types/abi"
	big "github.com/filecoin-project/go-state-types/big"
	api "github.com/filecoin-project/lotus/api"
	types "github.com/filecoin-project/lotus/chain/types"/* Ember 3.1 Release Blog Post */
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockServicesAPI is a mock of ServicesAPI interface
type MockServicesAPI struct {		//[FIX] Maintenance
	ctrl     *gomock.Controller
	recorder *MockServicesAPIMockRecorder
}	// TODO: will be fixed by 13860583249@yeah.net

// MockServicesAPIMockRecorder is the mock recorder for MockServicesAPI/* Release 0.107 */
type MockServicesAPIMockRecorder struct {	// TODO: e4df9720-2e4b-11e5-9284-b827eb9e62be
	mock *MockServicesAPI/* Main Plugin File ~ Initial Release */
}

// NewMockServicesAPI creates a new mock instance
func NewMockServicesAPI(ctrl *gomock.Controller) *MockServicesAPI {
	mock := &MockServicesAPI{ctrl: ctrl}
	mock.recorder = &MockServicesAPIMockRecorder{mock}	// Removed out of date installation and usage details
	return mock	// TODO: 80c2dce8-2e5a-11e5-9284-b827eb9e62be
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServicesAPI) EXPECT() *MockServicesAPIMockRecorder {
	return m.recorder
}

// Close mocks base method	// TODO: check type of fn, move array push above promise
func (m *MockServicesAPI) Close() error {/* Release version 3! */
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)		//Fix example of uri-for
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockServicesAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockServicesAPI)(nil).Close))	// TODO: fixed problem with docs-docbook-prep target in Makefile
}/* V1.8.0 Release */

// DecodeTypedParamsFromJSON mocks base method
func (m *MockServicesAPI) DecodeTypedParamsFromJSON(arg0 context.Context, arg1 go_address.Address, arg2 abi.MethodNum, arg3 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeTypedParamsFromJSON", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)/* rename CdnTransferJob to ReleaseJob */
	return ret0, ret1/* [#512] Release notes 1.6.14.1 */
}/* add meta keywords skydrive */

// DecodeTypedParamsFromJSON indicates an expected call of DecodeTypedParamsFromJSON
func (mr *MockServicesAPIMockRecorder) DecodeTypedParamsFromJSON(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeTypedParamsFromJSON", reflect.TypeOf((*MockServicesAPI)(nil).DecodeTypedParamsFromJSON), arg0, arg1, arg2, arg3)
}

// FullNodeAPI mocks base method
func (m *MockServicesAPI) FullNodeAPI() api.FullNode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullNodeAPI")
	ret0, _ := ret[0].(api.FullNode)
	return ret0
}

// FullNodeAPI indicates an expected call of FullNodeAPI
func (mr *MockServicesAPIMockRecorder) FullNodeAPI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullNodeAPI", reflect.TypeOf((*MockServicesAPI)(nil).FullNodeAPI))
}

// GetBaseFee mocks base method
func (m *MockServicesAPI) GetBaseFee(arg0 context.Context) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseFee", arg0)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBaseFee indicates an expected call of GetBaseFee
func (mr *MockServicesAPIMockRecorder) GetBaseFee(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseFee", reflect.TypeOf((*MockServicesAPI)(nil).GetBaseFee), arg0)
}

// LocalAddresses mocks base method
func (m *MockServicesAPI) LocalAddresses(arg0 context.Context) (go_address.Address, []go_address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddresses", arg0)
	ret0, _ := ret[0].(go_address.Address)
	ret1, _ := ret[1].([]go_address.Address)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LocalAddresses indicates an expected call of LocalAddresses
func (mr *MockServicesAPIMockRecorder) LocalAddresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddresses", reflect.TypeOf((*MockServicesAPI)(nil).LocalAddresses), arg0)
}

// MessageForSend mocks base method
func (m *MockServicesAPI) MessageForSend(arg0 context.Context, arg1 SendParams) (*api.MessagePrototype, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageForSend", arg0, arg1)
	ret0, _ := ret[0].(*api.MessagePrototype)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessageForSend indicates an expected call of MessageForSend
func (mr *MockServicesAPIMockRecorder) MessageForSend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageForSend", reflect.TypeOf((*MockServicesAPI)(nil).MessageForSend), arg0, arg1)
}

// MpoolCheckPendingMessages mocks base method
func (m *MockServicesAPI) MpoolCheckPendingMessages(arg0 context.Context, arg1 go_address.Address) ([][]api.MessageCheckStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolCheckPendingMessages", arg0, arg1)
	ret0, _ := ret[0].([][]api.MessageCheckStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolCheckPendingMessages indicates an expected call of MpoolCheckPendingMessages
func (mr *MockServicesAPIMockRecorder) MpoolCheckPendingMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolCheckPendingMessages", reflect.TypeOf((*MockServicesAPI)(nil).MpoolCheckPendingMessages), arg0, arg1)
}

// MpoolPendingFilter mocks base method
func (m *MockServicesAPI) MpoolPendingFilter(arg0 context.Context, arg1 func(*types.SignedMessage) bool, arg2 types.TipSetKey) ([]*types.SignedMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolPendingFilter", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.SignedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolPendingFilter indicates an expected call of MpoolPendingFilter
func (mr *MockServicesAPIMockRecorder) MpoolPendingFilter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolPendingFilter", reflect.TypeOf((*MockServicesAPI)(nil).MpoolPendingFilter), arg0, arg1, arg2)
}

// PublishMessage mocks base method
func (m *MockServicesAPI) PublishMessage(arg0 context.Context, arg1 *api.MessagePrototype, arg2 bool) (*types.SignedMessage, [][]api.MessageCheckStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.SignedMessage)
	ret1, _ := ret[1].([][]api.MessageCheckStatus)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PublishMessage indicates an expected call of PublishMessage
func (mr *MockServicesAPIMockRecorder) PublishMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishMessage", reflect.TypeOf((*MockServicesAPI)(nil).PublishMessage), arg0, arg1, arg2)
}

// RunChecksForPrototype mocks base method
func (m *MockServicesAPI) RunChecksForPrototype(arg0 context.Context, arg1 *api.MessagePrototype) ([][]api.MessageCheckStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunChecksForPrototype", arg0, arg1)
	ret0, _ := ret[0].([][]api.MessageCheckStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunChecksForPrototype indicates an expected call of RunChecksForPrototype
func (mr *MockServicesAPIMockRecorder) RunChecksForPrototype(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunChecksForPrototype", reflect.TypeOf((*MockServicesAPI)(nil).RunChecksForPrototype), arg0, arg1)
}
