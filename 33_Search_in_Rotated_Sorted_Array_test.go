package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func SearchInRotatedArray(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] != target {
			return -1
		}
		return 0
	}
	return search1(nums, 0, length-1, target)
}

func search1(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	if left == right {
		if nums[left] == target {
			return left
		}
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}

	if nums[left] < nums[mid] && nums[left] <= target && nums[mid] > target {
		return binarySearch(nums, left, mid-1, target)
	}
	if nums[right] > nums[mid] && nums[mid] < target && nums[right] >= target {
		return binarySearch(nums, mid+1, right, target)
	}
	if nums[left] > nums[mid] {
		return search1(nums, left, mid-1, target)
	}
	if nums[right] < nums[mid] {
		return search1(nums, mid+1, right, target)
	}
	return -1
}

func binarySearch(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}
	if target < nums[mid] {
		return binarySearch(nums, left, mid-1, target)
	}
	return binarySearch(nums, mid+1, right, target)

}

func search(nums []int, target int) int {
	rotated := indexOfMin(nums) /* 数组旋转了的距离 */
	size := len(nums)
	left, right := 0, size-1

	for left <= right {
		mid := (left + right) / 2
		/* nums 是 rotated，所以需要使用 rotatedMid 来获取 mid 的值 */
		rotatedMid := (rotated + mid) % size
		switch {
		case nums[rotatedMid] < target:
			left = mid + 1
		case target < nums[rotatedMid]:
			right = mid - 1
		default:
			return rotatedMid
		}
	}

	return -1
}

/* nums 是被旋转了的递增数组 */
func indexOfMin(nums []int) int {
	size := len(nums)
	left, right := 0, size-1
	for left < right {
		mid := (left + right) / 2
		if nums[right] < nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func BenchmarkSearch(b *testing.B) {
	b.Run("search", func(b *testing.B) {
		b.Run("", func(b *testing.B) {
			arr := []int{4, 5, 6, 7, 0, 1, 2}
			target := 0
			for n := 0; n < b.N; n++ {
				search(arr, target)
			}
		})
	})
	b.Run("search1", func(b *testing.B) {
		b.Run("", func(b *testing.B) {
			arr := []int{4, 5, 6, 7, 0, 1, 2}
			target := 0
			for n := 0; n < b.N; n++ {
				SearchInRotatedArray(arr, target)
			}
		})
	})
}

func TestSearchInRotatedArray(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			arr:    []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
	}

	for i, tc := range testCases {
		actual := SearchInRotatedArray(tc.arr, tc.target)
		assert.Equalf(t, tc.want, actual, "failed TestSearchInRotatedArray i: %d", i)
	}
}
