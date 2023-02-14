// Code generated by MockGen. DO NOT EDIT.
// Source: namesrv.go

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	primitive "github.com/A-Zee029/rocketmq-client-go/v2/primitive"
	gomock "github.com/golang/mock/gomock"
)

// MockNamesrvs is a mock of Namesrvs interface.
type MockNamesrvs struct {
	ctrl     *gomock.Controller
	recorder *MockNamesrvsMockRecorder
}

// MockNamesrvsMockRecorder is the mock recorder for MockNamesrvs.
type MockNamesrvsMockRecorder struct {
	mock *MockNamesrvs
}

// NewMockNamesrvs creates a new mock instance.
func NewMockNamesrvs(ctrl *gomock.Controller) *MockNamesrvs {
	mock := &MockNamesrvs{ctrl: ctrl}
	mock.recorder = &MockNamesrvsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamesrvs) EXPECT() *MockNamesrvsMockRecorder {
	return m.recorder
}

// AddBroker mocks base method.
func (m *MockNamesrvs) AddBroker(routeData *TopicRouteData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBroker", routeData)
}

// AddBroker indicates an expected call of AddBroker.
func (mr *MockNamesrvsMockRecorder) AddBroker(routeData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBroker", reflect.TypeOf((*MockNamesrvs)(nil).AddBroker), routeData)
}

// AddrList mocks base method.
func (m *MockNamesrvs) AddrList() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrList")
	ret0, _ := ret[0].([]string)
	return ret0
}

// AddrList indicates an expected call of AddrList.
func (mr *MockNamesrvsMockRecorder) AddrList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrList", reflect.TypeOf((*MockNamesrvs)(nil).AddrList))
}

// FetchAllTopicList mocks base method.
func (m *MockNamesrvs) FetchAllTopicList(ctx context.Context) (*TopicList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllTopicList", ctx)
	ret0, _ := ret[0].(*TopicList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllTopicList indicates an expected call of FetchAllTopicList.
func (mr *MockNamesrvsMockRecorder) FetchAllTopicList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllTopicList", reflect.TypeOf((*MockNamesrvs)(nil).FetchAllTopicList), ctx)
}

// FetchPublishMessageQueues mocks base method.
func (m *MockNamesrvs) FetchPublishMessageQueues(topic string) ([]*primitive.MessageQueue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPublishMessageQueues", topic)
	ret0, _ := ret[0].([]*primitive.MessageQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPublishMessageQueues indicates an expected call of FetchPublishMessageQueues.
func (mr *MockNamesrvsMockRecorder) FetchPublishMessageQueues(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPublishMessageQueues", reflect.TypeOf((*MockNamesrvs)(nil).FetchPublishMessageQueues), topic)
}

// FetchSubscribeMessageQueues mocks base method.
func (m *MockNamesrvs) FetchSubscribeMessageQueues(topic string) ([]*primitive.MessageQueue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSubscribeMessageQueues", topic)
	ret0, _ := ret[0].([]*primitive.MessageQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSubscribeMessageQueues indicates an expected call of FetchSubscribeMessageQueues.
func (mr *MockNamesrvsMockRecorder) FetchSubscribeMessageQueues(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSubscribeMessageQueues", reflect.TypeOf((*MockNamesrvs)(nil).FetchSubscribeMessageQueues), topic)
}

// FindBrokerAddrByName mocks base method.
func (m *MockNamesrvs) FindBrokerAddrByName(brokerName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBrokerAddrByName", brokerName)
	ret0, _ := ret[0].(string)
	return ret0
}

// FindBrokerAddrByName indicates an expected call of FindBrokerAddrByName.
func (mr *MockNamesrvsMockRecorder) FindBrokerAddrByName(brokerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBrokerAddrByName", reflect.TypeOf((*MockNamesrvs)(nil).FindBrokerAddrByName), brokerName)
}

// FindBrokerAddrByTopic mocks base method.
func (m *MockNamesrvs) FindBrokerAddrByTopic(topic string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBrokerAddrByTopic", topic)
	ret0, _ := ret[0].(string)
	return ret0
}

// FindBrokerAddrByTopic indicates an expected call of FindBrokerAddrByTopic.
func (mr *MockNamesrvsMockRecorder) FindBrokerAddrByTopic(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBrokerAddrByTopic", reflect.TypeOf((*MockNamesrvs)(nil).FindBrokerAddrByTopic), topic)
}

// FindBrokerAddressInSubscribe mocks base method.
func (m *MockNamesrvs) FindBrokerAddressInSubscribe(brokerName string, brokerId int64, onlyThisBroker bool) *FindBrokerResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBrokerAddressInSubscribe", brokerName, brokerId, onlyThisBroker)
	ret0, _ := ret[0].(*FindBrokerResult)
	return ret0
}

// FindBrokerAddressInSubscribe indicates an expected call of FindBrokerAddressInSubscribe.
func (mr *MockNamesrvsMockRecorder) FindBrokerAddressInSubscribe(brokerName, brokerId, onlyThisBroker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBrokerAddressInSubscribe", reflect.TypeOf((*MockNamesrvs)(nil).FindBrokerAddressInSubscribe), brokerName, brokerId, onlyThisBroker)
}

// UpdateNameServerAddress mocks base method.
func (m *MockNamesrvs) UpdateNameServerAddress() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateNameServerAddress")
}

// UpdateNameServerAddress indicates an expected call of UpdateNameServerAddress.
func (mr *MockNamesrvsMockRecorder) UpdateNameServerAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNameServerAddress", reflect.TypeOf((*MockNamesrvs)(nil).UpdateNameServerAddress))
}

// UpdateTopicRouteInfo mocks base method.
func (m *MockNamesrvs) UpdateTopicRouteInfo(topic string) (*TopicRouteData, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTopicRouteInfo", topic)
	ret0, _ := ret[0].(*TopicRouteData)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateTopicRouteInfo indicates an expected call of UpdateTopicRouteInfo.
func (mr *MockNamesrvsMockRecorder) UpdateTopicRouteInfo(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopicRouteInfo", reflect.TypeOf((*MockNamesrvs)(nil).UpdateTopicRouteInfo), topic)
}

// UpdateTopicRouteInfoWithDefault mocks base method.
func (m *MockNamesrvs) UpdateTopicRouteInfoWithDefault(topic, defaultTopic string, defaultQueueNum int) (*TopicRouteData, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTopicRouteInfoWithDefault", topic, defaultTopic, defaultQueueNum)
	ret0, _ := ret[0].(*TopicRouteData)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateTopicRouteInfoWithDefault indicates an expected call of UpdateTopicRouteInfoWithDefault.
func (mr *MockNamesrvsMockRecorder) UpdateTopicRouteInfoWithDefault(topic, defaultTopic, defaultQueueNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopicRouteInfoWithDefault", reflect.TypeOf((*MockNamesrvs)(nil).UpdateTopicRouteInfoWithDefault), topic, defaultTopic, defaultQueueNum)
}

// cleanOfflineBroker mocks base method.
func (m *MockNamesrvs) cleanOfflineBroker() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "cleanOfflineBroker")
}

// cleanOfflineBroker indicates an expected call of cleanOfflineBroker.
func (mr *MockNamesrvsMockRecorder) cleanOfflineBroker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cleanOfflineBroker", reflect.TypeOf((*MockNamesrvs)(nil).cleanOfflineBroker))
}
