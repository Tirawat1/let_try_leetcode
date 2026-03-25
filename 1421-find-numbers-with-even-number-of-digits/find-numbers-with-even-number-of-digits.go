func evenDigit(n int) bool {

	c := 0
	for n > 0 {
		c++
		n = n / 10
	}
	if c%2 == 0 {
		return true
	}

	return false
}

func findNumbers(nums []int) int {

	count := 0

	for _, v := range nums {
		if evenDigit(v) {
			count++
		}
	}

	return count
}