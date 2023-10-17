package gateway

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/asoltd/lancr/gen/go/lancr/v1"
)

func (gw *Gateway) AuthMiddleware(ctx context.Context, r *http.Request) error {
	// TODO here can also include auth scopes based on the token
	// the custom claims can be added
	// token, err := client.CustomTokenWithClaims(ctx, "some-uid", claims)
	// TODO ensure that custom claims are going to be persistant
	// param idtoken is the Firebase ID Token

	authheader := r.Header.Get("Authorization")
	if authheader == "" || len(authheader) < 7 {
		return fmt.Errorf("missing authorization header")
	}
	// split bearer token
	idtoken := authheader[7:]

	// TODO this token can be used here for some stuff, not relevant atm
	_, err := gw.auth.Authenticate(ctx, &v1.AuthenticateRequest{
		IdToken: idtoken,
	})
	if err != nil {
		return fmt.Errorf("failed to verify ID token: %w", err)
	}

	switch r.Method {
	case http.MethodGet:
		// TODO
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		// TODO
	}
	return nil
}
