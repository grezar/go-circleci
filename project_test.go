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

func Test_projects_GetAllCheckoutKeys(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	keyType := "deploy-key"

	mux.HandleFunc(fmt.Sprintf("/project/%s/checkout-key", projectSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"items": [{"type": "deploy-key"}], "next_page_token": "1"}`)
	})

	ctx := context.Background()
	pckl, err := client.Projects.GetAllCheckoutKeys(ctx, projectSlug)
	if err != nil {
		t.Errorf("Projects.GetAllCheckoutKeys got error: %v", err)
	}

	want := &ProjectCheckoutKeyList{
		Items: []*ProjectCheckoutKey{
			{
				Type: keyType,
			},
		},
		NextPageToken: "1",
	}

	if !cmp.Equal(pckl, want) {
		t.Errorf("Projects.GetAllCheckoutKeys got %+v, want %+v", pckl, want)
	}
}

func Test_projects_GetCheckoutKey(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	fingerprint := "xx:yy:zz"

	mux.HandleFunc(fmt.Sprintf("/project/%s/checkout-key/%s", projectSlug, fingerprint), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"fingerprint": "xx:yy:zz"}`)
	})

	ctx := context.Background()
	pck, err := client.Projects.GetCheckoutKey(ctx, projectSlug, fingerprint)
	if err != nil {
		t.Errorf("Projects.GetCheckoutKey got error: %v", err)
	}

	want := &ProjectCheckoutKey{
		Fingerprint: fingerprint,
	}

	if !cmp.Equal(pck, want) {
		t.Errorf("Projects.GetCheckoutKey got %+v, want %+v", pck, want)
	}
}

func Test_projects_DeletetCheckoutKey(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	fingerprint := "xx:yy:zz"

	mux.HandleFunc(fmt.Sprintf("/project/%s/checkout-key/%s", projectSlug, fingerprint), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"message": "string"}`)
	})

	ctx := context.Background()
	err := client.Projects.DeleteCheckoutKey(ctx, projectSlug, fingerprint)
	if err != nil {
		t.Errorf("Projects.DeleteCheckoutKey got error: %v", err)
	}
}

func Test_projects_CreateVariable(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	variableName := "ENV1"
	variableValue := "VAL1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/envvar", projectSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		testBody(t, r, `{"name":"ENV1","value":"VAL1"}`+"\n")
		fmt.Fprint(w, `{"name": "ENV1", "value": "VAL1"}`)
	})

	ctx := context.Background()
	pv, err := client.Projects.CreateVariable(ctx, projectSlug, ProjectCreateVariableOptions{
		Name:  String(variableName),
		Value: String(variableValue),
	})
	if err != nil {
		t.Errorf("Projects.CreateVariable got error: %v", err)
	}

	want := &ProjectVariable{
		Name:  variableName,
		Value: variableValue,
	}

	if !cmp.Equal(pv, want) {
		t.Errorf("Projects.CreateVariable got %+v, want %+v", pv, want)
	}
}
