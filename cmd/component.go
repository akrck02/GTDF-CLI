/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Create a component.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("component called")
	},
}

func init() {
	rootCmd.AddCommand(componentCmd)
}
