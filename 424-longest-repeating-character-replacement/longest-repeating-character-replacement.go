func characterReplacement(s string, k int) int {
	mapChar := make(map[byte]int)
	charsToReplace := 0
	maxFreq := 0
	left := 0
	max := 0

	for right := 0; right < len(s); right++ {

		mapChar[s[right]]++
		if mapChar[s[right]] > maxFreq {
			maxFreq = mapChar[s[right]]
		}
		charsToReplace = (right - left + 1) - maxFreq
		for charsToReplace > k {
			mapChar[s[left]]--
			left++
			charsToReplace = (right - left + 1) - maxFreq
		}

		// fmt.Println(charsToReplace)
		max = right - left + 1
	}
	return max
}