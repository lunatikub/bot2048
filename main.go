package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/lunatikub/bot2048/brain"
	"github.com/lunatikub/bot2048/eye"
	"github.com/lunatikub/bot2048/hand"
)

type options struct {
	botEnabled bool
	screenID   int
	depth      int
	log        string
	logEnabled bool
}

func getOptions() *options {
	opts := new(options)
	flag.BoolVar(&opts.botEnabled, "enableBot", false, "enable bot for play2048.co")
	flag.IntVar(&opts.screenID, "screenID", 1, "screen identifier 2048 tab")
	flag.IntVar(&opts.depth, "depth", 3, "depth of the algorithm")
	flag.StringVar(&opts.log, "log", "", "log file")

	flag.Parse()

	if opts.log != "" {
		file, err := os.OpenFile(opts.log, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
		opts.logEnabled = true
	} else {
		opts.logEnabled = false
		log.SetOutput(ioutil.Discard)
	}

	return opts
}

func move2str(m int) string {
	switch m {
	case brain.Left:
		return "left"
	case brain.Right:
		return "right"
	case brain.Up:
		return "up"
	case brain.Down:
		return "down"
	}
	panic("not an available move")
}

func main() {
	var board uint64 // main board
	var h *hand.Hand // bot hand
	var e *eye.Eye   // bot eye
	var move int     // next best move

	rand.Seed(time.Now().UTC().UnixNano())

	opts := getOptions()

	if opts.botEnabled {
		e = eye.Init(opts.screenID)
		board = e.FindNewTile(board)
		board = e.FindNewTile(board)
		h = hand.Init()
	} else {
		empty := brain.GetEmptyTiles(board)
		board = brain.SetRandomTile(board, empty)
		empty = brain.GetEmptyTiles(board)
		board = brain.SetRandomTile(board, empty)
	}

	for {
		brain.Dump(board)
		move = brain.GetBestMove(board, opts.depth)
		board = brain.Move(board, move)
		if opts.logEnabled {
			log.Printf("[brain] best move: %s", move2str(move))
		}
		if opts.botEnabled {
			h.PressKey(move)
			board = e.FindNewTile(board)
		} else {
			empty := brain.GetEmptyTiles(board)
			if len(empty) == 0 { // game over
				break
			}
			board = brain.SetRandomTile(board, empty)
		}
	}
	brain.EndGameDump(board)
}
