// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"

	environment "github.com/bucketeer-io/bucketeer/proto/environment"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ArchiveEnvironmentV2 mocks base method.
func (m *MockClient) ArchiveEnvironmentV2(ctx context.Context, in *environment.ArchiveEnvironmentV2Request, opts ...grpc.CallOption) (*environment.ArchiveEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ArchiveEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.ArchiveEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArchiveEnvironmentV2 indicates an expected call of ArchiveEnvironmentV2.
func (mr *MockClientMockRecorder) ArchiveEnvironmentV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveEnvironmentV2", reflect.TypeOf((*MockClient)(nil).ArchiveEnvironmentV2), varargs...)
}

// Close mocks base method.
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// ConvertTrialProject mocks base method.
func (m *MockClient) ConvertTrialProject(ctx context.Context, in *environment.ConvertTrialProjectRequest, opts ...grpc.CallOption) (*environment.ConvertTrialProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConvertTrialProject", varargs...)
	ret0, _ := ret[0].(*environment.ConvertTrialProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertTrialProject indicates an expected call of ConvertTrialProject.
func (mr *MockClientMockRecorder) ConvertTrialProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertTrialProject", reflect.TypeOf((*MockClient)(nil).ConvertTrialProject), varargs...)
}

// CreateEnvironment mocks base method.
func (m *MockClient) CreateEnvironment(ctx context.Context, in *environment.CreateEnvironmentRequest, opts ...grpc.CallOption) (*environment.CreateEnvironmentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEnvironment", varargs...)
	ret0, _ := ret[0].(*environment.CreateEnvironmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEnvironment indicates an expected call of CreateEnvironment.
func (mr *MockClientMockRecorder) CreateEnvironment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironment", reflect.TypeOf((*MockClient)(nil).CreateEnvironment), varargs...)
}

// CreateEnvironmentV2 mocks base method.
func (m *MockClient) CreateEnvironmentV2(ctx context.Context, in *environment.CreateEnvironmentV2Request, opts ...grpc.CallOption) (*environment.CreateEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.CreateEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEnvironmentV2 indicates an expected call of CreateEnvironmentV2.
func (mr *MockClientMockRecorder) CreateEnvironmentV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironmentV2", reflect.TypeOf((*MockClient)(nil).CreateEnvironmentV2), varargs...)
}

// CreateProject mocks base method.
func (m *MockClient) CreateProject(ctx context.Context, in *environment.CreateProjectRequest, opts ...grpc.CallOption) (*environment.CreateProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProject", varargs...)
	ret0, _ := ret[0].(*environment.CreateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockClientMockRecorder) CreateProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockClient)(nil).CreateProject), varargs...)
}

// CreateTrialProject mocks base method.
func (m *MockClient) CreateTrialProject(ctx context.Context, in *environment.CreateTrialProjectRequest, opts ...grpc.CallOption) (*environment.CreateTrialProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTrialProject", varargs...)
	ret0, _ := ret[0].(*environment.CreateTrialProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrialProject indicates an expected call of CreateTrialProject.
func (mr *MockClientMockRecorder) CreateTrialProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrialProject", reflect.TypeOf((*MockClient)(nil).CreateTrialProject), varargs...)
}

// DeleteEnvironment mocks base method.
func (m *MockClient) DeleteEnvironment(ctx context.Context, in *environment.DeleteEnvironmentRequest, opts ...grpc.CallOption) (*environment.DeleteEnvironmentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEnvironment", varargs...)
	ret0, _ := ret[0].(*environment.DeleteEnvironmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEnvironment indicates an expected call of DeleteEnvironment.
func (mr *MockClientMockRecorder) DeleteEnvironment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvironment", reflect.TypeOf((*MockClient)(nil).DeleteEnvironment), varargs...)
}

// DisableProject mocks base method.
func (m *MockClient) DisableProject(ctx context.Context, in *environment.DisableProjectRequest, opts ...grpc.CallOption) (*environment.DisableProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableProject", varargs...)
	ret0, _ := ret[0].(*environment.DisableProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableProject indicates an expected call of DisableProject.
func (mr *MockClientMockRecorder) DisableProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableProject", reflect.TypeOf((*MockClient)(nil).DisableProject), varargs...)
}

// EnableProject mocks base method.
func (m *MockClient) EnableProject(ctx context.Context, in *environment.EnableProjectRequest, opts ...grpc.CallOption) (*environment.EnableProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableProject", varargs...)
	ret0, _ := ret[0].(*environment.EnableProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableProject indicates an expected call of EnableProject.
func (mr *MockClientMockRecorder) EnableProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableProject", reflect.TypeOf((*MockClient)(nil).EnableProject), varargs...)
}

// GetEnvironment mocks base method.
func (m *MockClient) GetEnvironment(ctx context.Context, in *environment.GetEnvironmentRequest, opts ...grpc.CallOption) (*environment.GetEnvironmentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnvironment", varargs...)
	ret0, _ := ret[0].(*environment.GetEnvironmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment.
func (mr *MockClientMockRecorder) GetEnvironment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockClient)(nil).GetEnvironment), varargs...)
}

// GetEnvironmentByNamespace mocks base method.
func (m *MockClient) GetEnvironmentByNamespace(ctx context.Context, in *environment.GetEnvironmentByNamespaceRequest, opts ...grpc.CallOption) (*environment.GetEnvironmentByNamespaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnvironmentByNamespace", varargs...)
	ret0, _ := ret[0].(*environment.GetEnvironmentByNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironmentByNamespace indicates an expected call of GetEnvironmentByNamespace.
func (mr *MockClientMockRecorder) GetEnvironmentByNamespace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironmentByNamespace", reflect.TypeOf((*MockClient)(nil).GetEnvironmentByNamespace), varargs...)
}

// GetEnvironmentV2 mocks base method.
func (m *MockClient) GetEnvironmentV2(ctx context.Context, in *environment.GetEnvironmentV2Request, opts ...grpc.CallOption) (*environment.GetEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.GetEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironmentV2 indicates an expected call of GetEnvironmentV2.
func (mr *MockClientMockRecorder) GetEnvironmentV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironmentV2", reflect.TypeOf((*MockClient)(nil).GetEnvironmentV2), varargs...)
}

// GetProject mocks base method.
func (m *MockClient) GetProject(ctx context.Context, in *environment.GetProjectRequest, opts ...grpc.CallOption) (*environment.GetProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProject", varargs...)
	ret0, _ := ret[0].(*environment.GetProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockClientMockRecorder) GetProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockClient)(nil).GetProject), varargs...)
}

