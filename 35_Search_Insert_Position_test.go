package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func SearchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}

		if target < nums[i] {
			return i
		}

		if i == len(nums)-1 && target > nums[i] {
			return i + 1
		}

		if target > nums[i] && target < nums[i+1] {
			return i + 1
		}
	}

	return -1
}

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			arr:    []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
	}

	for i, tc := range testCases {
		actual := SearchInsert(tc.arr, tc.target)
		assert.Equalf(t, tc.want, actual, "failed TestSearchInRotatedArray i: %d", i)
	}
}
