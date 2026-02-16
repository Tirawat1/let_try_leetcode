func leastInterval(tasks []byte, n int) int {

	// tasks
	count := 0
	// make tasks in map
	mapChar := make(map[byte]int)

	for _, t := range tasks {
		mapChar[t]++
	}
	// find a max freq first
	maxFreq := 0
	for _, t := range mapChar {
		if t > maxFreq {
			maxFreq = t
		}
	}
	// find how many char is max freq
	maxChar := 0
	for _, t := range mapChar {
		if t == maxFreq {
			maxChar++
		}
	}

	// block fomula
	count = (maxFreq-1)*(n+1) + maxChar

	return max(count, len(tasks))
}