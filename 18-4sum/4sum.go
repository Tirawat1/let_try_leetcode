func fourSum(nums []int, target int) [][]int {
	answer := make([][]int, 0)
	if len(nums) < 4 {
		return answer
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			initLeft := i
			initLeftPlus := j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			runningLeft := initLeftPlus + 1
			runningRight := len(nums) - 1
			for runningLeft < runningRight {
				sum := nums[initLeft] + nums[initLeftPlus] + nums[runningLeft] + nums[runningRight]
				if sum == target {
					answer = append(answer, []int{nums[initLeft], nums[initLeftPlus], nums[runningLeft], nums[runningRight]})
					runningRight--
					runningLeft++
					for runningLeft < runningRight && nums[runningLeft] == nums[runningLeft-1] {
						runningLeft++
					}
					for runningRight > runningLeft && nums[runningRight] == nums[runningRight+1] {
						runningRight--
					}
				} else if sum > target {
					runningRight--
				} else {
					runningLeft++
				}
			}
		}
	}

	return answer
}
