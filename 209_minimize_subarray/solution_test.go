package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minSubArrayLen(s int, nums []int) int {
	var sum int
	var min, l, r int
	for r = 0; r < len(nums); r++ {
		sum += nums[r]
		for l <= r && sum >= s {
			tmp := r - l + 1
			if min == 0 {
				min = tmp
			}

			if min > tmp {
				min = tmp
			}

			sum -= nums[l]
			l++
		}
	}

	return min
}

func Test(t *testing.T) {
	testCases := []struct {
		arr  []int
		s    int
		want int
	}{
		{
			arr:  []int{1, 4, 4},
			s:    4,
			want: 1,
		},
		{
			arr:  []int{2, 3, 1, 2, 4, 3},
			s:    7,
			want: 2,
		},
		{
			arr:  []int{1, 2, 3, 4, 5},
			s:    11,
			want: 3,
		},
	}
	for i, tc := range testCases {
		got := minSubArrayLen(tc.s, tc.arr)
		assert.Equalf(t, tc.want, got, "failed isPalindrome with i: %d", i)
	}
}
