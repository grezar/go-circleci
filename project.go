package circleci

import (
	"context"
	"fmt"
	"time"
)

type Projects interface {
	Get(ctx context.Context, projectSlug string) (*Project, error)
	CreateCheckoutKey(ctx context.Context, projectSlug string, options ProjectCreateCheckoutKeyOptions) (*ProjectCheckoutKey, error)
	GetAllCheckoutKeys(ctx context.Context, projectSlug string) (*ProjectCheckoutKeyList, error)
	GetCheckoutKey(ctx context.Context, projectSlug, fingerprint string) (*ProjectCheckoutKey, error)
	DeleteCheckoutKey(ctx context.Context, projectSlug, fingerprint string) error
	CreateVariable(ctx context.Context, projectSlug string, options ProjectCreateVariableOptions) (*ProjectVariable, error)
	ListVariables(ctx context.Context, projectSlug string) (*ProjectVariableList, error)
	DeleteVariable(ctx context.Context, projectSlug, name string) error
	GetVariable(ctx context.Context, projectSlug, name string) (*ProjectVariable, error)
}

// projects implementes Projects interface
type projects struct {
	client *Client
}

type Project struct {
	Slug             string   `json:"slug"`
	Name             string   `json:"name"`
	OrganizationName string   `json:"organization_name"`
	VCSInfo          *VCSInfo `json:"vcs_info"`
}

type VCSInfo struct {
	VCSURL        string `json:"vcs_url"`
	Provider      string `json:"provider"`
	DefaultBranch string `json:"default_branch"`
}

func (s *projects) Get(ctx context.Context, projectSlug string) (*Project, error) {
	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	u := fmt.Sprintf("project/%s", projectSlug)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	p := &Project{}
	err = s.client.do(ctx, req, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

type ProjectCheckoutKey struct {
	PublicKey   string    `json:"public-key"`
	Type        string    `json:"type"`
	Fingerprint string    `json:"fingerprint"`
	Preferred   bool      `json:"preferred"`
	CreatedAt   time.Time `json:"created-at"`
}

type ProjectCreateCheckoutKeyOptions struct {
	Type *string `json:"type"`
}

func (o ProjectCreateCheckoutKeyOptions) valid() error {
	if !validString(o.Type) {
		return ErrRequiredProjectCheckoutKeyType
	}
	return nil
}

func (s *projects) CreateCheckoutKey(ctx context.Context, projectSlug string, options ProjectCreateCheckoutKeyOptions) (*ProjectCheckoutKey, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	u := fmt.Sprintf("project/%s/checkout-key", projectSlug)
	req, err := s.client.newRequest("POST", u, options)
	if err != nil {
		return nil, err
	}

	pck := &ProjectCheckoutKey{}
	err = s.client.do(ctx, req, pck)
	if err != nil {
		return nil, err
	}

	return pck, nil
}

type ProjectCheckoutKeyList struct {
	Items         []*ProjectCheckoutKey `json:"items"`
	NextPageToken string                `json:"next_page_token"`
}

func (s *projects) GetAllCheckoutKeys(ctx context.Context, projectSlug string) (*ProjectCheckoutKeyList, error) {
	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	u := fmt.Sprintf("project/%s/checkout-key", projectSlug)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	pckl := &ProjectCheckoutKeyList{}
	err = s.client.do(ctx, req, pckl)
	if err != nil {
		return nil, err
	}

	return pckl, nil
}

func (s *projects) GetCheckoutKey(ctx context.Context, projectSlug, fingerprint string) (*ProjectCheckoutKey, error) {
	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	if !validString(&fingerprint) {
		return nil, ErrRequiredProjectCheckoutKeyFingerprint
	}

	u := fmt.Sprintf("project/%s/checkout-key/%s", projectSlug, fingerprint)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	pck := &ProjectCheckoutKey{}
	err = s.client.do(ctx, req, pck)
	if err != nil {
		return nil, err
	}

	return pck, nil
}

func (s *projects) DeleteCheckoutKey(ctx context.Context, projectSlug, fingerprint string) error {
	if !validString(&projectSlug) {
		return ErrRequiredProjectSlug
	}

	if !validString(&fingerprint) {
		return ErrRequiredProjectCheckoutKeyFingerprint
	}

	u := fmt.Sprintf("project/%s/checkout-key/%s", projectSlug, fingerprint)
	req, err := s.client.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}

type ProjectVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProjectCreateVariableOptions struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

func (o ProjectCreateVariableOptions) valid() error {
	if !validString(o.Name) {
		return ErrRequiredProjectVariableName
	}

	if !validString(o.Value) {
		return ErrRequiredProjectVariableValue
	}

	return nil
}

func (s *projects) CreateVariable(ctx context.Context, projectSlug string, options ProjectCreateVariableOptions) (*ProjectVariable, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	u := fmt.Sprintf("project/%s/envvar", projectSlug)
	req, err := s.client.newRequest("POST", u, options)
	if err != nil {
		return nil, err
	}

	pv := &ProjectVariable{}
	err = s.client.do(ctx, req, pv)
	if err != nil {
		return nil, err
	}

	return pv, nil
}

type ProjectVariableList struct {
	Items         []*ProjectVariable `json:"items"`
	NextPageToken string             `json:"next_page_token"`
}

func (s *projects) ListVariables(ctx context.Context, projectSlug string) (*ProjectVariableList, error) {
	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	u := fmt.Sprintf("project/%s/envvar", projectSlug)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	pvl := &ProjectVariableList{}
	err = s.client.do(ctx, req, pvl)
	if err != nil {
		return nil, err
	}

	return pvl, nil
}

func (s *projects) DeleteVariable(ctx context.Context, projectSlug, name string) error {
	if !validString(&projectSlug) {
		return ErrRequiredProjectSlug
	}

	if !validString(&name) {
		return ErrRequiredProjectVariableName
	}

	u := fmt.Sprintf("project/%s/envvar/%s", projectSlug, name)
	req, err := s.client.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}

func (s *projects) GetVariable(ctx context.Context, projectSlug, name string) (*ProjectVariable, error) {
	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	if !validString(&name) {
		return nil, ErrRequiredProjectVariableName
	}

	u := fmt.Sprintf("project/%s/envvar/%s", projectSlug, name)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	pv := &ProjectVariable{}
	err = s.client.do(ctx, req, pv)
	if err != nil {
		return nil, err
	}

	return pv, nil
}
