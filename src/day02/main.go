package main

import (
	"fmt"
	"path/filepath"

	"advent2015/pkg/file"
)

func getTotalSquareFeet(s []string) int {
	var l, w, h int
	total := 0
	for _, line := range s {
		_, _ = fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		total += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}
	return total
}

func getTotalRibbonFeet(s []string) int {
	var l, w, h int
	total := 0
	for _, line := range s {
		_, _ = fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		total += 2*((l+w+h)-max(l, w, h)) + l*w*h
	}
	return total
}

func main() {
	absPathName, _ := filepath.Abs("src/day02/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getTotalSquareFeet(output))
	fmt.Println(getTotalRibbonFeet(output))
}
