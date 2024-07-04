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
		currNode := queue[0].node
		currCol := queue[0].col
		queue = queue[1:]

		colTable[currCol] = append(colTable[currCol], currNode.Val)

		if currCol < minCol {
			minCol = currCol
		}

		if currCol > maxCol {
			maxCol = currCol
		}

		if currNode.Left != nil {
			queue = append(queue, NodeCol{currNode.Left, currCol - 1})
		}
		if currNode.Right != nil {
			queue = append(queue, NodeCol{currNode.Right, currCol + 1})
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

