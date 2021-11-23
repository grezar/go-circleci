package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_users_Me(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"id": "1", "login": "login1", "name": "name1"}`)
	})

	ctx := context.Background()
	u, err := client.Users.Me(ctx)
	if err != nil {
		t.Errorf("Users.Me got error: %v", err)
	}

	want := &User{
		ID:    "1",
		Login: "login1",
		Name:  "name1",
	}

	if !cmp.Equal(u, want) {
		t.Errorf("Users.Me got %+v, want %+v", u, want)
	}
}

func Test_users_Collaborations(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/me/collaborations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `[{"vcs-type": "vcs1", "name": "name1", "avatar_url": "avatar1"}]`)
	})

	ctx := context.Background()
	cs, err := client.Users.Collaborations(ctx)
	if err != nil {
		t.Errorf("Users.Collaborations got error: %v", err)
	}

	want := []*Collaboration{
		{
			VcsType:   "vcs1",
			Name:      "name1",
			AvatarURL: "avatar1",
		},
	}

	if !cmp.Equal(cs, want) {
		t.Errorf("Users.Collaborations got %+v, want %+v", cs, want)
	}
}

func Test_users_GetUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	userID := "user1"

	mux.HandleFunc(fmt.Sprintf("/user/%s", userID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/vnd.api+json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"id": "1", "login": "login1", "name": "name1"}`)
	})

	ctx := context.Background()
	u, err := client.Users.GetUser(ctx, userID)
	if err != nil {
		t.Errorf("Users.GetUser got error: %v", err)
	}

	want := &User{
		ID:    "1",
		Login: "login1",
		Name:  "name1",
	}

	if !cmp.Equal(u, want) {
		t.Errorf("Users.GetUser got %+v, want %+v", u, want)
	}
}
