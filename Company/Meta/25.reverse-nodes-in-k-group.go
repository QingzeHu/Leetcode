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
	if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{0, head}
	prevGroupEnd := dummy

	for {
		kth := getKthNode(prevGroupEnd, k)
		if kth == nil {
			break
		}
		nextGroupStart := kth.Next

		// Reverse k nodes
		prev, curr := kth.Next, prevGroupEnd.Next
		for curr != nextGroupStart {
			temp := curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}

		// Connect with the previous part
		temp := prevGroupEnd.Next
		prevGroupEnd.Next = kth
		prevGroupEnd = temp
	}

	return dummy.Next
}

func getKthNode(curr *ListNode, k int) *ListNode {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	return curr
}

// @lc code=end

