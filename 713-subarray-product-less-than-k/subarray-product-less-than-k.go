func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	count := 0
	left := 0
	chkSum := 1
	for right := 0; right < len(nums); right++ {
		chkSum *= nums[right]
		for chkSum >= k {
			chkSum /= nums[left]
			left++
		}

		count += right - left + 1
	}
	return count
}