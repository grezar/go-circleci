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

func Test_pipelines_Continue(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/pipeline/continue", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		testBody(t, r, `{"continuation-key":"key1","configuration":"cfg1","parameters":{"deploy_prod":true}}`+"\n")
		fmt.Fprint(w, `{"message": "string"}`)
	})

	ctx := context.Background()
	err := client.Pipelines.Continue(ctx, PipelineContinueOptions{
		ContinuationKey: String("key1"),
		Configuration:   String("cfg1"),
		Parameters: map[string]interface{}{
			"deploy_prod": true,
		},
	})
	if err != nil {
		t.Errorf("Pipelines.Continue got error: %v", err)
	}
}

func Test_pipelines_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	pipelineID := "pipeline1"

	mux.HandleFunc(fmt.Sprintf("/pipeline/%s", pipelineID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"id": "1"}`)
	})

	ctx := context.Background()
	p, err := client.Pipelines.Get(ctx, pipelineID)
	if err != nil {
		t.Errorf("Pipeline.Get got error: %v", err)
	}

	want := &Pipeline{
		ID: "1",
	}

	if !cmp.Equal(p, want) {
		t.Errorf("Pipeline.Get got %+v, want %+v", p, want)
	}
}
