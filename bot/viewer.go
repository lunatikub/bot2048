package bot

import (
	"log"

	gc "github.com/gbin/goncurses"
)

// ViewerInit initialize viewer
func ViewerInit() *gc.Window {
	win, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	// Turn off character echo, hide the cursor and disable input buffering
	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)
	win.Clear()

	return win
}

const (
	sz     = 4
	xShift = 2
	yShift = 2
)

// ViewerRefresh dump the board
func ViewerRefresh(win *gc.Window, board uint64) {
	for i := 0; i < sz+1; i++ {
		win.MovePrint(i*2+yShift, xShift, "+----+----+----+----+")
	}
	for i := 0; i < sz; i++ {
		win.MovePrint(i*2+1+yShift, xShift, "|    |    |    |    |")
	}
	for y := 0; y < sz; y++ {
		for x := 0; x < sz; x++ {
			win.MovePrint(y*2+1+yShift, x+x*4+1+xShift,
				pow2[Get(board, uint8(y), uint8(x))])
		}
	}
	//time.Sleep(10 * time.Millisecond)
	win.Refresh()
}
