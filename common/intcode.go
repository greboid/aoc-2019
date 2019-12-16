package common

import (
	"fmt"
)

func Compute(programInput []int, input chan int, output chan int, halt chan bool) {
	program := make([]int, len(programInput))
	copy(program, programInput)
	idx := 0
	relativeBase := 0
	getPointer := func(index int) *int {
		for len(program) <= index {
			program = append(program, 0)
		}
		return &program[index]
	}
	getParam := func(pos int, instruction int) *int {
		parameter := program[idx+pos]
		mode := -1
		switch pos {
		case 1:
			mode = instruction / 100 % 10
		case 2:
			mode = instruction / 1000 % 10
		case 3:
			mode = instruction / 10000 % 10
		}
		switch mode {
		case 0:
			return getPointer(parameter)
		case 1:
			return &parameter
		case 2:
			return getPointer(relativeBase + parameter)
		default:
			panic(fmt.Sprintf("fault: invalid parameter mode: ip=%d instruction=%d offset=%d mode=%d", idx, instruction, pos, mode))
		}
	}
	for {
		instruction := program[idx]
		opCode := instruction % 100

		switch opCode {
		case 1:
			value1 := getParam(1, instruction)
			value2 := getParam(2, instruction)
			value3 := getParam(3, instruction)
			*value3 = *value1 + *value2
			idx += 4
		case 2:
			value1 := getParam(1, instruction)
			value2 := getParam(2, instruction)
			value3 := getParam(3, instruction)
			*value3 = *value1 * *value2
			idx += 4
		case 3:
			*getParam(1, instruction) = <-input
			idx += 2
		case 4:
			output <- *getParam(1, instruction)
			idx += 2
		case 5:
			value1 := getParam(1, instruction)
			value2 := getParam(2, instruction)
			if *value1 != 0 {
				idx = *value2
			} else {
				idx += 3
			}
		case 6:
			value1 := getParam(1, instruction)
			value2 := getParam(2, instruction)
			if *value1 == 0 {
				idx = *value2
			} else {
				idx += 3
			}
		case 7:
			value1 := getParam(1, instruction)
			value2 := getParam(2, instruction)
			value3 := getParam(3, instruction)
			if *value1 < *value2 {
				*value3 = 1
			} else {
				*value3 = 0
			}
			idx += 4
		case 8:
			value1 := getParam(1, instruction)
			value2 := getParam(2, instruction)
			value3 := getParam(3, instruction)
			if *value1 == *value2 {
				*value3 = 1
			} else {
				*value3 = 0
			}
			idx += 4
		case 9:
			relativeBase += *getParam(1, instruction)
			idx += 2
		case 99:
			halt <- true
			close(output)
			return
		default:
			panic(fmt.Sprintf("fault: invalid opcode: ip=%d instruction=%d opcode=%d", idx, instruction, opCode))
		}
	}
}
