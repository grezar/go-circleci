package circleci

import (
	"context"
	"fmt"
	"time"
)

type Workflows interface {
	Get(ctx context.Context, id string) (*Workflow, error)
	ApproveJob(ctx context.Context, id, approvalRequestID string) error
}

// workflows implements Workflows interface
type workflows struct {
	client *Client
}

type Workflow struct {
	PipelineID     string    `json:"pipeline_id"`
	CanceledBy     string    `json:"canceled_by"`
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	ProjectSlug    string    `json:"project_slug"`
	ErroredBy      string    `json:"errored_by"`
	Tag            string    `json:"tag"`
	Status         string    `json:"status"`
	StartedBy      string    `json:"started_by"`
	PipelineNumber int64     `json:"pipeline_number"`
	CreatedAt      time.Time `json:"created_at"`
	StoppedAt      time.Time `json:"stopped_at"`
}

func (s *workflows) Get(ctx context.Context, id string) (*Workflow, error) {
	if !validString(&id) {
		return nil, ErrRequiredWorkflowsWorkflowID
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
		return ErrRequiredWorkflowsWorkflowID
	}

	if !validString(&approvalRequestID) {
		return ErrRequiredWorkflowsApprovalRequestID
	}

	u := fmt.Sprintf("workflow/%s/approve/%s", id, approvalRequestID)
	req, err := s.client.newRequest("POST", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}
