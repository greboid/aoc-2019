package main

import (
	"aoc-2019/common"
	"fmt"
	"strings"
	"time"
)

func compute(opcodes []int) int {
	var idx, arg1, arg2, resultPos int
	for {
		switch opcodes[idx] {
		case 1:
			arg1 = opcodes[idx+1]
			arg2 = opcodes[idx+2]
			resultPos = opcodes[idx+3]
			opcodes[resultPos] = opcodes[arg1] + opcodes[arg2]
		case 2:
			arg1 = opcodes[idx+1]
			arg2 = opcodes[idx+2]
			resultPos = opcodes[idx+3]
			opcodes[resultPos] = opcodes[arg1] * opcodes[arg2]
		case 99:
			return opcodes[0]
		default:
			panic("Unknown opcode")
		}
		idx = idx + 4
	}
}

func part2(opcodes []int) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			checkOpcodes := make([]int, len(opcodes))
			copy(checkOpcodes, opcodes)

			checkOpcodes[1] = noun
			checkOpcodes[2] = verb
			result := compute(checkOpcodes)
			if result == 19690720 {
				return 100 * noun + verb
			}
		}
	}
	panic("Not found an answer")
}

func part1(opcodes []int) int {
	checkOpcodes := make([]int, len(opcodes))
	copy(checkOpcodes, opcodes)
	checkOpcodes[1] = 12
	checkOpcodes[2] = 2
	return compute(checkOpcodes)
}

func getInput(input string) []int {
	opcodesStr := strings.Split(strings.TrimSpace(input), ",")
	var opcodes []int
	for _, opcode := range opcodesStr {
		opcodes = append(opcodes, common.StringToInt(opcode))
	}
	return opcodes
}

func main() {
	defer common.OutputTimeTaken(time.Now())
	input := getInput(common.ReadInput("02/input.txt")[0])
	part1 := part1(input)
	part2 := part2(input)
	fmt.Printf("%d\n", part1)
	fmt.Printf("%d\n", part2)
}
