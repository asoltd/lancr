/*
Copyright © 2023 piotr.jp.ostrowski@gmail.com
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lancr",
	Short: "Executable for managing the services of lancr backend",
	Long:  `Longer description will be added later`,
	Run: func(cmd *cobra.Command, args []string) {
		// print help
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String(
		"gateway-addr",
		"0.0.0.0:4200",
		"Address of the gRPC gateway",
	)
	rootCmd.PersistentFlags().String(
		"grpc-server-addr",
		"0.0.0.0:4201",
		"Address of the gRPC server",
	)
}
