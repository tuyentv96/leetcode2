package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func max2nd(arr []int) int {
	var max, max2nd int
	for i := range arr {
		if arr[i] < arr[max2nd] {
			if arr[i] > arr[max] {
				max2nd = max
				max = i
			} else {
				max2nd = i
			}
		}
	}

	return arr[max2nd]
}

func TestMax2nd(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{5, 1, 6, 4, 9, 7, 2},
			want: 7,
		},
	}

	for i, tc := range testCases {
		got := max2nd(tc.arr)
		assert.Equalf(t, tc.want, got, "failed calculate twoSum with i: %d", i)
	}
}
