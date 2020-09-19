package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func uniquePathsOptimize(m int, n int) int {
	dp := make([]int, m)
	dp[0] = 1

	for i := 1; i < n; i++ {
		for i := 1; i < m; i++ {
			dp[i] += dp[i-1]
		}
	}

	return dp[m-1]
}

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		m    int
		n    int
		want int
	}{
		{
			m:    3,
			n:    2,
			want: 3,
		},
		//{
		//	m:    3,
		//	n:    7,
		//	want: 28,
		//},
	}

	for i, tc := range testCases {
		got := uniquePaths(tc.m, tc.n)
		assert.Equalf(t, tc.want, got, "failed uniquePaths with i: %d", i)
	}
}

func TestUniquePathsOptimize(t *testing.T) {
	testCases := []struct {
		m    int
		n    int
		want int
	}{
		{
			m:    3,
			n:    2,
			want: 3,
		},
		//{
		//	m:    3,
		//	n:    7,
		//	want: 28,
		//},
	}

	for i, tc := range testCases {
		got := uniquePathsOptimize(tc.m, tc.n)
		assert.Equalf(t, tc.want, got, "failed uniquePaths with i: %d", i)
	}
}
