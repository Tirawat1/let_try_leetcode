func singleNumber(nums []int) int {
	chk := make(map[int]int)

	for _, val := range nums {
		chk[val]++
	}
	ans := 0
	for key, val := range chk {
		if val == 1 {
			ans = key
		}
	}

	return ans
}