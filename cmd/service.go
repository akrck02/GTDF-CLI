package cmd

import (
	"github.com/akrck02/GTDF-CLI/api/core"
	"github.com/akrck02/GTDF-CLI/api/logger"
	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create a service.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			logger.Error("You must provide a name for the service.")
			return
		}

		delete, _ := cmd.Flags().GetBool("delete")

		if delete {
			core.DeleteService(args[0])
			return
		}

		core.NewService(args[0])
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.PersistentFlags().BoolP("delete", "d", false, "Delete the service.")
}
