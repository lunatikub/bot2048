package bot2048

const (
	sz = 4
)

type cell struct {
	y int
	x int
}

// Bot2048 Bot player of game 2048
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

func (b *Bot2048) resetMerge() {
	for y, row := range b.board {
		for x := range row {
			b.merged[y][x] = false
		}
	}
}

func (b *Bot2048) eq(c1, c2 cell) bool {
	return b.board[c1.y][c1.x] == b.board[c2.y][c2.x]
}

func (b *Bot2048) move(dst, src cell) {
	b.board[dst.y][dst.x] = b.board[src.y][src.x]
	b.board[src.y][src.x] = 0
}

func (b *Bot2048) merge(cond bool, dst, src cell) {
	if cond && !b.merged[src.y][src.x] && b.eq(src, dst) {
		b.board[dst.y][dst.x] = 2 * b.board[src.y][src.x]
		b.board[src.y][src.x] = 0
		b.merged[dst.y][dst.x] = true
	}
}
