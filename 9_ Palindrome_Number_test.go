package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var revertedNumber int
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x = x / 10
	}

	if x == revertedNumber || x == revertedNumber/10 {
		return true
	}

	return false
}

func TestPalindromeNumber(t *testing.T) {
	testCases := []struct {
		s    int
		want bool
	}{
		{
			s:    10,
			want: false,
		},
		{
			s:    1221,
			want: true,
		},
		{
			s:    121,
			want: true,
		},
		{
			s:    1212,
			want: false,
		},
		{
			s:    12,
			want: false,
		},
	}

	for i, tc := range testCases {
		actual := isPalindrome(tc.s)
		assert.Equalf(t, tc.want, actual, "failed TestLongestPalindrome i: %d", i)
	}
}
