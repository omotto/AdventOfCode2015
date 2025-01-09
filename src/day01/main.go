package main

import (
	"fmt"
	"path/filepath"

	"advent2015/pkg/file"
)

func getFloor(s []string) int {
	floor := 0
	for _, c := range s[0] {
		if c == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func getPosition(s []string) int {
	floor := 0
	for idx, c := range s[0] {
		if c == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return idx + 1
		}
	}
	return -1
}

func main() {
	absPathName, _ := filepath.Abs("src/day01/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getFloor(output))
	fmt.Println(getPosition(output))
}
