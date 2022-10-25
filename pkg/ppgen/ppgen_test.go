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
		{"expected default", args{4, "1"}, []string{"alcohol", "lucky", "draw", "ghost"}, assert.NoError},
		{"basic with custom seed", args{4, "cat"}, []string{"sudden", "dentist", "clerk", "practice"}, assert.NoError},
		{"negative number of words", args{-1, ""}, []string{}, assert.Error},
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

func TestGeneratePhrase(t *testing.T) {
	type args struct {
		n    int
		s    string
		seed string
	}

	tests := []struct {
		name      string
		args      args
		want      string
		assertion assert.ErrorAssertionFunc
	}{
		{"expected default", args{4, "_", "1"}, "alcohol_lucky_draw_ghost", assert.NoError},
		{"basic with custom seed", args{4, "_", "cat"}, "sudden_dentist_clerk_practice", assert.NoError},
		{"negative number of words", args{-1, "_", ""}, "", assert.Error},
		{"zero words", args{0, "_", ""}, "", assert.Error},
		{"empty separator", args{4, "", "cat"}, "suddendentistclerkpractice", assert.NoError},
		{"numeric separator", args{4, "3", "cat"}, "sudden3dentist3clerk3practice", assert.NoError},
		{"emoji separator", args{4, "🐈", "cat"}, "sudden🐈dentist🐈clerk🐈practice", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ppgen.GeneratePassPhrase(tt.args.n, tt.args.s, tt.args.seed)
			tt.assertion(t, err)

			if err == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
