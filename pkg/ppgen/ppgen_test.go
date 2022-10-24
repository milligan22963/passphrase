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
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "expected default",
			args: args{
				n: 4,
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ppgen.GeneratePhraseWords(tt.args.n)
			if !assert.Equal(t, err != nil, tt.wantErr, "words: %v, error = %v, wantErr %v", got, err, tt.wantErr, tt.want) {
				return
			}
			assert.Equal(t, len(got), tt.want)
		})
	}
}
