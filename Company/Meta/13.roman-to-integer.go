/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	sum := 0
	rm := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s); i++ {
		sum += rm[s[i]]
		if i != 0 {
			if rm[s[i-1]] < rm[s[i]] {
				sum -= rm[s[i-1]] * 2
			}
		}
	}
	return sum
}

// @lc code=end
