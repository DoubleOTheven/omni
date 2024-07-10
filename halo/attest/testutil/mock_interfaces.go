// Code generated by MockGen. DO NOT EDIT.
// Source: ./expected_interfaces.go
//
// Generated by this command:
//
//	mockgen -source ./expected_interfaces.go -package testutil -destination ./mock_interfaces.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	types "github.com/cosmos/cosmos-sdk/types"
	common "github.com/ethereum/go-ethereum/common"
	types0 "github.com/omni-network/omni/halo/attest/types"
	types1 "github.com/omni-network/omni/halo/valsync/types"
	xchain "github.com/omni-network/omni/lib/xchain"
	gomock "go.uber.org/mock/gomock"
)

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// GetPubKeyByConsAddr mocks base method.
func (m *MockStakingKeeper) GetPubKeyByConsAddr(arg0 context.Context, arg1 types.ConsAddress) (crypto.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKeyByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(crypto.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPubKeyByConsAddr indicates an expected call of GetPubKeyByConsAddr.
func (mr *MockStakingKeeperMockRecorder) GetPubKeyByConsAddr(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKeyByConsAddr", reflect.TypeOf((*MockStakingKeeper)(nil).GetPubKeyByConsAddr), arg0, arg1)
}

// MockVoter is a mock of Voter interface.
type MockVoter struct {
	ctrl     *gomock.Controller
	recorder *MockVoterMockRecorder
}

// MockVoterMockRecorder is the mock recorder for MockVoter.
type MockVoterMockRecorder struct {
	mock *MockVoter
}

// NewMockVoter creates a new mock instance.
func NewMockVoter(ctrl *gomock.Controller) *MockVoter {
	mock := &MockVoter{ctrl: ctrl}
	mock.recorder = &MockVoterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVoter) EXPECT() *MockVoterMockRecorder {
	return m.recorder
}

// GetAvailable mocks base method.
func (m *MockVoter) GetAvailable() []*types0.Vote {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailable")
	ret0, _ := ret[0].([]*types0.Vote)
	return ret0
}

// GetAvailable indicates an expected call of GetAvailable.
func (mr *MockVoterMockRecorder) GetAvailable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailable", reflect.TypeOf((*MockVoter)(nil).GetAvailable))
}

// LocalAddress mocks base method.
func (m *MockVoter) LocalAddress() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddress")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// LocalAddress indicates an expected call of LocalAddress.
func (mr *MockVoterMockRecorder) LocalAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddress", reflect.TypeOf((*MockVoter)(nil).LocalAddress))
}

// SetCommitted mocks base method.
func (m *MockVoter) SetCommitted(headers []*types0.BlockHeader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCommitted", headers)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCommitted indicates an expected call of SetCommitted.
func (mr *MockVoterMockRecorder) SetCommitted(headers any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCommitted", reflect.TypeOf((*MockVoter)(nil).SetCommitted), headers)
}

// SetProposed mocks base method.
func (m *MockVoter) SetProposed(headers []*types0.BlockHeader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProposed", headers)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProposed indicates an expected call of SetProposed.
func (mr *MockVoterMockRecorder) SetProposed(headers any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProposed", reflect.TypeOf((*MockVoter)(nil).SetProposed), headers)
}

// TrimBehind mocks base method.
func (m *MockVoter) TrimBehind(minsByChain map[xchain.ChainVersion]uint64) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrimBehind", minsByChain)
	ret0, _ := ret[0].(int)
	return ret0
}

// TrimBehind indicates an expected call of TrimBehind.
func (mr *MockVoterMockRecorder) TrimBehind(minsByChain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrimBehind", reflect.TypeOf((*MockVoter)(nil).TrimBehind), minsByChain)
}

// UpdateValidatorSet mocks base method.
func (m *MockVoter) UpdateValidatorSet(set *types1.ValidatorSetResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateValidatorSet", set)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateValidatorSet indicates an expected call of UpdateValidatorSet.
func (mr *MockVoterMockRecorder) UpdateValidatorSet(set any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValidatorSet", reflect.TypeOf((*MockVoter)(nil).UpdateValidatorSet), set)
}

// MockValProvider is a mock of ValProvider interface.
type MockValProvider struct {
	ctrl     *gomock.Controller
	recorder *MockValProviderMockRecorder
}

// MockValProviderMockRecorder is the mock recorder for MockValProvider.
type MockValProviderMockRecorder struct {
	mock *MockValProvider
}

// NewMockValProvider creates a new mock instance.
func NewMockValProvider(ctrl *gomock.Controller) *MockValProvider {
	mock := &MockValProvider{ctrl: ctrl}
	mock.recorder = &MockValProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValProvider) EXPECT() *MockValProviderMockRecorder {
	return m.recorder
}

// ActiveSetByHeight mocks base method.
func (m *MockValProvider) ActiveSetByHeight(ctx context.Context, height uint64) (*types1.ValidatorSetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveSetByHeight", ctx, height)
	ret0, _ := ret[0].(*types1.ValidatorSetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveSetByHeight indicates an expected call of ActiveSetByHeight.
func (mr *MockValProviderMockRecorder) ActiveSetByHeight(ctx, height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveSetByHeight", reflect.TypeOf((*MockValProvider)(nil).ActiveSetByHeight), ctx, height)
}

// MockChainNamer is a mock of ChainNamer interface.
type MockChainNamer struct {
	ctrl     *gomock.Controller
	recorder *MockChainNamerMockRecorder
}

// MockChainNamerMockRecorder is the mock recorder for MockChainNamer.
type MockChainNamerMockRecorder struct {
	mock *MockChainNamer
}

// NewMockChainNamer creates a new mock instance.
func NewMockChainNamer(ctrl *gomock.Controller) *MockChainNamer {
	mock := &MockChainNamer{ctrl: ctrl}
	mock.recorder = &MockChainNamerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainNamer) EXPECT() *MockChainNamerMockRecorder {
	return m.recorder
}

// ChainName mocks base method.
func (m *MockChainNamer) ChainName(chainVer xchain.ChainVersion) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainName", chainVer)
	ret0, _ := ret[0].(string)
	return ret0
}

// ChainName indicates an expected call of ChainName.
func (mr *MockChainNamerMockRecorder) ChainName(chainVer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainName", reflect.TypeOf((*MockChainNamer)(nil).ChainName), chainVer)
}
