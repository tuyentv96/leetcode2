package solution_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	//dp := make([]int, len(nums))
	//dp[0] = nums[0]
	//dp[1] = max(nums[0], nums[1])
	var prev int
	var cur int
	var tmp int
	//if nums[1] > nums[0] {
	//	cur = nums[1]
	//	prev = nums[0]
	//} else {
	//	cur = nums[0]
	//	prev = nums[1]
	//}

	for i := 0; i < len(nums); i++ {
		//prev = dp[i-1]
		//dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		tmp = cur
		cur = max(cur, prev+nums[i])
		prev = tmp
	}

	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestRob(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{1, 2, 3, 1},
			want: 4,
		},
		{
			arr:  []int{2, 7, 9, 3, 1},
			want: 12,
		},
	}

	for i, tc := range testCases {
		got := rob2(tc.arr)
		assert.Equalf(t, tc.want, got, "failed rob with i: %d", i)
	}
}
