/*
 * @lc app=leetcode id=314 lang=golang
 *
 * [314] Binary Tree Vertical Order Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeCol struct {
	node *TreeNode
	col  int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []TreeCol{{root, 1}}
	nodeTable := map[int][]int{}
	min, max := 0, 0
	for len(queue) != 0 {
		currNode := queue[0].node
		currCol := queue[0].col
		nodeTable[currCol] = append(nodeTable[currCol], currNode.Val)
		// dequeue
		queue = queue[1:]
		if currCol < min {
			min = currCol
		}
		if currCol > max {
			max = currCol
		}

		if currNode.Left != nil {
			queue = append(queue, TreeCol{currNode.Left, currCol - 1})
		}
		if currNode.Right != nil {
			queue = append(queue, TreeCol{currNode.Right, currCol + 1})
		}
	}
	result := [][]int{}
	for i := min; i <= max; i++ {
		if v, ok := nodeTable[i]; ok {
			result = append(result, v)
		}
	}
	return result
}

// @lc code=end

