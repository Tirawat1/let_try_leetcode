func mySqrt(x int) int {
	lo, hi := 0, x
	result := 0
	for lo <= hi {
		numX := lo + (hi-lo)/2
		sq := numX * numX
		if sq > x {
			hi = numX - 1
		} else if sq <= x {
			// fmt.Println(result)
			result = numX
			lo = numX + 1
		}
	}
	return result
}