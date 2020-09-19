package bot2048

func shiftBottom(b *Bot2048, c cell) {
	y := c.y + 1
	for {
		if y == sz || b.board[y][c.x] != 0 {
			break
		}
		b.move(cell{y, c.x}, cell{y - 1, c.x})
		y++
	}
	b.merge(y < sz, cell{y, c.x}, cell{y - 1, c.x})
}

func moveBottom(b *Bot2048) {
	b.resetMerge()
	for x := 0; x < sz; x++ {
		for y := sz - 2; y >= 0; y-- {
			if b.board[y][x] != 0 {
				shiftBottom(b, cell{y, x})
			}
		}
	}
}
