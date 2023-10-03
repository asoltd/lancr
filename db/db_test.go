package db_test

import (
	"testing"

	"github.com/asoltd/lancr/db"
)

func TestConnectClickhouse(t *testing.T) {
	_, err := db.ConnectClickhouse()
	if err != nil {
		t.Fatalf("ConnectClickhouse() failed: %v", err)
	}
}
