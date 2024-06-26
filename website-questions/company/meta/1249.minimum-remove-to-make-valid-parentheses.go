/*
 * @lc app=leetcode id=1249 lang=golang
 *
 * [1249] Minimum Remove to Make Valid Parentheses
 */

// @lc code=start
func minRemoveToMakeValid(s string) string {
	stack := []int{}
	chars := []rune(s)
	for i, char := range chars {
		if char == '(' {
			stack = append(stack, i)
		} else if char == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				chars[i] = 0
			}
		}
	}
	for _, index := range stack {
		chars[index] = 0
	}
	result := ""
	for _, char := range chars {
		if char != 0 {
			result += string(char)
		}
	}
	return result
}

// @lc code=end

