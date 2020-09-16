package bot2048

import (
	"fmt"
	"testing"
)

var b Bot2048

// Matrix type
type M [][]int

func setBoard(v M) {
	for y, row := range v {
		for x := range row {
			b.board[y][x] = v[y][x]
		}
	}
}

func eqBoard(v M) bool {
	for y, row := range v {
		for x := range row {
			if b.board[y][x] != v[y][x] {
				return false
			}
		}
	}
	return true
}

func dump() {
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
	setBoard(M{{2, 2, 0, 0}})
	b.merge(cell{0, 0}, cell{0, 1})
	if !eqBoard(M{{4, 0, 0, 0}}) {
		test.Errorf("expected: %d, got: %d",
			M{{4, 0, 0, 0}}, b.board)
	}
}

type moveFunc func(*Bot2048)

func testMove(test *testing.T, move moveFunc, init M, expected M) {
	setBoard(init)
	move(&b)
	if !eqBoard(expected) {
		test.Errorf("init: %d, expected %d, got: %d", init, expected, b.board)
	}
}

func TestMoveLeft(test *testing.T) {
	b.resetBoard()
	testMove(test, moveLeft, M{{0, 0, 0, 2}}, M{{2, 0, 0, 0}})
	testMove(test, moveLeft, M{{0, 0, 2, 0}}, M{{2, 0, 0, 0}})
	testMove(test, moveLeft, M{{0, 2, 0, 0}}, M{{2, 0, 0, 0}})
	testMove(test, moveLeft, M{{2, 0, 0, 0}}, M{{2, 0, 0, 0}})
	testMove(test, moveLeft, M{{4, 2, 0, 0}}, M{{4, 2, 0, 0}})
	testMove(test, moveLeft, M{{8, 0, 4, 2}}, M{{8, 4, 2, 0}})
	testMove(test, moveLeft, M{{2, 0, 4, 2}}, M{{2, 4, 2, 0}})
	testMove(test, moveLeft, M{{4, 4, 4, 4}}, M{{8, 8, 0, 0}})
}

func TestMoveRight(test *testing.T) {
	b.resetBoard()
	testMove(test, moveRight, M{{2, 0, 0, 0}}, M{{0, 0, 0, 2}})
	testMove(test, moveRight, M{{0, 2, 0, 0}}, M{{0, 0, 0, 2}})
	testMove(test, moveRight, M{{0, 0, 2, 0}}, M{{0, 0, 0, 2}})
	testMove(test, moveRight, M{{0, 0, 0, 2}}, M{{0, 0, 0, 2}})
	testMove(test, moveRight, M{{0, 0, 2, 4}}, M{{0, 0, 2, 4}})
	testMove(test, moveRight, M{{2, 4, 0, 8}}, M{{0, 2, 4, 8}})
	testMove(test, moveRight, M{{2, 4, 0, 2}}, M{{0, 2, 4, 2}})
	testMove(test, moveRight, M{{4, 4, 4, 4}}, M{{0, 0, 8, 8}})
}

func TestMoveTop(test *testing.T) {
	b.resetBoard()
	testMove(test, moveTop, M{{0}, {0}, {0}, {2}}, M{{2}, {0}, {0}, {0}})
	testMove(test, moveTop, M{{0}, {0}, {2}, {0}}, M{{2}, {0}, {0}, {0}})
	testMove(test, moveTop, M{{0}, {2}, {0}, {0}}, M{{2}, {0}, {0}, {0}})
	testMove(test, moveTop, M{{2}, {0}, {0}, {0}}, M{{2}, {0}, {0}, {0}})
	testMove(test, moveTop, M{{4}, {2}, {0}, {0}}, M{{4}, {2}, {0}, {0}})
	testMove(test, moveTop, M{{8}, {0}, {4}, {2}}, M{{8}, {4}, {2}, {0}})
	testMove(test, moveTop, M{{2}, {0}, {4}, {2}}, M{{2}, {4}, {2}, {0}})
	testMove(test, moveTop, M{{4}, {4}, {4}, {4}}, M{{8}, {8}, {0}, {0}})
}
