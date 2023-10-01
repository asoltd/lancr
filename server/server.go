// server package contains the backend implementation of the services defined in
// the proto files. The backend is responsible for implementing the business
// logic of the application, it is also responsible for interacting with the
// database and other external services.
package server

import (
	"context"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
)

type Backend struct {
	mu        *sync.RWMutex
	app       *firebase.App
	firestore *firestore.Client

	lancrv1.UnimplementedHeroServiceServer
}

// This assumes Workload Identity enabled with the authentication being stored
// as GOOGLE_APPLICATION_CREDENTIALS (the default)
func New() (*Backend, error) {
	app, err := firebase.NewApp(
		context.Background(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, err
	}

	return &Backend{
		mu:        &sync.RWMutex{},
		app:       app,
		firestore: client,
	}, nil
}
