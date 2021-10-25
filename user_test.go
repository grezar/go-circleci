package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_contexts_Me(t *testing.T) {
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
