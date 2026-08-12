func groupAnagrams(strs []string) [][]string {
	groupString := make(map[string][]string)
	for _, word := range strs {
		// sort string
		b := []byte(word)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)

		groupString[key] = append(groupString[key], word)
	}
	ans := make([][]string, 0)
	for _, word := range groupString {
		ans = append(ans, word)
	}

	return ans
}