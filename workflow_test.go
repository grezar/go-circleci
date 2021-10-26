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
