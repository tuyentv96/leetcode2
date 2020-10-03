package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isAlphaNumeric(s uint8) bool {
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
		return true
	}

	return false
}

func toLower(s uint8) uint8 {
	if s >= 'A' && s <= 'Z' {
		return s - 'A' + 'a'
	}

	return s
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	var left, right = 0, len(s) - 1
	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}

		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func Test(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
	}{
		{
			s:    ".,",
			want: true,
		},
		{
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			s:    "abba",
			want: true,
		},
		{
			s:    "abc",
			want: false,
		},
		{
			s:    "ab ba",
			want: true,
		},
	}
	for i, tc := range testCases {
		got := isPalindrome(tc.s)
		assert.Equalf(t, tc.want, got, "failed isPalindrome with i: %d", i)
	}
}
