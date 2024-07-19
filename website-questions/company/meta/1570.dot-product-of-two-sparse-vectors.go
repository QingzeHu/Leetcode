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
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			sv.elements[i] = nums[i]
		}
	}
	return sv
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	ans := 0
	for i, v := range this.elements {
		if v2, ok := vec.elements[i]; ok {
			ans += v * v2
		}
	}
	return ans
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
// @lc code=end

