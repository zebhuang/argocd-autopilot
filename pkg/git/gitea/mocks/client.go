// Code generated by MockGen. DO NOT EDIT.
// Source: ./provider_gitea.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitea "code.gitea.io/sdk/gitea"
	gomock "github.com/golang/mock/gomock"
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

// CreateOrgRepo mocks base method.
func (m *MockClient) CreateOrgRepo(org string, opt gitea.CreateRepoOption) (*gitea.Repository, *gitea.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrgRepo", org, opt)
	ret0, _ := ret[0].(*gitea.Repository)
	ret1, _ := ret[1].(*gitea.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateOrgRepo indicates an expected call of CreateOrgRepo.
func (mr *MockClientMockRecorder) CreateOrgRepo(org, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgRepo", reflect.TypeOf((*MockClient)(nil).CreateOrgRepo), org, opt)
}

// CreateRepo mocks base method.
func (m *MockClient) CreateRepo(opt gitea.CreateRepoOption) (*gitea.Repository, *gitea.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRepo", opt)
	ret0, _ := ret[0].(*gitea.Repository)
	ret1, _ := ret[1].(*gitea.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRepo indicates an expected call of CreateRepo.
func (mr *MockClientMockRecorder) CreateRepo(opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRepo", reflect.TypeOf((*MockClient)(nil).CreateRepo), opt)
}

// GetMyUserInfo mocks base method.
func (m *MockClient) GetMyUserInfo() (*gitea.User, *gitea.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyUserInfo")
	ret0, _ := ret[0].(*gitea.User)
	ret1, _ := ret[1].(*gitea.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMyUserInfo indicates an expected call of GetMyUserInfo.
func (mr *MockClientMockRecorder) GetMyUserInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyUserInfo", reflect.TypeOf((*MockClient)(nil).GetMyUserInfo))
}
