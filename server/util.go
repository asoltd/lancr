package server

import (
	"context"
	"fmt"
	"strings"
	"testing"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"google.golang.org/grpc/metadata"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetIDTokenFromMetadata(ctx context.Context) (*string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get grpc metadata from context")
	}

	xheader, ok := md["X-Firebase-ID-Token"]
	if !ok {
		return nil, fmt.Errorf("missing X-Firebase-ID-Token header")
	}

	idtoken := strings.Join(xheader, ",")
	return &idtoken, nil
}

func SetupTestDB(t *testing.T) *gorm.DB {
	// Open a test database connection (SQLite in-memory for this example)
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}

	err = db.AutoMigrate(&lancrv1.HeroORM{})
	if err != nil {
		t.Fatalf("Failed to auto-migrate: %v", err)
	}

	err = db.AutoMigrate(&lancrv1.ApprenticeORM{})
	if err != nil {
		t.Fatalf("Failed to auto-migrate: %v", err)
	}

	err = db.AutoMigrate(&lancrv1.QuestORM{})
	if err != nil {
		t.Fatalf("Failed to auto-migrate: %v", err)
	}

	return db
}
