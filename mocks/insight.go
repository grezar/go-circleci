// Code generated by MockGen. DO NOT EDIT.
// Source: insight.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	circleci "github.com/grezar/go-circleci"
)

// MockInsights is a mock of Insights interface.
type MockInsights struct {
	ctrl     *gomock.Controller
	recorder *MockInsightsMockRecorder
}

// MockInsightsMockRecorder is the mock recorder for MockInsights.
type MockInsightsMockRecorder struct {
	mock *MockInsights
}

// NewMockInsights creates a new mock instance.
func NewMockInsights(ctrl *gomock.Controller) *MockInsights {
	mock := &MockInsights{ctrl: ctrl}
	mock.recorder = &MockInsightsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInsights) EXPECT() *MockInsightsMockRecorder {
	return m.recorder
}

// GetTestMetricsForWorkflows mocks base method.
func (m *MockInsights) GetTestMetricsForWorkflows(ctx context.Context, projectSlug, workflowName string, options circleci.InsightsGetTestMetricsOptions) (*circleci.TestMetrics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestMetricsForWorkflows", ctx, projectSlug, workflowName, options)
	ret0, _ := ret[0].(*circleci.TestMetrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestMetricsForWorkflows indicates an expected call of GetTestMetricsForWorkflows.
func (mr *MockInsightsMockRecorder) GetTestMetricsForWorkflows(ctx, projectSlug, workflowName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestMetricsForWorkflows", reflect.TypeOf((*MockInsights)(nil).GetTestMetricsForWorkflows), ctx, projectSlug, workflowName, options)
}

// ListSummaryMetricsForWorkflows mocks base method.
func (m *MockInsights) ListSummaryMetricsForWorkflows(ctx context.Context, projectSlug string, options circleci.InsightsListSummaryMetricsOptions) (*circleci.SummaryMetricsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSummaryMetricsForWorkflows", ctx, projectSlug, options)
	ret0, _ := ret[0].(*circleci.SummaryMetricsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSummaryMetricsForWorkflows indicates an expected call of ListSummaryMetricsForWorkflows.
func (mr *MockInsightsMockRecorder) ListSummaryMetricsForWorkflows(ctx, projectSlug, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSummaryMetricsForWorkflows", reflect.TypeOf((*MockInsights)(nil).ListSummaryMetricsForWorkflows), ctx, projectSlug, options)
}

// ListSummaryMetricsForWorkflowJobs mocks base method.
func (m *MockInsights) ListSummaryMetricsForWorkflowJobs(ctx context.Context, projectSlug, workflowName string, options circleci.InsightsListSummaryMetricsOptions) (*circleci.SummaryMetricsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSummaryMetricsForWorkflowJobs", ctx, projectSlug, workflowName, options)
	ret0, _ := ret[0].(*circleci.SummaryMetricsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSummaryMetricsForWorkflowJobs indicates an expected call of ListSummaryMetricsForWorkflowJobs.
func (mr *MockInsightsMockRecorder) ListSummaryMetricsForWorkflowJobs(ctx, projectSlug, workflowName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSummaryMetricsForWorkflowJobs", reflect.TypeOf((*MockInsights)(nil).ListSummaryMetricsForWorkflowJobs), ctx, projectSlug, workflowName, options)
}

// ListWorkflowJobRuns mocks base method.
func (m *MockInsights) ListWorkflowJobRuns(ctx context.Context, projectSlug, workflowName, jobName string, options circleci.InsightsListWorkflowRunsOptions) (*circleci.WorkflowRunList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkflowJobRuns", ctx, projectSlug, workflowName, jobName, options)
	ret0, _ := ret[0].(*circleci.WorkflowRunList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkflowJobRuns indicates an expected call of ListWorkflowJobRuns.
func (mr *MockInsightsMockRecorder) ListWorkflowJobRuns(ctx, projectSlug, workflowName, jobName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflowJobRuns", reflect.TypeOf((*MockInsights)(nil).ListWorkflowJobRuns), ctx, projectSlug, workflowName, jobName, options)
}

// ListWorkflowRuns mocks base method.
func (m *MockInsights) ListWorkflowRuns(ctx context.Context, projectSlug, workflowName string, options circleci.InsightsListWorkflowRunsOptions) (*circleci.WorkflowRunList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkflowRuns", ctx, projectSlug, workflowName, options)
	ret0, _ := ret[0].(*circleci.WorkflowRunList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkflowRuns indicates an expected call of ListWorkflowRuns.
func (mr *MockInsightsMockRecorder) ListWorkflowRuns(ctx, projectSlug, workflowName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflowRuns", reflect.TypeOf((*MockInsights)(nil).ListWorkflowRuns), ctx, projectSlug, workflowName, options)
}
