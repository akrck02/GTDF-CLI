/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/akrck02/GTDF-CLI/api"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new GTDF project.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")

		if len(args) < 2 {
			fmt.Println("Please provide a project name and a project url.")
			return
		}

		// url or empty string
		url := args[0]

		// name or empty string
		name := args[1]

		api.New_project(url, name)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
