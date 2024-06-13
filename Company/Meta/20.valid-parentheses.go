/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := []rune{}
	sMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range s {
		if left, ok := sMap[c]; ok {
			topElement := '#'
			if len(stack) != 0 {
				topElement, stack = stack[len(stack)-1], stack[0:len(stack)-1]
			}
			if topElement != left {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

// @lc code=end

