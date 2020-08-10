// Code generated by MockGen. DO NOT EDIT.
// Source: main/internal/friends (interfaces: FriendUseCase)

// Package mock_friends is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "main/internal/models"
	reflect "reflect"
)

// MockFriendUseCase is a mock of FriendUseCase interface
type MockFriendUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockFriendUseCaseMockRecorder
}

// MockFriendUseCaseMockRecorder is the mock recorder for MockFriendUseCase
type MockFriendUseCaseMockRecorder struct {
	mock *MockFriendUseCase
}

// NewMockFriendUseCase creates a new mock instance
func NewMockFriendUseCase(ctrl *gomock.Controller) *MockFriendUseCase {
	mock := &MockFriendUseCase{ctrl: ctrl}
	mock.recorder = &MockFriendUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFriendUseCase) EXPECT() *MockFriendUseCaseMockRecorder {
	return m.recorder
}

// AddFriend mocks base method
func (m *MockFriendUseCase) AddFriend(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFriend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFriend indicates an expected call of AddFriend
func (mr *MockFriendUseCaseMockRecorder) AddFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFriend", reflect.TypeOf((*MockFriendUseCase)(nil).AddFriend), arg0, arg1)
}

// DeleteFriend mocks base method
func (m *MockFriendUseCase) DeleteFriend(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFriend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFriend indicates an expected call of DeleteFriend
func (mr *MockFriendUseCaseMockRecorder) DeleteFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFriend", reflect.TypeOf((*MockFriendUseCase)(nil).DeleteFriend), arg0, arg1)
}

// GetAllFriends mocks base method
func (m *MockFriendUseCase) GetAllFriends(arg0 string) ([]models.FriendLandingInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFriends", arg0)
	ret0, _ := ret[0].([]models.FriendLandingInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFriends indicates an expected call of GetAllFriends
func (mr *MockFriendUseCaseMockRecorder) GetAllFriends(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFriends", reflect.TypeOf((*MockFriendUseCase)(nil).GetAllFriends), arg0)
}

// GetUserLoginById mocks base method
func (m *MockFriendUseCase) GetUserLoginById(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLoginById", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLoginById indicates an expected call of GetUserLoginById
func (mr *MockFriendUseCaseMockRecorder) GetUserLoginById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLoginById", reflect.TypeOf((*MockFriendUseCase)(nil).GetUserLoginById), arg0)
}

// SearchFriends mocks base method
func (m *MockFriendUseCase) SearchFriends(arg0 int, arg1 string) ([]models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFriends", arg0, arg1)
	ret0, _ := ret[0].([]models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchFriends indicates an expected call of SearchFriends
func (mr *MockFriendUseCaseMockRecorder) SearchFriends(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFriends", reflect.TypeOf((*MockFriendUseCase)(nil).SearchFriends), arg0, arg1)
}