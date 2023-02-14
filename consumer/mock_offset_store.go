// Code generated by MockGen. DO NOT EDIT.
// Source: offset_store.go

// Package consumer is a generated GoMock package.
package consumer

import (
	reflect "reflect"

	primitive "github.com/A-Zee029/rocketmq-client-go/v2/primitive"
	gomock "github.com/golang/mock/gomock"
)

// MockOffsetStore is a mock of OffsetStore interface.
type MockOffsetStore struct {
	ctrl     *gomock.Controller
	recorder *MockOffsetStoreMockRecorder
}

// MockOffsetStoreMockRecorder is the mock recorder for MockOffsetStore.
type MockOffsetStoreMockRecorder struct {
	mock *MockOffsetStore
}

// NewMockOffsetStore creates a new mock instance.
func NewMockOffsetStore(ctrl *gomock.Controller) *MockOffsetStore {
	mock := &MockOffsetStore{ctrl: ctrl}
	mock.recorder = &MockOffsetStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOffsetStore) EXPECT() *MockOffsetStoreMockRecorder {
	return m.recorder
}

// persist mocks base method.
func (m *MockOffsetStore) persist(mqs []*primitive.MessageQueue) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "persist", mqs)
}

// persist indicates an expected call of persist.
func (mr *MockOffsetStoreMockRecorder) persist(mqs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "persist", reflect.TypeOf((*MockOffsetStore)(nil).persist), mqs)
}

// read mocks base method.
func (m *MockOffsetStore) read(mq *primitive.MessageQueue, t readType) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "read", mq, t)
	ret0, _ := ret[0].(int64)
	return ret0
}

// read indicates an expected call of read.
func (mr *MockOffsetStoreMockRecorder) read(mq, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "read", reflect.TypeOf((*MockOffsetStore)(nil).read), mq, t)
}

// readWithException mocks base method.
func (m *MockOffsetStore) readWithException(mq *primitive.MessageQueue, t readType) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "readWithException", mq, t)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// readWithException indicates an expected call of readWithException.
func (mr *MockOffsetStoreMockRecorder) readWithException(mq, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "readWithException", reflect.TypeOf((*MockOffsetStore)(nil).readWithException), mq, t)
}

// remove mocks base method.
func (m *MockOffsetStore) remove(mq *primitive.MessageQueue) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "remove", mq)
}

// remove indicates an expected call of remove.
func (mr *MockOffsetStoreMockRecorder) remove(mq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "remove", reflect.TypeOf((*MockOffsetStore)(nil).remove), mq)
}

// update mocks base method.
func (m *MockOffsetStore) update(mq *primitive.MessageQueue, offset int64, increaseOnly bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "update", mq, offset, increaseOnly)
}

// update indicates an expected call of update.
func (mr *MockOffsetStoreMockRecorder) update(mq, offset, increaseOnly interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "update", reflect.TypeOf((*MockOffsetStore)(nil).update), mq, offset, increaseOnly)
}
