package gateway

import (
	"context"
	"net/http"
)

func (gw *Gateway) AuthMiddleware(ctx context.Context, r *http.Request) error {
	// TODO here can also include auth scopes based on the token
	// the custom claims can be added
	// token, err := client.CustomTokenWithClaims(ctx, "some-uid", claims)
	// TODO ensure that custom claims are going to be persistant
	// param idtoken is the Firebase ID Token

	switch r.Method {
	case http.MethodGet:
		// TODO if this is a list request, then the user is allowed to see all of the
		// heroes, otherwise the user is only allowed to see their own hero
		return nil
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		idtoken := r.Header.Get("X-Firebase-ID-Token")
		token, err := gw.authclient.VerifyIDToken(ctx, idtoken)
		if err != nil {
			// TODO this is a firebase error, this might not be good to return,
			// at the moment this is written in the 403 not authorized
			// i would like to provide a nice response, in form
			// { 403, { error: "not authorized, principal is missing ${scope}" }
			return err
		}
		println(token.UID)
		// TODO if the userId matches the owner of the resource under
		// /v1/${resource_path}/${id} then the user is allowed to perform the
		// operation
		// gw.heroesclient.Get(ctx, ) ..
		return nil
	}
	return nil
}
