// server package contains the backend implementation of the services defined in
// the proto files. The backend is responsible for implementing the business
// logic of the application, it is also responsible for interacting with the
// database and other external services.
package server

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/asoltd/lancr/db"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"
)

type Backend struct {
	mu         *sync.RWMutex
	app        *firebase.App
	firestore  *firestore.Client
	authclient *auth.Client
	db         *gorm.DB

	lancrv1.UnimplementedHeroServiceServer
}

// This assumes Workload Identity enabled with the authentication being stored
// as GOOGLE_APPLICATION_CREDENTIALS (the default)
// TODO adjust the service account permissions to allow for auth operations
func New() (*Backend, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(
		context.Background(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	firestoreclient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	authclient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	return &Backend{
		mu:         &sync.RWMutex{},
		app:        app,
		firestore:  firestoreclient,
		authclient: authclient,
		db:         db,
	}, nil
}

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
