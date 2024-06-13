/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	sum := 0
	sybolmap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s); i++ {
		sum += sybolmap[s[i]]
		if i != 0 {
			if sybolmap[s[i-1]] < sybolmap[s[i]] {
				sum -= 2 * sybolmap[s[i-1]]
			}
		}
	}
	return sum
}

// @lc code=end

