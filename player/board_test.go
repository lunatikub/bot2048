package player

import (
	"fmt"
	"testing"
)

type M [][]uint16

func dump(b uint64) {
	var y, x uint8
	for y = 0; y < nrLine; y++ {
		for x = 0; x < nrCol; x++ {
			fmt.Printf("%-4d", get(b, y, x))
		}
		fmt.Println()
	}
}

func build(v []uint16) uint16 {
	var l uint16
	n := 12
	for i := range v {
		l |= v[i] << n
		n -= 4
	}
	return l
}

func getExpected(t *testing.T, b uint64, y, x uint8, e uint16) {
	if r := get(b, y, x); uint16(r) != e {
		t.Errorf("get(%d,%d), expected %d, got: %d", y, x, e, r)
	}
}

func lineExpected(t *testing.T, b uint64, y uint8, e []uint16) {
	for x := range e {
		getExpected(t, b, y, uint8(x), e[x])
	}
}

func TestSetGetTile(t *testing.T) {
	var b uint64
	b = set(b, 0, 0, 1)
	b = set(b, 3, 2, 2)
	getExpected(t, b, 0, 0, 1)
	getExpected(t, b, 3, 2, 2)
	getExpected(t, b, 2, 2, 0)
}

func TestSetGetLine(t *testing.T) {
	var b uint64
	b = setLine(b, 0, build([]uint16{1, 1, 2, 1}))
	lineExpected(t, b, 0, []uint16{1, 1, 2, 1})
}

func TestSetGetCol(t *testing.T) {
	var b uint64
	b = setCol(b, 2, build([]uint16{1, 2, 3, 1}))
	lineExpected(t, b, 0, []uint16{0, 0, 1, 0})
	lineExpected(t, b, 1, []uint16{0, 0, 2, 0})
	lineExpected(t, b, 2, []uint16{0, 0, 3, 0})
	lineExpected(t, b, 3, []uint16{0, 0, 1, 0})

	var e uint16 = 1<<12 | 2<<8 | 3<<4 | 1
	if n := getCol(b, 2); n != e {
		t.Errorf("get col(2), expected %d, got: %d", e, n)
	}
}

func setBoard(init M) uint64 {
	var b uint64
	for y, line := range init {
		b = setLine(b, uint8(y), build(line))
	}
	return b
}

type moveFunc func(uint64) uint64

func testMove(t *testing.T, move moveFunc, init M, e M) {
	b := setBoard(init)
	b = move(b)
	for y, line := range e {
		lineExpected(t, b, uint8(y), line)
	}
}

func TestMoveLeft(t *testing.T) {
	testMove(t, moveLeft, M{{0, 0, 1, 1}}, M{{2, 0, 0, 0}})
	testMove(t, moveLeft, M{{0, 0, 1, 0}}, M{{1, 0, 0, 0}})
	testMove(t, moveLeft, M{{0, 2, 0, 0}}, M{{2, 0, 0, 0}})
	testMove(t, moveLeft, M{{2, 0, 0, 0}}, M{{2, 0, 0, 0}})
	testMove(t, moveLeft, M{{4, 2, 0, 0}}, M{{4, 2, 0, 0}})
	testMove(t, moveLeft, M{{8, 0, 4, 2}}, M{{8, 4, 2, 0}})
	testMove(t, moveLeft, M{{2, 0, 4, 2}}, M{{2, 4, 2, 0}})
	testMove(t, moveLeft, M{{4, 4, 4, 4}}, M{{5, 5, 0, 0}})
	testMove(t, moveLeft, M{{2, 1, 1, 2}}, M{{2, 2, 2, 0}})
	testMove(t, moveLeft, M{{4, 4, 4, 2}}, M{{5, 4, 2, 0}})
}

func TestMoveRight(t *testing.T) {
	testMove(t, moveRight, M{{1, 1, 0, 0}}, M{{0, 0, 0, 2}})
	testMove(t, moveRight, M{{0, 0, 1, 0}}, M{{0, 0, 0, 1}})
	testMove(t, moveRight, M{{0, 2, 0, 0}}, M{{0, 0, 0, 2}})
	testMove(t, moveRight, M{{2, 0, 0, 0}}, M{{0, 0, 0, 2}})
	testMove(t, moveRight, M{{4, 2, 0, 0}}, M{{0, 0, 4, 2}})
	testMove(t, moveRight, M{{8, 0, 4, 2}}, M{{0, 8, 4, 2}})
	testMove(t, moveRight, M{{2, 0, 4, 2}}, M{{0, 2, 4, 2}})
	testMove(t, moveRight, M{{4, 4, 4, 4}}, M{{0, 0, 5, 5}})
	testMove(t, moveRight, M{{2, 1, 1, 2}}, M{{0, 2, 2, 2}})
	testMove(t, moveRight, M{{4, 4, 4, 2}}, M{{0, 4, 5, 2}})
}

func TestMoveUp(t *testing.T) {
	testMove(t, moveUp, M{{0}, {0}, {1}, {1}}, M{{2}, {0}, {0}, {0}})
	testMove(t, moveUp, M{{3}, {3}, {3}, {3}}, M{{4}, {4}, {0}, {0}})
	testMove(t, moveUp, M{{2}, {1}, {1}, {2}}, M{{2}, {2}, {2}, {0}})
	testMove(t, moveUp, M{{4}, {4}, {4}, {2}}, M{{5}, {4}, {2}, {0}})
}

func TestMoveDown(t *testing.T) {
	testMove(t, moveDown, M{{1}, {1}, {0}, {0}}, M{{0}, {0}, {0}, {2}})
	testMove(t, moveDown, M{{3}, {3}, {3}, {3}}, M{{0}, {0}, {4}, {4}})
	testMove(t, moveDown, M{{2}, {1}, {1}, {2}}, M{{0}, {2}, {2}, {2}})
	testMove(t, moveDown, M{{2}, {4}, {4}, {4}}, M{{0}, {2}, {4}, {5}})
}

func testMono(t *testing.T, e int, init M) {
	b := setBoard(init)
	if r := monotonicity(b); e != r {
		t.Errorf("monotonicity, expected %d, got: %d", e, r)
	}
}

func TestMonotonicity(t *testing.T) {
	// best cases
	testMono(t, 24, M{
		{4, 5, 6, 7},
		{3, 4, 5, 6},
		{2, 3, 4, 5},
		{1, 2, 3, 4}})
	testMono(t, 24, M{
		{4, 3, 2, 1},
		{5, 4, 3, 2},
		{6, 5, 4, 3},
		{7, 6, 5, 4}})

	// worth case
	testMono(t, 0, M{
		{1, 2, 1, 2},
		{2, 1, 2, 1},
		{1, 2, 1, 2},
		{2, 1, 2, 1}})
}
