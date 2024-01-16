package find_the_difference

import (
	"fmt"
	"testing"
)

func TestFindTheDifference(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want byte
	}{
		{"abcd", "abcde", byte('e')},
		{"", "y", byte('y')},
		{"abdc", "abced", byte('e')},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s-%s", test.s, test.t), func(t *testing.T) {
			if got := findTheDifference(test.s, test.t); got != test.want {
				t.Errorf("findTheDifference() got %q and want %q", got, test.want)
			}
		})
	}
}
