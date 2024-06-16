/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	write := 1
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[write] = nums[r]
			write++
		}
	}
	return write
}

// @lc code=end

