package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	var res int
	var prd = 1
	var left int

	if k <= 1 {
		return 0
	}

	for right := 0; right < len(nums); right++ {
		prd *= nums[right]
		for prd >= k {
			prd /= nums[left]
			left++
		}

		res += right - left + 1
	}

	return res
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	testCases := []struct {
		arr  []int
		k    int
		want int
	}{
		{
			arr:  []int{10, 5, 2, 6},
			k:    100,
			want: 8,
		},
	}
	for i, tc := range testCases {
		got := numSubarrayProductLessThanK(tc.arr, tc.k)
		assert.Equalf(t, tc.want, got, "failed numSubarrayProductLessThanK with i: %d", i)
	}
}
