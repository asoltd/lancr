package server

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/asoltd/lancr/db"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ContextWithAuthorizationMetadata creates a context with the authorization
// header for imitating gateway-incoming requests
func MockContextWithAuthorizationMetadata(token string) context.Context {
	metadataMap := make(map[string]string)
	metadataMap["authorization"] = token
	md := metadata.New(metadataMap)
	ctx := metadata.NewIncomingContext(context.Background(), md)
	return ctx
}

type MockAuthServiceClient struct {
	lancrv1.AuthServiceClient
}

func (m *MockAuthServiceClient) Authenticate(ctx context.Context, req *lancrv1.AuthenticateRequest, opts ...grpc.CallOption) (*lancrv1.AuthenticateResponse, error) {
	return &lancrv1.AuthenticateResponse{}, nil
}

func GetIDTokenFromMetadata(ctx context.Context) (*string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get grpc metadata from context")
	}
	log.Println(ctx)

	xheader, ok := md["authorization"]
	if !ok || len(xheader) < 1 {
		return nil, fmt.Errorf("missing Authorization header")
	}

	idtoken := xheader[0][7:]

	return &idtoken, nil
}

// Open an in-memory database connection SQLite3
func ConnectTestDB() (*gorm.DB, error) {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to open database connection: %v", err)
	}

	err = db.RunMigrations(gormDB)
	if err != nil {
		return nil, fmt.Errorf("Failed to run migrations: %v", err)
	}

	return gormDB, nil
}

func SetupTestDB(t *testing.T) *gorm.DB {
	gormDB, err := ConnectTestDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	return gormDB
}
