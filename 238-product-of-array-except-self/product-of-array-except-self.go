func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	left := make([]int, n)
	right := make([]int, n)

	// มันคือการแบ่ง index กันด้วยกฏการคูณ left[i] * right[n-1] = ans[i]
	// create left first
	// LEFT pass
	l := 1
	for i := 0; i < n; i++ {
		left[i] = l
		l *= nums[i]
	}

	// RIGHT pass
	r := 1
	for i := n - 1; i >= 0; i-- {
		right[i] = r
		r *= nums[i]
	}

	for i := 0; i < n; i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}