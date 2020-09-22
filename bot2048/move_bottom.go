package bot2048

func shiftBottom(b *B, c C) {
	y := c.Y + 1
	for {
		if y == SZ || b.Board[y][c.X] != 0 {
			break
		}
		b.swap(C{y, c.X}, C{y - 1, c.X})
		y++
	}
	b.merge(y < SZ, C{y, c.X}, C{y - 1, c.X})
}

func moveBottom(b *B) {
	b.resetMerge()
	for x := 0; x < SZ; x++ {
		for y := SZ - 2; y >= 0; y-- {
			if b.Board[y][x] != 0 {
				shiftBottom(b, C{y, x})
			}
		}
	}
}
