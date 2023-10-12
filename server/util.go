package server

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/asoltd/lancr/db"
	"google.golang.org/grpc/metadata"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetIDTokenFromMetadata(ctx context.Context) (*string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get grpc metadata from context")
	}

	xheader, ok := md["Authorization"]
	if !ok {
		return nil, fmt.Errorf("missing Authorization header")
	}

	joined := strings.Join(xheader, ",")
	if len(joined) < 7 {
		return nil, fmt.Errorf("invalid Authorization header")
	}

	idtoken := joined[7:]

	return &idtoken, nil
}

func SetupTestDB(t *testing.T) *gorm.DB {
	// Open a test database connection (SQLite in-memory for this example)
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}

	err = db.RunMigrations(gormDB)
	if err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	return gormDB
}
