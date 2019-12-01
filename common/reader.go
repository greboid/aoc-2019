package common

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadInput(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("Error reading input: %s", err)
	}
	return strings.Split(strings.TrimSpace(strings.Replace(string(data), "\r\n", "\n", -1)), "\n")
}
