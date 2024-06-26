/*
 * @lc app=leetcode id=408 lang=golang
 *
 * [408] Valid Word Abbreviation
 */

// @lc code=start
func validWordAbbreviation(word string, abbr string) bool {
	m, n := len(word), len(abbr)
	i, j := 0, 0
	for i < m && j < n {
		if word[i] == abbr[j] {
			i++
			j++
			continue
		}
		if abbr[j] <= '0' || abbr[j] > '9' {
			return false
		}
		strNumber := ""
		for j < n && abbr[j] >= '0' && abbr[j] <= '9' {
			strNumber += string(abbr[j])
			j++
		}
		number, _ := strconv.Atoi(strNumber)
		i += number
	}
	return i == m && j == n
}

// @lc code=end

