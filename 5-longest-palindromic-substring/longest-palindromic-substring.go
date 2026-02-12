func longestPalindrome(s string) string {
    start := 0
	lengMax := 0
	for i := 0; i < len(s); i++ {
		// even case
		s1, l1 := expand(s, i, i)
		// odd case
		s2, l2 := expand(s, i, i+1)

		if lengMax < l1 {
			lengMax = l1
			start = s1
		}
		if lengMax < l2 {
			lengMax = l2
			start = s2
		}

	}

	return s[start : start+lengMax]
}

func expand(sub string, left int, right int) (int, int) {

	// check left and right
	// left > 0 and right < len(s) , w[left] == w[right]
	// left -- , right ++
	// start fron center
	for left >= 0 && right < len(sub) && sub[left] == sub[right] {
		left--
		right++
	}

	start := left + 1
	length := right - left - 1

	return start, length

}