package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_workflows_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	workflowID := "workflow1"

	mux.HandleFunc(fmt.Sprintf("/workflow/%s", workflowID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"id": "1"}`)
	})

	ctx := context.Background()
	w, err := client.Workflows.Get(ctx, workflowID)
	if err != nil {
		t.Errorf("Workflows.Get got error: %v", err)
	}

	want := &Workflow{
		ID: "1",
	}

	if !cmp.Equal(w, want) {
		t.Errorf("Workflows.Get got %+v, want %+v", w, want)
	}
}

func Test_workflows_ApproveJob(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	workflowID := "workflow1"
	jobID := "job1"

	mux.HandleFunc(fmt.Sprintf("/workflow/%s/approve/%s", workflowID, jobID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"message": "string"}`)
	})

	ctx := context.Background()
	err := client.Workflows.ApproveJob(ctx, workflowID, jobID)
	if err != nil {
		t.Errorf("Workflows.ApproveJob got error: %v", err)
	}
}

func Test_workflows_Cancel(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	workflowID := "workflow1"

	mux.HandleFunc(fmt.Sprintf("/workflow/%s/cancel", workflowID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"message": "string"}`)
	})

	ctx := context.Background()
	err := client.Workflows.Cancel(ctx, workflowID)
	if err != nil {
		t.Errorf("Workflows.Cancel got error: %v", err)
	}
}

func Test_workflows_ListWorkflowJobs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	workflowID := "workflow1"

	mux.HandleFunc(fmt.Sprintf("/workflow/%s/job", workflowID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"items": [{"id": "1"}], "next_page_token": "1"}`)
	})

	ctx := context.Background()
	jl, err := client.Workflows.ListWorkflowJobs(ctx, workflowID)
	if err != nil {
		t.Errorf("Workflows.ListWorkflowJobs got error: %v", err)
	}

	want := &WorkflowJobList{
		Items: []*WorkflowJob{
			{
				ID: "1",
			},
		},
		NextPageToken: "1",
	}

	if !cmp.Equal(jl, want) {
		t.Errorf("Workflows.ListWorkflowJobs got %+v, want %+v", jl, want)
	}
}

func Test_workflows_Rerun(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	workflowID := "workflow1"

	mux.HandleFunc(fmt.Sprintf("/workflow/%s/rerun", workflowID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		testBody(t, r, `{"jobs":["xxx-yyy-zzz"],"from_failed":true,"sparse_tree":false}`+"\n")
		fmt.Fprint(w, `{"message": "string"}`)
	})

	ctx := context.Background()
	err := client.Workflows.Rerun(ctx, workflowID, WorkflowRerunOptions{
		Jobs: []*string{
			String("xxx-yyy-zzz"),
		},
		FromFailed: Bool(true),
		SparseTree: Bool(false),
	})
	if err != nil {
		t.Errorf("Workflows.Rerun got error: %v", err)
	}
}
