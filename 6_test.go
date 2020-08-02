package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func countSubstrings(s string) int {
	var res int
	for i := 0; i < len(s); i++ {

	}

	return res
}

func TestCountSubstrings(t *testing.T) {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "babad",
			want: 5,
		},
		{
			s:    "aabbcc",
			want: 9,
		},
	}

	for i, tc := range testCases {
		actual := longestPalindrome(tc.s)
		assert.Equalf(t, tc.want, actual, "failed TestLongestPalindrome i: %d", i)
	}
}
