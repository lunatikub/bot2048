package bot2048

func shiftLeft(b *B, c C) {
	x := c.X - 1
	for {
		if x == -1 || b.Board[c.Y][x] != 0 {
			break
		}
		b.swap(C{c.Y, x}, C{c.Y, x + 1})
		x--
	}
	b.merge(x >= 0, C{c.Y, x}, C{c.Y, x + 1})
}

func moveLeft(b *B) {
	b.resetMerge()
	for y := range b.Board {
		for x := 1; x < SZ; x++ {
			if b.Board[y][x] != 0 {
				shiftLeft(b, C{y, x})
			}
		}
	}

}
