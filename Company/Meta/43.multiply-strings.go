/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	pos := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1, p2 := i+j, i+j+1
			sum := mul + pos[p2]
			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}
	var result string
	p1 := 0
	for pos[p1] == 0 {
		p1++
	}
	for ; p1 < len(pos); p1++ {
		result += strconv.Itoa(pos[p1])
	}
	return result
}

// @lc code=end

