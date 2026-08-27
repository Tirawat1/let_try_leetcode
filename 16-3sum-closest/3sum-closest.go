
func threeSumClosest(arr []int, target int) int {
	sort.Ints(arr)
	minDiff := math.Inf(1)
	sum := 0
	for i := 0; i < len(arr); i++ {
		// skipped fixed index
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		left := i + 1
		right := len(arr) - 1
		for left < right {
			// fmt.Printf("fixed %d , left : %d, right : %d\n", arr[i], arr[left], arr[right])
			// find closest sum value
			closestSum := float64(arr[i] + arr[left] + arr[right])
			chkDiff := math.Abs(closestSum - float64(target))
			// check diff first if
			if chkDiff < minDiff {
				minDiff = chkDiff
				sum = int(closestSum)
			}

			// check min diff from target
			if closestSum < float64(target) {
				left++
			} else if closestSum > float64(target) {
				right--
			} else {
				return int(closestSum)
			}

		}

	}
	return sum
}