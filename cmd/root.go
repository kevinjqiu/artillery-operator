package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/kevinjqiu/artillery-operator/pkg/artillery"
)

cfg artillery.Config

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "artillery-operator",
	Short: "A brief description of your application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	flags := RootCmd.Flags()
	flags.StringVar(&cfg.Host, "apiserver", "", "The address of Kubernetes API server")
	flags.StringVar(&cfg.TLSConfig.CertFile, "cert-file", "", "Path to the client's public TLS cert file")
	flags.StringVar(&cfg.TLSConfig.KeyFile, "key-file", "", "Path to the client's private TLS cert file")
	flags.StringVar(&cfg.TLSConfig.CAFile, "ca-file", "", "Path to the TLS CA file")
}
