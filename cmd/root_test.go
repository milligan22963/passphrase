/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/

// Package cmd_test contains all unit tests for the base application command.
package cmd_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/milligan22963/passphrase/cmd"
)

func TestInputValidationErrors(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		assertion assert.ErrorAssertionFunc
		want      string
	}{
		{
			name:      "number flag too small",
			args:      []string{`-n=1`},
			assertion: assert.Error,
			want:      "invalid number of words:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cmd.NewRootCmd()

			b := bytes.NewBufferString("")
			c.SetOut(b)
			c.SetErr(b)
			c.SetArgs(tt.args)

			tt.assertion(t, c.Execute())

			out, err := io.ReadAll(b)
			require.NoError(t, err)

			assert.Contains(t, string(out), tt.want)
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
			name:      "long separator",
			args:      []string{`-s=FOO`},
			assertion: assert.NoError,
			want: wants{
				n: 4,
				s: "FOO",
			},
		},
		{
			name:      "explicit defaults",
			args:      []string{`-n=4`, `-s="_"`},
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
			name:      "dash separator",
			args:      []string{`-s="-"`},
			assertion: assert.NoError,
			want: wants{
				n: 4,
				s: "-",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cmd.NewRootCmd()

			b := bytes.NewBufferString("")
			c.SetOut(b)
			c.SetErr(b)
			c.SetArgs(tt.args)

			tt.assertion(t, c.Execute())

			out, err := io.ReadAll(b)
			require.NoError(t, err)

			assert.Len(t, strings.Split(string(out), tt.want.s), tt.want.n)
		})
	}
}

func TestIntegrationWithDefaults(t *testing.T) {
	c := cmd.NewRootCmd()

	b := bytes.NewBufferString("")
	c.SetOut(b)
	c.SetErr(b)
	c.SetArgs([]string{})

	assert.NoError(t, c.Execute())

	out, err := io.ReadAll(b)
	require.NoError(t, err)

	assert.Len(t, strings.Split(string(out), "_"), 4)
}

func ExecuteCommandC(t *testing.T, root *cobra.Command, args ...string) (*cobra.Command, string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err := root.ExecuteC()

	return c, buf.String(), err
}
