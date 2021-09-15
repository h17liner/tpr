package cmd

import (
	"github.com/h17liner/tpr/pkg/server"

	"github.com/spf13/cobra"
)

var servCmd = &cobra.Command{
	Use:   "serv",
	Short: "Run tpr server.",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(servCmd)
}
