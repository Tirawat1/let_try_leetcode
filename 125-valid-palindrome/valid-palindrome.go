func isPalindrome(s string) bool {

	newSentence := []rune{}
	for _, val := range s {
		if unicode.IsLetter(val) || unicode.IsDigit(val) {
			newSentence = append(newSentence, unicode.ToLower(val))
		}
	}
	left := 0
	right := len(newSentence) - 1
	for left < right {
		if newSentence[left] != newSentence[right] {
			return false
		}
		left++
		right--
	}

	return true
}
