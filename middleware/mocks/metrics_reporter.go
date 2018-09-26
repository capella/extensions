// Code generated by MockGen. DO NOT EDIT.
// Source: metrics_reporter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockMetricsReporter is a mock of MetricsReporter interface
type MockMetricsReporter struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsReporterMockRecorder
}

// MockMetricsReporterMockRecorder is the mock recorder for MockMetricsReporter
type MockMetricsReporterMockRecorder struct {
	mock *MockMetricsReporter
}

// NewMockMetricsReporter creates a new mock instance
func NewMockMetricsReporter(ctrl *gomock.Controller) *MockMetricsReporter {
	mock := &MockMetricsReporter{ctrl: ctrl}
	mock.recorder = &MockMetricsReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricsReporter) EXPECT() *MockMetricsReporterMockRecorder {
	return m.recorder
}

// Timing mocks base method
func (m *MockMetricsReporter) Timing(metric string, value time.Duration, tags ...string) error {
	varargs := []interface{}{metric, value}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Timing", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Timing indicates an expected call of Timing
func (mr *MockMetricsReporterMockRecorder) Timing(metric, value interface{}, tags ...interface{}) *gomock.Call {
	varargs := append([]interface{}{metric, value}, tags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timing", reflect.TypeOf((*MockMetricsReporter)(nil).Timing), varargs...)
}

// Gauge mocks base method
func (m *MockMetricsReporter) Gauge(metrics string, value float64, tags ...string) error {
	varargs := []interface{}{metrics, value}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Gauge", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Gauge indicates an expected call of Gauge
func (mr *MockMetricsReporterMockRecorder) Gauge(metrics, value interface{}, tags ...interface{}) *gomock.Call {
	varargs := append([]interface{}{metrics, value}, tags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockMetricsReporter)(nil).Gauge), varargs...)
}

// Increment mocks base method
func (m *MockMetricsReporter) Increment(metric string, tags ...string) error {
	varargs := []interface{}{metric}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Increment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Increment indicates an expected call of Increment
func (mr *MockMetricsReporterMockRecorder) Increment(metric interface{}, tags ...interface{}) *gomock.Call {
	varargs := append([]interface{}{metric}, tags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockMetricsReporter)(nil).Increment), varargs...)
}