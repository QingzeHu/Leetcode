/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    amap:=make(map[int]int)
	for i, num:=range nums{
		result := target-num
		if _,ok:=amap[result];ok{
			return []int{i, amap[result]}
		}
		amap[num]=i
	}
	return nil
}
// @lc code=end

