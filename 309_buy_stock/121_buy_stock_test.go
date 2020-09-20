package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

//func maxProfit(prices []int) int {
//	if len(prices) <= 1 {
//		return 0
//	}
//	buy, sell := make([]int, len(prices)), make([]int, len(prices))
//	for i := range buy {
//		buy[i] = math.MinInt64
//	}
//	buy[0] = -prices[0]
//	buy[1] = max(buy[0], -prices[1])
//	sell[1] = max(sell[0], buy[0]+prices[1])
//
//	for i := 2; i < len(prices); i++ {
//		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
//		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
//	}
//	//for i := 0; i < len(prices)-1; i++ {
//	//	if prev-prices[i] > stock {
//	//		stock = stock - prices[i]
//	//	}
//	//
//	//	prev = money
//	//
//	//	if stock+prices[i] > money {
//	//		money = stock + prices[i]
//	//	}
//	//}
//
//	return sell[len(sell)-1]
//	//return money
//}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var buy, sell int
	buy = -prices[0]
	var prevSell int

	for i := 1; i < len(prices); i++ {
		//buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		buy = max(buy, prevSell-prices[i])

		prevSell = sell

		//sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		sell = max(sell, buy+prices[i])
	}
	return sell
}

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{1, 2, 3, 0, 2},
			want: 3,
		},
		{
			arr:  []int{2, 1},
			want: 0,
		},
	}

	for i, tc := range testCases {
		got := maxProfit(tc.arr)
		assert.Equalf(t, tc.want, got, "failed maxProfit with i: %d", i)
	}
}
