//go:generate mockgen -source=$GOFILE -package=mock -destination=./mocks/$GOFILE
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
	Delete(ctx context.Context, contextID string) error
	ListVariables(ctx context.Context, contextID string) (*ContextVariableList, error)
	RemoveVariable(ctx context.Context, contextID string, variableName string) error
	AddOrUpdateVariable(ctx context.Context, contextID string, variableName string, options ContextAddOrUpdateVariableOptions) (*ContextVariable, error)
}

// contexts implements Contexts interface
type contexts struct {
	client *Client
}

type OwnerTypeType string

const (
	OwnerTypeOrganization OwnerTypeType = "organization"
	OwnerTypeAccount      OwnerTypeType = "account"
)

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
	OwnerID   *string        `url:"owner-id,omitempty"`
	OwnerSlug *string        `url:"owner-slug,omitempty"`
	OwnerType *OwnerTypeType `url:"owner-type,omitempty"`
	PageToken *string        `url:"page-token,omitempty"`
}

func (o ContextListOptions) valid() error {
	if !validString(o.OwnerID) && !validString(o.OwnerSlug) {
		return ErrRequiredEitherOrganizationIDOrSlug
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
	ID   *string        `json:"id,omitempty"`
	Slug *string        `json:"slug,omitempty"`
	Type *OwnerTypeType `json:"type,omitempty"`
}

func (o ContextCreateOptions) valid() error {
	if !validString(o.Owner.ID) && !validString(o.Owner.Slug) {
		return ErrRequiredEitherOrganizationIDOrSlug
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
	if !validString(&contextID) {
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

func (s *contexts) Delete(ctx context.Context, contextID string) error {
	if !validString(&contextID) {
		return ErrRequiredContextID
	}

	u := fmt.Sprintf("context/%s", contextID)
	req, err := s.client.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}

type ContextVariableList struct {
	Items         []*ContextVariable
	NextPageToken string `json:"next_page_token"`
}

type ContextVariable struct {
	Variable  string    `json:"variable"`
	CreatedAt time.Time `json:"created_at"`
	ContextID string    `json:"context_id"`
}

func (s *contexts) ListVariables(ctx context.Context, contextID string) (*ContextVariableList, error) {
	if !validString(&contextID) {
		return nil, ErrRequiredContextID
	}

	u := fmt.Sprintf("context/%s/environment-variable", contextID)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	cl := &ContextVariableList{}
	err = s.client.do(ctx, req, cl)
	if err != nil {
		return nil, err
	}

	return cl, nil
}

func (s *contexts) RemoveVariable(ctx context.Context, contextID, variableName string) error {
	if !validString(&contextID) {
		return ErrRequiredContextID
	}

	u := fmt.Sprintf("context/%s/environment-variable/%s", contextID, variableName)
	req, err := s.client.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}

type ContextAddOrUpdateVariableOptions struct {
	Value *string `json:"value"`
}

func (o ContextAddOrUpdateVariableOptions) valid() error {
	if !validString(o.Value) {
		return ErrRequiredEnvironmentVariableValue
	}
	return nil
}

func (s *contexts) AddOrUpdateVariable(ctx context.Context, contextID, variableName string, options ContextAddOrUpdateVariableOptions) (*ContextVariable, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	if !validString(&contextID) {
		return nil, ErrRequiredContextID
	}

	if !validString(&variableName) {
		return nil, ErrRequiredEnvironmentVariableName
	}

	u := fmt.Sprintf("context/%s/environment-variable/%s", contextID, variableName)
	req, err := s.client.newRequest("PUT", u, options)
	if err != nil {
		return nil, err
	}

	cv := &ContextVariable{}
	err = s.client.do(ctx, req, cv)
	if err != nil {
		return nil, err
	}

	return cv, nil
}
