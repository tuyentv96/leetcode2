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

func jump(nums []int) int {
	if len(nums) > 0 && nums[0] == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	var maxReach int
	var step int
	for i := 0; i < len(nums); i++ {
		if maxReach == i {
			step++
		}

		maxReach = max(maxReach, i+nums[i])
	}

	return step
}

func Test45(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		//{
		//	arr:  []int{2, 3, 1, 1, 4},
		//	want: 2,
		//},
		//{
		//	arr:  []int{2, 1},
		//	want: 1,
		//},
		//{
		//	arr:  []int{0},
		//	want: 0,
		//},
		{
			arr:  []int{1, 2},
			want: 1,
		},
	}

	for i, tc := range testCases {
		got := jump(tc.arr)
		assert.Equalf(t, tc.want, got, "failed 45 with i: %d", i)
	}
}
