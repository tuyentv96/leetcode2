package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if complementIndex, ok := m[complement]; ok {
			return []int{complementIndex, i}
		}

		m[v] = i
	}

	return nil
}

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 11, 7, 5}, 9, []int{0, 2}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for i, tc := range testCases {
		got := twoSum(tc.nums, tc.target)
		assert.Equalf(t, tc.want, got, "failed twoSum with i: %d", i)
	}
}
