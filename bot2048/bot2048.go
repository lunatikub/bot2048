package bot2048

const (
	sz = 4
)

type cell struct {
	y int
	x int
}

// Bot2048 Artifical player of game 2048
type Bot2048 struct {
	board  [sz][sz]int
	merged [sz][sz]bool
}

func (b *Bot2048) resetBoard() {
	for y, row := range b.board {
		for x := range row {
			b.board[y][x] = 0
		}
	}
}

func (b *Bot2048) resetMerged() {
	for y, row := range b.board {
		for x := range row {
			b.merged[y][x] = false
		}
	}
}

func (b *Bot2048) eq(c1 cell, c2 cell) bool {
	return b.board[c1.y][c1.x] == b.board[c2.y][c2.x]
}

func (b *Bot2048) merge(dst cell, src cell) {
	b.board[dst.y][dst.x] = 2 * b.board[src.y][src.x]
	b.board[src.y][src.x] = 0
	b.merged[dst.y][dst.x] = true
}

func (b *Bot2048) move(dst cell, src cell) {
	b.board[dst.y][dst.x] = b.board[src.y][src.x]
	b.board[src.y][src.x] = 0
}

func (b *Bot2048) moveLeft(c cell) {
	x := c.x - 1
	for {
		if x == -1 || b.board[c.y][x] != 0 {
			break
		}
		b.move(cell{c.y, x}, cell{c.y, x + 1})
		x--
	}
	if x >= 0 && !b.merged[c.y][x] && b.eq(cell{c.y, x}, cell{c.y, x + 1}) {
		b.merge(cell{c.y, x}, cell{c.y, x + 1})
	}
}

func (b *Bot2048) toLeft() {
	b.resetMerged()
	for y := range b.board {
		for x := 1; x < sz; x++ {
			if b.board[y][x] != 0 {
				b.moveLeft(cell{y, x})
			}
		}
	}
}
