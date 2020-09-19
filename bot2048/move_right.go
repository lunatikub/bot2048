package bot2048

func shiftRight(b *Bot2048, c cell) {
	x := c.x + 1
	for {
		if x == sz || b.board[c.y][x] != 0 {
			break
		}
		b.move(cell{c.y, x}, cell{c.y, x - 1})
		x++
	}
	b.merge(x < sz, cell{c.y, x}, cell{c.y, x - 1})
}

func moveRight(b *Bot2048) {
	b.resetMerge()
	for y := range b.board {
		for x := sz - 2; x >= 0; x-- {
			if b.board[y][x] != 0 {
				shiftRight(b, cell{y, x})
			}
		}
	}
}
