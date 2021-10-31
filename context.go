//go:generate mockgen -source=$GOFILE -package=mock -destination=./mocks/$GOFILE
package circleci

import (
	"context"
	"fmt"
	"time"
)

// Contexts describes all the context related methods that the CircleCI API
// supports.
//
// CircleCI API docs: https://circleci.com/docs/api/v2/#tag/Context
type Contexts interface {
	// List all contexts for an owner.
	List(ctx context.Context, options ContextListOptions) (*ContextList, error)

	// Returns basic information about a context.
	Get(ctx context.Context, contextID string) (*Context, error)

	// Create a new context.
	Create(ctx context.Context, options ContextCreateOptions) (*Context, error)

	// Delete a context.
	Delete(ctx context.Context, contextID string) error

	// List information about environment variables in a context, not including their values.
	ListVariables(ctx context.Context, contextID string) (*ContextVariableList, error)

	// Delete an environment variable from a context.
	RemoveVariable(ctx context.Context, contextID string, variableName string) error

	// Create or update an environment variable within a context.
	// Returns information about the environment variable, not including its value.
	AddOrUpdateVariable(ctx context.Context, contextID string, variableName string, options AddOrUpdateVariableOptions) (*ContextVariable, error)
}

// contexts implements Contexts interface
type contexts struct {
	client *Client
}

// Context represents a CircleCI context.
type Context struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// ContextList represents a list of contexts.
type ContextList struct {
	Items         []*Context `json:"items"`
	NextPageToken string     `json:"next_page_token"`
}

// ContextListOptions represents the options for listing contexts.
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

// List all the contexts based on the given options.
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

// ContextCreateOptions represents the options for creating a new context.
type ContextCreateOptions struct {
	Name  *string       `json:"name"`
	Owner *OwnerOptions `json:"owner"`
}

// OwnerOptions represents the owner options of a context.
type OwnerOptions struct {
	// The unique ID of the owner of the context. Specify either this or slug.
	ID *string `json:"id,omitempty"`

	// A string that represents an organization. Specify either this or id. Cannot be used for accounts.
	Slug *string `json:"slug,omitempty"`

	// TODO: Type should use an own type instead of string and also have an
	// default value "organization".
	//
	// If you specify an id.
	// Enum: "account", "organization"
	// The type of the owner. Defaults to "organization". Accounts are only used
	// as context owners in server.
	//
	// If you specify a slug.
	// Value: "organization" The type of owner. Defaults to "organization".
	// Accounts are only used as context owners in server and must be specified by
	// an id instead of a slug.
	Type *string `json:"type,omitempty"`
}

func (o ContextCreateOptions) valid() error {
	if !validString(o.Owner.ID) && !validString(o.Owner.Slug) {
		return ErrRequiredEitherIDOrSlug
	}
	return nil
}

// Create is used to create a new context.
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

// Get a context by its id.
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

// Delete a context by its id.
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

	u := fmt.Sprintf("context/%s", contextID)
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

type AddOrUpdateVariableOptions struct {
	Value *string `json:"value"`
}

func (o AddOrUpdateVariableOptions) valid() error {
	if !validString(o.Value) {
		return ErrRequiredContextVariableValue
	}
	return nil
}

func (s *contexts) AddOrUpdateVariable(ctx context.Context, contextID, variableName string, options AddOrUpdateVariableOptions) (*ContextVariable, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	if !validString(&contextID) {
		return nil, ErrRequiredContextID
	}

	if !validString(&variableName) {
		return nil, ErrRequiredContextVariableName
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
