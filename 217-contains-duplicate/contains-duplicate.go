func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i, val := range nums {
		if _, ok := m[val]; ok {
			return true
		}
		m[val] = i
	}
	return false
}