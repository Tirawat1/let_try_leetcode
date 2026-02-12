func longestPalindrome(s string) int {
    // test sync
	m := map[byte]int{}
	sum := 0
	hasOdd := false
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	// k : char , v : value
	// palindrome mean it is can count with a 1 max(even) and sum(odd)
	for _, v := range m {
		if v%2 == 0 {
			sum += v
		} else {
			sum += v - 1
			hasOdd = true
		}
	}

	if hasOdd {
		sum++
	}

	return sum
}
