// Code generated by MockGen. DO NOT EDIT.
// Source: feature.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	domain "github.com/bucketeer-io/bucketeer/pkg/feature/domain"
	mysql "github.com/bucketeer-io/bucketeer/pkg/storage/v2/mysql"
	feature "github.com/bucketeer-io/bucketeer/proto/feature"
)

// MockFeatureStorage is a mock of FeatureStorage interface.
type MockFeatureStorage struct {
	ctrl     *gomock.Controller
	recorder *MockFeatureStorageMockRecorder
}

// MockFeatureStorageMockRecorder is the mock recorder for MockFeatureStorage.
type MockFeatureStorageMockRecorder struct {
	mock *MockFeatureStorage
}

// NewMockFeatureStorage creates a new mock instance.
func NewMockFeatureStorage(ctrl *gomock.Controller) *MockFeatureStorage {
	mock := &MockFeatureStorage{ctrl: ctrl}
	mock.recorder = &MockFeatureStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeatureStorage) EXPECT() *MockFeatureStorageMockRecorder {
	return m.recorder
}

// CreateFeature mocks base method.
func (m *MockFeatureStorage) CreateFeature(ctx context.Context, feature *domain.Feature, environmentNamespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeature", ctx, feature, environmentNamespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFeature indicates an expected call of CreateFeature.
func (mr *MockFeatureStorageMockRecorder) CreateFeature(ctx, feature, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeature", reflect.TypeOf((*MockFeatureStorage)(nil).CreateFeature), ctx, feature, environmentNamespace)
}

// GetFeature mocks base method.
func (m *MockFeatureStorage) GetFeature(ctx context.Context, key, environmentNamespace string) (*domain.Feature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeature", ctx, key, environmentNamespace)
	ret0, _ := ret[0].(*domain.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeature indicates an expected call of GetFeature.
func (mr *MockFeatureStorageMockRecorder) GetFeature(ctx, key, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeature", reflect.TypeOf((*MockFeatureStorage)(nil).GetFeature), ctx, key, environmentNamespace)
}

// ListFeatures mocks base method.
func (m *MockFeatureStorage) ListFeatures(ctx context.Context, whereParts []mysql.WherePart, orders []*mysql.Order, limit, offset int) ([]*feature.Feature, int, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeatures", ctx, whereParts, orders, limit, offset)
	ret0, _ := ret[0].([]*feature.Feature)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListFeatures indicates an expected call of ListFeatures.
func (mr *MockFeatureStorageMockRecorder) ListFeatures(ctx, whereParts, orders, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeatures", reflect.TypeOf((*MockFeatureStorage)(nil).ListFeatures), ctx, whereParts, orders, limit, offset)
}

// ListFeaturesFilteredByExperiment mocks base method.
func (m *MockFeatureStorage) ListFeaturesFilteredByExperiment(ctx context.Context, whereParts []mysql.WherePart, orders []*mysql.Order, limit, offset int) ([]*feature.Feature, int, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeaturesFilteredByExperiment", ctx, whereParts, orders, limit, offset)
	ret0, _ := ret[0].([]*feature.Feature)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListFeaturesFilteredByExperiment indicates an expected call of ListFeaturesFilteredByExperiment.
func (mr *MockFeatureStorageMockRecorder) ListFeaturesFilteredByExperiment(ctx, whereParts, orders, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeaturesFilteredByExperiment", reflect.TypeOf((*MockFeatureStorage)(nil).ListFeaturesFilteredByExperiment), ctx, whereParts, orders, limit, offset)
}

// UpdateFeature mocks base method.
func (m *MockFeatureStorage) UpdateFeature(ctx context.Context, feature *domain.Feature, environmentNamespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFeature", ctx, feature, environmentNamespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFeature indicates an expected call of UpdateFeature.
func (mr *MockFeatureStorageMockRecorder) UpdateFeature(ctx, feature, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFeature", reflect.TypeOf((*MockFeatureStorage)(nil).UpdateFeature), ctx, feature, environmentNamespace)
}
