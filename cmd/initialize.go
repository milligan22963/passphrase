/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	NaughtyWordRepo = "List-of-Dirty-Naughty-Obscene-and-Otherwise-Bad-Words"
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

	// Read in the language and hash the naughty words in the given language
	// we can then compare hashes to rule out specific words in the target country
	// though we may want to hash all languages and not offend anyone...if possible...

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initializeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initializeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
