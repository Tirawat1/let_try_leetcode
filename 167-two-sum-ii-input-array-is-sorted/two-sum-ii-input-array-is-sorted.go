func twoSum(numbers []int, target int) []int {
    ans := make([]int, 2)

	// add logic left and right
	r := len(numbers) - 1
	l := 0

	for r >= l {
		if numbers[l]+numbers[r] == target {
			ans[0], ans[1] = l+1, r+1
			break
		}
		if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}

	}

	return ans
}