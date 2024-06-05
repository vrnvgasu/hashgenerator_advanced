// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/composer/hashcompose/api_grpc.pb.go

// Package hashcomposeMocks is a generated GoMock package.
package hashcomposeMocks

import (
	context "context"
	reflect "reflect"
	hashcompose "service1/pkg/composer/hashcompose"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockHashComposeServiceClient is a mock of HashComposeServiceClient interface.
type MockHashComposeServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockHashComposeServiceClientMockRecorder
}

// MockHashComposeServiceClientMockRecorder is the mock recorder for MockHashComposeServiceClient.
type MockHashComposeServiceClientMockRecorder struct {
	mock *MockHashComposeServiceClient
}

// NewMockHashComposeServiceClient creates a new mock instance.
func NewMockHashComposeServiceClient(ctrl *gomock.Controller) *MockHashComposeServiceClient {
	mock := &MockHashComposeServiceClient{ctrl: ctrl}
	mock.recorder = &MockHashComposeServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashComposeServiceClient) EXPECT() *MockHashComposeServiceClientMockRecorder {
	return m.recorder
}

// Gen mocks base method.
func (m *MockHashComposeServiceClient) Gen(ctx context.Context, in *hashcompose.StringList, opts ...grpc.CallOption) (hashcompose.HashComposeService_GenClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Gen", varargs...)
	ret0, _ := ret[0].(hashcompose.HashComposeService_GenClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Gen indicates an expected call of Gen.
func (mr *MockHashComposeServiceClientMockRecorder) Gen(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gen", reflect.TypeOf((*MockHashComposeServiceClient)(nil).Gen), varargs...)
}

// MockHashComposeService_GenClient is a mock of HashComposeService_GenClient interface.
type MockHashComposeService_GenClient struct {
	ctrl     *gomock.Controller
	recorder *MockHashComposeService_GenClientMockRecorder
}

// MockHashComposeService_GenClientMockRecorder is the mock recorder for MockHashComposeService_GenClient.
type MockHashComposeService_GenClientMockRecorder struct {
	mock *MockHashComposeService_GenClient
}

// NewMockHashComposeService_GenClient creates a new mock instance.
func NewMockHashComposeService_GenClient(ctrl *gomock.Controller) *MockHashComposeService_GenClient {
	mock := &MockHashComposeService_GenClient{ctrl: ctrl}
	mock.recorder = &MockHashComposeService_GenClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashComposeService_GenClient) EXPECT() *MockHashComposeService_GenClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockHashComposeService_GenClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockHashComposeService_GenClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockHashComposeService_GenClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockHashComposeService_GenClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockHashComposeService_GenClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockHashComposeService_GenClient)(nil).Context))
}

// Header mocks base method.
func (m *MockHashComposeService_GenClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockHashComposeService_GenClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockHashComposeService_GenClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockHashComposeService_GenClient) Recv() (*hashcompose.HashObj, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*hashcompose.HashObj)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockHashComposeService_GenClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockHashComposeService_GenClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockHashComposeService_GenClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockHashComposeService_GenClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockHashComposeService_GenClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockHashComposeService_GenClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockHashComposeService_GenClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockHashComposeService_GenClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockHashComposeService_GenClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockHashComposeService_GenClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockHashComposeService_GenClient)(nil).Trailer))
}

// MockHashComposeServiceServer is a mock of HashComposeServiceServer interface.
type MockHashComposeServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockHashComposeServiceServerMockRecorder
}

// MockHashComposeServiceServerMockRecorder is the mock recorder for MockHashComposeServiceServer.
type MockHashComposeServiceServerMockRecorder struct {
	mock *MockHashComposeServiceServer
}

// NewMockHashComposeServiceServer creates a new mock instance.
func NewMockHashComposeServiceServer(ctrl *gomock.Controller) *MockHashComposeServiceServer {
	mock := &MockHashComposeServiceServer{ctrl: ctrl}
	mock.recorder = &MockHashComposeServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashComposeServiceServer) EXPECT() *MockHashComposeServiceServerMockRecorder {
	return m.recorder
}

