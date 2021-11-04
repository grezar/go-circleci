//go:generate mockgen -source=$GOFILE -package=mock -destination=./mocks/$GOFILE
package circleci

import (
	"context"
	"fmt"
)

type Users interface {
	Me(ctx context.Context) (*User, error)
	Collaborations(ctx context.Context) ([]*Collaboration, error)
	GetUser(ctx context.Context, id string) (*User, error)
}

// users implements Users interface
type users struct {
	client *Client
}

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

func (s *users) Me(ctx context.Context) (*User, error) {
	u := "me"
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = s.client.do(ctx, req, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type Collaboration struct {
	VcsType   string `json:"vcs-type"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

func (s *users) Collaborations(ctx context.Context) ([]*Collaboration, error) {
	u := "me/collaborations"
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var cs []*Collaboration
	err = s.client.do(ctx, req, &cs)
	if err != nil {
		return nil, err
	}

	return cs, nil
}

func (s *users) GetUser(ctx context.Context, id string) (*User, error) {
	if !validString(&id) {
		return nil, ErrRequiredUserID
	}
	u := fmt.Sprintf("user/%s", id)
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = s.client.do(ctx, req, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
