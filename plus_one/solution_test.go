package buy_stock_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func plusOne(digits []int) []int {
	result := make([]int, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] += 1
		}

		tmp := digits[i] + result[i]
		result[i] = tmp % 10
		if tmp/10 > 0 {
			if i == 0 {
				result = append([]int{1}, result...)
				continue
			}

			result[i-1] = result[i-1] + 1
		}
	}

	return result
}
func Test(t *testing.T) {
	testCases := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{9, 9},
			want: []int{1, 0, 0},
		},
	}
	for i, tc := range testCases {
		got := plusOne(tc.arr)
		assert.Equalf(t, tc.want, got, "failed plusOne with i: %d", i)
	}
}
