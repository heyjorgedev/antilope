package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "antilope",
	Short: "Antilope is a CLI tool to manage your projects",
	Long: "Antilope is a CLI tool to manage your projects.",
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		panic(err)
	}
}