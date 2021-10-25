package circleci

import (
	"context"
	"fmt"
)

type Projects interface {
	Get(ctx context.Context, projectSlug string) (*Project, error)
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
