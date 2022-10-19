/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/

// Package cmd contains all CLI commands used by the application.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags
	dict        string
	separator   string
	phraseCount int
)

const (
	// MinimumWordCount is the minimum number of words in a passphrase.
	MinimumWordCount int = 4
)

const rootCommandLongDesc string = "passphrase is a password generator for " +
	"multi-word passphrases based on an XKCD comic (936). The list of seed words " +
	"can be customized by the user and length of password modified as needed."

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "passphrase [flags]",
	Version:           "1.0.0",
	Short:             "An app to generate a random passphrase",
	Long:              rootCommandLongDesc,
	Example:           `passphrase --separator='-' --number=4 --dictionary="./wordlist.txt"`,
	Args:              cobra.NoArgs,
	PersistentPreRunE: ValidateFlags,
	RunE:              RunRootCmd,
}

func init() {
	// global params/flags
	rootCmd.PersistentFlags().StringVarP(&dict, "dictionary", "d", "./wordlist.txt", "Path to word source (text file).")

	// main app params/flags
	rootCmd.Flags().IntVarP(&phraseCount, "number", "n", MinimumWordCount, "Number of words to include.")
	rootCmd.Flags().StringVarP(&separator, "separator", "s", "_", "Separator between words.")
}

// ValidateFlags checks that the flags are within expected boundaries.
func ValidateFlags(cmd *cobra.Command, args []string) error {
	if _, err := os.Stat(dict); err != nil {
		return fmt.Errorf("bad word list: %w", err)
	}

	if phraseCount < MinimumWordCount {
		return fmt.Errorf("invalid number of words: number = %d", phraseCount)
	}

	return nil
}

// GetRootCmd gets the application root command.
func GetRootCmd() *cobra.Command {
	return rootCmd
}

// RunRootCmd is executed when the application is run without any subcommands.
func RunRootCmd(cmd *cobra.Command, args []string) error {
	cmd.Printf("passphrase called: dictionary=%s; number=%d; separator=%s\n", dict, phraseCount, separator)
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
