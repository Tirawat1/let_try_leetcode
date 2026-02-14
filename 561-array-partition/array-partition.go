func arrayPairSum(nums []int) int {
	// [6,2,6,5,1,2]
	// [1,2,2,5,6,6]
	sort.Ints(nums)
	sum := 0

	for i := 0; i < len(nums)-1; i += 2 {
		sum += min(nums[i], nums[i+1])
	}

	return sum

}