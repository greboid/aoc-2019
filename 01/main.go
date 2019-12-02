package main

import (
	"aoc-2019/common"
	"fmt"
	"strconv"
)

func getFuelForModule(mass int64) int64 {
	return (mass / 3) - 2
}

func getAnswer(input []string) (int64, int64) {
	var (
		answer          int64 = 0
		recursiveAnswer int64 = 0
	)
	for _, line := range input {
		mass, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic("Input not an int")
		}
		fuelForModule := getFuelForModule(mass)
		answer += fuelForModule
		for fuelForModule > 0 {
			recursiveAnswer += fuelForModule
			fuelForModule = getFuelForModule(fuelForModule)
		}
	}
	return answer, recursiveAnswer
}

func main() {
	input := common.ReadInput("01/input.txt")
	part1, part2 := getAnswer(input)
	fmt.Println(part1)
	fmt.Println(part2)
}
