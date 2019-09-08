package main

import (
	"bufio"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

type Event int

const (
	Quit = iota
	Left
	Right
	Rotate
)

func (game *Game) getKeyboardInput(eventChan chan int) {
	for {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyCtrlC:
				eventChan <- 1
			case termbox.KeyArrowLeft:
				eventChan <- 2
				//game.moveLeft()
			case termbox.KeyArrowRight:
				game.moveRight()
			}

			//game.out.WriteString(fmt.Sprintf("Event Key detected: ", ev.Ch))
			//game.out.Flush()
		}
	}
}

type Coord struct {
	x int8
	y int8
}

type Shape interface {
	x() int8
	y() int8
	size() int8
	rotate()
}

type Game struct {
	out      *bufio.Writer
	shapeLoc Coord
	shape    Shape
	floorTop int8
	floor    [10]int8
	board    [20][10]int8
}

func (game Game) moveLeft() {
	game.shapeLoc.x--
}

func (game Game) moveRight() {
	game.shapeLoc.x++
}

func (game Game) drop() {

}

func (game Game) drawFrame() {
	game.out.WriteString("\033[2J\033[H")

	game.out.WriteString("+==========+\n")
	for y, _ := range game.board {
		y := int8(y)
		if int8(y) < game.shapeLoc.y || (int8(y) > game.shapeLoc.y && y < game.floorTop) {
			game.out.WriteString("|          |\n")
			continue
		}

		game.out.WriteString("|")
		for x, element := range game.board[y] {
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
	termbox.Init()
	defer termbox.Close()

	game := Game{
		out: bufio.NewWriterSize(os.Stdout, 1000),
		shapeLoc: Coord{
			x: 0,
			y: 0,
		},
		floorTop: 0,
		floor:    [10]int8{},
		board:    [20][10]int8{},
	}

	inputs := make(chan int)
	go game.getKeyboardInput(inputs)

	for {
		for j := 0; j < 20; j++ {
			game.drawFrame()
			time.Sleep(time.Millisecond * 30)

			for {
				select {
				case x, ok := <-inputs:
					if ok {
						switch x {
						case 1:
							return
						case 2:
							game.shapeLoc.x++
						}
					}
				default:
					goto EndFor
					//}
					//input := <- inputs
					//switch input{
					//case 1:
					//	return
					//case 2:
					//	game.shapeLoc.x++
					//case
				}
			}
		EndFor:
		}

		if game.shapeLoc.y < 19 {
			game.shapeLoc.y++
		}
	}
}
