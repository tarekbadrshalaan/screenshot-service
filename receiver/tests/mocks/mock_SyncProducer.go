// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Shopify/sarama (interfaces: SyncProducer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	"sync"

	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"
)

// MockSyncProducer is a mock of SyncProducer interface
type MockSyncProducer struct {
	ctrl     *gomock.Controller
	recorder *MockSyncProducerMockRecorder
	wg       *sync.WaitGroup
}

// MockSyncProducerMockRecorder is the mock recorder for MockSyncProducer
type MockSyncProducerMockRecorder struct {
	mock *MockSyncProducer
}

// NewMockSyncProducer creates a new mock instance
func NewMockSyncProducer(ctrl *gomock.Controller, wg *sync.WaitGroup) *MockSyncProducer {
	mock := &MockSyncProducer{ctrl: ctrl, wg: wg}
	mock.recorder = &MockSyncProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSyncProducer) EXPECT() *MockSyncProducerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSyncProducer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSyncProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSyncProducer)(nil).Close))
}

// SendMessage mocks base method
func (m *MockSyncProducer) SendMessage(arg0 *sarama.ProducerMessage) (int32, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	m.wg.Done()
	return ret0, ret1, ret2
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockSyncProducerMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockSyncProducer)(nil).SendMessage), arg0)
}

// SendMessages mocks base method
func (m *MockSyncProducer) SendMessages(arg0 []*sarama.ProducerMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessages", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessages indicates an expected call of SendMessages
func (mr *MockSyncProducerMockRecorder) SendMessages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessages", reflect.TypeOf((*MockSyncProducer)(nil).SendMessages), arg0)
}
