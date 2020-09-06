package remove_duplicates_from_sorted_array_26_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	var prev = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if prev == nums[i] {
			prev = nums[i]
			nums = append(nums[0:i], nums[i+1:]...)
			continue
		}

		prev = nums[i]
	}
	return len(nums)
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
	}

	for i, tc := range testCases {
		got := removeDuplicates(tc.nums)
		assert.Equalf(t, tc.want, got, "failed removeDuplicates with i: %d", i)
	}
}
