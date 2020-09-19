package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

			if i == 0 && j == 0 {
				dp[i][j] = 1 - obstacleGrid[0][0]
				continue
			}

			if i == 0 {
				dp[i][j] = dp[i][j-1]
				continue
			}

			if j == 0 {
				dp[i][j] = dp[i-1][j]
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func TestUniquePathsWithObstacles(t *testing.T) {
	testCases := []struct {
		arr  [][]int
		want int
	}{
		//{
		//	arr: [][]int{
		//		{
		//			0, 0, 0,
		//		},
		//		{
		//			0, 1, 0,
		//		},
		//		{
		//			0, 0, 0,
		//		},
		//	},
		//	want: 2,
		//},
		{
			arr: [][]int{
				{
					1, 0,
				},
				{
					0, 0,
				},
			},
			want: 0,
		},
		{
			arr: [][]int{
				{
					1, 0,
				},
			},
			want: 0,
		},
		{
			arr: [][]int{
				{
					0,
				},
			},
			want: 1,
		},
		{
			arr: [][]int{
				{
					1,
				},
			},
			want: 0,
		},
		{
			arr: [][]int{
				{
					0, 0,
				},
				{
					1, 0,
				},
			},
			want: 1,
		},
	}

	for i, tc := range testCases {
		got := uniquePathsWithObstacles(tc.arr)
		assert.Equalf(t, tc.want, got, "failed uniquePaths with i: %d", i)
	}
}
