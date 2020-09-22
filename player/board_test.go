package player

import (
	"fmt"
	"testing"
)

type M [][]uint16

const (
	sz = 4
)

func dump(b uint64) {
	var y, x uint8
	for y = 0; y < sz; y++ {
		for x = 0; x < sz; x++ {
			fmt.Print(getTile(b, y, x))
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
	//return v[0]<<12 | v[1]<<8 | v[2]<<4 | v[3]
}

func getExpected(t *testing.T, b uint64, y, x uint8, e uint16) {
	if r := getTile(b, y, x); uint16(r) != e {
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
	b = setTile(b, 0, 0, 1)
	b = setTile(b, 3, 2, 2)
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

type moveFunc func(uint64) uint64

func testMove(t *testing.T, move moveFunc, init M, e M) {
	var b uint64
	for y, line := range init {
		b = setLine(b, uint8(y), build(line))
	}
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

func TestMoveTop(t *testing.T) {
	testMove(t, moveTop, M{{0}, {0}, {1}, {1}}, M{{2}, {0}, {0}, {0}})
	testMove(t, moveTop, M{{3}, {3}, {3}, {3}}, M{{4}, {4}, {0}, {0}})
	testMove(t, moveTop, M{{2}, {1}, {1}, {2}}, M{{2}, {2}, {2}, {0}})
	testMove(t, moveTop, M{{4}, {4}, {4}, {2}}, M{{5}, {4}, {2}, {0}})
}

func TestMoveBottom(t *testing.T) {
	testMove(t, moveBottom, M{{1}, {1}, {0}, {0}}, M{{0}, {0}, {0}, {2}})
	testMove(t, moveBottom, M{{3}, {3}, {3}, {3}}, M{{0}, {0}, {4}, {4}})
	testMove(t, moveBottom, M{{2}, {1}, {1}, {2}}, M{{0}, {2}, {2}, {2}})
	testMove(t, moveBottom, M{{2}, {4}, {4}, {4}}, M{{0}, {2}, {4}, {5}})
}
