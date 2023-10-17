/*
Copyright © 2023 piotr.jp.ostrowski@gmail.com
*/
package cmd

import (
	"io"
	"net"
	"os"

	"github.com/asoltd/lancr/db"
	"github.com/asoltd/lancr/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs the gRPC server",
	Long:  `TODO(piotrostr): provide a longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
		grpclog.SetLoggerV2(log)

		grpcServerAddr := cmd.Flags().Lookup("grpc-server-addr").Value.String()
		authServiceAddr := cmd.Flags().Lookup("auth-service-addr").Value.String()
		local := cmd.Flags().Lookup("local").Value.String()

		lis, err := net.Listen("tcp", grpcServerAddr)
		if err != nil {
			log.Fatalf("failed to listen %w", err)
		}

		s := grpc.NewServer()

		var gormDB *gorm.DB
		if local == "true" {
			gormDB, err = server.ConnectTestDB()
			if err != nil {
				log.Fatalf("failed to connect to local DB %w", err)
			}
		} else {
			gormDB, err = db.Connect()
			if err != nil {
				log.Fatalf("failed to connect to TiDB %w", err)
			}
		}

		conn, err := grpc.DialContext(
			cmd.Context(),
			authServiceAddr,
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)
		if err != nil {
			log.Fatalf("failed to connect to AuthService %w", err)
		}

		auth := lancrv1.NewAuthServiceClient(conn)

		h := server.NewHeroServiceServer(gormDB, auth)
		a := server.NewApprenticeServiceServer(gormDB)
		q := server.NewQuestServiceServer(gormDB)
		t := server.NewTeamsServiceServer(gormDB)

		lancrv1.RegisterHeroServiceServer(s, h)
		lancrv1.RegisterApprenticeServiceServer(s, a)
		lancrv1.RegisterQuestServiceServer(s, q)
		lancrv1.RegisterTeamsServiceServer(s, t)

		reflection.Register(s)

		log.Infof("Serving gRPC on https://%s", grpcServerAddr)
		err = s.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
