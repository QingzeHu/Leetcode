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
	colTable := map[int][]int{}
	queue := []NodeCol{{root, 0}}
	minCol, maxCol := 0, 0
	for len(queue) != 0 {
		currentNode := queue[0]
		// dequeue
		queue = queue[1:]
		currVal := currentNode.node.Val
		currCol := currentNode.col

		colTable[currCol] = append(colTable[currCol], currVal)

		if currCol < minCol {
			minCol = currCol
		}
		if currCol > maxCol {
			maxCol = currCol
		}
		if currentNode.node.Left != nil {
			queue = append(queue, NodeCol{currentNode.node.Left, currCol - 1})
		}
		if currentNode.node.Right != nil {
			queue = append(queue, NodeCol{currentNode.node.Right, currCol + 1})
		}
	}

	result := [][]int{}
	for i := minCol; i <= maxCol; i++ {
		if vals, ok := colTable[i]; ok {
			result = append(result, vals)
		}
	}
	return result
}

// @lc code=end

