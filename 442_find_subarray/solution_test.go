package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findDuplicates(nums []int) []int {
	var res []int
	var sum int
	for i := 0; i < len(nums); i++ {
		sum ^= nums[i]
	}

	return res
}

func Test(t *testing.T) {
	testCases := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{3, 2},
		},
	}
	for i, tc := range testCases {
		got := findDuplicates(tc.arr)
		assert.Equalf(t, tc.want, got, "failed findDuplicates with i: %d", i)
	}
}
