func maxValue(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxSumSubarray(nums []int, k int) int {
	// step is init window and compare next window
	maxSum := 0
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum = sum

	for i := k; i < len(nums); i++ {

		sum = sum - nums[i-k] + nums[i]
		maxSum = maxValue(maxSum, sum)
	}

	return maxSum

}
