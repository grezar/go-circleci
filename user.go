package circleci

import (
	"context"
)

type Users interface {
	Me(ctx context.Context) (*User, error)
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
