/*
 * @lc app=leetcode id=1650 lang=golang
 *
 * [1650] Lowest Common Ancestor of a Binary Tree III
 */

// @lc code=start
/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
	a, b := p, q
	for a != b {
		if a != nil {
			a = a.Parent
		} else {
			a = q
		}
		if b != nil {
			b = b.Parent
		} else {
			b = p
		}
	}
	return a
}

// @lc code=end

