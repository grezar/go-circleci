package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_jobs_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	jobNumber := "1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/job/%s", projectSlug, jobNumber), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"name": "job1"}`)
	})

	ctx := context.Background()
	j, err := client.Jobs.Get(ctx, projectSlug, jobNumber)
	if err != nil {
		t.Errorf("Jobs.Get got error: %v", err)
	}

	want := &Job{
		Name: "job1",
	}

	if !cmp.Equal(j, want) {
		t.Errorf("Jobs.Get got %+v, want %+v", j, want)
	}
}

func Test_jobs_Cancel(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	jobNumber := "1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/job/%s/cancel", projectSlug, jobNumber), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"message": "success"}`)
	})

	ctx := context.Background()
	err := client.Jobs.Cancel(ctx, projectSlug, jobNumber)
	if err != nil {
		t.Errorf("Jobs.Cancel got error: %v", err)
	}
}

func Test_jobs_ListArtifacts(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	jobNumber := "1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/%s/artifacts", projectSlug, jobNumber), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"items": [{"path": "path", "node_index": 0, "url": "url"}]}`)
	})

	ctx := context.Background()
	j, err := client.Jobs.ListArtifacts(ctx, projectSlug, jobNumber)
	if err != nil {
		t.Errorf("Jobs.ListArtifacts got error: %v", err)
	}

	want := &ArtifactList{
		Items: []*Artifact{
			{
				Path:      "path",
				NodeIndex: 0,
				URL:       "url",
			},
		},
	}

	if !cmp.Equal(j, want) {
		t.Errorf("Jobs.ListArtifacts got %+v, want %+v", j, want)
	}
}

func Test_jobs_ListTestMetadata(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	jobNumber := "1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/%s/tests", projectSlug, jobNumber), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"items": [{"message": "message"}]}`)
	})

	ctx := context.Background()
	tml, err := client.Jobs.ListTestMetadata(ctx, projectSlug, jobNumber)
	if err != nil {
		t.Errorf("Jobs.ListTestMetadata got error: %v", err)
	}

	want := &TestMetadataList{
		Items: []*TestMetadata{
			{
				Message: "message",
			},
		},
	}

	if !cmp.Equal(tml, want) {
		t.Errorf("Jobs.ListTestMetadata got %+v, want %+v", tml, want)
	}
}
