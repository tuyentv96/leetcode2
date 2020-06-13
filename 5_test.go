package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestPalindrome(s string) string {
	var res string
	if len(s) < 3 {
		return res
	}

	var s1, s2 string
	var s1i, s2i int
	s1i = 0
	s2i = 2
	s1 = string(s[s1i])
	s2 = string(s[s2i])

	for i := 0; i < len(s); i++ {
		if s1 == s2 {
			return s[s1i : s2i+1]
		}
	}

	return res
}

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "babad",
			want: "bab",
		},
		{
			s:    "deabcdcbaef",
			want: "bcdcb",
		},
	}

	for i, tc := range testCases {
		actual := longestPalindrome(tc.s)
		assert.Equalf(t, tc.want, actual, "failed TestLongestPalindrome i: %d", i)
	}
}
