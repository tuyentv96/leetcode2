package climd_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func climbStairs(n int) int {
	dp := make([]int, n)
	if n <= 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    3,
			want: 3,
		},
	}

	for i, tc := range testCases {
		got := climbStairs(tc.n)
		assert.Equalf(t, tc.want, got, "failed climbStairs with i: %d", i)
	}
}
