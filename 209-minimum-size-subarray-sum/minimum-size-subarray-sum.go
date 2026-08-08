func minSubArrayLen(target int, nums []int) int {
	sum, left := 0, 0
	minLen := math.MaxInt

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			// fmt.Println("minLen updated to", minLen, "at right=", right, "left=", left)
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}