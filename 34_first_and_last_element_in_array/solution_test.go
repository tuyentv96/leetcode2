package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func Test(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   []int
	}{
		{
			arr:    []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
	}
	for i, tc := range testCases {
		got := searchRange(tc.arr, tc.target)
		assert.Equalf(t, tc.want, got, "failed isPalindrome with i: %d", i)
	}
}
