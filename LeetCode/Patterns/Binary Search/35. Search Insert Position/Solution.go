func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid >= len(nums) || nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo

}