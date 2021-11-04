//go:generate mockgen -source=$GOFILE -package=mock -destination=./mocks/$GOFILE
package circleci

import (
	"context"
	"fmt"
	"time"
)

type Workflows interface {
	Get(ctx context.Context, id string) (*Workflow, error)
	ApproveJob(ctx context.Context, id, approvalRequestID string) error
	Cancel(ctx context.Context, id string) error
	ListWorkflowJobs(ctx context.Context, id string) (*WorkflowJobList, error)
	Rerun(ctx context.Context, id string, options WorkflowRerunOptions) error
}

// workflows implements Workflows interface
type workflows struct {
	client *Client
}

type Workflow struct {
	PipelineID     string      `json:"pipeline_id"`
	CanceledBy     string      `json:"canceled_by"`
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	ProjectSlug    string      `json:"project_slug"`
	ErroredBy      string      `json:"errored_by"`
	Tag            string      `json:"tag"`
	Status         interface{} `json:"status"`
	StartedBy      string      `json:"started_by"`
	PipelineNumber int64       `json:"pipeline_number"`
	CreatedAt      time.Time   `json:"created_at"`
	StoppedAt      time.Time   `json:"stopped_at"`
}

func (s *workflows) Get(ctx context.Context, id string) (*Workflow, error) {
	if !validString(&id) {
		return nil, ErrRequiredWorkflowID
	}

	u := fmt.Sprintf("workflow/%s", id)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	w := &Workflow{}
	err = s.client.do(ctx, req, w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (s *workflows) ApproveJob(ctx context.Context, id, approvalRequestID string) error {
	if !validString(&id) {
		return ErrRequiredWorkflowID
	}

	if !validString(&approvalRequestID) {
		return ErrRequiredApprovalRequestID
	}

	u := fmt.Sprintf("workflow/%s/approve/%s", id, approvalRequestID)
	req, err := s.client.newRequest("POST", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}

func (s *workflows) Cancel(ctx context.Context, id string) error {
	if !validString(&id) {
		return ErrRequiredWorkflowID
	}

	u := fmt.Sprintf("workflow/%s/cancel", id)
	req, err := s.client.newRequest("POST", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}

type WorkflowJobList struct {
	Items         []*WorkflowJob `json:"items"`
	NextPageToken string         `json:"next_page_token"`
}

type WorkflowJob struct {
	ID                string    `json:"id"`
	CanceledBy        string    `json:"canceled_by"`
	Dependencies      []*string `json:"dependencies"`
	JobNumber         int64     `json:"job_number"`
	Name              string    `json:"name"`
	ApprovedBy        string    `json:"approved_by"`
	ProjectSlug       string    `json:"project_slug"`
	Status            string    `json:"status"`
	Type              string    `json:"type"`
	StartedAt         time.Time `json:"started_at"`
	StoppedAt         time.Time `json:"stopped_at"`
	ApprovalRequestID string    `json:"approval_request_id"`
}

func (s *workflows) ListWorkflowJobs(ctx context.Context, id string) (*WorkflowJobList, error) {
	if !validString(&id) {
		return nil, ErrRequiredWorkflowID
	}

	u := fmt.Sprintf("workflow/%s/job", id)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	jl := &WorkflowJobList{}
	err = s.client.do(ctx, req, jl)
	if err != nil {
		return nil, err
	}

	return jl, nil
}

type WorkflowRerunOptions struct {
	Jobs       []*string `json:"jobs,omitempty"`
	FromFailed *bool     `json:"from_failed,omitempty"`
	SparseTree *bool     `json:"sparse_tree,omitempty"`
}

func (o WorkflowRerunOptions) valid() error {
	// Nothing is required
	return nil
}

func (s *workflows) Rerun(ctx context.Context, id string, options WorkflowRerunOptions) error {
	if err := options.valid(); err != nil {
		return err
	}

	if !validString(&id) {
		return ErrRequiredWorkflowID
	}

	u := fmt.Sprintf("workflow/%s/rerun", id)
	req, err := s.client.newRequest("POST", u, options)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}
