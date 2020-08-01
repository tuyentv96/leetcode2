package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func reverse(x int) int {
	//const MaxUint = ^uint32(0)
	const MaxInt = 2147483647
	const MinInt = -2147483648

	var res int
	for {
		var pop = x % 10
		x = x / 10

		if res > MaxInt/10 || (res == MaxInt/10 && pop > 7) {
			return 0
		}

		if res < MinInt/10 || (res == MinInt/10 && pop < -8) {
			return 0
		}

		res = (res * 10) + pop
		if x == 0 {
			break
		}
	}

	return res
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		s    int
		want int
	}{
		{
			s:    123,
			want: 321,
		},
		{
			s:    -123,
			want: -321,
		},
		{
			s:    15342364690,
			want: 0,
		},
	}

	for i, tc := range testCases {
		actual := reverse(tc.s)
		assert.Equalf(t, tc.want, actual, "failed TestLongestPalindrome i: %d", i)
	}
}
