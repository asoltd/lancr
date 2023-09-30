package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"github.com/asoltd/lancr/gateway"
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
	return gateway.Run(ctx, grpcServerAddr, gatewayAddr)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
