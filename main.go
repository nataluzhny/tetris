package main

import (
	"fmt"
	"os"
	"os/exec"
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
	clearCmd *exec.Cmd
	shapeLoc Coord
	floorTop int8
	floor    [10]int8
	filled   [20][10]int8
}

func (game Game) drawFrame() {
	game.clearCmd.Run()

	fmt.Print("+==========+\n")
	for y, _ := range game.filled {
		y := int8(y)
		if int8(y) < game.shapeLoc.y || (int8(y) > game.shapeLoc.y && y < game.floorTop) {
			fmt.Println("|          |")
			continue
		}

		fmt.Print("|")
		for x, element := range game.filled[y] {
			x := int8(x)
			if x == game.shapeLoc.x && y == game.shapeLoc.y {
				fmt.Print("k")
			} else if element != 0 {
				fmt.Print("+")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("|\n")
	}

	fmt.Print("+==========+\n")
}

func main() {
	game := Game{
		clearCmd: exec.Command("clear"),
		shapeLoc: Coord{
			x: 0,
			y: 0,
		},
		floorTop: 0,
		floor:    [10]int8{},
		filled:   [20][10]int8{},
	}

	game.clearCmd.Stdout = os.Stdout

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
