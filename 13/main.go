package main

import (
	. "aoc-2019/common"
	"fmt"
	"time"
)

func main() {
	defer OutputTimeTaken(time.Now())
	input := GetCSVInputArray("13/input.txt")
	blockCount, width := getBlocks(input)
	fmt.Printf("%d\n", blockCount)
	fmt.Printf("%d\n", getScore(input, width))
}

func getBlocks(input []int) (int, int) {
	program1 := make([]int, len(input))
	copy(program1, input)
	compooter := NewCompooter(program1)
	compooter.Output = make(chan int, 3)
	blocks := 0
	var grid [100][100]int
	maxx := 0
	maxy := 0
	go func() {
		for {
			x := <- compooter.Output
			y := <- compooter.Output
			tile := <- compooter.Output
			if x > maxx {
				maxx = x
			}
			if y > maxy {
				maxy = y
			}
			if x != -1 {
				grid[y][x] = tile
			}
		}
	}()
	go compooter.Run()
	<-compooter.Halt
	for i := 0; i <= maxy; i++ {
		for j := 0; j <= maxx; j++ {
			switch grid[i][j] {
			case 2:
				blocks++
			}
		}
	}
	return blocks, maxx
}

func getScore(input []int, width int) int {
	score := 0
	program2 := make([]int, len(input))
	copy(program2, input)
	program2[0] = 2
	cheat(program2, width)
	compooter := NewCompooter(program2)
	compooter.Output = make(chan int, 3)
	go compooter.Run()
	go func() {
		for {
			compooter.Input <- 0
		}
	}()
	for {
		x, ok := <- compooter.Output
		if !ok {
			break
		}
		_ = <- compooter.Output
		tile := <- compooter.Output
		if x == -1 {
			score = tile
		}
	}
	<-compooter.Halt
	return score
}

//I really couldn't be bothered figuring out how to play the game...
//replace the paddle line with all paddles so it plays itself
func cheat(program []int, width int) {
	for i := len(program)-1; i > 200; i-- {
		if program[i] == 1 && program[i-width] == 1 {
			zeros := 0
			threes := 0
			for j := i - width; j < i; j++ {
				switch program[j] {
				case 0:
					zeros++
				case 3:
					threes++
				}
			}
			if threes == 1 && zeros == width - 2 {
				for j := i - width + 1; j < i; j++ {
					program[j] = 3
				}
			}
		}
	}
}
