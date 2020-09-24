package player

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	left = iota
	right
	up
	down
	nrMove
)

var pow2 = []int{
	0, 2, 4, 8, 16, 32, 64, 128, 256, 512,
	1024, 2048, 4096, 8192, 16384, 32768,
}

func move(b uint64, m int) uint64 {
	switch m {
	case left:
		return moveLeft(b)
	case right:
		return moveRight(b)
	case up:
		return moveUp(b)
	case down:
		return moveDown(b)
	}
	panic("not an available move")
}

// Dump a board
func Dump(b uint64) {
	var y, x uint8
	for y = 0; y < nrLine; y++ {
		for x = 0; x < nrCol; x++ {
			v := pow2[get(b, y, x)]
			fmt.Printf("%-5d", v)
		}
		fmt.Println()
	}
}

func nextMove(b uint64, d, md int) (int, int) {
	bestMove := 0
	bestScore := -1

	for m := 0; m < nrMove; m++ {
		bprime := b // copy the board
		bprime = move(bprime, m)
		if bprime != b {
			s, c := eval(bprime)
			bprime = set(bprime, c.y, c.x, 1)
			if d != 0 {
				_, myS := nextMove(bprime, d-1, md)
				s += int(float64(myS) * math.Pow(0.9, float64(md-d+1)))
			}
			if s > bestScore {
				bestMove = m
				bestScore = s
			}
		}
	}
	return bestMove, bestScore
}

// Play the next best move
func Play(b uint64) uint64 {
	m, _ := nextMove(b, 3, 3)
	b = move(b, m)
	return b
}

type tile struct {
	y uint8
	x uint8
}

// SetRmdTile Set a random tile with probabilities 2:90%, 4:10%
func SetRmdTile(b uint64) (uint64, bool) {
	var y, x uint8
	var empty []tile
	values := []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
	for y = 0; y < nrLine; y++ {
		for x = 0; x < nrCol; x++ {
			if get(b, y, x) == 0 {
				empty = append(empty, tile{y, x})
			}
		}
	}
	if len(empty) == 0 {
		return b, false
	}
	c := empty[rand.Intn(len(empty))]
	v := values[rand.Intn(len(values))]
	return set(b, c.y, c.x, v), true
}
