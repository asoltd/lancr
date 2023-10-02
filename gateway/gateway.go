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
	authclient *auth.Client
}

// LICENSE: MIT
// CREDIT: github.com/johanbrandhorst/grpc-gateway-boilerplate
// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
// there is the embed package used in reference implementation but it requires
// putting a file in a generated-code directory
func (gw *Gateway) getOpenAPIHandler() (http.Handler, error) {
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
	return &Gateway{authclient: authclient}, nil
}

// Setup dials the backend and sets up the gateway
func (gw *Gateway) Setup(ctx context.Context, dialAddr string, gatewayAddr string) (*http.Server, error) {
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	gwmux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "X-Firebase-ID-Token":
				return key, true
			default:
				return runtime.DefaultHeaderMatcher(key)
			}
		}),
	)
	err = lancrv1.RegisterHeroServiceHandler(ctx, gwmux, conn)
	if err != nil {
		return nil, err
	}

	oa, err := gw.getOpenAPIHandler()
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
				idtoken := r.Header.Get("X-Firebase-ID-Token")
				if idtoken == "" {
					w.WriteHeader(http.StatusForbidden)
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
			}

			gwmux.ServeHTTP(w, r)
		}),
	}, nil
}
