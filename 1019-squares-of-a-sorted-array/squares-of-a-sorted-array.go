// this is more optimize way
func sortedSquares(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	slices.Sort(nums)
	return nums
}