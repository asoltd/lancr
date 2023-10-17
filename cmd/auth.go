/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"net"
	"os"

	"github.com/asoltd/lancr/auth"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Run the AuthService gRPC Server",
	Run: func(cmd *cobra.Command, args []string) {
		log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
		grpclog.SetLoggerV2(log)

		authServiceAddr := cmd.Flags().Lookup("auth-service-addr").Value.String()

		lis, err := net.Listen("tcp", authServiceAddr)
		if err != nil {
			log.Fatalf("failed to listen %w", err)
		}

		a, err := auth.NewAuthService(cmd.Context())
		if err != nil {
			log.Fatalf("failed to create auth service %v", err)
		}

		s := grpc.NewServer()

		lancrv1.RegisterAuthServiceServer(s, a)

		reflection.Register(s)

		log.Infof("Serving gRPC on https://%s", authServiceAddr)
		err = s.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
