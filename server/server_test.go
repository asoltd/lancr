// server_test package only tests the backend implementation, it does not test
// the gateway nor the gRPC transport layer.
// all of the tests run on the staging environment, classic >>> mockist
package server_test

import (
	"testing"

	"github.com/asoltd/lancr/server"
)

func TestBackendNew(t *testing.T) {
	t.Parallel()

	_, err := server.New()
	if err != nil {
		t.Fatal(err)
	}
}