// ListEnvironments mocks base method.
func (m *MockClient) ListEnvironments(ctx context.Context, in *environment.ListEnvironmentsRequest, opts ...grpc.CallOption) (*environment.ListEnvironmentsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnvironments", varargs...)
	ret0, _ := ret[0].(*environment.ListEnvironmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironments indicates an expected call of ListEnvironments.
func (mr *MockClientMockRecorder) ListEnvironments(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironments", reflect.TypeOf((*MockClient)(nil).ListEnvironments), varargs...)
}

// ListEnvironmentsV2 mocks base method.
func (m *MockClient) ListEnvironmentsV2(ctx context.Context, in *environment.ListEnvironmentsV2Request, opts ...grpc.CallOption) (*environment.ListEnvironmentsV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnvironmentsV2", varargs...)
	ret0, _ := ret[0].(*environment.ListEnvironmentsV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironmentsV2 indicates an expected call of ListEnvironmentsV2.
func (mr *MockClientMockRecorder) ListEnvironmentsV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironmentsV2", reflect.TypeOf((*MockClient)(nil).ListEnvironmentsV2), varargs...)
}

// ListProjects mocks base method.
func (m *MockClient) ListProjects(ctx context.Context, in *environment.ListProjectsRequest, opts ...grpc.CallOption) (*environment.ListProjectsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjects", varargs...)
	ret0, _ := ret[0].(*environment.ListProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockClientMockRecorder) ListProjects(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockClient)(nil).ListProjects), varargs...)
}

// UnarchiveEnvironmentV2 mocks base method.
func (m *MockClient) UnarchiveEnvironmentV2(ctx context.Context, in *environment.UnarchiveEnvironmentV2Request, opts ...grpc.CallOption) (*environment.UnarchiveEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnarchiveEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.UnarchiveEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnarchiveEnvironmentV2 indicates an expected call of UnarchiveEnvironmentV2.
func (mr *MockClientMockRecorder) UnarchiveEnvironmentV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnarchiveEnvironmentV2", reflect.TypeOf((*MockClient)(nil).UnarchiveEnvironmentV2), varargs...)
}

// UpdateEnvironment mocks base method.
func (m *MockClient) UpdateEnvironment(ctx context.Context, in *environment.UpdateEnvironmentRequest, opts ...grpc.CallOption) (*environment.UpdateEnvironmentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEnvironment", varargs...)
	ret0, _ := ret[0].(*environment.UpdateEnvironmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEnvironment indicates an expected call of UpdateEnvironment.
func (mr *MockClientMockRecorder) UpdateEnvironment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvironment", reflect.TypeOf((*MockClient)(nil).UpdateEnvironment), varargs...)
}

// UpdateEnvironmentV2 mocks base method.
func (m *MockClient) UpdateEnvironmentV2(ctx context.Context, in *environment.UpdateEnvironmentV2Request, opts ...grpc.CallOption) (*environment.UpdateEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.UpdateEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEnvironmentV2 indicates an expected call of UpdateEnvironmentV2.
func (mr *MockClientMockRecorder) UpdateEnvironmentV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvironmentV2", reflect.TypeOf((*MockClient)(nil).UpdateEnvironmentV2), varargs...)
}

// UpdateProject mocks base method.
func (m *MockClient) UpdateProject(ctx context.Context, in *environment.UpdateProjectRequest, opts ...grpc.CallOption) (*environment.UpdateProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProject", varargs...)
	ret0, _ := ret[0].(*environment.UpdateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockClientMockRecorder) UpdateProject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockClient)(nil).UpdateProject), varargs...)
}
