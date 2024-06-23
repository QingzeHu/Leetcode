/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -1 * n
		x = 1.0 / x
	}
	var result float64 = 1
	for n != 0 {
		if n%2 == 1 {
			result = result * x
			n -= 1
		}
		x = x * x
		n = n / 2
	}
	return result
}

// @lc code=end

