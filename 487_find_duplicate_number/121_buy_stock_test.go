package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findDuplicate(nums []int) int {
	var i, j = 0, 0
	for true {
		i = nums[i]
		j = nums[nums[j]]
		if i == j {
			break
		}
	}

	i = 0
	for i != j {
		i = nums[i]
		j = nums[j]
	}

	return j
}

func Test(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		//{
		//	arr:  []int{1, 3, 4, 2, 2},
		//	want: 2,
		//},
		{
			arr:  []int{3, 1, 3, 4, 2},
			want: 3,
		},
	}

	for i, tc := range testCases {
		got := findDuplicate(tc.arr)
		assert.Equalf(t, tc.want, got, "failed findDuplicate with i: %d", i)
	}
}
