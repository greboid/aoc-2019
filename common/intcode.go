package common

import (
	"fmt"
)

func Compute(programInput []int, input chan int, output chan int, halt chan bool) {
	program := make([]int, len(programInput))
	copy(program, programInput)
	var idx int
	for {
		paddedIntCode := fmt.Sprintf("%05d", program[idx])
		opCode := StringToInt(paddedIntCode[3:5])
		paramModeMap := map[int]bool{
			1: paddedIntCode[2:3] == "0",
			2: paddedIntCode[1:2] == "0",
			3: paddedIntCode[0:1] == "0",
		}
		getParam := func(pos int) int {
			param := program[idx+pos]
			if paramModeMap[pos] {
				param = program[program[idx+pos]]
			}
			return param
		}
		switch opCode {
		case 1:
			value1 := getParam(1)
			value2 := getParam(2)
			program[program[idx+3]] = value1 + value2
			idx += 4
		case 2:
			value1 := getParam(1)
			value2 := getParam(2)
			program[program[idx+3]] = value1 * value2
			idx += 4
		case 3:
			program[program[idx+1]] = <-input
			idx += 2
		case 4:
			value1 := getParam(1)
			output <- value1
			idx += 2
		case 5:
			value1 := getParam(1)
			value2 := getParam(2)
			if value1 != 0 {
				idx = value2
			} else {
				idx += 3
			}
		case 6:
			value1 := getParam(1)
			value2 := getParam(2)
			if value1 == 0 {
				idx = value2
			} else {
				idx += 3
			}
		case 7:
			value1 := getParam(1)
			value2 := getParam(2)
			if value1 < value2 {
				program[program[idx+3]] = 1
			} else {
				program[program[idx+3]] = 0
			}
			idx += 4
		case 8:
			value1 := getParam(1)
			value2 := getParam(2)
			if value1 == value2 {
				program[program[idx+3]] = 1
			} else {
				program[program[idx+3]] = 0
			}
			idx += 4
		case 99:
			halt <- true
			close(output)
			return
		default:
			fmt.Printf("Unknown opcode %d\n", program[idx])
			panic("")
		}
	}
}
