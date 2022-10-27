/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/

// Package cmd_test contains all unit tests for the base application command.
package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/milligan22963/passphrase/cmd"
)

func TestInputValidationErrors(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		assertion     assert.ErrorAssertionFunc
		wantErrString string
	}{
		{"validation passes - explicit defaults", []string{"-n=4", "-s=_"}, assert.NoError, ""},
		{"number flag too small", []string{`-n=1`}, assert.Error, "invalid number of words:"},
		{"separator flag too long", []string{`-s=abcd`}, assert.Error, "separator must be a single-character string:"},
		{"invalid emoji separator", []string{"-s=🐈"}, assert.Error, "separator must be a single-character string:"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &cobra.Command{Use: "root", PreRunE: cmd.ValidateFlags, RunE: cmd.RunRootCmdE}
			cmd.RootCmdFlags(root)

			out, err := execute(t, root, tt.args...)
			tt.assertion(t, err)

			if err != nil {
				assert.Contains(t, out, tt.wantErrString)
			}
		})
	}
}

func TestIntegrationWithParams(t *testing.T) {
	type wants struct {
		n int
		s string
	}

	tests := []struct {
		name      string
		args      []string
		assertion assert.ErrorAssertionFunc
		want      wants
	}{
		{"big phrase", []string{`-n=42`}, assert.NoError, wants{42, "_"}},
		{"numeric separator", []string{`-s=3`}, assert.NoError, wants{4, "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &cobra.Command{Use: "root", PreRunE: cmd.ValidateFlags, RunE: cmd.RunRootCmdE}
			cmd.RootCmdFlags(root)

			out, err := execute(t, root, tt.args...)
			tt.assertion(t, err)

			if err == nil {
				assert.Len(t, strings.Split(out, tt.want.s), tt.want.n)
			}
		})
	}
}

func TestIntegrationWithDefaults(t *testing.T) {
	root := &cobra.Command{Use: "root", PreRunE: cmd.ValidateFlags, RunE: cmd.RunRootCmdE}
	cmd.RootCmdFlags(root)

	got, err := execute(t, root)

	assert.NoError(t, err)

	assert.Len(t, strings.Split(got, "_"), 4)
}

func execute(t *testing.T, cmd *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	err := cmd.Execute()

	return strings.TrimSpace(buf.String()), err
}
