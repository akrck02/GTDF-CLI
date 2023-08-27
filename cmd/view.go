package cmd

import (
	"github.com/akrck02/GTDF-CLI/api/core"
	"github.com/akrck02/GTDF-CLI/api/logger"
	middleware "github.com/akrck02/GTDF-CLI/api/middleware"
	"github.com/spf13/cobra"
)

// array of functions to be called before the command is executed
var viewMiddlewares = []func(cmd *cobra.Command, args []string){
	middleware.ProjectCheck,
}

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Create a view.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			logger.Error("You must provide a name for the view.")
			return
		}

		delete, _ := cmd.Flags().GetBool("delete")

		if delete {
			core.DeleteView(args[0])
			return
		}

		core.NewView(args[0])
	},
}

func init() {

	// add middlewares to the command
	for _, middleware := range viewMiddlewares {
		viewCmd.PreRun = middleware
	}

	// add command to root command
	rootCmd.AddCommand(viewCmd)
	viewCmd.PersistentFlags().BoolP("delete", "d", false, "Delete the view.")
}
