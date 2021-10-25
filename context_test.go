package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_contexts_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/context", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		testQuery(t, r, "owner-slug", "org")
		fmt.Fprint(w, `{"items": [{"id": "1"}], "next_page_token": "1"}`)
	})

	ctx := context.Background()
	cl, err := client.Contexts.List(ctx, ContextListOptions{
		OwnerSlug: String("org"),
	})
	if err != nil {
		t.Errorf("Contexts.List got error: %v", err)
	}

	want := &ContextList{
		Items: []*Context{
			{
				ID: "1",
			},
		},
		NextPageToken: "1",
	}

	if !cmp.Equal(cl, want) {
		t.Errorf("Contexts.List got %+v, want %+v", cl, want)
	}
}

func Test_contexts_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/context", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		testBody(t, r, `{"name":"ctx","owner":{"slug":"org"}}`+"\n")
		fmt.Fprint(w, `{"id": "1"}`)
	})

	ctx := context.Background()
	c, err := client.Contexts.Create(ctx, ContextCreateOptions{
		Name: String("ctx"),
		Owner: &OwnerOptions{
			Slug: String("org"),
		},
	})
	if err != nil {
		t.Errorf("Contexts.Create got error: %v", err)
	}

	want := &Context{
		ID: "1",
	}

	if !cmp.Equal(c, want) {
		t.Errorf("Contexts.Create got %+v, want %+v", c, want)
	}
}

func Test_contexts_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	contextID := "ctx1"

	mux.HandleFunc(fmt.Sprintf("/context/%s", contextID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"id": "1"}`)
	})

	ctx := context.Background()
	c, err := client.Contexts.Get(ctx, contextID)
	if err != nil {
		t.Errorf("Contexts.Get got error: %v", err)
	}

	want := &Context{
		ID: "1",
	}

	if !cmp.Equal(c, want) {
		t.Errorf("Contexts.Get got %+v, want %+v", c, want)
	}
}

func Test_contexts_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	contextID := "ctx1"

	mux.HandleFunc(fmt.Sprintf("/context/%s", contextID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"message": "string"}`)
	})

	ctx := context.Background()
	err := client.Contexts.Delete(ctx, contextID)
	if err != nil {
		t.Errorf("Contexts.Delete got error: %v", err)
	}
}

func Test_contexts_ListVariables(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	contextID := "ctx1"

	mux.HandleFunc(fmt.Sprintf("/context/%s", contextID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"items": [{"variable": "ENVVAR1", "context_id": "ctx1"}], "next_page_token": "1"}`)
	})

	ctx := context.Background()
	cl, err := client.Contexts.ListVariables(ctx, contextID)
	if err != nil {
		t.Errorf("Contexts.ListVariables got error: %v", err)
	}

	want := &ContextVariableList{
		Items: []*ContextVariable{
			{
				Variable:  "ENVVAR1",
				ContextID: "ctx1",
			},
		},
		NextPageToken: "1",
	}

	if !cmp.Equal(cl, want) {
		t.Errorf("Contexts.ListVariables got %+v, want %+v", cl, want)
	}
}

func Test_contexts_RemoveVariable(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	variableName := "envVar1"
	contextID := "ctx1"

	mux.HandleFunc(fmt.Sprintf("/context/%s/environment-variable/%s", contextID, variableName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"message": "string"}`)
	})

	ctx := context.Background()
	err := client.Contexts.RemoveVariable(ctx, contextID, variableName)
	if err != nil {
		t.Errorf("Contexts.RemoveVariable got error: %v", err)
	}
}
