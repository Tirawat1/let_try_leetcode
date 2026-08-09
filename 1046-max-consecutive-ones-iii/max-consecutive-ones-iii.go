func longestOnes(nums []int, k int) int {
	zeroCount := 0
	max := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		// update max
		if right-left+1 > max {
			max = right - left + 1
		}

	}
	return max
}