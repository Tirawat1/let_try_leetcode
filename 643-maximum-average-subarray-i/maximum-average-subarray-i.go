func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	// init first sum
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {

		// make sum = - start window + new value in window
		sum = sum - nums[i-k] + nums[i]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)

}