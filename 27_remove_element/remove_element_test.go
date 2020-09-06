package solution_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeElement(nums []int, val int) int {

}

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
	}

	for i, tc := range testCases {
		got := removeDuplicates(tc.nums)
		assert.Equalf(t, tc.want, got, "failed removeDuplicates with i: %d", i)
	}
}
