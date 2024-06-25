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

type ColumnNode struct {
	node *TreeNode
	col  int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	columnTable := map[int][]int{}
	minCol, maxCol := 0, 0
	queue := []ColumnNode{{root, 0}}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		currentCol := currentNode.col
		currentVal := currentNode.node.Val
		columnTable[currentCol] = append(columnTable[currentCol], currentVal)

		if currentCol < minCol {
			minCol = currentCol
		}
		if currentCol > maxCol {
			maxCol = currentCol
		}

		if currentNode.node.Left != nil {
			queue = append(queue, ColumnNode{currentNode.node.Left, currentCol - 1})
		}
		if currentNode.node.Right != nil {
			queue = append(queue, ColumnNode{currentNode.node.Right, currentCol + 1})
		}
	}

	result := [][]int{}
	for i := minCol; i <= maxCol; i++ {
		if vals, ok := columnTable[i]; ok {
			result = append(result, vals)
		}
	}
	return result
}

// @lc code=end

