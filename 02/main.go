package main

import (
	"aoc-2019/common"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type opcodeReturn struct {
	opcode int
	noun int
	verb int
}

func compute(opcodes []int, wg *sync.WaitGroup, results chan opcodeReturn) int {
	returnValue := opcodeReturn{0, opcodes[1], opcodes[2]}
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
			if wg != nil && results != nil {
				returnValue.opcode = opcodes[0]
				wg.Done()
				results <- returnValue
			}
			return opcodes[0]
		default:
			wg.Done()
			panic("Unknown opcode")
		}
		idx = idx + 4
	}
}

func part2(opcodes []int) int {
	results := make(chan opcodeReturn)
	var wg sync.WaitGroup
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			checkOpcodes := make([]int, len(opcodes))
			copy(checkOpcodes, opcodes)

			checkOpcodes[1] = noun
			checkOpcodes[2] = verb
			wg.Add(1)
			go compute(checkOpcodes, &wg, results)
		}
	}
	wg.Wait()
	for result := range results {
		if result.opcode == 19690720 {
			return 100 * result.noun + result.verb
		}
	}
	panic("Not found an answer")
}

func part1(opcodes []int) int {
	checkOpcodes := make([]int, len(opcodes))
	copy(checkOpcodes, opcodes)
	checkOpcodes[1] = 12
	checkOpcodes[2] = 2
	return compute(checkOpcodes, nil, nil)
}

func getInput(input string) []int {
	opcodesStr := strings.Split(strings.TrimSpace(input), ",")
	var opcodes []int
	for _, opcode := range opcodesStr {
		opcodes = append(opcodes, stringToInt(opcode))
	}
	return opcodes
}

func main() {
	input := getInput(common.ReadInput("02/input.txt")[0])
	part1 := part1(input)
	part2 := part2(input)
	fmt.Printf("%d\n", part1)
	fmt.Printf("%d\n", part2)
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Unable to convert to int")
	}
	return i
}
