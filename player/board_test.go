package player

import (
	"fmt"
	"testing"
)

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

func buildLine(v []uint16) uint16 {
	return v[0]<<12 | v[1]<<8 | v[2]<<4 | v[3]
}

func getExpected(t *testing.T, b uint64, y, x, e uint8) {
	if r := getTile(b, y, x); r != e {
		t.Errorf("get(%d,%d), expected %d, got: %d", y, x, e, r)
	}
}

func lineExpected(t *testing.T, b uint64, y uint8, e []uint8) {
	var x uint8
	for x = 0; x < sz; x++ {
		getExpected(t, b, y, x, e[x])
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
	b = setLine(b, 0, buildLine([]uint16{1, 1, 2, 1}))
	lineExpected(t, b, 0, []uint8{1, 1, 2, 1})
}

type moveFunc func(uint64) uint64

func testMove(t *testing.T, move moveFunc, l []uint16, e []uint8) {
	var b uint64
	b = setLine(b, 0, buildLine(l))
	b = move(b)
	lineExpected(t, b, 0, e)
}

func TestMoveLeft(t *testing.T) {
	testMove(t, moveLeft, []uint16{0, 0, 1, 1}, []uint8{2, 0, 0, 0})
	testMove(t, moveLeft, []uint16{0, 0, 1, 0}, []uint8{1, 0, 0, 0})
	testMove(t, moveLeft, []uint16{0, 2, 0, 0}, []uint8{2, 0, 0, 0})
	testMove(t, moveLeft, []uint16{2, 0, 0, 0}, []uint8{2, 0, 0, 0})
	testMove(t, moveLeft, []uint16{4, 2, 0, 0}, []uint8{4, 2, 0, 0})
	testMove(t, moveLeft, []uint16{8, 0, 4, 2}, []uint8{8, 4, 2, 0})
	testMove(t, moveLeft, []uint16{2, 0, 4, 2}, []uint8{2, 4, 2, 0})
	testMove(t, moveLeft, []uint16{4, 4, 4, 4}, []uint8{5, 5, 0, 0})
	testMove(t, moveLeft, []uint16{2, 1, 1, 2}, []uint8{2, 2, 2, 0})
	testMove(t, moveLeft, []uint16{4, 4, 4, 2}, []uint8{5, 4, 2, 0})
}

func TestMoveRight(t *testing.T) {
	testMove(t, moveRight, []uint16{1, 1, 0, 0}, []uint8{0, 0, 0, 2})
	testMove(t, moveRight, []uint16{0, 0, 1, 0}, []uint8{0, 0, 0, 1})
	testMove(t, moveRight, []uint16{0, 2, 0, 0}, []uint8{0, 0, 0, 2})
	testMove(t, moveRight, []uint16{2, 0, 0, 0}, []uint8{0, 0, 0, 2})
	testMove(t, moveRight, []uint16{4, 2, 0, 0}, []uint8{0, 0, 4, 2})
	testMove(t, moveRight, []uint16{8, 0, 4, 2}, []uint8{0, 8, 4, 2})
	testMove(t, moveRight, []uint16{2, 0, 4, 2}, []uint8{0, 2, 4, 2})
	testMove(t, moveRight, []uint16{4, 4, 4, 4}, []uint8{0, 0, 5, 5})
	testMove(t, moveRight, []uint16{2, 1, 1, 2}, []uint8{0, 2, 2, 2})
	testMove(t, moveRight, []uint16{4, 4, 4, 2}, []uint8{0, 4, 5, 2})
}
