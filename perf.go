package zerobitperf

func nzerobit_1(key []byte) int {
	var count int
	for _, b := range key {
		for i := 0; i < 8; i++ {
			if b&(1<<i) == 0 {
				count++
			} else {
				return count
			}
		}
	}
	return count
}

func nzerobit_2(key []byte) int {
	var count int
	for _, b := range key {
		for i := 0; i < 8; i++ {
			if (b>>i)&1 == 0 {
				count++
			} else {
				return count
			}
		}
	}
	return count
}

func nzerobit_3(key []byte) int {
	var count int
	for _, b := range key {
		if (b>>0)&1 == 0 {
			count++
		} else {
			return count
		}
		if (b>>1)&1 == 0 {
			count++
		} else {
			return count
		}
		if (b>>2)&1 == 0 {
			count++
		} else {
			return count
		}
		if (b>>3)&1 == 0 {
			count++
		} else {
			return count
		}
		if (b>>4)&1 == 0 {
			count++
		} else {
			return count
		}
		if (b>>5)&1 == 0 {
			count++
		} else {
			return count
		}
		if (b>>6)&1 == 0 {
			count++
		} else {
			return count
		}
		if (b>>7)&1 == 0 {
			count++
		} else {
			return count
		}
	}
	return count
}

var zeros = [256]int{
	8, 7, 6, 6, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

func nzerobit_4(key []byte) int {
	var count int
	for _, b := range key {
		count += zeros[b]
		if b != 0 {
			return count
		}
	}
	return count
}

var debruijn64 = [64]byte{
	0, 1, 2, 53, 3, 7, 54, 27, 4, 38, 41, 8, 34,
	55, 48, 28, 62, 5, 39, 46, 44, 42, 22, 9, 24,
	35, 59, 56, 49, 18, 29, 11, 63, 52, 6, 26,
	37, 40, 33, 47, 61, 45, 43, 21, 23, 58, 17,
	10, 51, 25, 36, 32, 60, 20, 57, 16, 50, 31,
	19, 15, 30, 14, 13, 12,
}

func nzerobit_5(key []byte) int {
	for i, b := range key {
		if b != 0 {
			x := uint64(b&-b) * 0x03f79d71b4ca8b09 >> 58
			return i*8 + int(debruijn64[x])
		}
	}
	return len(key) * 8
}
