package gateway

import (
	"context"
	"fmt"
	"io/fs"
	"mime"
	"net/http"
	"strings"

	"github.com/asoltd/lancr/gen"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

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

func Run(ctx context.Context, dialAddr string, gatewayAddr string) error {
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return err
	}

	gwmux := runtime.NewServeMux()
	err = lancrv1.RegisterHeroServiceHandler(ctx, gwmux, conn)
	if err != nil {
		return err
	}

	oa, err := getOpenAPIHandler()
	if err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// could use /api here rather than /v1, I think is smarter to do long-term
			if strings.HasPrefix(r.URL.Path, "/v1") {
				gwmux.ServeHTTP(w, r)
				return
			}
			oa.ServeHTTP(w, r)
		}),
	}

	return fmt.Errorf("gRPC-Gateway server: %w", gwServer.ListenAndServe())
}
