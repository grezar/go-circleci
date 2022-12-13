package circleci

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	url, err := url.Parse(server.URL)
	if err != nil {
		panic(fmt.Sprintf("failed to parse test server URL: %s, error: %v", server.URL, err))
	}

	cfg := DefaultConfig()
	cfg.Token = "fake-token"
	client, err = NewClient(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize client, error: %v", err))
	}
	client.baseURL = url

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Method got %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	t.Helper()
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) got %q, want %q", header, got, want)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	t.Helper()
	var mgot map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&mgot)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	var mwant map[string]interface{}
	json.Unmarshal([]byte(want), &mwant)

	if !cmp.Equal(mgot, mwant) {
		t.Errorf("request Body is %s, want %s", mgot, mwant)
	}
}

func testQuery(t *testing.T, r *http.Request, key, want string) {
	t.Helper()
	got := r.URL.Query().Get(key)
	if got != want {
		t.Errorf("URL.Query(%q) got %q, want %q", key, got, want)
	}
}
