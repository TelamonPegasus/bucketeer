// Code generated by MockGen. DO NOT EDIT.
// Source: webhook.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	domain "github.com/bucketeer-io/bucketeer/pkg/autoops/domain"
	mysql "github.com/bucketeer-io/bucketeer/pkg/storage/v2/mysql"
	autoops "github.com/bucketeer-io/bucketeer/proto/autoops"
)

// MockWebhookStorage is a mock of WebhookStorage interface.
type MockWebhookStorage struct {
	ctrl     *gomock.Controller
	recorder *MockWebhookStorageMockRecorder
}

// MockWebhookStorageMockRecorder is the mock recorder for MockWebhookStorage.
type MockWebhookStorageMockRecorder struct {
	mock *MockWebhookStorage
}

// NewMockWebhookStorage creates a new mock instance.
func NewMockWebhookStorage(ctrl *gomock.Controller) *MockWebhookStorage {
	mock := &MockWebhookStorage{ctrl: ctrl}
	mock.recorder = &MockWebhookStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebhookStorage) EXPECT() *MockWebhookStorageMockRecorder {
	return m.recorder
}

// CreateWebhook mocks base method.
func (m *MockWebhookStorage) CreateWebhook(ctx context.Context, webhook *domain.Webhook, environmentNamespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWebhook", ctx, webhook, environmentNamespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWebhook indicates an expected call of CreateWebhook.
func (mr *MockWebhookStorageMockRecorder) CreateWebhook(ctx, webhook, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWebhook", reflect.TypeOf((*MockWebhookStorage)(nil).CreateWebhook), ctx, webhook, environmentNamespace)
}

// DeleteWebhook mocks base method.
func (m *MockWebhookStorage) DeleteWebhook(ctx context.Context, id, environmentNamespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWebhook", ctx, id, environmentNamespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWebhook indicates an expected call of DeleteWebhook.
func (mr *MockWebhookStorageMockRecorder) DeleteWebhook(ctx, id, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWebhook", reflect.TypeOf((*MockWebhookStorage)(nil).DeleteWebhook), ctx, id, environmentNamespace)
}

// GetWebhook mocks base method.
func (m *MockWebhookStorage) GetWebhook(ctx context.Context, id, environmentNamespace string) (*domain.Webhook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWebhook", ctx, id, environmentNamespace)
	ret0, _ := ret[0].(*domain.Webhook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebhook indicates an expected call of GetWebhook.
func (mr *MockWebhookStorageMockRecorder) GetWebhook(ctx, id, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebhook", reflect.TypeOf((*MockWebhookStorage)(nil).GetWebhook), ctx, id, environmentNamespace)
}

// ListWebhooks mocks base method.
func (m *MockWebhookStorage) ListWebhooks(ctx context.Context, whereParts []mysql.WherePart, orders []*mysql.Order, limit, offset int) ([]*autoops.Webhook, int, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWebhooks", ctx, whereParts, orders, limit, offset)
	ret0, _ := ret[0].([]*autoops.Webhook)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListWebhooks indicates an expected call of ListWebhooks.
func (mr *MockWebhookStorageMockRecorder) ListWebhooks(ctx, whereParts, orders, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebhooks", reflect.TypeOf((*MockWebhookStorage)(nil).ListWebhooks), ctx, whereParts, orders, limit, offset)
}

// UpdateWebhook mocks base method.
func (m *MockWebhookStorage) UpdateWebhook(ctx context.Context, webhook *domain.Webhook, environmentNamespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWebhook", ctx, webhook, environmentNamespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWebhook indicates an expected call of UpdateWebhook.
func (mr *MockWebhookStorageMockRecorder) UpdateWebhook(ctx, webhook, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWebhook", reflect.TypeOf((*MockWebhookStorage)(nil).UpdateWebhook), ctx, webhook, environmentNamespace)
}
