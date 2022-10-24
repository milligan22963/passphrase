/*
  Copyright © 2022 DW Milligan dwm@afmsoftware.com
*/

// Package ppgen creates a passphrase from a list of words.
package ppgen

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"
	"time"

	"github.com/tyler-smith/go-bip39/wordlists"
)

// GeneratePhraseWords picks 'n' number of words at random. Function is deterministic for repeat uses of the same seed.
func GeneratePhraseWords(n int, seed string) ([]string, error) {
	if n < 1 {
		return nil, fmt.Errorf("number of words must be greater than 0, n = %d", n)
	}

	wordList := wordlists.English
	max := len(wordList)
	var words []string

	if len(seed) > 0 {
		seedHash, err := hash(seed)
		if err != nil {
			return nil, fmt.Errorf("failed to hash seed: %w", err)
		}

		rand.Seed(seedHash)
	} else {
		rand.Seed(time.Now().UnixNano())
	}

	// Initialize pseudorandom with time if no value is provided.
	for i := 0; i < n; i++ {
		wordIndex := rand.Intn(max) //nolint:gosec // we can probably ignore this as a security issue

		words = append(words, wordList[wordIndex])
	}

	return words, nil
}

// GeneratePassPhrase creates a string with `n` words, separated by `separator`.
func GeneratePassPhrase(n int, separator string, seed string) (string, error) {
	words, err := GeneratePhraseWords(n, seed)
	if err != nil {
		return "", err
	}

	return strings.Join(words, separator), nil
}

func hash(s string) (int64, error) {
	h := fnv.New64a()

	_, err := h.Write([]byte(s))
	if err != nil {
		return 0, err
	}

	return int64(h.Sum64()), nil
}
