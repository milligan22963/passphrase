/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/

// Package ppgen creates a passphrase from a list of words.
package ppgen

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/tyler-smith/go-bip39/wordlists"
)

// GeneratePhraseWords picks 'n' number of words at random. Function is deterministic for repeat uses of the same seed.
func GeneratePhraseWords(n int) ([]string, error) {
	wordList := wordlists.English
	max := len(wordList)
	var words []string

	// Initialize pseudorandom with time if no value is provided.
	for i := 0; i < n; i++ {
		wordIndex, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			return nil, fmt.Errorf("failed to get word %d of %d: %w", i+1, n, err)
		}

		words = append(words, wordList[wordIndex.Int64()])
	}

	return words, nil
}

// GeneratePassPhrase creates a string with `n` words, separated by `separator`.
func GeneratePassPhrase(n int, separator string) (string, error) {
	words, err := GeneratePhraseWords(n)
	if err != nil {
		return "", fmt.Errorf("failed to generate passphrase with %d words, separator=\"%s\": %w", n, separator, err)
	}
	return strings.Join(words, separator), nil
}
