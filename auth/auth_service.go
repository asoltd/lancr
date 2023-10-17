package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
)

type AuthService struct {
	lancrv1.UnimplementedAuthServiceServer
	authclient *auth.Client
}

var claims = map[string]interface{}{}

func NewAuthService(ctx context.Context) (*AuthService, error) {
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "asoltd-guilds",
	})
	if err != nil {
		return nil, err
	}

	authclient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &AuthService{
		authclient: authclient,
	}, nil
}

func (a *AuthService) Authenticate(ctx context.Context, req *lancrv1.AuthenticateRequest) (*lancrv1.AuthenticateResponse, error) {
	idtoken := req.GetIdToken()
	token, err := a.authclient.VerifyIDToken(ctx, idtoken)
	if err != nil {
		return nil, err
	}

	// this might be useful, consider using kv store
	claims[token.UID] = token.Claims

	// user also has some other fields that might be useful
	// log.Printf("user: %+v", user.UserMetadata)
	// log.Printf("user: %+v", user.ProviderUserInfo)

	log.Println("authenticated user:", token.UID)

	return &lancrv1.AuthenticateResponse{
		Token: &lancrv1.Token{
			AuthTime: token.AuthTime,
			Issuer:   token.Issuer,
			Audience: token.Audience,
			Expires:  token.Expires,
			IssuedAt: token.IssuedAt,
			Subject:  token.Subject,
			Uid:      token.UID,
			FirebaseInfo: &lancrv1.FirebaseInfo{
				SignInProvider: token.Firebase.SignInProvider,
				Tenant:         token.Firebase.Tenant,
			},
		},
	}, nil
}
