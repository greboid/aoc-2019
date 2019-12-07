package main

import (
	"aoc-2019/common"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	defer common.OutputTimeTaken(time.Now())
	input := common.ReadInput("03/input.txt")
	var wire1 = drawInstruction(strings.Split(input[0], ","))
	var wire2 = drawInstruction(strings.Split(input[1], ","))
	fmt.Printf("%d\n", leastDistance(wire1, wire2))
	fmt.Printf("%d\n", leastSteps(wire1, wire2))
}

func leastSteps(wire1 wire, wire2 wire) int {
	shortestLatency := math.MaxInt64
	for key := range wire1 {
		_, ok := wire2[key]
		if ok {
			steps1 := wire1[key]
			steps2 := wire2[key]
			currentLatency := steps1 + steps2
			if currentLatency < shortestLatency {
				shortestLatency = currentLatency
			}
		}
	}
	return shortestLatency
}

func leastDistance(wire1 wire, wire2 wire) int {
	shortestDistance := math.MaxInt64
	for point := range wire1 {
		_, ok := wire2[point]
		if ok {
			currentDistance := Abs(point.x) + Abs(point.y)
			if currentDistance < shortestDistance {
				shortestDistance = currentDistance
			}
		}
	}
	return shortestDistance
}

func drawInstruction(instructions []string) wire {
	wire := make(wire)
	current := point{0, 0}
	steps := 0
	for _, instruction := range instructions {
		magnitude := common.StringToInt(instruction[1:])
		switch instruction[0] {
		case 'U':
			for i := current.y + 1; i < current.y+1+magnitude; i++ {
				steps++
				_, ok := wire[point{current.x, i}]
				if !ok {
					wire[point{current.x, i}] = steps
				}
			}
			current.y += magnitude
		case 'R':
			for i := current.x + 1; i < current.x+1+magnitude; i++ {
				steps++
				_, ok := wire[point{i, current.y}]
				if !ok {
					wire[point{i, current.y}] = steps
				}
			}
			current.x += magnitude
		case 'D':
			for i := current.y - 1; i > current.y-1-magnitude; i-- {
				steps++
				_, ok := wire[point{current.x, i}]
				if !ok {
					wire[point{current.x, i}] = steps
				}
			}
			current.y -= magnitude
		case 'L':
			for i := current.x - 1; i > current.x-1-magnitude; i-- {
				steps++
				_, ok := wire[point{i, current.y}]
				if !ok {
					wire[point{i, current.y}] = steps
				}
			}
			current.x -= magnitude
		}
	}
	return wire
}

type wire map[point]int

type point struct {
	x int
	y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
