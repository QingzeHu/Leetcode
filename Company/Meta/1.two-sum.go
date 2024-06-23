/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	nm := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := nm[complement]; ok {
			return []int{i, j}
		}
		nm[num] = i
	}
	return nil
}

// @lc code=end

