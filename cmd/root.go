package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "skylark",
	Short: "skylark ",
	Long:  `skylark is a alert server`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
