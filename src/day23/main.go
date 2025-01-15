package main

import (
	"fmt"
	"math"
	"path/filepath"
	"slices"
	"strconv"

	"advent2015/pkg/file"
)

func findMultiples(v int) []int {
	var multiple []int
	for c := 1; c <= int(math.Sqrt(float64(v))); c++ {
		if v%c == 0 {
			multiple = append(multiple, c, v/c)
		}
	}
	slices.Sort(multiple)
	return slices.Compact(multiple)
}

func getHouseNumber(s string) int {
	value, _ := strconv.Atoi(s)
	for house := 1; house < value; house++ {
		currentValue := 0
		for _, m := range findMultiples(house) {
			currentValue += m * 10
		}
		if currentValue >= value {
			return house
		}
	}
	return 0
}

func getHouseNumber2(s string) int {
	value, _ := strconv.Atoi(s)
	for house := 1; house < value; house++ {
		currentValue := 0
		for _, m := range findMultiples(house) {
			if house/m <= 50 {
				currentValue += m * 11
			}
		}
		if currentValue >= value {
			return house
		}
	}
	return 0
}

func main() {
	absPathName, _ := filepath.Abs("src/day20/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getHouseNumber(output[0]))
	fmt.Println(getHouseNumber2(output[0]))
}
