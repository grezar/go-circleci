package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_webhooks_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	webhookID := "webhook1"
	mux.HandleFunc(fmt.Sprintf("/webhook/%s", webhookID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"id": "1"}`)
	})

	ctx := context.Background()
	w, err := client.Webhooks.Get(ctx, webhookID)
	if err != nil {
		t.Errorf("Webhooks.Get got error: %v", err)
	}

	want := &Webhook{ID: "1"}

	if !cmp.Equal(w, want) {
		t.Errorf("Webhooks.Get got %+v, want %+v", w, want)
	}
}

func Test_webhooks_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	scopeID := "project1"
	scopeType := "project"
	mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Circle-Token", client.token)
		testQuery(t, r, "scope-id", scopeID)
		testQuery(t, r, "scope-type", scopeType)
		fmt.Fprint(w, `{"items": [{"id": "1"}], "next_page_token": "1"}`)
	})

	ctx := context.Background()
	wl, err := client.Webhooks.List(ctx, WebhookListOptions{
		ScopeID:   String(scopeID),
		ScopeType: String(scopeType),
	})
	if err != nil {
		t.Errorf("Webhooks.List got error: %v", err)
	}

	want := &WebhookList{
		Items: []*Webhook{
			{
				ID: "1",
			},
		},
		NextPageToken: "1",
	}

	if !cmp.Equal(wl, want) {
		t.Errorf("Webhooks.List got %+v, want %+v", wl, want)
	}
}

func Test_webhooks_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	body := `{"name":"webhook","events":["workflow-completed"],"url":"example.com","verify-tls":false,"signing-secret":"xyz","scope":{"id":"123","type":"project"}}`
	mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Circle-Token", client.token)
		testBody(t, r, body+"\n")
		fmt.Fprint(w, `{"id": "1"}`)
	})

	ctx := context.Background()
	wb, err := client.Webhooks.Create(ctx, WebhookCreateOptions{
		Name:          String("webhook"),
		URL:           String("example.com"),
		SigningSecret: String("xyz"),
		Scope: &Scope{
			ID:   "123",
			Type: "project",
		},
		Events:    []*Event{EventType(EventWorkflowCompleted)},
		VerifyTLS: Bool(false),
	})
	if err != nil {
		t.Errorf("Webhooks.Create got error: %v", err)
	}

	want := &Webhook{
		ID: "1",
	}

	if !cmp.Equal(wb, want) {
		t.Errorf("Webhooks.Create got %+v, want %+v", wb, want)
	}
}
