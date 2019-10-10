package main

import (
	"bufio"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

type KeyboardEvent int

const (
	Quit KeyboardEvent = iota
	Left
	Right
	Rotate
	Drop
)

func (game *Game) getKeyboardInput(eventChan chan KeyboardEvent) {
	for {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyCtrlC:
				eventChan <- Quit
			case termbox.KeyArrowLeft:
				eventChan <- Left
				//game.moveLeft()
			case termbox.KeyArrowRight:
				eventChan <- Right
			case termbox.KeyArrowDown:
				eventChan <- Drop
			}

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
	out         *bufio.Writer
	shapeLoc    Coord
	shape       Shape
	floorTop    int8
	floor       [10]int8
	board       [20][10]int8
	boardHeight int8
}

func (game Game) drawFrame() {
	game.out.WriteString("\033[2J\033[H")

	game.out.WriteString("+==========+\n")
	for y, _ := range game.board {
		y := int8(y)
		if y < game.shapeLoc.y {
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

//func (game Game) isSpaceBelow() {
//	if

func main() {
	termbox.Init()
	defer termbox.Close()

	floor := [10]int8{}
	for i := range floor {
		floor[i] = 20
	}

	game := Game{
		out: bufio.NewWriterSize(os.Stdout, 1000),
		shapeLoc: Coord{
			x: 0,
			y: 0,
		},
		floorTop: 0,
		floor:    floor,
		board:    [20][10]int8{},
	}

	inputs := make(chan KeyboardEvent)
	go game.getKeyboardInput(inputs)

	for {
		for j := 0; j < 20; j++ {
			game.drawFrame()
			time.Sleep(time.Millisecond * 10)

			for {
				select {
				case x, ok := <-inputs:
					if ok {
						switch x {
						case Quit:
							return
						case Left:
							game.shapeLoc.x--
						case Right:
							game.shapeLoc.x++
						}
					}
				default:
					goto EndFor
				}
			}
		EndFor:
		}

		if game.shapeLoc.y+1 == game.floor[game.shapeLoc.x] {
			game.floor[game.shapeLoc.x]--
			game.board[game.shapeLoc.y][game.shapeLoc.x] = 1
			game.shapeLoc = Coord{0, 0}
		} else {
			game.shapeLoc.y++
		}
	}
}
