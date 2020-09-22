package bot2048

import (
	"testing"
)

var b B

// Matrix type
type M [][]int

func setBoard(v M) {
	for y, row := range v {
		for x := range row {
			b.Board[y][x] = v[y][x]
		}
	}
}

func eqBoard(v M) bool {
	for y, row := range v {
		for x := range row {
			if b.Board[y][x] != v[y][x] {
				return false
			}
		}
	}
	return true
}

type moveFunc func(*B)

func testMove(test *testing.T, move moveFunc, init M, expected M) {
	setBoard(init)
	move(&b)
	if !eqBoard(expected) {
		test.Errorf("init: %d, expected %d, got: %d", init, expected, b.Board)
	}
}

func TestMoveLeft(test *testing.T) {
	b.reset()
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
	b.reset()
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
	b.reset()
	testMove(test, moveTop, M{{0}, {0}, {0}, {2}}, M{{2}, {0}, {0}, {0}})
	testMove(test, moveTop, M{{0}, {0}, {2}, {0}}, M{{2}, {0}, {0}, {0}})
	testMove(test, moveTop, M{{0}, {2}, {0}, {0}}, M{{2}, {0}, {0}, {0}})
	testMove(test, moveTop, M{{2}, {0}, {0}, {0}}, M{{2}, {0}, {0}, {0}})
	testMove(test, moveTop, M{{4}, {2}, {0}, {0}}, M{{4}, {2}, {0}, {0}})
	testMove(test, moveTop, M{{8}, {0}, {4}, {2}}, M{{8}, {4}, {2}, {0}})
	testMove(test, moveTop, M{{2}, {0}, {4}, {2}}, M{{2}, {4}, {2}, {0}})
	testMove(test, moveTop, M{{4}, {4}, {4}, {4}}, M{{8}, {8}, {0}, {0}})
}

func TestMoveBottom(test *testing.T) {
	b.reset()
	testMove(test, moveBottom, M{{2}, {0}, {0}, {0}}, M{{0}, {0}, {0}, {2}})
	testMove(test, moveBottom, M{{0}, {2}, {0}, {0}}, M{{0}, {0}, {0}, {2}})
	testMove(test, moveBottom, M{{0}, {0}, {2}, {0}}, M{{0}, {0}, {0}, {2}})
	testMove(test, moveBottom, M{{0}, {0}, {0}, {2}}, M{{0}, {0}, {0}, {2}})
	testMove(test, moveBottom, M{{0}, {0}, {2}, {4}}, M{{0}, {0}, {2}, {4}})
	testMove(test, moveBottom, M{{2}, {4}, {0}, {8}}, M{{0}, {2}, {4}, {8}})
	testMove(test, moveBottom, M{{2}, {4}, {0}, {2}}, M{{0}, {2}, {4}, {2}})
	testMove(test, moveBottom, M{{4}, {4}, {4}, {4}}, M{{0}, {0}, {8}, {8}})
}

func TestScore(test *testing.T) {
	b.reset()
	setBoard(M{{2, 2}})
	moveLeft(&b)
	b.Board[1][0] = 2
	moveRight(&b)
	b.Board[1][0] = 2
	moveRight(&b)
	b.Board[2][0] = 2
	moveTop(&b)
	b.Board[0][1] = 2
	moveRight(&b)
	scoreExpected := 20
	if b.Score != scoreExpected {
		test.Errorf("score expected: %d, got %d", scoreExpected, b.Score)
	}
}
