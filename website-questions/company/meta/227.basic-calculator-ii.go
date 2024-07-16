/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 */

// @lc code=start
func calculate(s string) int {
	stack := []int{0}
	num := 0
	n := len(s)
	op := '+'
	for i := 0; i < n; i++ {
		currnCh := s[i]
		if currnCh >= '0' && currnCh <= '9' {
			num = num*10 + int(currnCh-'0')
		}
		if currnCh == '+' || currnCh == '-' || currnCh == '*' || currnCh == '/' || i == n-1 {
			m := len(stack)
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[m-1] = stack[m-1] * num
			case '/':
				stack[m-1] = stack[m-1] / num
			}
			op = rune(currnCh)
			num = 0
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}

// @lc code=end

