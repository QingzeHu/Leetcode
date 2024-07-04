/*
 * @lc app=leetcode id=528 lang=golang
 *
 * [528] Random Pick with Weight
 */

// @lc code=start
type Solution struct {
	prefixsums []int
	totalsum   int
}

func Constructor(w []int) Solution {
	sums := make([]int, len(w))
	sums[0] = w[0]
	for i := 1; i < len(w); i++ {
		sums[i] = sums[i-1] + w[i]
	}
	return Solution{sums, sums[len(sums)-1]}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.totalsum) + 1
	return sort.Search(len(this.prefixsums), func(i int) bool {
		return this.prefixsums[i] >= target
	})
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end

