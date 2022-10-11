/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initializeCmd represents the initialize command
var initializeCmd = &cobra.Command{
	Use:   "initialize",
	Short: "Used to initialize the application for later usage.",
	Long: `Initializes the database that will be used later when
	generating the passphrases.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initialize called")
	},
}

func init() {
	rootCmd.AddCommand(initializeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initializeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initializeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
