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
	nm := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ans := make([]string, 0)
	var backtrack func(combination string, nextDigits string)
	backtrack = func(combination string, nextDigits string) {
		if nextDigits == "" {
			ans = append(ans, combination)
		} else {
			letters := nm[nextDigits[0]-'2']
			for i := 0; i < len(letters); i++ {
				backtrack(combination+string(letters[i]), nextDigits[1:])
			}
		}
	}
	// "23"
	backtrack("", digits)
	return ans
}

// @lc code=end

