package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findLength(A []int, B []int) int {
	var max int
	dp := make([][]int, len(A)+1)
	for i := range B {
		dp[i] = make([]int, len(B)+1)
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				dp[i][j] += 1
				if i > 0 && j > 0 {
					dp[i][j] += dp[i-1][j-1]
				}
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}

	return max
}

func Test(t *testing.T) {
	testCases := []struct {
		a    []int
		b    []int
		want int
	}{
		{
			a:    []int{1, 2, 3, 2, 1},
			b:    []int{3, 2, 1, 4, 7},
			want: 3,
		},
	}
	for i, tc := range testCases {
		got := findLength(tc.a, tc.b)

		assert.Equalf(t, tc.want, got, "failed findLength with i: %d", i)
	}
}
