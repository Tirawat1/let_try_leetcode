

type MovingAverage struct {
	size int
	sum  int
	nums []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		sum:  0,
		nums: []int{},
	}
}
func (this *MovingAverage) Next(val int) float64 {
	this.nums = append(this.nums, val)
	// check if more thant size or not if more than erase first and + new value
	if len(this.nums) > this.size {
		this.sum = this.sum + val - this.nums[0]
		this.nums = this.nums[1:] // use this to remove index 0
	} else {
		this.sum += val
	}

	return float64(this.sum) / float64(len(this.nums))
}

