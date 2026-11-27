func isPerfectSquare(num int) bool {
	// your code here
	lo, hi := 0, num
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sq := mid * mid
		if sq == num {
			return true
		} else if sq > num {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}
