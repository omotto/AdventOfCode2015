package main

import (
	"fmt"
	"path/filepath"

	"advent2015/pkg/file"
)

func getNumHouses(s []string) int {
	type coord struct {
		x, y int
	}
	houses := 1
	current := coord{0, 0}
	visited := map[coord]struct{}{current: struct{}{}}
	for _, d := range s[0] {
		switch d {
		case '^':
			current.y--
		case '<':
			current.x--
		case '>':
			current.x++
		case 'v':
			current.y++
		}
		if _, ok := visited[current]; !ok {
			houses++
			visited[current] = struct{}{}
		}
	}
	return houses
}

func getNumHouses2(s []string) int {
	type coord struct {
		x, y int
	}
	houses := 1
	current1 := &coord{0, 0}
	current2 := &coord{0, 0}
	visited := map[coord]struct{}{*current1: struct{}{}}
	var current *coord
	for idx, d := range s[0] {
		if idx%2 == 0 {
			current = current1
		} else {
			current = current2
		}
		switch d {
		case '^':
			current.y--
		case '<':
			current.x--
		case '>':
			current.x++
		case 'v':
			current.y++
		}
		if _, ok := visited[*current]; !ok {
			houses++
			visited[*current] = struct{}{}
		}
	}
	return houses
}

func main() {
	absPathName, _ := filepath.Abs("src/day03/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumHouses(output))
	fmt.Println(getNumHouses2(output))
}
