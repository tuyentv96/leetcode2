package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func canJump(nums []int) bool {
	var maxReach int
	for i := 0; i < len(nums); i++ {
		if maxReach < i {
			return false
		}

		maxReach = max(maxReach, i+nums[i])
	}

	return true
}

func Test55(t *testing.T) {
	testCases := []struct {
		arr  []int
		want bool
	}{
		{
			arr:  []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			arr:  []int{3, 2, 1, 0, 4},
			want: false,
		},
	}

	for i, tc := range testCases {
		got := canJump(tc.arr)
		assert.Equalf(t, tc.want, got, "failed 55 with i: %d", i)
	}
}
