package main

import (
	"aoc-2019/common"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer common.OutputTimeTaken(time.Now())
	input := getOrbitMap(common.ReadInput("06/input.txt"))
	getTotalOrbits(input)
	calculateSantaOrbit(input)
}

func getOrbitMap(input []string) map[string]string {
	orbits := make(map[string]string, 0)
	for _, line := range input {
		o := strings.Split(line, ")")
		orbits[o[1]] = o[0]
	}
	return orbits
}

func getTotalOrbits(input map[string]string) {
	total := 0
	for _, object := range input {
		for {
			total++
			parent, ok := input[object]
			if !ok {
				break
			}
			object = parent
		}
	}
	fmt.Printf("%d\n", total)
}

func calculateSantaOrbit(input map[string]string) {
	youPath := make([]string, 0)
	sanPath := make([]string, 0)
	commonOrbits := 0
	o := "YOU"
	ok := false
	for {
		o, ok = input[o]
		if !ok {
			break
		}
		youPath = append(youPath, o)
	}
	o = "SAN"
	for {
		o, ok = input[o]
		if !ok {
			break
		}
		sanPath = append(sanPath, o)
	}
	for i, j := len(youPath)-1, len(sanPath)-1; i >= 0 && j >= 0; {
		if youPath[i] == sanPath[j] {
			commonOrbits++
		} else {
			break
		}
		i--
		j--
	}
	fmt.Printf("%d\n", len(youPath) + len(sanPath) - commonOrbits*2)
}