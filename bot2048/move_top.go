package bot2048

func shiftTop(b *B, c C) {
	y := c.Y - 1
	for {
		if y == -1 || b.Board[y][c.X] != 0 {
			break
		}
		b.swap(C{y, c.X}, C{y + 1, c.X})
		y--
	}
	b.merge(y >= 0, C{y, c.X}, C{y + 1, c.X})
}

func moveTop(b *B) {
	b.resetMerge()
	for x := 0; x < SZ; x++ {
		for y := 1; y < SZ; y++ {
			if b.Board[y][x] != 0 {
				shiftTop(b, C{y, x})
			}
		}
	}
}
