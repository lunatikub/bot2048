package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	bot2048 "github.com/lunatikub/bot2048/bot2048"
)

func setRmd(b *bot2048.B) {
	var empty []bot2048.C
	values := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 4}

	for y, row := range b.Board {
		for x, v := range row {
			if v == 0 {
				empty = append(empty, bot2048.C{Y: y, X: x})
			}
		}
	}
	c := empty[rand.Intn(len(empty))]
	v := values[rand.Intn(len(values))]
	log.Println("[set]", c, v)
	b.Board[c.Y][c.X] = v
}

func setLog() {
	file, err := os.OpenFile("bot2048.log",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	//log.SetOutput(ioutil.Discard)
}

func main() {
	setLog()
	rand.Seed(time.Now().UnixNano())

	w := []int{1, 1}

	var b bot2048.B
	setRmd(&b)
	setRmd(&b)

	for {
		b.Dump("main")
		if r := b.Play(w); !r {
			break
		}
		setRmd(&b)
	}
	fmt.Println(b.Score, b.MaxVal)
}
