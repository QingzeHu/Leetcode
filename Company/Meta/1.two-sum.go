/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	resultMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := resultMap[complement]; ok {
			return []int{i, j}
		}
		resultMap[num] = i
	}
	return nil
}

// @lc code=end

