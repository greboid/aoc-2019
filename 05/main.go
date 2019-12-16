package main

import (
	"aoc-2019/common"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer common.OutputTimeTaken(time.Now())
	program := getInput(common.ReadInput("05/input.txt")[0])
	part1(program)
	part2(program)
}

func part2(program []int) {
	memory := make([]int, len(program))
	copy(memory, program)
	compooter := common.NewCompooter(memory)
	compooter.Input <- 5
	go compooter.Run()
	<-compooter.Halt
	result := 0
	for value := range compooter.Output {
		result = value
	}
	fmt.Printf("%d\n", result)
}

func part1(program []int) {
	memory := make([]int, len(program))
	copy(memory, program)
	compooter := common.NewCompooter(memory)
	compooter.Input <- 8
	go compooter.Run()
	<-compooter.Halt
	result := 0
	for value := range compooter.Output {
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
