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
func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
		
	}

	// 字典，用于存储每个列的节点值
	m := make(map[int][]int)

	// 两个队列，一个用于存储节点，一个用于存储节点对应的列索引
	q1 := []*TreeNode{root}
	q2 := []int{0}

	// 记录最小和最大的列索引
	var min, max int

	// 广度优先搜索（BFS）
	for len(q1) > 0 {
		cur := q1[0] // 取出队列中的第一个节点
		q1 = q1[1:]  // 移除第一个节点
		col := q2[0] // 取出对应的列索引
		q2 = q2[1:]  // 移除第一个列索引

		// 将节点值追加到对应列的列表中
		m[col] = append(m[col], cur.Val)

		// 更新最小和最大列索引
		if min > col {
			min = col
		}
		if max < col {
			max = col
		}

		// 将左子节点加入队列，并更新列索引
		if cur.Left != nil {
			q1 = append(q1, cur.Left)
			q2 = append(q2, col-1)
		}
		// 将右子节点加入队列，并更新列索引
		if cur.Right != nil {
			q1 = append(q1, cur.Right)
			q2 = append(q2, col+1)
		}
	}

	// 构建结果列表
	result := make([][]int, 0, max-min+1)
	for i := min; i <= max; i++ {
		result = append(result, m[i])
	}
	return result
}

// @lc code=end

