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
		numStr := ""
		for j < n && abbr[j] >= '0' && abbr[j] <= '9' {
			numStr += string(abbr[j])
			j++
		}
		num, _ := strconv.Atoi(numStr)
		i += num
	}
	return i == m && j == n
}

// @lc code=end

