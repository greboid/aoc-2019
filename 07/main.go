package main

import (
	. "aoc-2019/common"
	"fmt"
	"time"
)

func main() {
	defer OutputTimeTaken(time.Now())
	program := GetCSVInputArray("07/input.txt")
	fmt.Printf("%d\n", maxOutput(program, []int{0, 1, 2, 3, 4}, false))
	fmt.Printf("%d\n", maxOutput(program, []int{5, 6, 7, 8, 9}, true))
}

func maxOutput(program []int, phases []int, feedback bool) int {
	max := 0
	for _, permutation := range permutations(phases) {
		val := runAmplifiers(program, permutation, feedback)
		if val > max {
			max = val
		}
	}
	return max
}

func runAmplifiers(program []int, phases []int, feedback bool) int {
	compooters := make([]*Compooter, 5)
	for i := 0; i < len(compooters); i++ {
		compooters[i] = setupVM(program, IntToString(i))
	}
	for i, compooter := range compooters {
		if i > 0 {
			compooter.Input = compooters[i-1].Output
		} else if feedback {
			compooter.Input = compooters[len(compooters)-1].Output
		} else {
			compooter.Input = make(chan int, 2)
		}
		compooter.Input <- phases[i]
	}
	compooters[0].Input <- 0
	for _, compooter := range compooters {
		go compooter.Run()
	}
	<-compooters[4].Halt
	return <-compooters[4].Output
}

func setupVM(program []int, name string) *Compooter {
	memory := make([]int, len(program))
	copy(memory, program)
	compooter := NewCompooter(memory)
	compooter.Name = name
	compooter.Output = make(chan int, 2)
	return compooter
}

//https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
func permutations(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}