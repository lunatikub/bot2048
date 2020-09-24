package brain

const (
	nrLine     = 4
	nrCol      = 4
	nibbleBits = 4
	lineBits   = 16
	boardBits  = 64
	nibbleMask = 0xf
	lineMask   = 0xffff
)

// return the shift to get a tile
func shiftTile(y, x uint8) uint8 {
	return boardBits - y*lineBits - (x+1)*nibbleBits
}

// Set a tile
func Set(b uint64, y, x, v uint8) uint64 {
	n := shiftTile(y, x)
	b &= ^(nibbleMask << n) // clear the nibble
	return b | uint64(v)<<n // set the nibble
}

// Get a tile
func Get(b uint64, y, x uint8) uint8 {
	return uint8(b >> shiftTile(y, x) & nibbleMask)
}

// return the shift to get a line
func shiftLine(y uint8) uint8 {
	return boardBits - (y+1)*lineBits
}

func setLine(b uint64, y uint8, l uint16) uint64 {
	n := shiftLine(y)
	b &= ^(lineMask << n)   // clear the line
	return b | uint64(l)<<n // set the uint16
}

func getLine(b uint64, y uint8) uint16 {
	return uint16(b >> shiftLine(y))
}

func setCol(b uint64, x uint8, c uint16) uint64 {
	b = Set(b, 0, x, uint8(c>>(lineBits-nibbleBits)&nibbleMask))
	b = Set(b, 1, x, uint8(c>>(lineBits-2*nibbleBits)&nibbleMask))
	b = Set(b, 2, x, uint8(c>>(lineBits-3*nibbleBits)&nibbleMask))
	b = Set(b, 3, x, uint8(c&nibbleMask))
	return b
}

func getCol(b uint64, x uint8) uint16 {
	return uint16(Get(b, 0, x))<<(lineBits-nibbleBits) |
		uint16(Get(b, 1, x))<<(lineBits-2*nibbleBits) |
		uint16(Get(b, 2, x))<<(lineBits-3*nibbleBits) |
		uint16(Get(b, 3, x))
}

func moveLeft(b uint64) uint64 {
	b = setLine(b, 0, transLeftUp[getLine(b, 0)])
	b = setLine(b, 1, transLeftUp[getLine(b, 1)])
	b = setLine(b, 2, transLeftUp[getLine(b, 2)])
	b = setLine(b, 3, transLeftUp[getLine(b, 3)])
	return b
}

func moveRight(b uint64) uint64 {
	b = setLine(b, 0, transRightDown[getLine(b, 0)])
	b = setLine(b, 1, transRightDown[getLine(b, 1)])
	b = setLine(b, 2, transRightDown[getLine(b, 2)])
	b = setLine(b, 3, transRightDown[getLine(b, 3)])
	return b
}

func moveUp(b uint64) uint64 {
	b = setCol(b, 0, transLeftUp[getCol(b, 0)])
	b = setCol(b, 1, transLeftUp[getCol(b, 1)])
	b = setCol(b, 2, transLeftUp[getCol(b, 2)])
	b = setCol(b, 3, transLeftUp[getCol(b, 3)])
	return b
}

func moveDown(b uint64) uint64 {
	b = setCol(b, 0, transRightDown[getCol(b, 0)])
	b = setCol(b, 1, transRightDown[getCol(b, 1)])
	b = setCol(b, 2, transRightDown[getCol(b, 2)])
	b = setCol(b, 3, transRightDown[getCol(b, 3)])
	return b
}
