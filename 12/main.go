package main

import (
	. "aoc-2019/common"
	"fmt"
	"time"
)

func main() {
	defer OutputTimeTaken(time.Now())
	input := parseMoons(ReadInput("12/input.txt"))
	fmt.Printf("%d\n", calculateEnergy(input))
	fmt.Printf("%d\n", getNumberOfSteps(input))
}

//.........
func getNumberOfSteps(input []moon) int {
	moons := make([]moon, len(input))
	copy(moons, input)
	steps, xSteps, ySteps, zSteps := 1, 0, 0, 0
	for ; xSteps == 0 || ySteps == 0 || zSteps == 0; steps++ {
		simulate(moons)
		if xSteps == 0 {
			found := true
			for i, moon := range moons {
				if moon.pos.X != input[i].pos.X || moon.vel.X != input[i].vel.X {
					found = false
					break
				}
			}
			if found {
				xSteps = steps
			}
		}
		if ySteps == 0 {
			found := true
			for i, moon := range moons {
				if moon.pos.Y != input[i].pos.Y || moon.vel.Y != input[i].vel.Y {
					found = false
					break
				}
			}
			if found {
				ySteps = steps
			}
		}
		if zSteps == 0 {
			found := true
			for i, moon := range moons {
				if moon.pos.Z != input[i].pos.Z || moon.vel.Z != input[i].vel.Z {
					found = false
					break
				}
			}
			if found {
				zSteps = steps
			}
		}
	}
	return Lcm(Lcm(xSteps, ySteps), zSteps)
}

func simulate(moons []moon) {
	for i, a := range moons {
		for j, b := range moons {
			if j == i {
				continue
			}
			a.vel = a.vel.Plus(b.pos.Minus(a.pos).Sign())
		}
		moons[i] = a
	}
	for i, moon := range moons {
		moons[i].pos = moon.pos.Plus(moon.vel)
	}
}

func calculateEnergy(input []moon) (totalEnergy int) {
	moons := make([]moon, len(input))
	copy(moons, input)
	for step := 0; step < 1000; step++ {
		simulate(moons)
	}
	for _, moon := range moons {
		potentialEnergy := moon.pos.Abs()
		kineticEnergy := moon.vel.Abs()
		totalEnergy += potentialEnergy * kineticEnergy
	}
	return
}

func parseMoons(input []string) (moons []moon) {
	for _, line := range input {
		var x, y, z int
		_, _ = fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moon := moon{
			pos: Vector{x, y, z},
			vel: Vector{0, 0, 0},
		}
		moons = append(moons, moon)
	}
	return
}

type moon struct {
	pos, vel Vector
}
