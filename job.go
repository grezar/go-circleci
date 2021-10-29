package circleci

import (
	"context"
	"fmt"
	"time"
)

type Jobs interface {
	Get(ctx context.Context, projectSlug string, jobNumber string) (*Job, error)
}

type jobs struct {
	client *Client
}

type Job struct {
	WebURL         string          `json:"web_url"`
	Project        *JobProject     `json:"project"`
	ParallelRuns   []*ParallelRuns `json:"parallel_runs"`
	StartedAt      time.Time       `json:"started_at"`
	LatestWorkflow *LatestWorkflow `json:"latest_workflow"`
	Name           string          `json:"name"`
	Executor       *Executor       `json:"executor"`
	Parallelism    int             `json:"parallelism"`
	Status         interface{}     `json:"status"`
	Number         int             `json:"number"`
	Pipeline       *JobPipeline    `json:"pipeline"`
	Duration       int             `json:"duration"`
	CreatedAt      time.Time       `json:"created_at"`
	Messages       []*JobMessage   `json:"messages"`
	Contexts       []*Context      `json:"contexts"`
	Organization   Organization    `json:"organization"`
	QueuedAt       time.Time       `json:"queued_at"`
	StoppedAt      time.Time       `json:"stopped_at"`
}

type JobProject struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	ExternalURL string `json:"external_url"`
}

type ParallelRuns struct {
	Index  int    `json:"index"`
	Status string `json:"status"`
}

type LatestWorkflow struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Executor struct {
	Type          string `json:"type"`
	ResourceClass string `json:"resource_class"`
}

type JobPipeline struct {
	ID string `json:"id"`
}

type JobMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

type JobContext struct {
	Name string `json:"name"`
}

type Organization struct {
	Name string `json:"name"`
}

func (s *jobs) Get(ctx context.Context, projectSlug string, jobNumber string) (*Job, error) {
	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	if !validString(&jobNumber) {
		return nil, ErrRequiredJobNumber
	}

	u := fmt.Sprintf("project/%s/job/%s", projectSlug, jobNumber)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	j := &Job{}
	err = s.client.do(ctx, req, j)
	if err != nil {
		return nil, err
	}

	return j, nil
}