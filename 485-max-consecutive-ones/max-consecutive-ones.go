func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0

	for _, i := range nums {
		if i == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count > max {
		max = count
	}
	return max
}