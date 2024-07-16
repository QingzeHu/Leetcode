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

type ColNode struct {
	node *TreeNode
	col  int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []ColNode{{root, 0}}
	treeMap := make(map[int][]int)
	minCol, maxCol := 0, 0

	for len(queue) != 0 {
		currentNode := queue[0].node
		currentCol := queue[0].col
		queue = queue[1:]

		treeMap[currentCol] = append(treeMap[currentCol], currentNode.Val)
		if currentCol < minCol {
			minCol = currentCol
		}
		if currentCol > maxCol {
			maxCol = currentCol
		}
		if currentNode.Left != nil {
			queue = append(queue, ColNode{currentNode.Left, currentCol - 1})
		}
		if currentNode.Right != nil {
			queue = append(queue, ColNode{currentNode.Right, currentCol + 1})
		}
	}

	result := [][]int{}
	for i := minCol; i <= maxCol; i++ {
		if v, ok := treeMap[i]; ok {
			result = append(result, v)
		}
	}
	return result
}

// @lc code=end

