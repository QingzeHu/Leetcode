/*
 * @lc app=leetcode id=1570 lang=golang
 *
 * [1570] Dot Product of Two Sparse Vectors
 */

// @lc code=start
type SparseVector struct {
	elements map[int]int
}

func Constructor(nums []int) SparseVector {
	sv := SparseVector{elements: make(map[int]int)}
	for i, num := range nums {
		if num != 0 {
			sv.elements[i] = num
		}
	}
	return sv
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	result := 0
	for i, e := range vec.elements {
		if v2, ok := this.elements[i]; ok {
			result += v2 * e
		}
	}
	return result
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
// @lc code=end

