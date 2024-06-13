/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	ans := []string{}
	rm := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	var trackback func(combination, nextDegits string)
	trackback = func(combination, nextDegits string) {
		if nextDegits == "" {
			ans = append(ans, combination)
		} else {
			letters := rm[nextDegits[0]-'2']
			for i := 0; i < len(letters); i++ {
				trackback(combination+string(letters[i]), nextDegits[1:])
			}
		}
	}
	trackback("", digits)
	return ans
}

// @lc code=end

