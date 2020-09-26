package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func combinationSum3(k int, n int) [][]int {
	c, res := []int{}, [][]int{}
	findcombinationSum(0, 1, k, n, c, &res)
	return res
}

func findcombinationSum(level int, start int, k int, n int, c []int, res *[][]int) {
	if level == k || n < 0 {
		if n == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}

		return
	}

	for i := start; i <= 9; i++ {
		if i > n {
			break
		}

		c = append(c, i)
		findcombinationSum(level+1, i+1, k, n-i, c, res)
		c = c[:len(c)-1]
	}
}

func TestCombinationSum(t *testing.T) {
	testCases := []struct {
		k    int
		n    int
		want int
	}{
		{
			k:    3,
			n:    9,
			want: 3,
		},
	}
	for i, tc := range testCases {
		got := combinationSum3(tc.k, tc.n)
		assert.Equalf(t, tc.want, len(got), "failed CombinationSum with i: %d", i)
	}
}
