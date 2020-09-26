package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	c, res := []int{}, [][]int{}
	res = append(res, []int{})
	findsubsets(0, nums, c, &res)
	return res
}

func findsubsets(start int, nums []int, c []int, res *[][]int) {
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		c = append(c, nums[i])
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)

		findsubsets(i+1, nums, c, res)
		c = c[:len(c)-1]
	}
}

func TestSubsets(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 2, 2},
			want: 9,
		},
	}
	for i, tc := range testCases {
		got := subsetsWithDup(tc.nums)
		assert.Equalf(t, tc.want, len(got), "failed subset with i: %d", i)
	}
}
