package player

const (
	nrLine = 4
	nrCol  = 4
)

func shiftTile(y, x uint8) uint8 {
	return 60 - y*16 - x*4
}

func setTile(b uint64, y, x uint8, v uint64) uint64 {
	n := shiftTile(y, x)
	b &= ^(0xf << n) // clear the nibble
	return b | v<<n  // set the nibble
}

func getTile(b uint64, y, x uint8) uint8 {
	return uint8(b >> shiftTile(y, x) & 0xf)
}

func shiftLine(y uint8) uint8 {
	return 48 - y*16
}

func setLine(b uint64, y uint8, l uint16) uint64 {
	n := shiftLine(y)
	b &= ^(0xffff << n)     // clear the line
	return b | uint64(l)<<n // set the line
}

func getLine(b uint64, y uint8) uint16 {
	return uint16(b >> shiftLine(y) & 0xffff)
}

func moveLeft(b uint64) uint64 {
	b = setLine(b, 0, transLeftTop[getLine(b, 0)])
	b = setLine(b, 1, transLeftTop[getLine(b, 1)])
	b = setLine(b, 2, transLeftTop[getLine(b, 2)])
	b = setLine(b, 3, transLeftTop[getLine(b, 3)])
	return b
}
