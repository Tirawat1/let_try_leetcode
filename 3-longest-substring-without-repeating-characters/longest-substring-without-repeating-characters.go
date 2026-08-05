func lengthOfLongestSubstring(s string) int {
	lastSeen := map[byte]int{}
	start := 0
	longest := 0
	for i := 0; i < len(s); i++ {
		// check first char
		c := s[i]
		// if duplicate stop longest start = last index
		// สิ่งที่ต้องทำตรงนี้คือตรวจก่อนว่ามีตัวช้ำกันไหมคือดูจาก lastseen ว่า c ล่าสุดที่เข้ามา
		// มีตัวช้ำกันแล้วไหมใน lastSeen ก็คือสมมติวิ่ง 1 2 3 4 มาจะหยุด

		// 1R1T7
		if latestValue, ok := lastSeen[c]; ok && start <= latestValue {
			start = latestValue + 1
		}
		space := i - start + 1
		if space > longest {
			longest = space
		}
		// length := i - start + 1
		// if length > longest {
		// 	longest = length
		// }
		lastSeen[c] = i

	}

	return longest
}
