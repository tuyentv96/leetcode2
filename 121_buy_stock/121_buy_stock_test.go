package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var min = prices[0]
	var max int

	for i := 1; i < len(prices); i++ {
		diff := prices[i] - min
		if diff > max {
			max = diff
		}

		if prices[i] < min {
			min = prices[i]
		}
	}

	return max
}

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
	}

	for i, tc := range testCases {
		got := maxProfit(tc.arr)
		assert.Equalf(t, tc.want, got, "failed maxProfit with i: %d", i)
	}
}
