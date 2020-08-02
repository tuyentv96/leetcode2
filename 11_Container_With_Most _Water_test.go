package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxArea(height []int) int {
	var left, right = 0, len(height) - 1
	var max = 0
	for left < right {
		var y int
		if height[left] < height[right] {
			y = height[left]
		} else {
			y = height[right]
		}

		if y*(right-left) > max {
			max = y * (right - left)
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want: 49,
		},
	}

	for i, tc := range testCases {
		actual := maxArea(tc.arr)
		assert.Equalf(t, tc.want, actual, "failed TestLongestPalindrome i: %d", i)
	}
}
