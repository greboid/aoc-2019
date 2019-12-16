package common

import (
	"fmt"
)

type Compooter struct {
	idx          int
	memory       []int
	relativeBase int
	Output       chan int
	Input        chan int
	Halt         chan bool
}

func NewCompooter(program []int) *Compooter {
	return &Compooter{
		idx:    0,
		memory: program,
		Output: make(chan int, 100),
		Input:  make(chan int, 1),
		Halt:   make(chan bool, 1),
	}
}

func (compooter *Compooter) getParam(pos int, instruction int) *int {
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
		for len(compooter.memory) <= compooter.memory[compooter.idx+pos] {
			compooter.memory = append(compooter.memory, make([]int, 1024)...)
		}
		return &compooter.memory[compooter.memory[compooter.idx+pos]]
	case 1:
		return &compooter.memory[compooter.idx+pos]
	case 2:
		for len(compooter.memory) <= compooter.memory[compooter.idx+pos] {
			compooter.memory = append(compooter.memory, make([]int, 1024)...)
		}
		return &compooter.memory[compooter.relativeBase + compooter.memory[compooter.idx+pos]]
	default:
		panic(fmt.Sprintf("fault: invalid parameter mode: ip=%d instruction=%d offset=%d mode=%d", compooter.idx, instruction, pos, mode))
	}
}

func (compooter *Compooter) Run() {
	for {
		instruction := compooter.memory[compooter.idx]
		opCode := instruction % 100
		switch opCode {
		case 1:
			value1 := compooter.getParam(1, instruction)
			value2 := compooter.getParam(2, instruction)
			value3 := compooter.getParam(3, instruction)
			*value3 = *value1 + *value2
			compooter.idx += 4
		case 2:
			value1 := compooter.getParam(1, instruction)
			value2 := compooter.getParam(2, instruction)
			value3 := compooter.getParam(3, instruction)
			*value3 = *value1 * *value2
			compooter.idx += 4
		case 3:
			*compooter.getParam(1, instruction) = <-compooter.Input
			compooter.idx += 2
		case 4:
			compooter.Output <- *compooter.getParam(1, instruction)
			compooter.idx += 2
		case 5:
			value1 := compooter.getParam(1, instruction)
			value2 := compooter.getParam(2, instruction)
			if *value1 != 0 {
				compooter.idx = *value2
			} else {
				compooter.idx += 3
			}
		case 6:
			value1 := compooter.getParam(1, instruction)
			value2 := compooter.getParam(2, instruction)
			if *value1 == 0 {
				compooter.idx = *value2
			} else {
				compooter.idx += 3
			}
		case 7:
			value1 := compooter.getParam(1, instruction)
			value2 := compooter.getParam(2, instruction)
			value3 := compooter.getParam(3, instruction)
			if *value1 < *value2 {
				*value3 = 1
			} else {
				*value3 = 0
			}
			compooter.idx += 4
		case 8:
			value1 := compooter.getParam(1, instruction)
			value2 := compooter.getParam(2, instruction)
			value3 := compooter.getParam(3, instruction)
			if *value1 == *value2 {
				*value3 = 1
			} else {
				*value3 = 0
			}
			compooter.idx += 4
		case 9:
			compooter.relativeBase += *compooter.getParam(1, instruction)
			compooter.idx += 2
		case 99:
			compooter.Halt <- true
			close(compooter.Output)
			return
		default:
			panic(fmt.Sprintf("fault: invalid opcode: ip=%d instruction=%d opcode=%d", compooter.idx, instruction, opCode))
		}
	}
}
