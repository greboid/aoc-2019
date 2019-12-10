package main

import (
	"aoc-2019/common"
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	defer common.OutputTimeTaken(time.Now())
	asteroids := getAsteroids(common.ReadInput("10/input.txt"))
	var bestVisible, bestLocation = getBestVisible(asteroids)
	fmt.Printf("%d\n", bestVisible)
	winner := getPewOrder(asteroids, bestLocation)
	fmt.Printf("%d\n", winner[199].x*100+winner[199].y)
}

func (a asteroid) calculateAngle(bestLocation asteroid) float64 {
	dist := a.Minus(bestLocation)
	return 2.0*math.Pi - (math.Atan2(float64(dist.x), float64(dist.y)) + math.Pi)
}

func (a asteroid) findVisibleAsteroids(asteroids map[asteroid]bool) []asteroid {
	visible := make(map[asteroid]asteroid)
	for asteroid := range asteroids {
		if asteroid == a {
			continue
		}
		dist := asteroid.Minus(a)
		dir := dist.DividedBy(greatestCommonDivisor(abs(dist.x), abs(dist.y)))
		if occluder, ok := visible[dir]; ok {
			occluderDist := occluder.Minus(a)
			if (dir.x != 0 && occluderDist.x/dir.x < dist.x/dir.x) || (dir.y != 0 && occluderDist.y/dir.y < dist.y/dir.y) {
				continue
			}
		}
		visible[dir] = asteroid
	}
	var result []asteroid
	for _, asteroid := range visible {
		result = append(result, asteroid)
	}
	return result
}

func (a asteroid) Minus(b asteroid) asteroid {
	return asteroid{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}

func (a asteroid) DividedBy(factor int) asteroid {
	return asteroid{
		x: a.x / factor,
		y: a.y / factor,
	}
}

func getPewOrder(asteroids map[asteroid]bool, location asteroid) []asteroid {
	pewOrder := make([]asteroid, 0)
	for len(asteroids) > 1 {
		list := location.findVisibleAsteroids(asteroids)
		sort.Slice(list, func(i, j int) bool {
			return list[i].calculateAngle(location) < list[j].calculateAngle(location)
		})
		pewOrder = append(pewOrder, list...)
		for _, asteroid := range list {
			delete(asteroids, asteroid)
		}
	}
	return pewOrder
}

func getBestVisible(asteroids map[asteroid]bool) (int, asteroid) {
	var bestVisible int
	var bestLocation asteroid
	for location := range asteroids {
		visible := len(location.findVisibleAsteroids(asteroids))
		if visible > bestVisible {
			bestVisible = visible
			bestLocation = location
		}
	}
	return bestVisible, bestLocation
}

func getAsteroids(input []string) map[asteroid]bool {
	asteroids := make(map[asteroid]bool)
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				asteroids[asteroid{x, y}] = true
			}
		}
	}
	return asteroids
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type asteroid struct {
	x int
	y int
}
