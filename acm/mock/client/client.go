// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/awslabs/fargatecli/acm (interfaces: Client)

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"

	acm "github.com/awslabs/fargatecli/acm"
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

// DeleteCertificate mocks base method.
func (m *MockClient) DeleteCertificate(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCertificate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCertificate indicates an expected call of DeleteCertificate.
func (mr *MockClientMockRecorder) DeleteCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCertificate", reflect.TypeOf((*MockClient)(nil).DeleteCertificate), arg0)
}

// ImportCertificate mocks base method.
func (m *MockClient) ImportCertificate(arg0, arg1, arg2 []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportCertificate", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportCertificate indicates an expected call of ImportCertificate.
func (mr *MockClientMockRecorder) ImportCertificate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportCertificate", reflect.TypeOf((*MockClient)(nil).ImportCertificate), arg0, arg1, arg2)
}

// InflateCertificate mocks base method.
func (m *MockClient) InflateCertificate(arg0 *acm.Certificate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InflateCertificate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InflateCertificate indicates an expected call of InflateCertificate.
func (mr *MockClientMockRecorder) InflateCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InflateCertificate", reflect.TypeOf((*MockClient)(nil).InflateCertificate), arg0)
}

// ListCertificates mocks base method.
func (m *MockClient) ListCertificates() (acm.Certificates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCertificates")
	ret0, _ := ret[0].(acm.Certificates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCertificates indicates an expected call of ListCertificates.
func (mr *MockClientMockRecorder) ListCertificates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificates", reflect.TypeOf((*MockClient)(nil).ListCertificates))
}

// RequestCertificate mocks base method.
func (m *MockClient) RequestCertificate(arg0 string, arg1 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestCertificate", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestCertificate indicates an expected call of RequestCertificate.
func (mr *MockClientMockRecorder) RequestCertificate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCertificate", reflect.TypeOf((*MockClient)(nil).RequestCertificate), arg0, arg1)
}
