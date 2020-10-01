package bot

import (
	"log"
	"math/rand"
)

// BoardSZ number of lines and colums of a board
const BoardSZ = 4

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

var nrBoardEvaluated int // stats

var moves = []int{Left, Right, Up, Down}

// Power of 2 (except 0 -> 0 instead of 1)
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

func expectIMaxPlayer(board uint64, depth int) (int, int) {
	score := -1
	move := 0

	for m := range moves {
		newBoard := Move(board, m)
		if newBoard == board {
			continue
		}
		newScore, _ := expectIMax(newBoard, depth-1, Board)
		if newScore > score {
			score = newScore
			move = m
		}
	}
	return score, move
}

func expectIMaxBoard(board uint64, depth int) (int, int) {
	score := -1
	tiles := GetEmptyTiles(board)

	for _, tile := range tiles {
		newBoard := Set(board, tile.y, tile.x, 2)
		newScore, _ := expectIMax(newBoard, depth-1, Player)
		if newScore != -1 {
			score += (newScore * 10) / 100 // 10% to pop 4
		}
		newBoard = Set(board, tile.y, tile.x, 1)
		newScore, _ = expectIMax(newBoard, depth-1, Player)
		if newScore != -1 {
			score += (newScore * 90) / 100 // 90% to pop 2
		}
	}
	score /= len(tiles)
	return score, 0
}

func expectIMax(board uint64, depth, agent int) (int, int) {
	if depth == 0 {
		nrBoardEvaluated++
		return eval(board), 0
	}
	if agent == Player {
		return expectIMaxPlayer(board, depth)
	}
	return expectIMaxBoard(board, depth)
}

// GetBestMove return the best next move
func GetBestMove(board uint64, depth int) int {
	_, move := expectIMax(board, depth, Player)
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

// GetStats return the maximum value of the board, the score
// and the number of board evaluated
func GetStats(board uint64) (int, int, int) {
	score := 0
	var y, x, max uint8
	for y = 0; y < nrLine; y++ {
		for x = 0; x < nrCol; x++ {
			v := Get(board, y, x)
			if v > max {
				max = v
			}
			score += pow2[v]
		}
	}
	return pow2[max], score, nrBoardEvaluated
}

// EndGameDump dump the board, score and the max tile value
func EndGameDump(board uint64) {
	max, score, _ := GetStats(board)
	log.Println("[brain] Game over")
	log.Println("[brain] score", score)
	log.Println("[brain] best tile", max)
	log.Println("[brain] number of board evaluated", nrBoardEvaluated)
}
