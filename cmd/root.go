/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	core "github.com/akrck02/GTDF-CLI/api/core"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gtdf",
	Short: "GTD Framework CLI toolkit",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		//if flag version is set, print version and exit
		if version, _ := cmd.Flags().GetBool("version"); version {
			fmt.Println("GTDF-CLI Version : v" + core.CLI_VERSION)
			os.Exit(0)
		}

		fmt.Println(GTDF_MESSAGE)
		cmd.Help()
	},
}

var GTDF_MESSAGE = `

_____  _____  ____   _____    _____  _____  _____  _____  _____  _ _ _  _____  _____  _____ 
|   __||_   _||    \ |   __|  |   __|| __  ||  _  ||     ||   __|| | | ||     || __  ||  |  |
|  |  |  | |  |  |  ||   __|  |   __||    -||     || | | ||   __|| | | ||  |  ||    -||    -|
|_____|  |_|  |____/ |__|     |__|   |__|__||__|__||_|_|_||_____||_____||_____||__|__||__|__|
																								 

------------------------------------------------------------------------------------------------------
~ What is GTDF-CLI? 
------------------------------------------------------------------------------------------------------
GTDF-CLI is a command line tool to help you to use of the GTD Framework.

------------------------------------------------------------------------------------------------------
~ What are the commands I can use?
------------------------------------------------------------------------------------------------------
`

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// add flags to root command
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print version")

	// add commands to root command
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
