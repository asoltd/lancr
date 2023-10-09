package gateway

import (
	"context"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/asoltd/lancr/gen"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Gateway struct {
	Handler            http.Handler
	authclient         *auth.Client
	gwmux              *runtime.ServeMux
	client             lancrv1.HeroServiceClient
	connectedToBackend bool
}

// LICENSE: MIT
// CREDIT: github.com/johanbrandhorst/grpc-gateway-boilerplate
// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
// there is the embed package used in reference implementation but it requires
// putting a file in a generated-code directory
func getOpenAPIHandler() (http.Handler, error) {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(gen.OpenAPI, "openapiv2")
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(subFS)), nil
}

// verify that the gateway is running with the right service account
func New(ctx context.Context) (*Gateway, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}

	authclient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	gwmux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "Authorization":
				return key, true
			default:
				return runtime.DefaultHeaderMatcher(key)
			}
		}),
	)

	return &Gateway{
		authclient:         authclient,
		gwmux:              gwmux,
		connectedToBackend: false,
	}, nil
}

// ConnectToBackend assumes that all of the below services are under the same
// gRPC server, in this case that would be the Hero, Apprentice and Quest
// services
func (gw *Gateway) ConnectToBackend(ctx context.Context, dialAddr string) error {
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return err
	}

	err = lancrv1.RegisterHeroServiceHandler(ctx, gw.gwmux, conn)
	if err != nil {
		return err
	}

	err = lancrv1.RegisterApprenticeServiceHandler(ctx, gw.gwmux, conn)
	if err != nil {
		return err
	}

	err = lancrv1.RegisterQuestServiceHandler(ctx, gw.gwmux, conn)
	if err != nil {
		return err
	}

	gw.client = lancrv1.NewHeroServiceClient(conn)
	gw.connectedToBackend = true
	// TODO add a client here? I am thinking of a better way to see if someone can
	// do POST/PUT/DELETEs on their own data

	return nil
}

// SetupHandler sets up the gateway handler, in prod always requires calling ConnectToBackend prior
// Otherwise, all of the calls to the backend server will fail
func (gw *Gateway) SetupHandler(ctx context.Context) error {
	oa, err := getOpenAPIHandler()
	if err != nil {
		return err
	}
	gw.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/v1") {
			oa.ServeHTTP(w, r)
			return
		}

		// only accept GET, POST, PUT, DELETE to the gateway
		// what's more, don't accept any of POST/PUT/DELETE without the
		// X-Firebase-ID-Token header set
		switch r.Method {
		case http.MethodGet:
		case http.MethodPost, http.MethodPut, http.MethodDelete:
			idtoken := r.Header.Get("Authorization")
			if idtoken == "" {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Authorization header is required"))
				return
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// authenticate and authorize the request, return 403 if unauthorized
		err := gw.AuthMiddleware(ctx, r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Printf("Failed to write response: %s", err.Error())
			}
			return
		}

		// this is important for separating the dependency of the gateway on the server
		// in testing, in prod ConnectBackend is always called before SetupHandler
		//
		// in production this is not needed, because the gateway is always going to be
		// connected to the backend
		if gw.connectedToBackend {
			gw.gwmux.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMovedPermanently)
			_, err = w.Write([]byte("Backend is not connected"))
			if err != nil {
				log.Printf("Failed to write repsonse: %s", err.Error())
			}
		}
	})

	return nil
}

// Run is an all-in-one function that sets up the gateway and runs it, designed
// for running the prod version with backend connected
func (gw *Gateway) Run(ctx context.Context, dialAddr string, gatewayAddr string) error {
	err := gw.ConnectToBackend(ctx, dialAddr)
	if err != nil {
		return err
	}

	err = gw.SetupHandler(ctx)
	if err != nil {
		return err
	}

	gwServer := http.Server{
		Addr:    gatewayAddr,
		Handler: gw.Handler,
	}

	err = gwServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
