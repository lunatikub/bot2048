package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	gc "github.com/gbin/goncurses"
	bot "github.com/lunatikub/bot2048/bot"
)

type options struct {
	depth      int
	log        string
	logEnabled bool
	pretty     bool
}

func getOptions() *options {
	opts := new(options)
	flag.IntVar(&opts.depth, "depth", 3, "depth of the algorithm")
	flag.StringVar(&opts.log, "log", "", "log file")
	flag.BoolVar(&opts.pretty, "pretty", false, "dump the doard with ncurses")

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
	case bot.Left:
		return "left"
	case bot.Right:
		return "right"
	case bot.Up:
		return "up"
	case bot.Down:
		return "down"
	}
	panic("not an available move")
}

func main() {
	var board uint64   // main board
	var move int       // next best move
	var win *gc.Window // window for ncurses

	rand.Seed(time.Now().UTC().UnixNano())

	opts := getOptions()

	if opts.pretty {
		win = bot.ViewerInit()
	}

	empty := bot.GetEmptyTiles(board)
	board = bot.SetRandomTile(board, empty)
	empty = bot.GetEmptyTiles(board)
	board = bot.SetRandomTile(board, empty)

	for {
		move = bot.GetBestMove(board, opts.depth)
		board = bot.Move(board, move)
		if opts.logEnabled {
			log.Printf("[bot] move: %s", move2str(move))
		}
		if opts.pretty {
			bot.ViewerRefresh(win, board)
		}
		empty := bot.GetEmptyTiles(board)
		if len(empty) == 0 { // game over
			break
		}
		board = bot.SetRandomTile(board, empty)
	}
	bot.EndGameDump(board)
}
