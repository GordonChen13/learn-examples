// Code generated by MockGen. DO NOT EDIT.
// Source: ../match.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	models "github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockmatch is a mock of match interface
type Mockmatch struct {
	ctrl     *gomock.Controller
	recorder *MockmatchMockRecorder
}

// MockmatchMockRecorder is the mock recorder for Mockmatch
type MockmatchMockRecorder struct {
	mock *Mockmatch
}

// NewMockmatch creates a new mock instance
func NewMockmatch(ctrl *gomock.Controller) *Mockmatch {
	mock := &Mockmatch{ctrl: ctrl}
	mock.recorder = &MockmatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockmatch) EXPECT() *MockmatchMockRecorder {
	return m.recorder
}

// create mocks base method
func (m *Mockmatch) create(ctx context.Context, match *models.Match) (*models.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "create", ctx, match)
	ret0, _ := ret[0].(*models.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// create indicates an expected call of create
func (mr *MockmatchMockRecorder) create(ctx, match interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "create", reflect.TypeOf((*Mockmatch)(nil).create), ctx, match)
}
