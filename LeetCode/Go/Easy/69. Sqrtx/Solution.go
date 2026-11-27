// better verison
func mySqrt(x int) int {
	lo, hi := 0, x
	for lo <= hi {
		numX := lo + (hi-lo)/2
		sq := numX * numX
		if sq == x {
			return numX
		} else if sq > x {
			hi = numX - 1
		} else if sq <= x {
			lo = numX + 1
		}
	}
	return lo - 1
}
