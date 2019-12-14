package main

import (
	"aoc-2019/common"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer common.OutputTimeTaken(time.Now())
	input := getReactionsMap(common.ReadInput("14/input.txt"))
	required := map[string]int{"FUEL": 1}
	react(required, input)
	fmt.Printf("%d\n", required["ORE"])
	fmt.Printf("%d\n", testMaxFuel(input, 1000000000000, required["ORE"]))
}

func testMaxFuel(input map[string]reactions, availableOre int, oreRequiredForOneFuel int) int {
	fuel := 0
	step := availableOre / oreRequiredForOneFuel
	required := make(map[string]int)
	for {
		test := make(map[string]int)
		for name, amount := range required {
			test[name] = amount
		}
		test["FUEL"] += step
		react(test, input)
		if test["ORE"] <= availableOre {
			fuel += step
			required = test
			continue
		}
		if step > 1 {
			step = step / 2
			continue
		}
		break
	}
	return fuel
}

func react(required map[string]int, reactions map[string]reactions) {
	for {
		changed := false
		for name, amount := range required {
			if amount > 0 {
				if reaction, ok := reactions[name]; ok {
					changed = true
					factor := (amount + reaction.product.Quantity - 1) / reaction.product.Quantity
					required[name] -= factor * reaction.product.Quantity
					for _, input := range reaction.reactants {
						required[input.Name] += factor * input.Quantity
					}
				}
			}
		}
		if !changed {
			return
		}
	}
}

func getReactionsMap(input []string) map[string]reactions {
	reactions := make(map[string]reactions)
	for _, line := range input {
		var reaction = inputToReaction(line)
		reactions[reaction.product.Name] = reaction
	}
	return reactions
}

func inputToReaction(input string) reactions {
	bits := strings.Split(input, " => ")
	var reaction reactions
	for _, input := range strings.Split(bits[0], ", ") {
		reaction.reactants = append(reaction.reactants, inputToReactant(input))
	}
	reaction.product = inputToReactant(bits[1])
	return reaction
}

func inputToReactant(input string) reactant {
	bits := strings.Split(input, " ")
	return reactant{common.StringToInt(bits[0]), bits[1]}
}

type reactant struct {
	Quantity int
	Name     string
}

type reactions struct {
	reactants []reactant
	product   reactant
}
