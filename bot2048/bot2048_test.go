package bot2048

import (
	"fmt"
	"testing"
)

var b Bot2048

type M [][]int

func (b *Bot2048) setBoard(v M) {
	for y, row := range v {
		for x := range row {
			b.board[y][x] = v[y][x]
		}
	}
}

func (b *Bot2048) eqBoard(v M) bool {
	for y, row := range v {
		for x := range row {
			if b.board[y][x] != v[y][x] {
				return false
			}
		}
	}
	return true
}

func (b *Bot2048) dump() {
	fmt.Println()
	for _, row := range b.board {
		for _, v := range row {
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func TestMerge(test *testing.T) {
	b.resetBoard()
	b.setBoard(M{{2, 2, 0, 0}})
	b.merge(cell{0, 0}, cell{0, 1})
	if !b.eqBoard(M{{4, 0, 0, 0}}) {
		test.Errorf("merge: expected: %d, got: %d",
			M{{4, 0, 0, 0}}, b.board)
	}
}

func testMoveLeft(test *testing.T, b *Bot2048, init M, expected M) {
	b.setBoard(init)
	b.toLeft()
	if !b.eqBoard(expected) {
		test.Errorf("toLeft: expected %d, got: %d", expected, b.board)
	}
	b.resetBoard()
}

func TestToLeft(test *testing.T) {
	testMoveLeft(test, &b, M{{0, 0, 0, 2}}, M{{2, 0, 0, 0}})
	testMoveLeft(test, &b, M{{0, 0, 2, 0}}, M{{2, 0, 0, 0}})
	testMoveLeft(test, &b, M{{0, 2, 0, 0}}, M{{2, 0, 0, 0}})
	testMoveLeft(test, &b, M{{2, 0, 0, 0}}, M{{2, 0, 0, 0}})
	testMoveLeft(test, &b, M{{4, 2, 0, 0}}, M{{4, 2, 0, 0}})
	testMoveLeft(test, &b, M{{8, 0, 4, 2}}, M{{8, 4, 2, 0}})
	testMoveLeft(test, &b, M{{2, 0, 4, 2}}, M{{2, 4, 2, 0}})
	testMoveLeft(test, &b, M{{2, 0, 4, 2}}, M{{2, 4, 2, 0}})
	testMoveLeft(test, &b, M{{4, 4, 4, 4}}, M{{8, 8, 0, 0}})
}
