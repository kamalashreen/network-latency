package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "network-latency",
	Short: "network-latency demonstrates the number of bytes incurred by data sent using HTTP 1.1, HTTP 2 and gRPC protocols",
}

func Execute(versionInfo string) {
	rootCmd.Version = versionInfo

	rootCmd.AddCommand(&cobra.Command{
		Use:     "xml",
		Short:   "Display number of bytes with XML",
		Run: func(_ *cobra.Command, args []string) {

		},
	})

	_ = rootCmd.Execute()
}
