package solution_test

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	var ret [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++

				// increase cause to duplicate
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return ret
}

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for i, tc := range testCases {
		got := threeSum(tc.nums)
		assert.Equalf(t, tc.want, got, "failed threeSum with i: %d", i)
	}
}
