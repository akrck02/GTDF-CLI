package cmd

import (
	"github.com/akrck02/GTDF-CLI/api/core"
	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create a service.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		core.NewService(args[0])
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
