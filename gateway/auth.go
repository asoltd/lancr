package gateway

import (
	"context"
	"fmt"
	"net/http"
)

func (gw *Gateway) AuthMiddleware(ctx context.Context, r *http.Request) error {
	// TODO here can also include auth scopes based on the token
	// the custom claims can be added
	// token, err := client.CustomTokenWithClaims(ctx, "some-uid", claims)
	// TODO ensure that custom claims are going to be persistant
	// param idtoken is the Firebase ID Token

	idtoken := r.Header.Get("X-Firebase-ID-Token")
	/*token*/ _, err := gw.authclient.VerifyIDToken(ctx, idtoken)
	if err != nil {
		// TODO this is a firebase error, this might not be good to return,
		// at the moment this is written in the 403 not authorized
		// i would like to provide a nice response, in form
		// { 403, { error: "not authorized, principal is missing ${scope}" }
		return fmt.Errorf("failed to verify ID token")
	}

	switch r.Method {
	case http.MethodGet:
		// TODO
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		// TODO
	}
	return nil
}
