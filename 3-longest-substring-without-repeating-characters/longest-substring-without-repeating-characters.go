func lengthOfLongestSubstring(s string) int {

	lastSeen := map[byte]int{}

	start := 0
	longest := 0

	for i := 0; i < len(s); i++ {

		c := s[i]

		// check duplicate ถ้าเจออันช้ำให้ return ok เข้า if และ last >= start กัน case มันเลยจุดนับไปแล้ว
		if last, ok := lastSeen[c]; ok && last >= start {
			start = last + 1
		}
		// update longest
		length := i - start + 1
		if length > longest {
			longest = length
		}
		// update map
		lastSeen[c] = i
	}

	return longest
}