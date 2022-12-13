package circleci

import (
	"context"
	"fmt"
)

type Webhooks interface {
	Get(ctx context.Context, id string) (*Webhook, error)
	List(ctx context.Context, options WebhookListOptions) (*WebhookList, error)
	Create(ctx context.Context, options WebhookCreateOptions) (*Webhook, error)
}

type webhooks struct {
	client *Client
}

type WebhookList struct {
	Items         []*Webhook `json:"items"`
	NextPageToken string     `json:"next_page_token"`
}

type Webhook struct {
	ID            string   `json:"id"`
	URL           string   `json:"url"`
	Name          string   `json:"name"`
	SigningSecret string   `json:"signing-secret"`
	Scope         Scope    `json:"scope"`
	Events        []string `json:"events"`
	VerifyTLS     bool     `json:"verify-tls"`
}

type Scope struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (s *webhooks) Get(ctx context.Context, id string) (*Webhook, error) {
	if !validString(&id) {
		return nil, ErrRequiredWebhookID
	}

	u := fmt.Sprintf("webhook/%s", id)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	w := &Webhook{}
	err = s.client.do(ctx, req, w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

type WebhookListOptions struct {
	ScopeID   *string `url:"scope-id,omitempty"`
	ScopeType *string `url:"scope-type,omitempty"`
}

func (o WebhookListOptions) valid() error {
	if !validString(o.ScopeID) {
		return ErrRequiredWebhookScopeID
	}

	if !validString(o.ScopeType) {
		return ErrRequiredWebhookScopeType
	}

	return nil
}

func (w *webhooks) List(ctx context.Context, options WebhookListOptions) (*WebhookList, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	u := "webhook"
	req, err := w.client.newRequest("GET", u, options)
	if err != nil {
		return nil, err
	}

	wb := &WebhookList{}
	err = w.client.do(ctx, req, wb)
	if err != nil {
		return nil, err
	}

	return wb, nil
}

type Event string

const (
	EventWorkflowCompleted Event = "workflow-completed"
	EventJobCompleted      Event = "job-completed"
)

type WebhookCreateOptions struct {
	Name          *string  `json:"name"`
	Events        []*Event `json:"events"`
	URL           *string  `json:"url"`
	VerifyTLS     *bool    `json:"verify-tls"`
	SigningSecret *string  `json:"signing-secret"`
	Scope         *Scope   `json:"scope"`
}

func (o WebhookCreateOptions) valid() error {
	if !validString(o.Name) {
		return ErrRequiredWebhookName
	}

	if !validArrayOfEvent(o.Events) {
		return ErrRequiredWebhookEvents
	}

	if !validString(o.URL) {
		return ErrRequiredWebhookURL
	}

	if !validBool(o.VerifyTLS) {
		return ErrRequiredWebhookVerifyTLS
	}

	if !validString(o.SigningSecret) {
		return ErrRequiredWebhookSigningSecret
	}

	if !validString(&o.Scope.ID) {
		return ErrRequiredWebhookScopeID
	}

	if !validString(&o.Scope.Type) {
		return ErrRequiredWebhookScopeType
	}

	return nil
}

func (w *webhooks) Create(ctx context.Context, options WebhookCreateOptions) (*Webhook, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	u := "webhook"
	req, err := w.client.newRequest("POST", u, &options)
	if err != nil {
		return nil, err
	}

	wb := &Webhook{}
	err = w.client.do(ctx, req, wb)
	if err != nil {
		return nil, err
	}

	return wb, nil
}
