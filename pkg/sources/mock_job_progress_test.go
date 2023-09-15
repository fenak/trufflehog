// Code generated by MockGen. DO NOT EDIT.
// Source: ./job_progress.go

// Package sources is a generated GoMock package.
package sources

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockJobProgressHook is a mock of JobProgressHook interface.
type MockJobProgressHook struct {
	ctrl     *gomock.Controller
	recorder *MockJobProgressHookMockRecorder
}

// MockJobProgressHookMockRecorder is the mock recorder for MockJobProgressHook.
type MockJobProgressHookMockRecorder struct {
	mock *MockJobProgressHook
}

// NewMockJobProgressHook creates a new mock instance.
func NewMockJobProgressHook(ctrl *gomock.Controller) *MockJobProgressHook {
	mock := &MockJobProgressHook{ctrl: ctrl}
	mock.recorder = &MockJobProgressHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobProgressHook) EXPECT() *MockJobProgressHookMockRecorder {
	return m.recorder
}

// End mocks base method.
func (m *MockJobProgressHook) End(arg0 JobProgressRef, arg1 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "End", arg0, arg1)
}

// End indicates an expected call of End.
func (mr *MockJobProgressHookMockRecorder) End(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "End", reflect.TypeOf((*MockJobProgressHook)(nil).End), arg0, arg1)
}

// EndEnumerating mocks base method.
func (m *MockJobProgressHook) EndEnumerating(arg0 JobProgressRef, arg1 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EndEnumerating", arg0, arg1)
}

// EndEnumerating indicates an expected call of EndEnumerating.
func (mr *MockJobProgressHookMockRecorder) EndEnumerating(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndEnumerating", reflect.TypeOf((*MockJobProgressHook)(nil).EndEnumerating), arg0, arg1)
}

// EndUnitChunking mocks base method.
func (m *MockJobProgressHook) EndUnitChunking(arg0 JobProgressRef, arg1 SourceUnit, arg2 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EndUnitChunking", arg0, arg1, arg2)
}

// EndUnitChunking indicates an expected call of EndUnitChunking.
func (mr *MockJobProgressHookMockRecorder) EndUnitChunking(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndUnitChunking", reflect.TypeOf((*MockJobProgressHook)(nil).EndUnitChunking), arg0, arg1, arg2)
}

// Finish mocks base method.
func (m *MockJobProgressHook) Finish(arg0 JobProgressRef) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finish", arg0)
}

// Finish indicates an expected call of Finish.
func (mr *MockJobProgressHookMockRecorder) Finish(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockJobProgressHook)(nil).Finish), arg0)
}

// ReportChunk mocks base method.
func (m *MockJobProgressHook) ReportChunk(arg0 JobProgressRef, arg1 SourceUnit, arg2 *Chunk) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportChunk", arg0, arg1, arg2)
}

// ReportChunk indicates an expected call of ReportChunk.
func (mr *MockJobProgressHookMockRecorder) ReportChunk(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportChunk", reflect.TypeOf((*MockJobProgressHook)(nil).ReportChunk), arg0, arg1, arg2)
}

// ReportError mocks base method.
func (m *MockJobProgressHook) ReportError(arg0 JobProgressRef, arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportError", arg0, arg1)
}

// ReportError indicates an expected call of ReportError.
func (mr *MockJobProgressHookMockRecorder) ReportError(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportError", reflect.TypeOf((*MockJobProgressHook)(nil).ReportError), arg0, arg1)
}

// ReportUnit mocks base method.
func (m *MockJobProgressHook) ReportUnit(arg0 JobProgressRef, arg1 SourceUnit) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportUnit", arg0, arg1)
}

// ReportUnit indicates an expected call of ReportUnit.
func (mr *MockJobProgressHookMockRecorder) ReportUnit(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportUnit", reflect.TypeOf((*MockJobProgressHook)(nil).ReportUnit), arg0, arg1)
}

// Start mocks base method.
func (m *MockJobProgressHook) Start(arg0 JobProgressRef, arg1 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0, arg1)
}

// Start indicates an expected call of Start.
func (mr *MockJobProgressHookMockRecorder) Start(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockJobProgressHook)(nil).Start), arg0, arg1)
}

// StartEnumerating mocks base method.
func (m *MockJobProgressHook) StartEnumerating(arg0 JobProgressRef, arg1 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartEnumerating", arg0, arg1)
}

// StartEnumerating indicates an expected call of StartEnumerating.
func (mr *MockJobProgressHookMockRecorder) StartEnumerating(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartEnumerating", reflect.TypeOf((*MockJobProgressHook)(nil).StartEnumerating), arg0, arg1)
}

// StartUnitChunking mocks base method.
func (m *MockJobProgressHook) StartUnitChunking(arg0 JobProgressRef, arg1 SourceUnit, arg2 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartUnitChunking", arg0, arg1, arg2)
}

// StartUnitChunking indicates an expected call of StartUnitChunking.
func (mr *MockJobProgressHookMockRecorder) StartUnitChunking(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartUnitChunking", reflect.TypeOf((*MockJobProgressHook)(nil).StartUnitChunking), arg0, arg1, arg2)
}
