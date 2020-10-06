package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeDuplicates(nums []int) int {
	var begin int
	for i := 1; i < len(nums); i++ {
		if begin == 0 || nums[i] != nums[begin] || nums[i] != nums[begin-1] {
			begin++
			nums[begin], nums[i] = nums[i], nums[begin]
		}
	}

	return begin + 1
}

func TestExist(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		//{
		//	arr:  []int{1, 1, 1, 2, 2, 3},
		//	want: 5,
		//},
		{
			arr:  []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: 7,
		},
		//{
		//	arr:  []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
		//	want: 7,
		//},
	}
	for i, tc := range testCases {
		got := removeDuplicates(tc.arr)
		assert.Equalf(t, tc.want, got, "failed removeDuplicates with i: %d", i)
	}
}
