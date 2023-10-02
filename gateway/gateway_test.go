package gateway_test

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Fatal("TODO (cant let this through CI :P)")
}

// Tests if the token is required for POST, PUT, DELETE
func TestRequiresFirebaseIDToken(t *testing.T) {}

func TestDoesNotRequireFirebaseIDTokenForGET(t *testing.T) {}

// Tests if the only allowed methods are GET, POST, PUT, DELETE
func TestOnlyAllowsAllowedMethods(t *testing.T) {}

func TestAllowsAnyMethodForOpenAPI(t *testing.T) {}

func TestAuthenticate(t *testing.T) {}

func TestAuthorize(t *testing.T) {}
