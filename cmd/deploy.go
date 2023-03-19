package cmd

import (
	"fmt"
	"log"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/jorgemurta/antilope/internal/config"
	"github.com/jorgemurta/antilope/internal/hetzner"
	"github.com/jorgemurta/antilope/internal/server"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new version of the application",
	Long:  `Deploy a new version of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		apiKey, _ := config.HetznerApiToken()
		client := hcloud.NewClient(hcloud.WithToken(apiKey))

		srv, err := hetzner.EnsureServerExists(ctx, client, "hello");
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(srv.IpAddress)

		err = server.EnsureSoftwareIsInstalled(srv)
		if err != nil {
			panic(err)
		}

		// Copy files into the machine excluding the gitignore files

		// Bootstrap Spryker Docker

		// Start Spryker Docker
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}