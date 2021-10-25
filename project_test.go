package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_projects_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"

	mux.HandleFunc(fmt.Sprintf("/project/%s", projectSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"slug": "gh/org1/prj1"}`)
	})

	ctx := context.Background()
	p, err := client.Projects.Get(ctx, projectSlug)
	if err != nil {
		t.Errorf("Projects.Get got error: %v", err)
	}

	want := &Project{
		Slug: projectSlug,
	}

	if !cmp.Equal(p, want) {
		t.Errorf("Projects.Get got %+v, want %+v", p, want)
	}
}

func Test_projects_CreateCheckoutKey(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	keyType := "deploy-key"

	mux.HandleFunc(fmt.Sprintf("/project/%s/checkout-key", projectSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		testBody(t, r, `{"type":"deploy-key"}`+"\n")
		fmt.Fprint(w, `{"type": "deploy-key", "preferred": true}`)
	})

	ctx := context.Background()
	pck, err := client.Projects.CreateCheckoutKey(ctx, projectSlug, ProjectCreateCheckoutKeyOptions{
		Type: String(keyType),
	})
	if err != nil {
		t.Errorf("Projects.CreateCheckoutKey got error: %v", err)
	}

	want := &ProjectCheckoutKey{
		Type:      keyType,
		Preferred: true,
	}

	if !cmp.Equal(pck, want) {
		t.Errorf("Projects.CreateCheckoutKey got %+v, want %+v", pck, want)
	}
}
