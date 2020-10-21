package solution

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var min = 9999999
	var res int
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if min >= abs(sum, target) {
				res = sum
				min = abs(sum, target)
			}

			if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
func Test(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			arr:    []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
	}
	for i, tc := range testCases {
		got := threeSumClosest(tc.arr, tc.target)
		assert.Equalf(t, tc.want, got, "failed threeSumClosest with i: %d", i)
	}
}
