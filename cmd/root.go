/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "passphrase",
	Short: "An app to generate a random passphrase",
	Long: `An application that can generate a random passphrase to be used
	as a passphrase for various applications or websites.
	
For example:

	passphrase --separator '-' --number 4 --language en`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("passphrase called")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var language = "en"
var separator = ""
var phraseCount = 4

func init() {
	// global params/flags
	rootCmd.PersistentFlags().StringVar(&language, "language", "l", "language to utilize such as en, it, etc")

	// main app params/flags
	rootCmd.Flags().IntVar(&phraseCount, "number", 4, "Number of phrases to include")
	rootCmd.Flags().StringVar(&separator, "separator", "", "Separator between phrases")
}
