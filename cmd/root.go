package cmd

import (
	"github.com/kamalashreen/network-latency/protocol"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "network-latency",
	Short: "network-latency demonstrates the number of bytes incurred by data sent using HTTP 1.1, HTTP 2 and gRPC protocols",
}

// Execute defines root commands for different network protocols
func Execute(versionInfo string) {
	rootCmd.Version = versionInfo

	rootCmd.AddCommand(&cobra.Command{
		Use:   "http",
		Short: "Display number of bytes of XML document using HTTP 1.1 protocol",
		Run: func(_ *cobra.Command, args []string) {
			protocol.HTTP()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "http2",
		Short: "Display number of bytes of XML document using HTTP 2 protocol with binary encoding",
		Run: func(_ *cobra.Command, args []string) {
			protocol.HTTP2()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "grpc",
		Short: "Display number of bytes of Protocol buffer using gRPC protocol",
		Run: func(_ *cobra.Command, args []string) {
			protocol.GRPC()
		},
	})

	_ = rootCmd.Execute()
}
