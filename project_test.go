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

func Test_projects_ListCheckoutKeys(t *testing.T) {
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
	pckl, err := client.Projects.ListCheckoutKeys(ctx, projectSlug)
	if err != nil {
		t.Errorf("Projects.ListCheckoutKeys got error: %v", err)
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
		t.Errorf("Projects.ListCheckoutKeys got %+v, want %+v", pckl, want)
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

func Test_projects_ListVariables(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/envvar", projectSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"items": [{"name": "ENV1"}], "next_page_token": "1"}`)
	})

	ctx := context.Background()
	pvl, err := client.Projects.ListVariables(ctx, projectSlug)
	if err != nil {
		t.Errorf("Projects.ListVariables got error: %v", err)
	}

	want := &ProjectVariableList{
		Items: []*ProjectVariable{
			{
				Name: "ENV1",
			},
		},
		NextPageToken: "1",
	}

	if !cmp.Equal(pvl, want) {
		t.Errorf("Projects.ListVariables got %+v, want %+v", pvl, want)
	}
}

func Test_projects_DeleteVariable(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	variableName := "ENV1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/envvar/%s", projectSlug, variableName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"message": "string"}`)
	})

	ctx := context.Background()
	err := client.Projects.DeleteVariable(ctx, projectSlug, variableName)
	if err != nil {
		t.Errorf("Projects.DeleteVariable got error: %v", err)
	}
}

func Test_projects_GetVariable(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"
	variableName := "ENV1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/envvar/%s", projectSlug, variableName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"name": "ENV1", "value": "VAL1"}`)
	})

	ctx := context.Background()
	pv, err := client.Projects.GetVariable(ctx, projectSlug, variableName)
	if err != nil {
		t.Errorf("Projects.GetVariable got error: %v", err)
	}

	want := &ProjectVariable{
		Name:  variableName,
		Value: "VAL1",
	}

	if !cmp.Equal(pv, want) {
		t.Errorf("Projects.GetVariable got %+v, want %+v", pv, want)
	}
}

func Test_projects_TriggerPipeline(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/pipeline", projectSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		testBody(t, r, `{"branch":"main","tag":"v0.1.0","parameters":{"deploy_prod":true}}`+"\n")
		fmt.Fprint(w, `{"id": "1","state": "created", "number": 0}`)
	})

	ctx := context.Background()
	p, err := client.Projects.TriggerPipeline(ctx, projectSlug, ProjectTriggerPipelineOptions{
		Branch: String("main"),
		Tag:    String("v0.1.0"),
		Parameters: map[string]interface{}{
			"deploy_prod": true,
		},
	})

	if err != nil {
		t.Errorf("Projects.TriggerPipeline got error: %v", err)
	}

	want := &Pipeline{
		ID:     "1",
		State:  "created",
		Number: 0,
	}

	if !cmp.Equal(p, want) {
		t.Errorf("Projects.TriggerPipeline got %+v, want %+v", p, want)
	}
}
