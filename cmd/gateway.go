/*
Copyright © 2023 piotr.jp.ostrowski@gmail.com
*/
package cmd

import (
	"context"
	"io"
	"os"

	"github.com/asoltd/lancr/gateway"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/grpclog"
)

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Runs the gRPC gateway",
	Long: `Starts a gateway that runs on port 4200 and proxies the REST requests
	transcribed into gRPC to the gRPC server running on port 4201. Hitting the '/'
	route leads the user to the OpenAPI documentation.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
		grpclog.SetLoggerV2(log)

		dialAddress := cmd.Flags().Lookup("grpc-server-addr").Value.String()
		gatewayAddr := cmd.Flags().Lookup("gateway-addr").Value.String()
		log.Infof("Serving gRPC-Gateway and OpenAPI Documentation on https://%s", gatewayAddr)
		gw, err := gateway.New(ctx)
		if err != nil {
			log.Fatalf("Failed to create gateway: %v", err)
		}
		err = gw.Run(ctx, dialAddress, gatewayAddr)
		if err != nil {
			log.Fatalf("Failed to run the gateway: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(gatewayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gatewayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gatewayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
