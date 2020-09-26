package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minimumTotal2(triangle [][]int) int {
	dp := make([][]int, len(triangle)+1)
	for i := 0; i < len(triangle)+1; i++ {
		dp[i] = make([]int, len(triangle)+1)
	}

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}

	return dp[0][0]
}

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle)+1)
	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Test(t *testing.T) {
	testCases := []struct {
		arr  [][]int
		want int
	}{
		{
			arr: [][]int{
				{
					2,
				},
				{
					3, 4,
				},
				{
					6, 5, 7,
				},
				{
					4, 1, 8, 3,
				},
			},
			want: 11,
		},
	}
	for i, tc := range testCases {
		got := minimumTotal(tc.arr)
		assert.Equalf(t, tc.want, got, "failed with i: %d", i)
	}
}
