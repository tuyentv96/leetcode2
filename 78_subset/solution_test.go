package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func subsets(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	res = append(res, []int{})
	findsubsets(0, nums, c, &res)
	return res
}

func findsubsets(start int, nums []int, c []int, res *[][]int) {
	for i := start; i < len(nums); i++ {
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
			nums: []int{1, 2, 3},
			want: 8,
		},
	}
	for i, tc := range testCases {
		got := subsets(tc.nums)
		assert.Equalf(t, tc.want, len(got), "failed subset with i: %d", i)
	}
}
