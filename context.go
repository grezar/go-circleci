package circleci

import (
	"context"
	"fmt"
	"time"
)

type Contexts interface {
	List(ctx context.Context, options ContextListOptions) (*ContextList, error)
	Get(ctx context.Context, contextID string) (*Context, error)
	Create(ctx context.Context, options ContextCreateOptions) (*Context, error)
}

// contexts implements Contexts interface
type contexts struct {
	client *Client
}

type Context struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ContextList struct {
	Items         []*Context `json:"items"`
	NextPageToken string     `json:"next_page_token"`
}

type ContextListOptions struct {
	OwnerID   *string `url:"owner-id,omitempty"`
	OwnerSlug *string `url:"owner-slug,omitempty"`
	OwnerType *string `url:"owner-type,omitempty"`
	PageToken *string `url:"page-token,omitempty"`
}

func (o ContextListOptions) valid() error {
	if !validString(o.OwnerID) && !validString(o.OwnerSlug) {
		return ErrRequiredEitherIDOrSlug
	}
	return nil
}

func (s *contexts) List(ctx context.Context, options ContextListOptions) (*ContextList, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	u := "context"
	req, err := s.client.newRequest("GET", u, &options)
	if err != nil {
		return nil, err
	}

	cl := &ContextList{}
	err = s.client.do(ctx, req, cl)
	if err != nil {
		return nil, err
	}

	return cl, nil
}

type ContextCreateOptions struct {
	Name  *string       `json:"name"`
	Owner *OwnerOptions `json:"owner"`
}

type OwnerOptions struct {
	ID   *string `json:"id,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Type *string `json:"type,omitempty"`
}

func (o ContextCreateOptions) valid() error {
	if !validString(o.Owner.ID) && !validString(o.Owner.Slug) {
		return ErrRequiredEitherIDOrSlug
	}
	return nil
}

func (s *contexts) Create(ctx context.Context, options ContextCreateOptions) (*Context, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	u := "context"
	req, err := s.client.newRequest("POST", u, &options)
	if err != nil {
		return nil, err
	}

	c := &Context{}
	err = s.client.do(ctx, req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *contexts) Get(ctx context.Context, contextID string) (*Context, error) {
	if contextID == "" {
		return nil, ErrRequiredContextID
	}

	u := fmt.Sprintf("context/%s", contextID)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := &Context{}
	err = s.client.do(ctx, req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
