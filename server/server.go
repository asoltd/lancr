// server package contains the backend implementation of the services defined in
// the proto files. The backend is responsible for implementing the business
// logic of the application, it is also responsible for interacting with the
// database and other external services.
package server

import (
	"context"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	lancrv1 "github.com/asoltd/lancr/gen/go/proto/lancr/v1"
	"google.golang.org/api/option"
)

type Backend struct {
	mu        *sync.RWMutex
	app       *firebase.App
	firestore *firestore.Client

	lancrv1.UnimplementedHeroServiceServer
}

// New sets up the backend with authenticated firestore SDK client If keypath
// parameter is left specified, it authenticates with a given service account
// JSON, otherwise with application default credentials
func New(keypath *string) (*Backend, error) {
	var app *firebase.App
	var err error
	if keypath != nil {
		if _, err := os.Stat(*keypath); os.IsNotExist(err) {
			return nil, err
		}

		sa := option.WithCredentialsFile(*keypath)
		app, err = firebase.NewApp(
			context.Background(),
			nil,
			sa,
		)
	} else {
		app, err = firebase.NewApp(
			context.Background(),
			nil,
		)
	}
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

func NewLocal(keypath string) (*Backend, error) {
	return New(&keypath)
}

func NewRemote() (*Backend, error) {
	return New(nil)
}
