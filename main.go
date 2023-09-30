package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"github.com/asoltd/lancr/gen"
	lancrv1 "github.com/asoltd/lancr/gen/go/proto/lancr/v1"
	"github.com/asoltd/lancr/server"
)

// TODO's
// - [ ] Add TLS certs for proxy to server
// - [ ] Move config to environment variables / YAML + viper
// - [ ] Move command-line flags to cobra
// - [ ] Separate the backend and frontend and while sticking to monorepo
// - [ ] Perfect the schema mapping between the Firestore existing types and the
// protobufs declarations TODO(oliwierost)
// - [ ] Add Authentication
// - [ ] Add Authorization
// - [ ] Add CI
// - [ ] Deploy, add CD
// - [ ] Implement the remainder of the CRUD operations for all of the items
// - [ ] Ensure idempotency of the application (no double withdrawals),
// transactional support should be enabled by default in Firestore since it is
// based in Spanner

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

func runGateway(ctx context.Context, dialAddr string, gatewayAddr string) error {
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

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	grpcServerAddr := "0.0.0.0:4201"
	gatewayAddr := "0.0.0.0:4200"

	lis, err := net.Listen("tcp", grpcServerAddr)
	if err != nil {
		return fmt.Errorf("failed to listen %w", err)
	}

	s := grpc.NewServer()
	backend, err := server.New(
		"./asoltd-lancr-firebase-admin-service-account.json",
	)
	if err != nil {
		return err
	}
	lancrv1.RegisterHeroServiceServer(s, backend)

	// register reflection service on gRPC server
	reflection.Register(s)

	log.Infof("Serving gRPC on https://%s", grpcServerAddr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Infof("Serving gRPC-Gateway and OpenAPI Documentation on https://%s", gatewayAddr)
	return runGateway(ctx, grpcServerAddr, gatewayAddr)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