// Gen mocks base method.
func (m *MockHashComposeServiceServer) Gen(arg0 *hashcompose.StringList, arg1 hashcompose.HashComposeService_GenServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gen", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Gen indicates an expected call of Gen.
func (mr *MockHashComposeServiceServerMockRecorder) Gen(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gen", reflect.TypeOf((*MockHashComposeServiceServer)(nil).Gen), arg0, arg1)
}

// mustEmbedUnimplementedHashComposeServiceServer mocks base method.
func (m *MockHashComposeServiceServer) mustEmbedUnimplementedHashComposeServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedHashComposeServiceServer")
}

// mustEmbedUnimplementedHashComposeServiceServer indicates an expected call of mustEmbedUnimplementedHashComposeServiceServer.
func (mr *MockHashComposeServiceServerMockRecorder) mustEmbedUnimplementedHashComposeServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedHashComposeServiceServer", reflect.TypeOf((*MockHashComposeServiceServer)(nil).mustEmbedUnimplementedHashComposeServiceServer))
}

// MockUnsafeHashComposeServiceServer is a mock of UnsafeHashComposeServiceServer interface.
type MockUnsafeHashComposeServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeHashComposeServiceServerMockRecorder
}

// MockUnsafeHashComposeServiceServerMockRecorder is the mock recorder for MockUnsafeHashComposeServiceServer.
type MockUnsafeHashComposeServiceServerMockRecorder struct {
	mock *MockUnsafeHashComposeServiceServer
}

// NewMockUnsafeHashComposeServiceServer creates a new mock instance.
func NewMockUnsafeHashComposeServiceServer(ctrl *gomock.Controller) *MockUnsafeHashComposeServiceServer {
	mock := &MockUnsafeHashComposeServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeHashComposeServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeHashComposeServiceServer) EXPECT() *MockUnsafeHashComposeServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedHashComposeServiceServer mocks base method.
func (m *MockUnsafeHashComposeServiceServer) mustEmbedUnimplementedHashComposeServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedHashComposeServiceServer")
}

// mustEmbedUnimplementedHashComposeServiceServer indicates an expected call of mustEmbedUnimplementedHashComposeServiceServer.
func (mr *MockUnsafeHashComposeServiceServerMockRecorder) mustEmbedUnimplementedHashComposeServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedHashComposeServiceServer", reflect.TypeOf((*MockUnsafeHashComposeServiceServer)(nil).mustEmbedUnimplementedHashComposeServiceServer))
}

// MockHashComposeService_GenServer is a mock of HashComposeService_GenServer interface.
type MockHashComposeService_GenServer struct {
	ctrl     *gomock.Controller
	recorder *MockHashComposeService_GenServerMockRecorder
}

// MockHashComposeService_GenServerMockRecorder is the mock recorder for MockHashComposeService_GenServer.
type MockHashComposeService_GenServerMockRecorder struct {
	mock *MockHashComposeService_GenServer
}

// NewMockHashComposeService_GenServer creates a new mock instance.
func NewMockHashComposeService_GenServer(ctrl *gomock.Controller) *MockHashComposeService_GenServer {
	mock := &MockHashComposeService_GenServer{ctrl: ctrl}
	mock.recorder = &MockHashComposeService_GenServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashComposeService_GenServer) EXPECT() *MockHashComposeService_GenServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockHashComposeService_GenServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockHashComposeService_GenServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockHashComposeService_GenServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockHashComposeService_GenServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockHashComposeService_GenServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockHashComposeService_GenServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockHashComposeService_GenServer) Send(arg0 *hashcompose.HashObj) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockHashComposeService_GenServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockHashComposeService_GenServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockHashComposeService_GenServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockHashComposeService_GenServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockHashComposeService_GenServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockHashComposeService_GenServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockHashComposeService_GenServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockHashComposeService_GenServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockHashComposeService_GenServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockHashComposeService_GenServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockHashComposeService_GenServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockHashComposeService_GenServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockHashComposeService_GenServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockHashComposeService_GenServer)(nil).SetTrailer), arg0)
}
