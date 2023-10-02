// due to the nature of the gateway, the successful authentication +
// authorization cannot be tested, will be checking for the forbidden stuff this
// way, the core functionalities of gateway can be still be tested, without the
// backend connection
package gateway_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/asoltd/lancr/gateway"
)

func TestNew(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := gateway.New(ctx)
	if err != nil {
		t.Fatalf("Failed to create gateway: %v", err)
	}
}

func TestSetupHandler(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	gw, err := gateway.New(ctx)
	if err != nil {
		t.Fatalf("Failed to create gateway: %v", err)
	}
	err = gw.SetupHandler(ctx)
	if err != nil {
		t.Fatalf("Failed to setup handler: %v", err)
	}
}

func TestSwaggerUI(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	gw, err := gateway.New(ctx)
	if err != nil {
		t.Fatalf("Failed to create gateway: %v", err)
	}

	err = gw.SetupHandler(ctx)
	if err != nil {
		t.Errorf("Failed to setup handler: %v", err)
	}

	s := httptest.NewServer(gw.Handler)
	defer s.Close()

	res, err := s.Client().Get(s.URL + "/")
	if err != nil {
		t.Errorf("Failed to GET /: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected to get 200, got %d", res.StatusCode)
	}
}

// Tests if the token is required for POST, PUT, DELETE
func TestRequiresFirebaseIDToken(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	gw, err := gateway.New(ctx)
	if err != nil {
		t.Fatalf("Failed to create gateway: %v", err)
	}

	err = gw.SetupHandler(ctx)
	if err != nil {
		t.Errorf("Failed to setup handler: %v", err)
	}

	s := httptest.NewServer(gw.Handler)
	defer s.Close()

	res, err := s.Client().Post(s.URL+"/v1/heroes", "", nil)
	if err != nil {
		t.Errorf("Failed to POST /v1/heroes: %v", err)
	}

	if res.StatusCode != http.StatusForbidden {
		t.Errorf("Expected to get 403, got %d", res.StatusCode)
	}
}

func TestDoesNotRequireFirebaseIDTokenForGET(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	gw, err := gateway.New(ctx)
	if err != nil {
		t.Fatalf("Failed to create gateway: %v", err)
	}

	err = gw.SetupHandler(ctx)
	if err != nil {
		t.Errorf("Failed to setup handler: %v", err)
	}

	s := httptest.NewServer(gw.Handler)
	defer s.Close()

	res, err := s.Client().Get(s.URL + "/v1/heroes")
	if err != nil {
		t.Errorf("Failed to GET /v1/heroes: %v", err)
	}

	// the response here will be 301 likely, which is fine
	if res.StatusCode == http.StatusForbidden {
		t.Errorf("Expected not to get 403, got %d", res.StatusCode)
	}
}

// Tests if the only allowed methods are GET, POST, PUT, DELETE
func TestOnlyAllowsAllowedMethods(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	gw, err := gateway.New(ctx)
	if err != nil {
		t.Fatalf("Failed to create gateway: %v", err)
	}
	err = gw.SetupHandler(ctx)
	if err != nil {
		t.Fatalf("Failed to setup handler: %v", err)
	}

	s := httptest.NewServer(gw.Handler)
	defer s.Close()

	res, err := s.Client().Do(&http.Request{
		URL: &url.URL{
			Path:   "/v1/heroes",
			Host:   strings.Split(s.URL, "://")[1],
			Scheme: "http",
		},
		Method: http.MethodPatch,
	})
	if err != nil {
		t.Errorf("Failed to PATCH /v1/heroes: %v", err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected to get 405, got %d", res.StatusCode)
	}
}
