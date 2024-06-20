/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	previous := "1"
	for i := 2; i <= n; i++ {
		previous = formatString(previous)
	}
	return previous
}

func formatString(s string) string {
	count := 0
	var str strings.Builder
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		count++
		if i+1 == sLen || s[i+1] != s[i] {
			str.WriteString(strconv.Itoa(count))
			str.WriteByte(s[i])
			count = 0
		}
	}
	return str.String()
}

// @lc code=end

