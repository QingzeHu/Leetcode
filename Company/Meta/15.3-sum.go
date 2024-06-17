/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				k--
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			}
		}
	}
	return ans
}

// @lc code=end

