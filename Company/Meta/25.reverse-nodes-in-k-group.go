/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	point, ktail := head, &ListNode{}
	var newHead *ListNode = nil
	for point != nil {
		count := 0
		point = head
		for count < k && point != nil {
			point = point.Next
			count++
		}
		if count == k {
			revHead := reversList(head, k)
			if newHead == nil {
				newHead = revHead
			}
			if ktail != nil {
				ktail.Next = revHead
			}
			ktail = head
			head = point
		}
	}
	if ktail != nil {
		ktail.Next = head
	}
	if newHead == nil {
		return head
	} else {
		return newHead
	}
}

func reversList(head *ListNode, k int) *ListNode {
	var newHead, point *ListNode = nil, head
	for k > 0 {
		nextNode := point.Next
		point.Next = newHead
		newHead = point
		point = nextNode
		k--
	}
	return newHead
}

// @lc code=end

