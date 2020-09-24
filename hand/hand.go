package hand

import (
	"log"
	"time"

	"github.com/lunatikub/bot2048/brain"
	"github.com/micmonay/keybd_event"
)

// Hand simulate key press
type Hand struct {
	kb keybd_event.KeyBonding
}

// Init the hand
func Init() *Hand {
	var err error
	h := new(Hand)
	h.kb, err = keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	//	For linux, it is very important to wait almost 2 seconds
	log.Println("game will start in 3 seconds...")
	time.Sleep(3 * time.Second)
	return h
}

// PressKey simulate press key
func (h *Hand) PressKey(move int) {
	h.kb.Clear()
	switch move {
	case brain.Right:
		h.kb.SetKeys(keybd_event.VK_RIGHT)
	case brain.Left:
		h.kb.SetKeys(keybd_event.VK_LEFT)
	case brain.Up:
		h.kb.SetKeys(keybd_event.VK_UP)
	case brain.Down:
		h.kb.SetKeys(keybd_event.VK_DOWN)
	}

	h.kb.Press()
	time.Sleep(50 * time.Millisecond)
	h.kb.Release()
}
