func threeSum(nums []int) [][]int {

	// sort array first
	// [-1, 0, 1, 2, -1, -4] -> [-4,-1,-1,0,1,2]
	// use 3 num to += 0
	sort.Ints(nums)
	answer := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		// i = fixed
		left := i + 1
		right := len(nums) - 1
		var sum int
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for right > left {
			//sum = fixed + left + right
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 {
				// append to array
				answer = append(answer, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else { // sum < 0
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}
	return answer
}
