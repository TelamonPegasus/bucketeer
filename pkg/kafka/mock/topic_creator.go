// Code generated by MockGen. DO NOT EDIT.
// Source: topic_creator.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTopicCreator is a mock of TopicCreator interface.
type MockTopicCreator struct {
	ctrl     *gomock.Controller
	recorder *MockTopicCreatorMockRecorder
}

// MockTopicCreatorMockRecorder is the mock recorder for MockTopicCreator.
type MockTopicCreatorMockRecorder struct {
	mock *MockTopicCreator
}

// NewMockTopicCreator creates a new mock instance.
func NewMockTopicCreator(ctrl *gomock.Controller) *MockTopicCreator {
	mock := &MockTopicCreator{ctrl: ctrl}
	mock.recorder = &MockTopicCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopicCreator) EXPECT() *MockTopicCreatorMockRecorder {
	return m.recorder
}

// CreateTopics mocks base method.
func (m *MockTopicCreator) CreateTopics(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopics", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTopics indicates an expected call of CreateTopics.
func (mr *MockTopicCreatorMockRecorder) CreateTopics(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopics", reflect.TypeOf((*MockTopicCreator)(nil).CreateTopics), ctx)
}
