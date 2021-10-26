package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_pipelines_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/pipeline/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"items": [{"id": "1", "trigger": {"type": "explicit"}}], "next_page_token": "1"}`)
	})

	ctx := context.Background()
	pl, err := client.Pipelines.List(ctx, PipelineListOptions{
		OrgSlug: String("org1"),
		Mine:    Bool(true),
	})
	if err != nil {
		t.Errorf("Pipelines.List got error: %v", err)
	}

	want := &PipelineList{
		Items: []*Pipeline{
			{
				ID: "1",
				Trigger: &Trigger{
					Type: "explicit",
				},
			},
		},
		NextPageToken: "1",
	}

	if !cmp.Equal(pl, want) {
		t.Errorf("Pipelines.List got %+v, want %+v", pl, want)
	}
}