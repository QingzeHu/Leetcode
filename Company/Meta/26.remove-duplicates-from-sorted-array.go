/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	write := 1
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			if nums[i-1] != nums[i] {
				nums[write] = nums[i]
				write++
			}
		}
	}
	return write
}

// @lc code=end

