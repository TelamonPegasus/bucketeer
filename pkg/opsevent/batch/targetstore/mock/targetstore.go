// Code generated by MockGen. DO NOT EDIT.
// Source: targetstore.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	domain "github.com/bucketeer-io/bucketeer/pkg/autoops/domain"
	domain0 "github.com/bucketeer-io/bucketeer/pkg/environment/domain"
)

// MockEnvironmentLister is a mock of EnvironmentLister interface.
type MockEnvironmentLister struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentListerMockRecorder
}

// MockEnvironmentListerMockRecorder is the mock recorder for MockEnvironmentLister.
type MockEnvironmentListerMockRecorder struct {
	mock *MockEnvironmentLister
}

// NewMockEnvironmentLister creates a new mock instance.
func NewMockEnvironmentLister(ctrl *gomock.Controller) *MockEnvironmentLister {
	mock := &MockEnvironmentLister{ctrl: ctrl}
	mock.recorder = &MockEnvironmentListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironmentLister) EXPECT() *MockEnvironmentListerMockRecorder {
	return m.recorder
}

// GetEnvironments mocks base method.
func (m *MockEnvironmentLister) GetEnvironments(ctx context.Context) []*domain0.Environment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironments", ctx)
	ret0, _ := ret[0].([]*domain0.Environment)
	return ret0
}

// GetEnvironments indicates an expected call of GetEnvironments.
func (mr *MockEnvironmentListerMockRecorder) GetEnvironments(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironments", reflect.TypeOf((*MockEnvironmentLister)(nil).GetEnvironments), ctx)
}

// MockAutoOpsRuleLister is a mock of AutoOpsRuleLister interface.
type MockAutoOpsRuleLister struct {
	ctrl     *gomock.Controller
	recorder *MockAutoOpsRuleListerMockRecorder
}

// MockAutoOpsRuleListerMockRecorder is the mock recorder for MockAutoOpsRuleLister.
type MockAutoOpsRuleListerMockRecorder struct {
	mock *MockAutoOpsRuleLister
}

// NewMockAutoOpsRuleLister creates a new mock instance.
func NewMockAutoOpsRuleLister(ctrl *gomock.Controller) *MockAutoOpsRuleLister {
	mock := &MockAutoOpsRuleLister{ctrl: ctrl}
	mock.recorder = &MockAutoOpsRuleListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutoOpsRuleLister) EXPECT() *MockAutoOpsRuleListerMockRecorder {
	return m.recorder
}

// GetAutoOpsRules mocks base method.
func (m *MockAutoOpsRuleLister) GetAutoOpsRules(ctx context.Context, environmentNamespace string) []*domain.AutoOpsRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutoOpsRules", ctx, environmentNamespace)
	ret0, _ := ret[0].([]*domain.AutoOpsRule)
	return ret0
}

// GetAutoOpsRules indicates an expected call of GetAutoOpsRules.
func (mr *MockAutoOpsRuleListerMockRecorder) GetAutoOpsRules(ctx, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoOpsRules", reflect.TypeOf((*MockAutoOpsRuleLister)(nil).GetAutoOpsRules), ctx, environmentNamespace)
}

// MockTargetStore is a mock of TargetStore interface.
type MockTargetStore struct {
	ctrl     *gomock.Controller
	recorder *MockTargetStoreMockRecorder
}

// MockTargetStoreMockRecorder is the mock recorder for MockTargetStore.
type MockTargetStoreMockRecorder struct {
	mock *MockTargetStore
}

// NewMockTargetStore creates a new mock instance.
func NewMockTargetStore(ctrl *gomock.Controller) *MockTargetStore {
	mock := &MockTargetStore{ctrl: ctrl}
	mock.recorder = &MockTargetStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTargetStore) EXPECT() *MockTargetStoreMockRecorder {
	return m.recorder
}

// GetAutoOpsRules mocks base method.
func (m *MockTargetStore) GetAutoOpsRules(ctx context.Context, environmentNamespace string) []*domain.AutoOpsRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutoOpsRules", ctx, environmentNamespace)
	ret0, _ := ret[0].([]*domain.AutoOpsRule)
	return ret0
}

// GetAutoOpsRules indicates an expected call of GetAutoOpsRules.
func (mr *MockTargetStoreMockRecorder) GetAutoOpsRules(ctx, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoOpsRules", reflect.TypeOf((*MockTargetStore)(nil).GetAutoOpsRules), ctx, environmentNamespace)
}

// GetEnvironments mocks base method.
func (m *MockTargetStore) GetEnvironments(ctx context.Context) []*domain0.Environment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironments", ctx)
	ret0, _ := ret[0].([]*domain0.Environment)
	return ret0
}

// GetEnvironments indicates an expected call of GetEnvironments.
func (mr *MockTargetStoreMockRecorder) GetEnvironments(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironments", reflect.TypeOf((*MockTargetStore)(nil).GetEnvironments), ctx)
}

// Run mocks base method.
func (m *MockTargetStore) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run.
func (mr *MockTargetStoreMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockTargetStore)(nil).Run))
}

// Stop mocks base method.
func (m *MockTargetStore) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockTargetStoreMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTargetStore)(nil).Stop))
}
