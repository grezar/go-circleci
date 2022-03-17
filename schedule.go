//go:generate mockgen -source=$GOFILE -package=mock -destination=./mocks/$GOFILE
package circleci

import (
	"context"
	"fmt"
	"time"
)

type Schedules interface {
	List(ctx context.Context, projectSlug string, options ScheduleListOptions) (*ScheduleList, error)
	Get(ctx context.Context, scheduleID string) (*Schedule, error)
}

// schedules implements Contexts interface
type schedules struct {
	client *Client
}

type Schedule struct {
	ID          string            `json:"id"`
	Timetable   Timetable         `json:"timetable"`
	UpdatedAt   time.Time         `json:"updated-at"`
	Name        string            `json:"name"`
	CreatedAt   time.Time         `json:"created-at"`
	ProjectSlug string            `json:"project-slug"`
	Parameters  map[string]string `json:"parameters"`
	Actor       Actor             `json:"actor"`
	Description string            `json:"description"`
}

type Timetable struct {
	PerHour    int      `json:"per-hour"`
	HoursOfDay []int    `json:"hours-of-day"`
	DaysOfWeek []string `json:"days-of-week"`
}

type ScheduleList struct {
	Items         []*Schedule `json:"items"`
	NextPageToken string      `json:"next_page_token"`
}

type ScheduleListOptions struct {
	PageToken *string `url:"page-token,omitempty"`
}

func (o ScheduleListOptions) valid() error {
	// Nothing is required
	return nil
}

func (s *schedules) List(ctx context.Context, projectSlug string, options ScheduleListOptions) (*ScheduleList, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	if !validString(&projectSlug) {
		return nil, ErrRequiredProjectSlug
	}

	u := fmt.Sprintf("project/%s/schedule", projectSlug)
	req, err := s.client.newRequest("GET", u, &options)
	if err != nil {
		return nil, err
	}

	sl := &ScheduleList{}
	err = s.client.do(ctx, req, sl)
	if err != nil {
		return nil, err
	}

	return sl, nil
}

func (s *schedules) Get(ctx context.Context, scheduleID string) (*Schedule, error) {
	if !validString(&scheduleID) {
		return nil, ErrRequiredScheduleID
	}

	u := fmt.Sprintf("schedule/%s", scheduleID)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	schedule := &Schedule{}
	err = s.client.do(ctx, req, schedule)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}
