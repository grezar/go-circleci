package circleci

import (
	"context"
	"fmt"
)

type Webhooks interface {
	List(ctx context.Context, projectSlug string) (*WebhookList, error)
	Create(ctx context.Context, webhook Webhook) (*Webhook, error)
}
type webhooks struct {
	client *Client
}

type WebhookList struct {
	Items         []*Webhook `json:"items"`
	NextPageToken string     `json:"next_page_token"`
}
type Webhook struct {
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

func (w *webhooks) List(ctx context.Context, projectID string) (*WebhookList, error) {
	if !validString(&projectID) {
		return nil, ErrRequiredProjectID
	}

	u := fmt.Sprintf("webhook?scope-id=%s&scope-type=project", projectID)
	req, err := w.client.newRequest("GET", u, nil)
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

func (w *webhooks) Create(ctx context.Context, webhook Webhook) (*Webhook, error) {
	req, err := w.client.newRequest("POST", "webhook", webhook)
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
