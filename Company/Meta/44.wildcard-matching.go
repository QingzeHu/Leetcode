/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	if p == s {
		return true
	}

	if pLen > 0 && strings.Count(p, "*") == pLen {
		return true
	}

	if pLen == 0 || sLen == 0 {
		return false
	}

	dp := make([][]bool, pLen+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sLen+1)
	}
	dp[0][0] = true
	for j := 1; j <= pLen; j++ {
		if p[j-1] == '*' {
			i := 1
			for !dp[j-1][i-1] && i <= sLen {
				i++
			}
			dp[j][i-1] = dp[j-1][i-1]
			for i <= sLen {
				dp[j][i] = true
				i++
			}
		} else if p[j-1] == '?' {
			for s := 1; s <= sLen; s++ {
				dp[j][s] = dp[j-1][s-1]
			}
		} else {
			for i := 1; i <= sLen; i++ {
				dp[j][i] = dp[j-1][i-1] && p[j-1] == s[i-1]
			}
		}
	}
	return dp[pLen][sLen]
}

// @lc code=end

