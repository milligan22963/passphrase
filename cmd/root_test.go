/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/

// Package cmd_test contains all unit tests for the base application command.
package cmd_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/milligan22963/passphrase/cmd"
)

func TestRunRootCmd(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		assertion assert.ErrorAssertionFunc
		want      string
	}{
		{
			name:      "default",
			args:      []string{},
			assertion: assert.NoError,
			want:      "correct_horse_battery_staple",
		},
		{
			name:      "number flag too small",
			args:      []string{"-n=1"},
			assertion: assert.Error,
			want:      "invalid number of words:",
		},
		{
			name:      "invalid wordlist file",
			args:      []string{`-d="test.foo"`},
			assertion: assert.Error,
			want:      "bad word list:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cmd.GetRootCmd()

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
