func canConstruct(ransomNote string, magazine string) bool {
	var list [26]int
	var sum int
	// find all char in ransomNote
	for _, ch := range ransomNote {
		low := unicode.ToLower(ch)
		list[low-'a']++
		sum++
	}
	// find duplicate to minus
	for _, ch := range magazine {
		low := unicode.ToLower(ch)
		if list[low-'a'] != 0 {
			list[low-'a']--
			sum--
			if sum == 0 {
				return true
			}
		}
	}
	return false
}