package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func subarraySum(nums []int, k int) int {
	sums := make(map[int]int)
	sums[0] = 1
	sum := 0
	res := 0
	for _, num := range nums {
		sum += num
		sums[sum]++
		res += sums[sum-k]
	}

	return res
}

func TestSubarraySum(t *testing.T) {
	testCases := []struct {
		arr  []int
		k    int
		want int
	}{
		{
			arr:  []int{1, 2, 3},
			k:    3,
			want: 2,
		},
		{
			arr:  []int{1, 1, 1},
			k:    2,
			want: 2,
		},
	}
	for i, tc := range testCases {
		got := subarraySum(tc.arr, tc.k)
		assert.Equalf(t, tc.want, got, "failed subarraySum with i: %d", i)
	}
}
