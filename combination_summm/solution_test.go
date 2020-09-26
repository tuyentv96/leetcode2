package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findcombinationSum(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
	}

	for i := index; i < len(nums); i++ {
		if nums[i] > target {
			break
		}

		c = append(c, nums[i])
		findcombinationSum(nums, target-nums[i], i, c, res)
		c = c[:len(c)-1]
	}
}

func TestCombinationSum(t *testing.T) {
	testCases := []struct {
		arr  []int
		sum  int
		want int
	}{
		{
			arr:  []int{2, 3, 6, 7},
			sum:  7,
			want: 2,
		},
	}
	for i, tc := range testCases {
		got := combinationSum(tc.arr, tc.sum)
		assert.Equalf(t, tc.want, len(got), "failed CombinationSum with i: %d", i)
	}
}
