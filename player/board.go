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

func setCol(b uint64, x uint8, c uint16) uint64 {
	b = setTile(b, 0, x, uint64(c>>12&0xf))
	b = setTile(b, 1, x, uint64(c>>8&0xf))
	b = setTile(b, 2, x, uint64(c>>4&0xf))
	b = setTile(b, 3, x, uint64(c&0xf))
	return b
}

func getCol(b uint64, x uint8) uint16 {
	return uint16(getTile(b, 0, x))<<12 |
		uint16(getTile(b, 1, x))<<8 |
		uint16(getTile(b, 2, x))<<4 |
		uint16(getTile(b, 3, x))
}

func moveLeft(b uint64) uint64 {
	b = setLine(b, 0, transLeftTop[getLine(b, 0)])
	b = setLine(b, 1, transLeftTop[getLine(b, 1)])
	b = setLine(b, 2, transLeftTop[getLine(b, 2)])
	b = setLine(b, 3, transLeftTop[getLine(b, 3)])
	return b
}

func moveRight(b uint64) uint64 {
	b = setLine(b, 0, transRightBottom[getLine(b, 0)])
	b = setLine(b, 1, transRightBottom[getLine(b, 1)])
	b = setLine(b, 2, transRightBottom[getLine(b, 2)])
	b = setLine(b, 3, transRightBottom[getLine(b, 3)])
	return b
}

func moveTop(b uint64) uint64 {
	b = setCol(b, 0, transLeftTop[getCol(b, 0)])
	b = setCol(b, 1, transLeftTop[getCol(b, 1)])
	b = setCol(b, 2, transLeftTop[getCol(b, 2)])
	b = setCol(b, 3, transLeftTop[getCol(b, 3)])
	return b
}

func moveBottom(b uint64) uint64 {
	b = setCol(b, 0, transRightBottom[getCol(b, 0)])
	b = setCol(b, 1, transRightBottom[getCol(b, 1)])
	b = setCol(b, 2, transRightBottom[getCol(b, 2)])
	b = setCol(b, 3, transRightBottom[getCol(b, 3)])
	return b
}
