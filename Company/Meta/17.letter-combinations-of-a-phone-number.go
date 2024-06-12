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
	phoneMap := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	var output []string
	var trackback func(combination, nextDigits string)
	trackback = func(combination, nextDigits string) {
		if nextDigits == "" {
			output = append(output, combination)
		} else {
			letters := phoneMap[nextDigits[0]-'2']
			for _, letter := range letters {
				trackback(combination+string(letter), nextDigits[1:])
			}
		}
	}
	trackback("", digits)
	return output
}

// @lc code=end

