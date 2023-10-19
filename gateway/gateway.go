package gateway

import (
	"context"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/asoltd/lancr/gen"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Gateway struct {
	Handler            http.Handler
	gwmux              *runtime.ServeMux
	connectedToBackend bool
	auth               lancrv1.AuthServiceClient
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
func New(auth lancrv1.AuthServiceClient) (*Gateway, error) {
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
		gwmux:              gwmux,
		connectedToBackend: false,
		auth:               auth,
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

	err = lancrv1.RegisterTeamsServiceHandler(ctx, gw.gwmux, conn)
	if err != nil {
		return err
	}

	return nil
}

// SetupHandler sets up the gateway handler, in prod always requires calling ConnectToBackend prior
// Otherwise, all of the calls to the backend server will fail
func (gw *Gateway) SetupHandler(ctx context.Context) error {
	// nginx rewrites the request so gateway doesn't work :(
	oa, err := getOpenAPIHandler()
	if err != nil {
		return err
	}
	gw.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/v1") && os.Getenv("RUN_OPENAPI") == "true" {
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

		gw.gwmux.ServeHTTP(w, r)
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
