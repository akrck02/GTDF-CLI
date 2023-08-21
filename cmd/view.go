package cmd

import (
	"github.com/akrck02/GTDF-CLI/api/core"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Create a view.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		core.NewView(args[0])
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
