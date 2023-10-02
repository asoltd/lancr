package server

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"google.golang.org/grpc/metadata"
)

type Operation int

// operation enum with values create, read, update, delete
const (
	Create Operation = iota
	Read
	Update
	Delete
)

func (o Operation) String() string {
	switch o {
	case Create:
		return "create"
	case Read:
		return "read"
	case Update:
		return "update"
	case Delete:
		return "delete"
	}
	return ""
}

// I am generally debating whether the authentication/authorization should be on the backend
// on the gateway the operation can be inferred from the HTTP method + the path to determine the resource
// this will be simpler to implement this way, since the backend is a client
// say someone wants to update a hero
//
// UPDATE /v1/heroes/123 { hero: { "name": "new name" } }
//
// then we know that the ID 123 has to be GetHero(123)'ed
// then the response can be used to check if the update can be performed
//
// GetHeroResponse { hero: { id: '123', ownerId: '123' } }
//
// assert token.UID == res.hero.ownerId, otherwise 403
//
// with this flow, the backend can operate without the need to always use the ID
// gateway can be tested separately and the actual HeroService does not require
// any authentication/authorization work, so the tests can also be ran separately
// at the moment, using the below methods prevents testing the business logic of
// the backend without stubs from firebase
//
// this is important, since later on the heroes API will be used by other
// services, stuff like the search service or payment service
//
// those services likely won't have access to the firebase auth
func (b *Backend) getIDTokenFromMetadata(ctx context.Context) (*string, error) {
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

// this can also include auth scopes based on the token
// the custom claims can be added
// token, err := client.CustomTokenWithClaims(ctx, "some-uid", claims)
// TODO ensure that this is going to be persistant
func (b *Backend) Authenticate(ctx context.Context) (*auth.Token, error) {
	idtoken, err := b.getIDTokenFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	token, err := b.authclient.VerifyIDToken(ctx, *idtoken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// for now is read all so only need to lock out the users outside of Firebase +
// prevent edition of non-owned items
// TODO write out this flow for all of the operations user will perform and base
// the implementation of this function around it
func (b *Backend) Authorize(
	token *auth.Token,
	snapshot *firestore.DocumentSnapshot,
	operation Operation,
) error {
	// parse ownerID from snapshot
	ownerId, ok := snapshot.Data()["ownerId"].(string)
	if !ok {
		return fmt.Errorf("failed to parse ownerId from snapshot")
	}
	switch operation {
	case Create:
		// TODO this will require something else
		return nil
	case Read:
		return nil
	case Update:
		if token.UID != ownerId {
			return fmt.Errorf("unauthorized")
		}
	case Delete:
		if token.UID != ownerId {
			return fmt.Errorf("unauthorized")
		}
	}
	return nil
}
