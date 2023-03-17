package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Antilope",
	Long:  `All software has versions. This is Antilope's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Antilope version 0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}