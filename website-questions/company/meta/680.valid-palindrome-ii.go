/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 */

// @lc code=start
func validPalindrome(s string) bool {
	isPalidrome := func(iss string, left, right int) bool {
		for left < right {
			if iss[left] != iss[right] {
				return false
			}
			left++
			right--

		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalidrome(s, left+1, right) || isPalidrome(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

// @lc code=end

