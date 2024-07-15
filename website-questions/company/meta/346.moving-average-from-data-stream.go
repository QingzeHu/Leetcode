/*
 * @lc app=leetcode id=346 lang=golang
 *
 * [346] Moving Average from Data Stream
 */

// @lc code=start
type MovingAverage struct {
	size   int
	window []int
	sum    int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:   size,
		window: make([]int, 0, size),
		sum:    0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.size == len(this.window) {
		this.sum -= this.window[0]
		this.window = this.window[1:]
	}
	this.window = append(this.window, val)
	this.sum += val
	return float64(this.sum) / float64(len(this.window))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
// @lc code=end

