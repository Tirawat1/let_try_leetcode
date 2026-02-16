func reorganizeString(s string) string {

	// step 1 : check a most repeated
	mapChar := make(map[rune]int)
	for _, v := range s {
		mapChar[v]++
	}
	mostRep := 0
	for _, v := range mapChar {
		if v > mostRep {
			mostRep = v
		}
	}

	// step 2 : check if most left more than any another left rune
	if mostRep > (len(s)+1)/2 {
		return ""
	}
	// step 2.5 find most char
	mostChar := rune(0)
	for k, v := range mapChar {
		if v == mostRep {
			mostChar = k
		}
	}

	// Example aaaa b c d
	// step 3 : make a new string with a_a_a_a
	ans := make([]rune, len(s))
	idx := 0
	for mapChar[mostChar] > 0 {
		if idx >= len(s) {
			idx = 1
		}
		ans[idx] = mostChar
		idx += 2
		mapChar[mostChar]--
	}

	for char, count := range mapChar {
		for count > 0 {
			if idx >= len(s) {
				idx = 1
			}
			ans[idx] = char
			idx += 2
			count--
		}
	}

	return string(ans)
}
