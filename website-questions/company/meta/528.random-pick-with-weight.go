/*
 * @lc app=leetcode id=528 lang=golang
 *
 * [528] Random Pick with Weight
 */

// @lc code=start
type Solution struct {
	prefixSum []int
	totalSum  int
}

func Constructor(w []int) Solution {
	prefixSums := make([]int, len(w))
	prefixSums[0] = w[0]
	for i := 1; i < len(w); i++ {
		prefixSums[i] = prefixSums[i-1] + w[i]
	}
	return Solution{prefixSums, prefixSums[len(w)-1]}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.totalSum) + 1
	return sort.Search(len(this.prefixSum), func(i int) bool {
		return this.prefixSum[i] >= target
	})
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end

