package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minWindow(s string, t string) string {
	need := make([]int, 255)
	have := make([]int, 255)

	for i := range t {
		need[t[i]]++
	}

	var left, right int
	var count int
	var result string
	var min = len(s) + 1
	for right = 0; right < len(s); right++ {
		if have[s[right]] < need[s[right]] {
			count++
		}

		have[s[right]]++
		for count == len(t) {
			if min > right-left+1 {
				min = right - left + 1
				result = s[left : right+1]
			}

			have[s[left]]--
			if have[s[left]] < need[s[left]] {
				count--
			}

			left++
		}
	}

	return result
}

func Test(t *testing.T) {
	testCases := []struct {
		s    string
		t    string
		want string
	}{
		{
			s:    "DBACDAB",
			t:    "ABC",
			want: "BAC",
		},
		{
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
	}
	for i, tc := range testCases {
		got := minWindow(tc.s, tc.t)
		assert.Equalf(t, tc.want, got, "failed minWindow with i: %d", i)
	}
}
