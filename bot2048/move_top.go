package bot2048

func shiftTop(b *Bot2048, c cell) {
	y := c.y - 1
	for {
		if y == -1 || b.board[y][c.x] != 0 {
			break
		}
		b.move(cell{y, c.x}, cell{y + 1, c.x})
		y--
	}
	b.merge(y >= 0, cell{y, c.x}, cell{y + 1, c.x})
}

func moveTop(b *Bot2048) {
	b.resetMerge()
	for x := 0; x < sz; x++ {
		for y := 1; y < sz; y++ {
			if b.board[y][x] != 0 {
				shiftTop(b, cell{y, x})
			}
		}
	}
}
