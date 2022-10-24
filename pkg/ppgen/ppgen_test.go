/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/

// Package ppgen_test contains unit tests for ppgen package.
package ppgen_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/milligan22963/passphrase/pkg/ppgen"
)

func TestGeneratePhraseWords(t *testing.T) {
	type args struct {
		n    int
		seed string
	}

	tests := []struct {
		name      string
		args      args
		want      []string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "expected default",
			args: args{
				n:    4,
				seed: "1",
			},
			want:      []string{"alcohol", "lucky", "draw", "ghost"},
			assertion: assert.NoError,
		},
		{
			name: "basic with custom seed",
			args: args{
				n:    4,
				seed: "cat",
			},
			want:      []string{"sudden", "dentist", "clerk", "practice"},
			assertion: assert.NoError,
		},
		{
			name: "negative number of words",
			args: args{
				n:    -1,
				seed: "",
			},
			want:      []string{},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ppgen.GeneratePhraseWords(tt.args.n, tt.args.seed)
			tt.assertion(t, err)

			if err == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
