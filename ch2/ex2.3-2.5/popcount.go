package ex23

var pc [256]byte

type PopCountFunc func(uint64) int

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopCountCircle(x uint64) int {
	var n = 0
	for i := uint(0); i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

func PopCountBit(x uint64) int {
	var n = 0
	for i := uint(0); i < 64; i++ {
		n += int((x >> i) & 1)
	}
	return n
}

func PopCountClearRight(x uint64) int {
	var n = 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
