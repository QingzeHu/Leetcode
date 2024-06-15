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
	numbers := []string{
		"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	combinations := []string{}
	var backtrack func(combination string, nextDigits string)
	backtrack = func(combination string, nextDigits string) {
		if nextDigits == "" {
			combinations = append(combinations, combination)
		} else {
			for _, letter := range numbers[nextDigits[0]-'2'] {
				backtrack(combination+string(letter), nextDigits[1:])
			}
		}

	}
	backtrack("", digits)
	return combinations
}

// @lc code=end

