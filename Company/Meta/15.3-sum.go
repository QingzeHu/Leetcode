/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	ans := [][]int{}
	for i := 0; i < l-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, l-1
		for j < k {
			v := nums[i] + nums[j] + nums[k]
			if v < 0 {
				j++
			} else if v > 0 {
				k--
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
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

