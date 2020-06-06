package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var iterator *ListNode
	var root *ListNode
	var quotient int
	for l1 != nil || l2 != nil {
		newLn := &ListNode{
			Val:  0,
			Next: nil,
		}

		if l1 != nil {
			newLn.Val += l1.Val
		}
		if l2 != nil {
			newLn.Val += l2.Val
		}

		if quotient > 0 {
			newLn.Val += quotient
			quotient = 0
		}

		if root == nil {
			root = newLn
		}

		if iterator == nil {
			iterator = newLn
		} else {
			iterator.Next = newLn
			iterator = iterator.Next
		}

		if newLn.Val > 9 {
			quotient = newLn.Val / 10
			newLn.Val = newLn.Val % 10
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if quotient > 0 && iterator != nil {
		iterator.Next = &ListNode{
			Val:  quotient,
			Next: nil,
		}
		iterator = iterator.Next
	}

	return root
}

func buildListNode(list []int) *ListNode {
	var ln *ListNode
	var start *ListNode
	for _, v := range list {
		newLn := &ListNode{
			Val:  v,
			Next: nil,
		}

		if ln == nil {
			ln = &ListNode{}
		}

		if start == nil {
			start = newLn
		}

		ln.Next = newLn
		ln = ln.Next
	}

	return start
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{[]int{1, 9, 3}, []int{1, 4, 3}, []int{2, 3, 7}},
		{[]int{5}, []int{5}, []int{0, 1}},
		{[]int{1, 8}, []int{0}, []int{1, 8}},
	}

	for i, tc := range testCases {
		l1 := buildListNode(tc.l1)
		l2 := buildListNode(tc.l2)
		got := addTwoNumbers(l1, l2)
		assert.Equalf(t, buildListNode(tc.want), got, "failed calculate twoSum with i: %d", i)
	}
}
