/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/

// Package cmd contains all CLI commands used by the application.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/milligan22963/passphrase/pkg/ppgen"
)

var (
	rootCmd *cobra.Command

	// Used for flags
	separator   string
	phraseCount int
)

const (
	// MinimumWordCount is the minimum number of words in a passphrase.
	MinimumWordCount int = 4
)

const rootCommandLongDesc string = "passphrase is a password generator for " +
	"multi-word passphrases based on an XKCD comic (936)."

	// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "passphrase [flags]",
	Version: "1.0.0",
	Short:   "An app to generate a random passphrase",
	Long:    rootCommandLongDesc,
	Example: `passphrase --separator='-' --number=4`,
	Args:    cobra.NoArgs,
	PreRunE: ValidateFlags,
	RunE:    RunRootCmdE,
}

func init() {
	RootCmdFlags(rootCmd)
}

// RunRootCmdE is the main entry point for the root command.
func RunRootCmdE(cmd *cobra.Command, args []string) error {
	out, err := ppgen.GeneratePassPhrase(phraseCount, separator)
	if err != nil {
		return fmt.Errorf("failed to generate passphrase")
	}

	cmd.Println(out)

	return nil
}

// RootCmdFlags adds flags to the root command.
func RootCmdFlags(cmd *cobra.Command) {
	// main app params/flags
	cmd.Flags().IntVarP(&phraseCount, "number", "n", MinimumWordCount, "Number of words to include.")
	cmd.Flags().StringVarP(&separator, "separator", "s", "_", "Separator between words.")
}

// ValidateFlags checks that the flags are within expected boundaries.
func ValidateFlags(cmd *cobra.Command, args []string) error {
	if phraseCount < MinimumWordCount {
		return fmt.Errorf("invalid number of words: number = %d", phraseCount)
	}

	if len(separator) != 1 {
		return fmt.Errorf("separator must be a single-character string: separator = %s, length=%d", separator, len(separator))
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
