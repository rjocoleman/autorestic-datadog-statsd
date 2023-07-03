// Code generated by MockGen. DO NOT EDIT.
// Source: statsd.go

// Package mock_statsd is a generated GoMock package.
package mock_statsd

import (
	reflect "reflect"
	time "time"

	statsd "github.com/DataDog/datadog-go/v5/statsd"
	gomock "go.uber.org/mock/gomock"
)

// MockClientInterface is a mock of ClientInterface interface.
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface.
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance.
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClientInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientInterface)(nil).Close))
}

// Count mocks base method.
func (m *MockClientInterface) Count(name string, value int64, tags []string, rate float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockClientInterfaceMockRecorder) Count(name, value, tags, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockClientInterface)(nil).Count), name, value, tags, rate)
}

// CountWithTimestamp mocks base method.
func (m *MockClientInterface) CountWithTimestamp(name string, value int64, tags []string, rate float64, timestamp time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountWithTimestamp", name, value, tags, rate, timestamp)
	ret0, _ := ret[0].(error)
	return ret0
}

// CountWithTimestamp indicates an expected call of CountWithTimestamp.
func (mr *MockClientInterfaceMockRecorder) CountWithTimestamp(name, value, tags, rate, timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountWithTimestamp", reflect.TypeOf((*MockClientInterface)(nil).CountWithTimestamp), name, value, tags, rate, timestamp)
}

// Decr mocks base method.
func (m *MockClientInterface) Decr(name string, tags []string, rate float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decr", name, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decr indicates an expected call of Decr.
func (mr *MockClientInterfaceMockRecorder) Decr(name, tags, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockClientInterface)(nil).Decr), name, tags, rate)
}

// Distribution mocks base method.
func (m *MockClientInterface) Distribution(name string, value float64, tags []string, rate float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Distribution", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Distribution indicates an expected call of Distribution.
func (mr *MockClientInterfaceMockRecorder) Distribution(name, value, tags, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Distribution", reflect.TypeOf((*MockClientInterface)(nil).Distribution), name, value, tags, rate)
}

// Event mocks base method.
func (m *MockClientInterface) Event(e *statsd.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Event", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Event indicates an expected call of Event.
func (mr *MockClientInterfaceMockRecorder) Event(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Event", reflect.TypeOf((*MockClientInterface)(nil).Event), e)
}

// Flush mocks base method.
func (m *MockClientInterface) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockClientInterfaceMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockClientInterface)(nil).Flush))
}

// Gauge mocks base method.
func (m *MockClientInterface) Gauge(name string, value float64, tags []string, rate float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gauge", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Gauge indicates an expected call of Gauge.
func (mr *MockClientInterfaceMockRecorder) Gauge(name, value, tags, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockClientInterface)(nil).Gauge), name, value, tags, rate)
}

// GaugeWithTimestamp mocks base method.
func (m *MockClientInterface) GaugeWithTimestamp(name string, value float64, tags []string, rate float64, timestamp time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GaugeWithTimestamp", name, value, tags, rate, timestamp)
	ret0, _ := ret[0].(error)
	return ret0
}

// GaugeWithTimestamp indicates an expected call of GaugeWithTimestamp.
func (mr *MockClientInterfaceMockRecorder) GaugeWithTimestamp(name, value, tags, rate, timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GaugeWithTimestamp", reflect.TypeOf((*MockClientInterface)(nil).GaugeWithTimestamp), name, value, tags, rate, timestamp)
}

// GetTelemetry mocks base method.
func (m *MockClientInterface) GetTelemetry() statsd.Telemetry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTelemetry")
	ret0, _ := ret[0].(statsd.Telemetry)
	return ret0
}

// GetTelemetry indicates an expected call of GetTelemetry.
func (mr *MockClientInterfaceMockRecorder) GetTelemetry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTelemetry", reflect.TypeOf((*MockClientInterface)(nil).GetTelemetry))
}

// Histogram mocks base method.
func (m *MockClientInterface) Histogram(name string, value float64, tags []string, rate float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Histogram", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Histogram indicates an expected call of Histogram.
func (mr *MockClientInterfaceMockRecorder) Histogram(name, value, tags, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Histogram", reflect.TypeOf((*MockClientInterface)(nil).Histogram), name, value, tags, rate)
}

// Incr mocks base method.
func (m *MockClientInterface) Incr(name string, tags []string, rate float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", name, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Incr indicates an expected call of Incr.
func (mr *MockClientInterfaceMockRecorder) Incr(name, tags, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockClientInterface)(nil).Incr), name, tags, rate)
}

// IsClosed mocks base method.
func (m *MockClientInterface) IsClosed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClosed indicates an expected call of IsClosed.
func (mr *MockClientInterfaceMockRecorder) IsClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosed", reflect.TypeOf((*MockClientInterface)(nil).IsClosed))
}

// ServiceCheck mocks base method.
func (m *MockClientInterface) ServiceCheck(sc *statsd.ServiceCheck) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceCheck", sc)
	ret0, _ := ret[0].(error)
	return ret0
}

// ServiceCheck indicates an expected call of ServiceCheck.
func (mr *MockClientInterfaceMockRecorder) ServiceCheck(sc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceCheck", reflect.TypeOf((*MockClientInterface)(nil).ServiceCheck), sc)
}

// Set mocks base method.
func (m *MockClientInterface) Set(name, value string, tags []string, rate float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockClientInterfaceMockRecorder) Set(name, value, tags, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockClientInterface)(nil).Set), name, value, tags, rate)
}

// SimpleEvent mocks base method.
func (m *MockClientInterface) SimpleEvent(title, text string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimpleEvent", title, text)
	ret0, _ := ret[0].(error)
	return ret0
}

// SimpleEvent indicates an expected call of SimpleEvent.
func (mr *MockClientInterfaceMockRecorder) SimpleEvent(title, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimpleEvent", reflect.TypeOf((*MockClientInterface)(nil).SimpleEvent), title, text)
}

// SimpleServiceCheck mocks base method.
func (m *MockClientInterface) SimpleServiceCheck(name string, status statsd.ServiceCheckStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimpleServiceCheck", name, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// SimpleServiceCheck indicates an expected call of SimpleServiceCheck.
func (mr *MockClientInterfaceMockRecorder) SimpleServiceCheck(name, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimpleServiceCheck", reflect.TypeOf((*MockClientInterface)(nil).SimpleServiceCheck), name, status)
}

// TimeInMilliseconds mocks base method.
func (m *MockClientInterface) TimeInMilliseconds(name string, value float64, tags []string, rate float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeInMilliseconds", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// TimeInMilliseconds indicates an expected call of TimeInMilliseconds.
func (mr *MockClientInterfaceMockRecorder) TimeInMilliseconds(name, value, tags, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeInMilliseconds", reflect.TypeOf((*MockClientInterface)(nil).TimeInMilliseconds), name, value, tags, rate)
}

// Timing mocks base method.
func (m *MockClientInterface) Timing(name string, value time.Duration, tags []string, rate float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timing", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Timing indicates an expected call of Timing.
func (mr *MockClientInterfaceMockRecorder) Timing(name, value, tags, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timing", reflect.TypeOf((*MockClientInterface)(nil).Timing), name, value, tags, rate)
}
