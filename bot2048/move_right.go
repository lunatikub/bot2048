package bot2048

func shiftRight(b *B, c C) {
	x := c.X + 1
	for {
		if x == SZ || b.Board[c.Y][x] != 0 {
			break
		}
		b.swap(C{c.Y, x}, C{c.Y, x - 1})
		x++
	}
	b.merge(x < SZ, C{c.Y, x}, C{c.Y, x - 1})
}

func moveRight(b *B) {
	b.resetMerge()
	for y := range b.Board {
		for x := SZ - 2; x >= 0; x-- {
			if b.Board[y][x] != 0 {
				shiftRight(b, C{y, x})
			}
		}
	}
}
