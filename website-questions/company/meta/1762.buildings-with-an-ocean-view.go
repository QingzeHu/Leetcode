/*
 * @lc app=leetcode id=1762 lang=golang
 *
 * [1762] Buildings With an Ocean View
 */

// @lc code=start
func findBuildings(heights []int) []int {
	maxHeight := 0
	result := []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			result = append(result, i)
			maxHeight = heights[i]
		}
	}
	result2 := make([]int, len(result))
	for i, v := range result {
		result2[len(result)-i-1] = v
	}
	return result2
}

// @lc code=end

