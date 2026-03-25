func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		a1 := nums[i] * nums[i]
		ans[i] = int(a1)
	}
	sort.Ints(ans)
	return ans
}