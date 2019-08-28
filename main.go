package main

import (
	"bufio"
	//"fmt"
	//"io"
	"os"
	"time"
)

const boardsize int = 200

type Coord struct {
	x int8
	y int8
}

type Shape struct {
	bool
}

type Game struct {
	out *bufio.Writer
	shapeLoc Coord
	floorTop int8
	floor    [10]int8
	filled   [20][10]int8
}

func (game Game) drawFrame() {
	game.out.WriteString("\033[2J\033[H")

	game.out.WriteString("+==========+\n")
	for y, _ := range game.filled {
		y := int8(y)
		if int8(y) < game.shapeLoc.y || (int8(y) > game.shapeLoc.y && y < game.floorTop) {
			game.out.WriteString("|          |\n")
			continue
		}

		game.out.WriteString("|")
		for x, element := range game.filled[y] {
			x := int8(x)
			if x == game.shapeLoc.x && y == game.shapeLoc.y {
				game.out.WriteString("k")
			} else if element != 0 {
				game.out.WriteString("+")
			} else {
				game.out.WriteString(" ")
			}
		}
		game.out.WriteString("|\n")
	}

	game.out.WriteString("+==========+\n")
	game.out.Flush()
}

func main() {
	game := Game{
		out: bufio.NewWriterSize(os.Stdout, 1000),
		shapeLoc: Coord{
			x: 0,
			y: 0,
		},
		floorTop: 0,
		floor:    [10]int8{},
		filled:   [20][10]int8{},
	}

	for {
		for j := 0; j < 20; j++ {
			game.drawFrame()
			time.Sleep(time.Millisecond * 30)
		}

		if game.shapeLoc.y < 19 {
			game.shapeLoc.y++
		}
	}
}
