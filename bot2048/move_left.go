package bot2048

func shiftLeft(b *Bot2048, c cell) {
	x := c.x - 1
	for {
		if x == -1 || b.board[c.y][x] != 0 {
			break
		}
		b.move(cell{c.y, x}, cell{c.y, x + 1})
		x--
	}
	b.merge(x >= 0, cell{c.y, x}, cell{c.y, x + 1})
}

func moveLeft(b *Bot2048) {
	b.resetMerge()
	for y := range b.board {
		for x := 1; x < sz; x++ {
			if b.board[y][x] != 0 {
				shiftLeft(b, cell{y, x})
			}
		}
	}
}
