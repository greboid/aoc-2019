package main

import (
	"aoc-2019/common"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer common.OutputTimeTaken(time.Now())
	input := strings.Split(common.ReadInput("04/input.txt")[0], "-")
	start := common.StringToInt(input[0])
	end := common.StringToInt(input[1])
	count1 := 0
	count2 := 0

	for i := start; i <= end; i++ {
		password := i
		var digits [6]int
		for j := 5; j >= 0; j-- {
			digits[j] = password % 10
			password /= 10
		}
		if digits[0] <= digits[1] &&
			digits[1] <= digits[2] &&
			digits[2] <= digits[3] &&
			digits[3] <= digits[4] &&
			digits[4] <= digits[5] &&
			(digits[0] == digits[1] ||
				digits[1] == digits[2] ||
				digits[2] == digits[3] ||
				digits[3] == digits[4] ||
				digits[4] == digits[5]) {
			count1++
			pairs := 0
			if digits[0] == digits[1] &&
				digits[1] != digits[2] {
				pairs++
			}
			if digits[0] != digits[1] &&
				digits[1] == digits[2] &&
				digits[2] != digits[3] {
				pairs++
			}
			if digits[1] != digits[2] &&
				digits[2] == digits[3] &&
				digits[3] != digits[4] {
				pairs++
			}
			if digits[2] != digits[3] &&
				digits[3] == digits[4] &&
				digits[4] != digits[5] {
				pairs++
			}
			if digits[3] != digits[4] &&
				digits[4] == digits[5] {
				pairs++
			}
			if pairs >= 1 {
				count2++
			}
		}
	}
	fmt.Printf("%d\n", count1)
	fmt.Printf("%d\n", count2)
}
