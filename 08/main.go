package main

import (
	"aoc-2019/common"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	defer common.OutputTimeTaken(time.Now())
	input := common.ReadInput("08/input.txt")[0]
	width := 25
	height := 6
	imageDigits := extractLayers(input, width, height)
	leastZeros := leastDigitsInLayer(imageDigits, 0)
	fmt.Printf("%d\n", countDigitInLayer(leastZeros, 1)*countDigitInLayer(leastZeros, 2))
	imageData := buildImage(imageDigits)
	outputASCII(imageData, width, height)
}

func outputASCII(imageData [][]int, width int, height int) {
	white := " "
	black := "█"
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch imageData[x][y] {
			case 0:
				fmt.Printf("%s", white)
			case 1:
				fmt.Printf("%s", black)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func outputPng(imageData [][]int, width int, height int) {
	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})
	for x := range imageData {
		for y, value := range imageData[x] {
			switch value {
			case 0:
				img.Set(x, y, color.White)
			case 1:
				img.Set(x, y, color.Black)
			}
		}
	}
	f, err := os.OpenFile("output.png", os.O_RDWR, 644)
	if err != nil {
		panic("Unable to output part2")
	}
	err = png.Encode(f, img)
	if err != nil {
		panic("Unable to output part2")
	}
	fmt.Println("See output.png")
}

func buildImage(imageLayers imageLayers) [][]int {
	width := 25
	height := 6

	outputData := make([][]int, width)
	for x := range outputData {
		outputData[x] = make([]int, height)
		for y := range outputData[x] {
			outputData[x][y] = 2
		}
	}

	for _, layer := range imageLayers {
		for y, row := range layer {
			for x, digit := range row {
				if outputData[x][y] == 2 {
					outputData[x][y] = digit
				}
			}
		}
	}
	return outputData
}

func leastDigitsInLayer(image imageLayers, target int) layer {
	min := math.MaxInt64
	var minLayer layer
	for _, layer := range image {
		count := countDigitInLayer(layer, target)
		if count < min {
			min = count
			minLayer = layer
		}
	}
	return minLayer
}

func countDigitInLayer(layer layer, target int) int {
	sum := 0
	for _, row := range layer {
		for _, digit := range row {
			if digit == target {
				sum++
			}
		}
	}
	return sum
}

func extractLayers(input string, width int, height int) imageLayers {
	digits := make([]int, len(input))
	layers := make([]layer, 0)
	for i, digit := range strings.Split(input, "") {
		digits[i] = common.StringToInt(digit)
	}
	for i := 0; i < len(digits); i += width * height {
		layer := make([][]int, 0)
		for j := i; j < i+width*height; j += width {
			row := digits[j : j+width]
			layer = append(layer, row)
		}
		layers = append(layers, layer)
	}
	return layers
}

type layer [][]int

type imageLayers []layer
