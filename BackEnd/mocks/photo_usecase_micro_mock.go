// Code generated by MockGen. DO NOT EDIT.
// Source: main/internal/microservices/photos/delivery (interfaces: PhotoCheckerClient)

// Package mock_delivery is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	delivery "main/internal/microservices/photos/delivery"
	reflect "reflect"
)

// MockPhotoCheckerClient is a mock of PhotoCheckerClient interface
type MockPhotoCheckerClient struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoCheckerClientMockRecorder
}

// MockPhotoCheckerClientMockRecorder is the mock recorder for MockPhotoCheckerClient
type MockPhotoCheckerClientMockRecorder struct {
	mock *MockPhotoCheckerClient
}

// NewMockPhotoCheckerClient creates a new mock instance
func NewMockPhotoCheckerClient(ctrl *gomock.Controller) *MockPhotoCheckerClient {
	mock := &MockPhotoCheckerClient{ctrl: ctrl}
	mock.recorder = &MockPhotoCheckerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPhotoCheckerClient) EXPECT() *MockPhotoCheckerClientMockRecorder {
	return m.recorder
}

// GetPhotosFromAlbum mocks base method
func (m *MockPhotoCheckerClient) GetPhotosFromAlbum(arg0 context.Context, arg1 *delivery.AlbumId, arg2 ...grpc.CallOption) (*delivery.Photos, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPhotosFromAlbum", varargs...)
	ret0, _ := ret[0].(*delivery.Photos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotosFromAlbum indicates an expected call of GetPhotosFromAlbum
func (mr *MockPhotoCheckerClientMockRecorder) GetPhotosFromAlbum(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotosFromAlbum", reflect.TypeOf((*MockPhotoCheckerClient)(nil).GetPhotosFromAlbum), varargs...)
}

// UploadPhotoToAlbum mocks base method
func (m *MockPhotoCheckerClient) UploadPhotoToAlbum(arg0 context.Context, arg1 *delivery.PhotoInAlbum, arg2 ...grpc.CallOption) (*delivery.Dummy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadPhotoToAlbum", varargs...)
	ret0, _ := ret[0].(*delivery.Dummy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadPhotoToAlbum indicates an expected call of UploadPhotoToAlbum
func (mr *MockPhotoCheckerClientMockRecorder) UploadPhotoToAlbum(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPhotoToAlbum", reflect.TypeOf((*MockPhotoCheckerClient)(nil).UploadPhotoToAlbum), varargs...)
}