/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	amonut := len(lists)
	interval := 1
	for interval < amonut {
		for i := 0; i < amonut-interval; i += interval * 2 {
			lists[i] = merge2KlLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	if amonut != 0 {
		return lists[0]
	}
	return nil
}

func merge2KlLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	point := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			point.Next = l1
			l1 = l1.Next
		} else {
			point.Next = l2
			l2 = l2.Next
		}
		point = point.Next
	}
	if l1 == nil {
		point.Next = l2
	} else {
		point.Next = l1
	}
	return head.Next
}

// @lc code=end

