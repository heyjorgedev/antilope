package cmd

import (
	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/jorgemurta/antilope/internal/config"
	"github.com/jorgemurta/antilope/internal/hetzner"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a server",
	Long:  `Delete a server`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		apiKey, _ := config.HetznerApiToken()
		client := hcloud.NewClient(hcloud.WithToken(apiKey))

		err := hetzner.DeleteServer(ctx, client, "hello");

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}