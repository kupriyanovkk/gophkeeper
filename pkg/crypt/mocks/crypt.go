// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/crypt/crypt.go

// Package mock_crypt is a generated GoMock package.
package mock_crypt

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCryptAbstract is a mock of CryptAbstract interface.
type MockCryptAbstract struct {
	ctrl     *gomock.Controller
	recorder *MockCryptAbstractMockRecorder
}

// MockCryptAbstractMockRecorder is the mock recorder for MockCryptAbstract.
type MockCryptAbstractMockRecorder struct {
	mock *MockCryptAbstract
}

// NewMockCryptAbstract creates a new mock instance.
func NewMockCryptAbstract(ctrl *gomock.Controller) *MockCryptAbstract {
	mock := &MockCryptAbstract{ctrl: ctrl}
	mock.recorder = &MockCryptAbstractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptAbstract) EXPECT() *MockCryptAbstractMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockCryptAbstract) Decode(sha string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", sha)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockCryptAbstractMockRecorder) Decode(sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockCryptAbstract)(nil).Decode), sha)
}

// Encode mocks base method.
func (m *MockCryptAbstract) Encode(payload string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", payload)
	ret0, _ := ret[0].(string)
	return ret0
}

// Encode indicates an expected call of Encode.
func (mr *MockCryptAbstractMockRecorder) Encode(payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockCryptAbstract)(nil).Encode), payload)
}
