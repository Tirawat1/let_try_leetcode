func isValid(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{}

	for _, r := range s {
		b := byte(r)
		if b == '(' || b == '{' || b == '[' {
			stack = append(stack, r)
			continue
		}
		if open, ok := pairs[b]; ok {
			if len(stack) == 0 || rune(open) != stack[len(stack)-1]  {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}