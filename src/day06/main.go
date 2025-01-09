package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"advent2015/pkg/file"
)

const (
	turnOn  = 0
	turnOff = 1
	toggle  = 2
)

type coord struct {
	x, y int
}

func parseLine(line string) (op int, begin, end coord) {
	if strings.HasPrefix(line, "turn on") {
		op = turnOn
		_, _ = fmt.Sscanf(line, "turn on %d,%d through %d,%d", &begin.x, &begin.y, &end.x, &end.y)
	} else if strings.HasPrefix(line, "turn off") {
		op = turnOff
		_, _ = fmt.Sscanf(line, "turn off %d,%d through %d,%d", &begin.x, &begin.y, &end.x, &end.y)
	} else {
		op = toggle
		_, _ = fmt.Sscanf(line, "toggle %d,%d through %d,%d", &begin.x, &begin.y, &end.x, &end.y)
	}
	return op, begin, end
}

func setLights(lights [][]bool, op int, begin, end coord) [][]bool {
	for y := begin.y; y < end.y+1; y++ {
		for x := begin.x; x < end.x+1; x++ {
			switch op {
			case turnOff:
				lights[y][x] = false
			case turnOn:
				lights[y][x] = true
			case toggle:
				lights[y][x] = !lights[y][x]
			}
		}
	}
	return lights
}

func getNumLightsLit(s []string) int {
	lights := make([][]bool, 1000, 1000)
	for idx := 0; idx < len(lights); idx++ {
		lights[idx] = make([]bool, 1000, 1000)
	}
	for _, line := range s {
		op, begin, end := parseLine(line)
		lights = setLights(lights, op, begin, end)
	}
	numLights := 0
	for y := 0; y < len(lights); y++ {
		for x := 0; x < len(lights[y]); x++ {
			if lights[y][x] {
				numLights++
			}
		}
	}
	return numLights
}

func setLightBrightness(lights [][]int, op int, begin, end coord) [][]int {
	for y := begin.y; y < end.y+1; y++ {
		for x := begin.x; x < end.x+1; x++ {
			switch op {
			case turnOff:
				if lights[y][x] > 0 {
					lights[y][x]--
				}
			case turnOn:
				lights[y][x]++
			case toggle:
				lights[y][x] += 2
			}
		}
	}
	return lights
}

func getNumLightsBrightness(s []string) int {
	lights := make([][]int, 1000, 1000)
	for idx := 0; idx < len(lights); idx++ {
		lights[idx] = make([]int, 1000, 1000)
	}
	for _, line := range s {
		op, begin, end := parseLine(line)
		lights = setLightBrightness(lights, op, begin, end)
	}
	brightness := 0
	for y := 0; y < len(lights); y++ {
		for x := 0; x < len(lights[y]); x++ {
			brightness += lights[y][x]
		}
	}
	return brightness
}

func main() {
	absPathName, _ := filepath.Abs("src/day06/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumLightsLit(output))
	fmt.Println(getNumLightsBrightness(output))
}
