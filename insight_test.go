package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_insights_ListSummaryMetrics(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"

	mux.HandleFunc(fmt.Sprintf("/insights/%s/workflows", projectSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		testQuery(t, r, "page-token", "1")
		testQuery(t, r, "all-branches", "true")
		testQuery(t, r, "reporting-window", "last-90-days")
		fmt.Fprint(w, `{"items": [{"name": "build-and-test"}], "next_page_token": "2"}`)
	})

	ctx := context.Background()
	sml, err := client.Insights.ListSummaryMetrics(ctx, projectSlug, InsightsListSummaryMetricsOptions{
		PageToken:       String("1"),
		AllBranches:     Bool(true),
		ReportingWindow: ReportingWindow(Last90Days),
	})
	if err != nil {
		t.Errorf("Insights.ListSummaryMetrics got error: %v", err)
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
		t.Errorf("Insights.ListSummaryMetrics got %+v, want %+v", sml, want)
	}
}
