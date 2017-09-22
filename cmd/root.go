package cmd

import (
	"context"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/kevinjqiu/artillery-operator/pkg/artillery"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var cfg artillery.Config

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "artillery-operator",
	Short: "A brief description of your application",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info("Starting...")
		ctx, cancel := context.WithCancel(context.Background())
		wg, ctx := errgroup.WithContext(ctx)

		term := make(chan os.Signal)

		select {
		case <-term:
			log.Info("received SIGTERM, exiting gracefully...")
		case <-ctx.Done():
		}

		cancel()

		if err := wg.Wait(); err != nil {
			return err
		}
		return nil
	},
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
