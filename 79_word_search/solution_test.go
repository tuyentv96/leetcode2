package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] && dfs(i, j, 0, board, word) {
				return true
			}
		}
	}

	return false
}

func dfs(i, j, count int, board [][]byte, word string) bool {
	if count == len(word) {
		return true
	}

	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 || word[count] != board[i][j] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = ' '
	found := dfs(i, j-1, count+1, board, word) ||
		dfs(i+1, j, count+1, board, word) ||
		dfs(i, j+1, count+1, board, word) ||
		dfs(i-1, j, count+1, board, word)
	board[i][j] = tmp

	return found
}

func TestExist(t *testing.T) {
	testCases := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		//{
		//	board: [][]byte{
		//		{'A', 'B', 'C', 'E'},
		//		{'S', 'F', 'C', 'S'},
		//		{'A', 'D', 'E', 'E'},
		//	},
		//	word: "ABCCED",
		//	want: true,
		//},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
	}
	for i, tc := range testCases {
		got := exist(tc.board, tc.word)
		assert.Equalf(t, tc.want, got, "failed exist with i: %d", i)
	}
}
