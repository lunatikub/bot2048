package bot

import (
	"log"
	"math/rand"
)

// Tile a cell board
type Tile struct {
	y uint8
	x uint8
}

// enumeration of the agent type
const (
	Board = iota
	Player
)

// enumeration of the move type
const (
	Left = iota
	Right
	Up
	Down
	nrMove
)

var moves = []int{Left, Right, Up, Down}

var pow2 = []int{
	0, 2, 4, 8, 16, 32, 64, 128, 256, 512,
	1024, 2048, 4096, 8192, 16384, 32768,
}

// Move play a move
func Move(board uint64, move int) uint64 {
	switch move {
	case Left:
		return moveLeft(board)
	case Right:
		return moveRight(board)
	case Up:
		return moveUp(board)
	case Down:
		return moveDown(board)
	}
	panic("not an available move")
}

// GetEmptyTiles get the empty tile list
func GetEmptyTiles(board uint64) []Tile {
	var y, x uint8
	var empty []Tile
	for y = 0; y < nrLine; y++ {
		for x = 0; x < nrCol; x++ {
			if Get(board, y, x) == 0 {
				empty = append(empty, Tile{y, x})
			}
		}
	}
	return empty
}

const (
	ratio  = 2
	weight = 100000
)

// the optimal setup is given by a linear and monotonic
// decreasing order of the tile values
func eval(board uint64) int {
	maxScore := -1

	for _, path := range paths {
		w := weight
		score := 0
		for _, tile := range path {
			v := Get(board, tile.y, tile.x)
			score += pow2[v] * w
			w /= ratio
		}
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

// Agent player
func expectIMaxPlayer(board uint64, depth int) int {
	score := -1

	for m := range moves {
		newBoard := Move(board, m)
		if newBoard == board {
			continue
		}
		newScore := expectIMax(newBoard, depth-1, Board)
		if newScore > score {
			score = newScore
		}
	}
	return score
}

// Agent board
func expectIMaxBoard(board uint64, depth int) int {
	score := -1
	tiles := GetEmptyTiles(board)

	for _, tile := range tiles {
		newBoard := Set(board, tile.y, tile.x, 2)
		newScore := expectIMax(newBoard, depth-1, Player)
		if newScore != -1 {
			score += (newScore * 10) / 100 // 10% to pop 4
		}
		newBoard = Set(board, tile.y, tile.x, 1)
		newScore = expectIMax(newBoard, depth-1, Player)
		if newScore != -1 {
			score += (newScore * 90) / 100 // 90% to pop 2
		}
	}
	score /= len(tiles)
	return score
}

func expectIMax(board uint64, depth, agent int) int {
	if depth == 0 {
		return eval(board)
	}
	if agent == Player {
		return expectIMaxPlayer(board, depth)
	}
	return expectIMaxBoard(board, depth)
}

// GetBestMove return the best next move
func GetBestMove(board uint64, depth int) int {
	move := 0
	score := -1

	for m := range moves {
		newBoard := Move(board, m)
		if newBoard == board {
			continue
		}
		newScore := expectIMax(newBoard, depth-1, Board)
		if newScore > score {
			move = m
			score = newScore
		}
	}
	return move
}

// SetRandomTile Set a random tile with probabilities 2:90%, 4:10%
func SetRandomTile(board uint64, empty []Tile) uint64 {
	values := []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 4}
	c := empty[rand.Intn(len(empty))]
	v := values[rand.Intn(len(values))]
	log.Printf("[bot] set tile (y:%d,x:%d): %d", c.y, c.x, v)
	return Set(board, c.y, c.x, v)
}

// EndGameDump dump the board, score and the max tile value
func EndGameDump(b uint64) {
	n := 0
	var y, x, m uint8
	for y = 0; y < nrLine; y++ {
		for x = 0; x < nrCol; x++ {
			v := Get(b, y, x)
			if v > m {
				m = v
			}
			n += pow2[v]
		}
	}
	log.Println("[brain] Game over")
	log.Println("[brain] score", n)
	log.Println("[brain] best tile", pow2[m])
}
