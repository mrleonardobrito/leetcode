package merge_strings_alternately

import (
	"fmt"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s+%s", test.word1, test.word2), func(t *testing.T) {
			if got := mergeAlternately(test.word1, test.word2); got != test.want {
				t.Errorf("mergeAlternately() got %q, and want %q", got, test.want)
			}
		})
	}
}
