package circleci

import (
	"context"
	"time"
)

type Pipelines interface {
	List(ctx context.Context, options PipelineListOptions) (*PipelineList, error)
	Continue(ctx context.Context, options PipelineContinueOptions) error
}

type pipelines struct {
	client *Client
}

type PipelineList struct {
	Items         []*Pipeline `json:"items"`
	NextPageToken string      `json:"next_page_token"`
}

type Pipeline struct {
	ID          string           `json:"id"`
	ProjectSlug string           `json:"project_slug"`
	State       string           `json:"state"`
	Number      int64            `json:"number"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at,omitempty"`
	Trigger     *Trigger         `json:"trigger"`
	Vcs         *VCS             `json:"vcs"`
	Errors      []*PipelineError `json:"errors"`
}

type Trigger struct {
	Type       string    `json:"type"`
	ReceivedAt time.Time `json:"received_at"`
	Actor      *Actor    `json:"actor"`
}

type Actor struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

type VCS struct {
	ProviderName        string  `json:"provider_name"`
	TargetRepositoryURL string  `json:"target_repository_url"`
	Branch              string  `json:"branch,omitempty"`
	ReviewID            string  `json:"review_id,omitempty"`
	ReviewURL           string  `json:"review_url,omitempty"`
	Revision            string  `json:"revision"`
	Tag                 string  `json:"tag,omitempty"`
	OriginRepositoryURL string  `json:"origin_repository_url"`
	Commit              *Commit `json:"commit"`
}

type Commit struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type PipelineError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type PipelineListOptions struct {
	OrgSlug   *string `json:"org-slug,omitempty"`
	Mine      *bool   `json:"mine,omitempty"`
	PageToken *string `json:"page-token,omitempty"`
}

func (o PipelineListOptions) valid() error {
	// Nothing is required
	return nil
}

func (s *pipelines) List(ctx context.Context, options PipelineListOptions) (*PipelineList, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	u := "pipeline"
	req, err := s.client.newRequest("GET", u, &options)
	if err != nil {
		return nil, err
	}

	pl := &PipelineList{}
	err = s.client.do(ctx, req, pl)
	if err != nil {
		return nil, err
	}

	return pl, nil
}

type PipelineContinueOptions struct {
	ContinuationKey *string                `json:"continuation-key"`
	Configuration   *string                `json:"configuration"`
	Parameters      map[string]interface{} `json:"parameters,omitempty"`
}

func (o PipelineContinueOptions) valid() error {
	if !validString(o.ContinuationKey) {
		return ErrRequiredPipelineContinuationKey
	}

	if !validString(o.Configuration) {
		return ErrRequiredPipelineConfiguration
	}

	return nil
}

func (s *pipelines) Continue(ctx context.Context, options PipelineContinueOptions) error {
	if err := options.valid(); err != nil {
		return err
	}

	u := "pipeline/continue"
	req, err := s.client.newRequest("POST", u, &options)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}
