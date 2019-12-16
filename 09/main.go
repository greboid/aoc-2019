package main

import (
	"aoc-2019/common"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer common.OutputTimeTaken(time.Now())
	program := getInput(common.ReadInput("09/input.txt")[0])
	part1(program)
	part2(program)
}

func part1(program []int) {
	input := make(chan int, 1)
	output := make(chan int, 100)
	halt := make(chan bool, 1)
	go common.Compute(program, input, output, halt)
	input <- 1
	<-halt
	result := 0
	for value := range output {
		result = value
	}
	fmt.Printf("%d\n", result)
}

func part2(program []int) {
	input := make(chan int, 1)
	output := make(chan int, 100)
	halt := make(chan bool, 1)
	go common.Compute(program, input, output, halt)
	input <- 2
	<-halt
	result := 0
	for value := range output {
		result = value
	}
	fmt.Printf("%d\n", result)
}

func getInput(input string) []int {
	opcodesStr := strings.Split(strings.TrimSpace(input), ",")
	var opcodes []int
	for _, opcode := range opcodesStr {
		opcodes = append(opcodes, common.StringToInt(opcode))
	}
	return opcodes
}