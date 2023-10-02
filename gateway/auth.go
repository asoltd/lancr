package gateway

import (
	"context"
	"net/http"

	"firebase.google.com/go/auth"
)

// this can also include auth scopes based on the token
// the custom claims can be added
// token, err := client.CustomTokenWithClaims(ctx, "some-uid", claims)
// TODO ensure that custom claims are going to be persistant
// param idtoken is the Firebase ID Token
func (gw *Gateway) Authenticate(ctx context.Context, idtoken string) (*auth.Token, error) {
	token, err := gw.authclient.VerifyIDToken(ctx, idtoken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// TODO write out this flow for all of the operations user will perform and base
// the implementation of this function around it
func (gw *Gateway) Authorize(userId string, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		// TODO if this is a list request, then the user is allowed to see all of the
		// heroes, otherwise the user is only allowed to see their own hero
		return nil
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		// TODO if the userId matches the owner of the resource under
		// /v1/${resource_path}/${id} then the user is allowed to perform the
		// operation
		return nil
	}
	return nil
}

func (gw *Gateway) AuthMiddleware(ctx context.Context, r *http.Request) error {
	idtoken := r.Header.Get("X-Firebase-ID-Token")
	token, err := gw.Authenticate(ctx, idtoken)
	if err != nil {
		// TODO this is a firebase error, this might not be good to return,
		// at the moment this is written in the 403 not authorized
		// i would like to provide a nice response, in form
		// { 403, { error: "not authorized, principal is missing ${scope}" }
		return err
	}

	err = gw.Authorize(token.UID, r)
	if err != nil {
		return err
	}

	return nil
}
