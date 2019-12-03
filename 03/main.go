package main

import (
	"aoc-2019/common"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := common.ReadInput("03/input.txt")
	crosses := make([]point, 0)
	visited := make([]map[point]bool, len(input))
	for lineNumber, line := range input {
		visited[lineNumber] = make(map[point]bool)
		current := point{x: 0, y: 0}
		instructions := strings.Split(strings.TrimSpace(line), ",")
		for _, instruction := range instructions {
			magnitude := common.StringToInt(instruction[1:])
			switch instruction[0] {
			case 'U':
				for i := 0; i < magnitude; i++ {
					current.y++
					if lineNumber == 1 && visited[0][current] {
						crosses = append(crosses, current)
					}
					visited[lineNumber][current] = true
				}
			case 'R':
				for i := 0; i < magnitude; i++ {
					current.x++
					if lineNumber == 1 && visited[0][current] {
						crosses = append(crosses, current)
					}
					visited[lineNumber][current] = true
				}
			case 'D':
				for i := 0; i < magnitude; i++ {
					current.y--
					if lineNumber == 1 && visited[0][current] {
						crosses = append(crosses, current)
					}
					visited[lineNumber][current] = true
				}
			case 'L':
				for i := 0; i < magnitude; i++ {
					current.x--
					if lineNumber == 1 && visited[0][current] {
						crosses = append(crosses, current)
					}
					visited[lineNumber][current] = true
				}
			}
		}
	}
	distances := make(map[point]int)
	for _, distance := range crosses {
		distances[distance] = Abs(distance.x) + Abs(distance.y)
	}
	smallestDistance := math.MaxInt64
	for _, distance := range distances {
		if distance < smallestDistance {
			smallestDistance = distance
		}
	}
	fmt.Printf("%d\n", smallestDistance)
}

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
