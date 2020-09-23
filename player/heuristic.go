package player

func abs(n int16) int16 {
	y := n >> 15
	return (n ^ y) - y
}

func incDec(t1, t2 uint8, i, d int16) (int16, int16) {
	if t1 < t2 {
		i++
	} else if t1 > t2 {
		d++
	}
	return i, d
}

func monotonicityTopDown(b uint64) int16 {
	var i, d int16 // increment/decrement
	var y, x uint8
	for x = 0; x < nrCol; x++ {
		for y = 0; y < nrLine-1; y++ {
			i, d = incDec(get(b, y, x), get(b, y+1, x), i, d)
		}
	}
	return abs(i - d)
}

func monotonicityLeftRight(b uint64) int16 {
	var i, d int16
	var y, x uint8
	for y = 0; y < nrLine; y++ {
		for x = 0; x < nrCol-1; x++ {
			i, d = incDec(get(b, y, x), get(b, y, x+1), i, d)
		}
	}
	return abs(i - d)
}

func monotonicity(b uint64) int16 {
	n := monotonicityLeftRight(b)
	n += monotonicityTopDown(b)
	return n
}

func smoothness(b uint64) int16 {
	var n int16
	var y, x uint8
	for y = 0; y < nrLine; y++ {
		for x = 0; x < nrCol-1; x++ {
			n += abs(int16(get(b, y, x)) - int16(get(b, y, x+1)))
		}
	}
	for x = 0; x < nrCol; x++ {
		for y = 0; y < nrLine-1; y++ {
			n += abs(int16(get(b, y, x)) - int16(get(b, y+1, x)))
		}
	}
	return n
}

func freeTiles(b uint64) int16 {
	var y, x uint8
	var n int16
	for y = 0; y < nrLine; y++ {
		for x = 0; x < nrCol; x++ {
			if get(b, y, x) == 0 {
				n++
			}
		}
	}
	return n
}
