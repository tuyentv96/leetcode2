package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxSub(arr []int) int {
	if len(arr) < 1 {
		return 0
	}

	var right, left = len(arr) - 1, 0
	var sum int
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > arr[right] {
			right = i
		}

		sum = arr[right] - arr[left]

		if arr[right]-arr[i] > sum {
			left = i
		}
	}

	return arr[right] - arr[left]
}

func TestMaxSub(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{5, 1, 6, 4, 9, 7, 2},
			want: 8,
		},
		{
			arr:  []int{0, 7, 6, 4, 8, 7, 2},
			want: 8,
		},
		{
			arr:  []int{7},
			want: 0,
		},
		{
			arr:  []int{},
			want: 0,
		},
	}

	for i, tc := range testCases {
		got := maxSub(tc.arr)
		assert.Equalf(t, tc.want, got, "failed calculate twoSum with i: %d", i)
	}
}
