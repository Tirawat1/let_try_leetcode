func runningSum(nums []int) []int {

	n := len(nums)
	ans := make([]int, n)
	ans[0] = nums[0]
	// ans = {a1 , a1+a2 , a1+a2+a3 . a1+a2+a3+a4}
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + nums[i]
	}
	return ans

}