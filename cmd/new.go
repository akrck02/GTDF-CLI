package cmd

import (
	"fmt"

	"github.com/akrck02/GTDF-CLI/api/core"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new GTDF project.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Please provide a project name and a project url.")
			return
		}

		// url or empty string
		url := "./"

		// name or empty string
		name := args[0]

		core.NewProject(url, name)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
