package main

import (
	"aoc-2019/common"
	"fmt"
	"time"
)

var (
	up    = common.Point{Y: -1}
	right = common.Point{X: 1}
	down  = common.Point{Y: 1}
	left  = common.Point{X: -1}
	white = " "
	black = "█"
)

func main()  {
	defer common.OutputTimeTaken(time.Now())
	program := common.GetCSVInputArray("11/input.txt")
	paintCount, _ := run(program, 0)
	_, grid := run(program, 1)
	fmt.Printf("%d\n", paintCount)
	drawGrid(grid)
}

func run(program []int, startColour int) (int, map[common.Point]int) {
	compooter := common.NewCompooter(program)
	compooter.Output = make(chan int, 100000000)
	grid := make(map[common.Point]int, 100)
	pos := common.Point{}
	colour := startColour
	direction := up
	paintCount := 0
	go compooter.Run()
	go func() {
		for {
			if colour == 1 {
				compooter.Input <- 1
			} else {
				compooter.Input <- 0
			}
			paint := <-compooter.Output
			turn := <-compooter.Output
			if _, ok := grid[pos]; !ok {
				paintCount++
			}
			grid[pos] = paint
			direction = getTurn(direction, turn)
			pos.X += direction.X
			pos.Y += direction.Y
			colour = grid[pos]
		}
	}()
	<-compooter.Halt
	return paintCount, grid
}

func drawGrid(grid map[common.Point]int) {
	var min common.Point
	var max common.Point
	for pos := range grid {
		min = min.Min(pos)
		max = max.Max(pos)
	}
	for j := 0; j <= 5; j++ {
		for i := 0; i <= 50; i++ {
			if grid[common.Point{i, j}] == 1 {
				fmt.Printf(black)
			} else {
				fmt.Printf(white)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getTurn(point common.Point, rotation int) common.Point {
	if rotation == 1 {
		switch point {
		case up:
			return right
		case right:
			return down
		case down:
			return left
		case left:
			return up
		}
	} else {
		switch point {
		case up:
			return left
		case left:
			return down
		case down:
			return right
		case right:
			return up
		}
	}
	panic("Unable to turn")
}