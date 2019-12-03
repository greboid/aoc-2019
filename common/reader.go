package common

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic("Unable to open file.")
	}
	defer func() { _ = file.Close() }()
	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return lines
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Unable to convert to int")
	}
	return i
}