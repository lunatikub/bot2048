package bot2048

import (
	"log"
	"math"
	"strconv"
)

// Number of colums and rows of a board
const (
	SZ = 4
)

// C Cooridinate of a cell
type C struct {
	Y int
	X int
}

// B Bot player of game 2048
type B struct {
	Board   [SZ][SZ]int
	merged  [SZ][SZ]bool
	Score   int
	MaxVal  int
	nrMerge int
}

func (b *B) resetBoard() {
	for y, row := range b.Board {
		for x := range row {
			b.Board[y][x] = 0
		}
	}
}

func (b *B) resetMerge() {
	b.nrMerge = 0
	for y, row := range b.Board {
		for x := range row {
			b.merged[y][x] = false
		}
	}
}

func (b *B) reset() {
	b.resetBoard()
	b.resetMerge()
	b.Score = 0
}

func (b *B) eq(c1, c2 C) bool {
	return b.Board[c1.Y][c1.X] == b.Board[c2.Y][c2.X]
}

func (b *B) swap(dst, src C) {
	b.Board[dst.Y][dst.X] = b.Board[src.Y][src.X]
	b.Board[src.Y][src.X] = 0
}

func (b *B) merge(cond bool, dst, src C) {
	if cond && !b.merged[src.Y][src.X] && b.eq(src, dst) {
		b.Board[dst.Y][dst.X] = 2 * b.Board[src.Y][src.X]
		b.Board[src.Y][src.X] = 0
		b.merged[dst.Y][dst.X] = true
		b.nrMerge++
		b.Score += b.Board[dst.Y][dst.X]
		if b.MaxVal < b.Board[dst.Y][dst.X] {
			b.MaxVal = b.Board[dst.Y][dst.X]
		}
	}
}

const (
	left = iota
	right
	top
	bottom
	nrMove
)

func move2str(m int) string {
	switch m {
	case left:
		return "left"
	case right:
		return "right"
	case top:
		return "top"
	case bottom:
		return "bottom"
	}
	log.Panic("not available move")
	panic("not an available move")
}

func (b *B) move(m int) {
	switch m {
	case left:
		moveLeft(b)
	case right:
		moveRight(b)
	case top:
		moveTop(b)
	case bottom:
		moveBottom(b)
	}
}

func copy(bDst *B, bSrc *B) {
	for y, row := range bSrc.Board {
		for x, v := range row {
			bDst.Board[y][x] = v
		}
	}
	bDst.Score = bSrc.Score
}

func eq(b1 *B, b2 *B) bool {
	for y, row := range b2.Board {
		for x, v := range row {
			if b1.Board[y][x] != v {
				return false
			}
		}
	}
	return true
}

func evalMove(b *B, m int, w []int) (bool, int) {
	var newB B
	copy(&newB, b)
	newB.move(m)
	return !eq(b, &newB), newB.eval(w)
}

// Play the best move
func (b *B) Play(w []int) bool {
	bestMove := 0
	bestScore := math.MinInt64
	gameOver := true

	for m := 0; m < nrMove; m++ {
		if swap, score := evalMove(b, m, w); swap {
			log.Printf("[eval][%s] score: %d", move2str(m), score)
			if score > bestScore {
				bestMove = m
				bestScore = score
			}
			gameOver = false
		}
	}

	if gameOver {
		log.Println("!!! Game Over !!!!")
		return false
	}

	log.Printf("best move: %s", move2str(bestMove))
	b.move(bestMove)
	return true
}

// Dump the board
func (b *B) Dump(prefix string) {
	log.Println("[board]", prefix)
	for _, row := range b.Board {
		s := ""
		for _, v := range row {
			s = s + strconv.Itoa(v) + " "
		}
		log.Println(s)
	}
}
