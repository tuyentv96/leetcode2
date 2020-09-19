package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])

	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
				continue
			}

			if i == 0 {
				dp[0][j] = dp[0][j-1] + grid[i][j]
				continue
			}

			if j == 0 {
				dp[i][0] = dp[i-1][0] + grid[i][j]
				continue
			}

			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}

func TestMinPathSum(t *testing.T) {
	testCases := []struct {
		arr  [][]int
		want int
	}{
		{
			arr: [][]int{
				{
					1, 3, 1,
				},
				{
					1, 5, 1,
				},
				{
					4, 2, 1,
				},
			},
			want: 7,
		},
	}

	for i, tc := range testCases {
		got := minPathSum(tc.arr)
		assert.Equalf(t, tc.want, got, "failed uniquePaths with i: %d", i)
	}
}
