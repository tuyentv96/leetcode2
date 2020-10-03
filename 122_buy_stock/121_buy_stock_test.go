package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var max int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			max += prices[i+1] - prices[i]
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
			want: 7,
		},
	}

	for i, tc := range testCases {
		got := maxProfit(tc.arr)
		assert.Equalf(t, tc.want, got, "failed maxProfit with i: %d", i)
	}
}
