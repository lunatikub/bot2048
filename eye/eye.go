package eye

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/kbinani/screenshot"
	brain "github.com/lunatikub/bot2048/brain"
)

// 4x4 tiles 2048 board game
const (
	sz = 4
)

var borderColor = color.RGBA{187, 173, 160, 255}

type block struct {
	val uint8
	col color.Color
}

var blockColor = []block{
	{2, color.RGBA{238, 228, 218, 255}},
	{4, color.RGBA{237, 224, 200, 255}},
}

// Eye on a tab browser 2048 board game.
type Eye struct {
	tileSz int             // tile size
	w      int             // width
	h      int             // height
	shift  int             // shift coordinate to get a color from a tile
	xTLC   int             // x coordinate of the top left corner
	yTLC   int             // y coordinate of the top left corner
	xBRC   int             // x coordinate of the bottom right corner
	yBRC   int             // y coordinate of the bottom right corner
	bounds image.Rectangle // bounds of the screen
}

// find the top left corner
func findTLC(img *image.RGBA, bounds *image.Rectangle) (int, int, error) {
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			if borderColor == img.At(x, y) {
				return y, x, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("cannot find top left corner")
}

// find the top right corner
func findTRC(img *image.RGBA, bounds *image.Rectangle,
	yTLC, xTLC int) (int, int, error) {
	x := xTLC
	for {
		if borderColor != img.At(x, yTLC) {
			return yTLC, x - 1, nil
		}
		if x == bounds.Dx() {
			break
		}
		x++
	}
	return 0, 0, fmt.Errorf("cannot find top right corner")
}

// find the bottom right corner
func findBRC(img *image.RGBA, bounds *image.Rectangle,
	yTRC, xTRC int) (int, int, error) {
	y := yTRC
	for {
		if borderColor != img.At(xTRC, y) {
			return y - 1, xTRC, nil
		}
		if y == bounds.Dy() {
			break
		}
		y++
	}
	return 0, 0, fmt.Errorf("cannot find bottom right corner")
}

func (e *Eye) findCorners(img *image.RGBA, bounds *image.Rectangle) error {
	var err error
	var yTRC, xTRC int
	if e.yTLC, e.xTLC, err = findTLC(img, bounds); err != nil {
		return err
	}
	if yTRC, xTRC, err = findTRC(img, bounds, e.yTLC, e.xTLC); err != nil {
		return err
	}
	if e.yBRC, e.xBRC, err = findBRC(img, bounds, yTRC, xTRC); err != nil {
		return err
	}
	return nil
}

func (e *Eye) getTile(img *image.RGBA, y, x int) uint8 {
	x = e.xTLC + x*e.tileSz + e.shift
	y = e.yTLC + y*e.tileSz + e.shift

	for _, v := range blockColor {
		if v.col == img.At(x, y) {
			return v.val
		}
	}
	return 0
}

func (e *Eye) imgEq(img1, img2 *image.RGBA) bool {
	for x := 0; x < e.bounds.Dx(); x++ {
		for y := 0; y < e.bounds.Dy(); y++ {
			if img1.At(x, y) != img2.At(x, y) {
				return false
			}
		}
	}
	return true
}

func (e *Eye) findTile(b uint64) (uint8, uint8, uint8) {
	var y, x uint8
	img, _ := screenshot.CaptureRect(e.bounds)
	for y = 0; y < sz; y++ {
		for x = 0; x < sz; x++ {
			if brain.Get(b, y, x) == 0 {
				if v := e.getTile(img, int(y), int(x)); v != 0 {
					return y, x, v
				}
			}
		}
	}
	return 255, 255, 255
}

// FindNewTile find a new tile (2 or 4) on the 2048 game board
func (e *Eye) FindNewTile(b uint64) uint64 {
	n := 0
	for {
		y1, x1, v1 := e.findTile(b)
		time.Sleep(1 * time.Millisecond)
		y2, x2, v2 := e.findTile(b)
		if y2 == y1 || x2 == x1 || v1 == v2 && y1 != 255 {
			n++
		} else {
			n = 0
		}
		if n == 10 {
			log.Printf("detect new tile {y:%d, x:%d}: %d", y1, x1, v1)
			return brain.Set(b, y1, x1, v1/2)
		}
	}
}

func (e *Eye) setProperties() {
	e.w = e.xBRC - e.xTLC
	e.h = e.yBRC - e.yTLC
	e.tileSz = e.w / 4
	e.shift = e.tileSz / 3
}

// Init the eye
func Init(screenID int) *Eye {
	e := new(Eye)
	screenshot.NumActiveDisplays()
	e.bounds = screenshot.GetDisplayBounds(screenID)
	img, err := screenshot.CaptureRect(e.bounds)
	if err != nil {
		panic(err)
	}
	if err = e.findCorners(img, &e.bounds); err != nil {
		panic(err)
	}
	e.setProperties()
	log.Printf("[eye] top left corner: {y:%d,x:%d}", e.yTLC, e.xTLC)
	log.Printf("[eye] bottom right corner: {y:%d,x:%d}", e.yBRC, e.xBRC)
	log.Printf("[eye] width: %d, height: %d", e.w, e.h)
	log.Printf("[eye] tile size: %d, tile shift: %d", e.tileSz, e.shift)
	return e
}
