package circleci

import (
	"context"
	"fmt"
	"time"
)

type Insights interface {
	ListSummaryMetrics(ctx context.Context, projectSlug string, options InsightsListSummaryMetricsOptions) (*SummaryMetricsList, error)
	ListWorkflowRuns(ctx context.Context, projectSlug, workflowName string, options InsightsListWorkflowRunsOptions) (*WorkflowRunList, error)
}

// insights implementes Insights interface
type insights struct {
	client *Client
}

type reportingWindow string

const (
	Last7Days   reportingWindow = "last-7-days"
	Last30Days  reportingWindow = "last-30-days"
	Last60Days  reportingWindow = "last-60-days"
	Last90Days  reportingWindow = "last-90-days"
	Last24Hours reportingWindow = "last-24-hours"
)

type SummaryMetricsList struct {
	Items         []*SummaryMetrics `json:"items"`
	NextPageToken string            `json:"next_page_token"`
}

type SummaryMetrics struct {
	Name        string    `json:"name"`
	WindowStart time.Time `json:"window_start"`
	WindowEnd   time.Time `json:"window_end"`
	Metrics     *Metrics  `json:"metrics"`
}

type Metrics struct {
	TotalRuns        int              `json:"total_runs"`
	SuccessfulRuns   int              `json:"successful_runs"`
	Mttr             int              `json:"mttr"`
	TotalCreditsUsed int              `json:"total_credits_used"`
	FailedRuns       int              `json:"failed_runs"`
	SuccessRate      int              `json:"success_rate"`
	DurationMetrics  *DurationMetrics `json:"duration_metrics"`
	TotalRecoveries  int              `json:"total_recoveries"`
	Throughput       int              `json:"throughput"`
}

type DurationMetrics struct {
	Min               int `json:"min"`
	Mean              int `json:"mean"`
	Median            int `json:"median"`
	P95               int `json:"p95"`
	Max               int `json:"max"`
	StandardDeviation int `json:"standard_deviation"`
}

type InsightsListSummaryMetricsOptions struct {
	ReportingWindow *reportingWindow `url:"reporting-window,omitempty"`
	AllBranches     *bool            `url:"all-branches,omitempty"`
	Branch          *string          `url:"branch,omitempty"`
	PageToken       *string          `url:"page-token,omitempty"`
}

func (o InsightsListSummaryMetricsOptions) valid() error {
	// Nothing is required
	return nil
}

func (s *insights) ListSummaryMetrics(ctx context.Context, projectSlug string, options InsightsListSummaryMetricsOptions) (*SummaryMetricsList, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	u := fmt.Sprintf("insights/%s/workflows", projectSlug)
	req, err := s.client.newRequest("GET", u, &options)
	if err != nil {
		return nil, err
	}

	sml := &SummaryMetricsList{}
	err = s.client.do(ctx, req, sml)
	if err != nil {
		return nil, err
	}

	return sml, nil
}

type WorkflowRunList struct {
	Items         []*WorkflowRun `json:"items"`
	NextPageToken string         `json:"next_page_token"`
}

type WorkflowRun struct {
	ID          string    `json:"id"`
	Branch      string    `json:"branch"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
	StoppedAt   time.Time `json:"stopped_at"`
	CreditsUsed int       `json:"credits_used"`
	Status      string    `json:"status"`
}

type InsightsListWorkflowRunsOptions struct {
	AllBranches *bool      `url:"all-branches,omitempty"`
	Branch      *string    `url:"branch,omitempty"`
	StartDate   *time.Time `url:"start-date,omitempty"`
	EndDate     *time.Time `url:"end-date,omitempty"`
	PageToken   *string    `url:"page-token,omitempty"`
}

func (o InsightsListWorkflowRunsOptions) valid() error {
	// Nothing is required
	return nil
}

func (s *insights) ListWorkflowRuns(ctx context.Context, projectSlug, workflowName string, options InsightsListWorkflowRunsOptions) (*WorkflowRunList, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	if !validString(&workflowName) {
		return nil, ErrRequiredWorkflowName
	}

	u := fmt.Sprintf("insights/%s/workflows/%s", projectSlug, workflowName)
	req, err := s.client.newRequest("GET", u, &options)
	if err != nil {
		return nil, err
	}

	wil := &WorkflowRunList{}
	err = s.client.do(ctx, req, wil)
	if err != nil {
		return nil, err
	}

	return wil, nil
}
