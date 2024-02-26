package validanagram

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"ac", "bb", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s-%s", test.s, test.t), func(t *testing.T) {
			if got := validAnagram(test.s, test.t); got != test.want {
				t.Error("validAnagram(", test.s, ",", test.t, ") got ", got, " and want ", test.want)
			}
		})
	}
}
