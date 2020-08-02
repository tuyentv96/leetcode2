package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func SearchInRotatedArray(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] != target {
			return -1
		}
		return 0
	}
	return search1(nums, 0, length-1, target)
}

func search1(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	if left == right {
		if nums[left] == target {
			return left
		}
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}

	if nums[left] < nums[mid] && nums[left] <= target && nums[mid] > target {
		return binarySearch(nums, left, mid-1, target)
	}
	if nums[right] > nums[mid] && nums[mid] < target && nums[right] >= target {
		return binarySearch(nums, mid+1, right, target)
	}
	if nums[left] > nums[mid] {
		return search1(nums, left, mid-1, target)
	}
	if nums[right] < nums[mid] {
		return search1(nums, mid+1, right, target)
	}
	return -1
}

func binarySearch(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}
	if target < nums[mid] {
		return binarySearch(nums, left, mid-1, target)
	}
	return binarySearch(nums, mid+1, right, target)

}

func TestSearchInRotatedArray(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			arr:    []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
	}

	for i, tc := range testCases {
		actual := SearchInRotatedArray(tc.arr, tc.target)
		assert.Equalf(t, tc.want, actual, "failed TestSearchInRotatedArray i: %d", i)
	}
}
