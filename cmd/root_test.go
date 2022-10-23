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
		{
			name:          "validation passes - explicit defaults",
			args:          []string{"-n=4", "-s=_"},
			assertion:     assert.NoError,
			wantErrString: "",
		},
		{
			name:          "number flag too small",
			args:          []string{`-n=1`},
			assertion:     assert.Error,
			wantErrString: "invalid number of words:",
		},
		{
			name:          "separator flag too long",
			args:          []string{`-s=abcd`},
			assertion:     assert.Error,
			wantErrString: "separator must be a single-character string:",
		},
		{
			name:          "invalid emoji separator",
			args:          []string{"-s=🐈"},
			assertion:     assert.Error,
			wantErrString: "separator must be a single-character string:",
		},
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
		{
			name:      "letter separator",
			args:      []string{`-s=F`},
			assertion: assert.NoError,
			want: wants{
				n: 4,
				s: "F",
			},
		},
		{
			name:      "explicit defaults",
			args:      []string{`-n=4`, `-s=_`},
			assertion: assert.NoError,
			want: wants{
				n: 4,
				s: "_",
			},
		},
		{
			name:      "big phrase",
			args:      []string{`-n=42`},
			assertion: assert.NoError,
			want: wants{
				n: 42,
				s: "_",
			},
		},
		{
			name:      "numeric separator",
			args:      []string{`-s=3`},
			assertion: assert.NoError,
			want: wants{
				n: 4,
				s: "3",
			},
		},
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
