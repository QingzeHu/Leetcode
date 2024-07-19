/*
 * @lc app=leetcode id=1249 lang=golang
 *
 * [1249] Minimum Remove to Make Valid Parentheses
 */

// @lc code=start
func minRemoveToMakeValid(s string) string {
	stack := []int{}
	chars := []rune(s)
	for i, v := range chars {
		if v == '(' {
			stack = append(stack, i)
		} else if v == ')' {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else {
				chars[i] = 0
			}
		}
	}
	for i := 0; i < len(stack); i++ {
		chars[stack[i]] = 0
	}
	result := ""
	for i := 0; i < len(chars); i++ {
		if chars[i] != 0 {
			result += string(chars[i])
		}
	}

	return result
}

// @lc code=end

