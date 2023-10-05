package db_test

import (
	"testing"

	"github.com/asoltd/lancr/db"
)

func TestConnectClickhouse(t *testing.T) {
	t.Skip()
	_, err := db.ConnectClickhouse()
	if err != nil {
		t.Fatalf("ConnectClickhouse() failed: %v", err)
	}
}

func TestConnectTiDB(t *testing.T) {
	_, err := db.ConnectTiDB()
	if err != nil {
		t.Fatalf("ConnectTiDB() failed: %v", err)
	}
}
