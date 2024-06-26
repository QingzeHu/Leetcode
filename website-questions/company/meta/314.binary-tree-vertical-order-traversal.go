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

type NodeCol struct {
	node *TreeNode
	col  int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	nodeTable := map[int][]int{}
	minCol, maxCol := 0, 0
	queue := []NodeCol{{root, 0}}

	for len(queue) != 0 {
		currentNode := queue[0]
		// dequeue
		queue = queue[1:]
		node := currentNode.node
		col := currentNode.col
		nodeTable[col] = append(nodeTable[col], node.Val)

		if col < minCol {
			minCol = col
		}
		if col > maxCol {
			maxCol = col
		}
		if node.Left != nil {
			queue = append(queue, NodeCol{node.Left, col - 1})
		}
		if node.Right != nil {
			queue = append(queue, NodeCol{node.Right, col + 1})
		}
	}
	result := [][]int{}
	for i := minCol; i <= maxCol; i++ {
		if vals, ok := nodeTable[i]; ok {
			result = append(result, vals)
		}
	}
	return result
}

// @lc code=end

