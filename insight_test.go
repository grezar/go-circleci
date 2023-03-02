package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_insights_ListSummaryMetricsForWorkflows(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"

	mux.HandleFunc(fmt.Sprintf("/insights/%s/workflows", projectSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Circle-Token", client.token)
		testQuery(t, r, "page-token", "1")
		testQuery(t, r, "all-branches", "true")
		testQuery(t, r, "reporting-window", "last-90-days")
		fmt.Fprint(w, `{"items": [{"name": "build-and-test"}], "next_page_token": "2"}`)
	})

	ctx := context.Background()
	sml, err := client.Insights.ListSummaryMetricsForWorkflows(ctx, projectSlug, InsightsListSummaryMetricsOptions{
		PageToken:       String("1"),
		AllBranches:     Bool(true),
		ReportingWindow: ReportingWindow(Last90Days),
	})
	if err != nil {
		t.Errorf("Insights.ListSummaryMetricsForWorkflows got error: %v", err)
	}

	want := &SummaryMetricsList{
		Items: []*SummaryMetrics{
			{
				Name: "build-and-test",
			},
		},
		NextPageToken: "2",
	}

	if !cmp.Equal(sml, want) {
		t.Errorf("Insights.ListSummaryMetricsForWorkflows got %+v, want %+v", sml, want)
	}
}

func Test_insights_ListSummaryMetricsForWorkflowJobs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	workflosName := "workflow1"

	mux.HandleFunc(fmt.Sprintf("/insights/%s/workflows/%s/jobs", projectSlug, workflosName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Circle-Token", client.token)
		testQuery(t, r, "page-token", "1")
		testQuery(t, r, "all-branches", "true")
		testQuery(t, r, "reporting-window", "last-90-days")
		fmt.Fprint(w, `{"items": [{"name": "build-and-test"}], "next_page_token": "2"}`)
	})

	ctx := context.Background()
	sml, err := client.Insights.ListSummaryMetricsForWorkflowJobs(ctx, projectSlug, workflosName, InsightsListSummaryMetricsOptions{
		PageToken:       String("1"),
		AllBranches:     Bool(true),
		ReportingWindow: ReportingWindow(Last90Days),
	})
	if err != nil {
		t.Errorf("Insights.ListSummaryMetricsForWorkflowJobs got error: %v", err)
	}

	want := &SummaryMetricsList{
		Items: []*SummaryMetrics{
			{
				Name: "build-and-test",
			},
		},
		NextPageToken: "2",
	}

	if !cmp.Equal(sml, want) {
		t.Errorf("Insights.ListSummaryMetricsForWorkflowJobs got %+v, want %+v", sml, want)
	}
}

func Test_insights_GetTestMetricsForWorkflows(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	workflosName := "workflow1"

	mux.HandleFunc(fmt.Sprintf("/insights/%s/workflows/%s/test-metrics", projectSlug, workflosName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Circle-Token", client.token)
		testQuery(t, r, "all-branches", "true")
		fmt.Fprint(w, `{"average_test_count": 0, "most_failed_tests": [{"failed_runs": 0}]}`)
	})

	ctx := context.Background()
	wrl, err := client.Insights.GetTestMetricsForWorkflows(ctx, projectSlug, workflosName, InsightsGetTestMetricsOptions{
		AllBranches: Bool(true),
	})
	if err != nil {
		t.Errorf("Insights.GetTestMetricsForWorkflows got error: %v", err)
	}

	want := &TestMetrics{
		AverageTestCount: 0,
		MostFailedTests: []*MostFailedTest{
			{
				FailedRuns: 0,
			},
		},
	}

	if !cmp.Equal(wrl, want) {
		t.Errorf("Insights.GetTestMetricsForWorkflows got %+v, want %+v", wrl, want)
	}
}

func Test_insights_ListWorkflowRuns(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	workflosName := "workflow1"

	mux.HandleFunc(fmt.Sprintf("/insights/%s/workflows/%s", projectSlug, workflosName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Circle-Token", client.token)
		testQuery(t, r, "page-token", "1")
		testQuery(t, r, "all-branches", "true")
		testQuery(t, r, "start-date", "2020-08-21T13:26:29Z")
		testQuery(t, r, "end-date", "2020-09-04T13:26:29Z")
		fmt.Fprint(w, `{"items": [{"id": "1", "credits_used": 0}], "next_page_token": "2"}`)
	})

	ctx := context.Background()
	startDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-08-21T13:26:29Z")
	endDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-09-04T13:26:29Z")
	wrl, err := client.Insights.ListWorkflowRuns(ctx, projectSlug, workflosName, InsightsListWorkflowRunsOptions{
		PageToken:   String("1"),
		AllBranches: Bool(true),
		StartDate:   Time(startDate),
		EndDate:     Time(endDate),
	})
	if err != nil {
		t.Errorf("Insights.ListWorkflowRuns got error: %v", err)
	}

	want := &WorkflowRunList{
		Items: []*WorkflowRun{
			{
				ID:          "1",
				CreditsUsed: 0,
			},
		},
		NextPageToken: "2",
	}

	if !cmp.Equal(wrl, want) {
		t.Errorf("Insights.ListWorkflowRuns got %+v, want %+v", wrl, want)
	}
}

func Test_insights_ListWorkflowJobRuns(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	workflosName := "workflow1"
	jobName := "job1"

	mux.HandleFunc(fmt.Sprintf("/insights/%s/workflows/%s/jobs/%s", projectSlug, workflosName, jobName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Circle-Token", client.token)
		testQuery(t, r, "page-token", "1")
		testQuery(t, r, "all-branches", "true")
		testQuery(t, r, "start-date", "2020-08-21T13:26:29Z")
		testQuery(t, r, "end-date", "2020-09-04T13:26:29Z")
		fmt.Fprint(w, `{"items": [{"id": "1", "credits_used": 0}], "next_page_token": "2"}`)
	})

	ctx := context.Background()
	startDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-08-21T13:26:29Z")
	endDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-09-04T13:26:29Z")
	wrl, err := client.Insights.ListWorkflowJobRuns(ctx, projectSlug, workflosName, jobName, InsightsListWorkflowRunsOptions{
		PageToken:   String("1"),
		AllBranches: Bool(true),
		StartDate:   Time(startDate),
		EndDate:     Time(endDate),
	})
	if err != nil {
		t.Errorf("Insights.ListWorkflowJobRuns got error: %v", err)
	}

	want := &WorkflowRunList{
		Items: []*WorkflowRun{
			{
				ID:          "1",
				CreditsUsed: 0,
			},
		},
		NextPageToken: "2",
	}

	if !cmp.Equal(wrl, want) {
		t.Errorf("Insights.ListWorkflowJobRuns got %+v, want %+v", wrl, want)
	}
}
