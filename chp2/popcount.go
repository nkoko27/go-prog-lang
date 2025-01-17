package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Popcount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		[byte(x>>(1*8))] +
	 	[byte(x>>(2*8))] +
		[byte(x>>(3*8))] +
		[byte(x>>(4*8))] +
		[byte(x>>(5*8))] +
		[byte(x>>(6*8))] +
		[byte(x>>(7*8))])
}

// exercises to be done here