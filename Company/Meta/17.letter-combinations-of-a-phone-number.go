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
	pn := []string{
		"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	ans := []string{}
	var backtrack func(combination, nextDigits string)
	backtrack = func(combination, nextDigits string) {
		if nextDigits == "" {
			ans = append(ans, combination)
		} else {
			letters := pn[nextDigits[0]-'2']
			for i := 0; i < len(letters); i++ {
				backtrack(combination+string(letters[i]), nextDigits[1:])
			}
		}
	}
	backtrack("", digits)
	return ans
}

// @lc code=end

