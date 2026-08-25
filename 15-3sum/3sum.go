func threeSum(arr []int) [][]int {
	answer := make([][]int, 0)
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		// i = fixed value
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		left := i + 1
		right := len(arr) - 1
		for left < right {
			sum := arr[left] + arr[right]
			if sum == -arr[i] {
				// if found duplicate
				// เจอ match แล้ว บันทึกคำตอบ
				val := []int{arr[left], arr[right], arr[i]}
				answer = append(answer, val)
				left++
				right--
				for left < right && arr[left] == arr[left-1] {
					left++
				}
				for right > left && arr[right] == arr[right+1] {
					right--
				}
			} else if sum < -arr[i] {
				left++ // sum น้อยไป ต้องเพิ่ม

			} else {
				right-- // sum มากไป ต้องลด

			}
		}

	}
	return answer

}