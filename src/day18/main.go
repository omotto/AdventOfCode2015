package main

import (
	"bytes"
	"fmt"
	"path/filepath"

	"advent2015/pkg/file"
)

func getLightsOn(m []string) int {
	result := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] == '#' {
				result++
			}
		}
	}
	return result
}

func getNumLightOn(s []string, steps int) int {
	var m []string
	for idx := 0; idx < steps; idx++ {
		m = []string{}
		for y := 0; y < len(s); y++ {
			var buffer bytes.Buffer
			for x := 0; x < len(s[y]); x++ {
				turnedOn := 0
				if y > 0 && x > 0 && s[y-1][x-1] == '#' {
					turnedOn++
				}
				if y < len(s)-1 && x < len(s[y])-1 && s[y+1][x+1] == '#' {
					turnedOn++
				}
				if y < len(s)-1 && x > 0 && s[y+1][x-1] == '#' {
					turnedOn++
				}
				if y > 0 && x < len(s[y])-1 && s[y-1][x+1] == '#' {
					turnedOn++
				}
				if y > 0 && s[y-1][x] == '#' {
					turnedOn++
				}
				if y < len(s)-1 && s[y+1][x] == '#' {
					turnedOn++
				}
				if x > 0 && s[y][x-1] == '#' {
					turnedOn++
				}
				if x < len(s[y])-1 && s[y][x+1] == '#' {
					turnedOn++
				}
				if s[y][x] == '.' {
					if turnedOn == 3 {
						buffer.WriteString("#")
					} else {
						buffer.WriteString(".")
					}
				} else {
					if turnedOn == 3 || turnedOn == 2 {
						buffer.WriteString("#")
					} else {
						buffer.WriteString(".")
					}
				}
			}
			m = append(m, buffer.String())
		}
		s = m
	}
	return getLightsOn(m)
}

func getNumLightOn2(s []string, steps int) int {
	var m []string
	for idx := 0; idx < steps; idx++ {
		m = []string{}
		for y := 0; y < len(s); y++ {
			var buffer bytes.Buffer
			for x := 0; x < len(s[y]); x++ {
				if (y == 0 && x == 0) || (y == 0 && x == len(s[y])-1) || (x == 0 && y == len(s)-1) || (x == len(s[y])-1 && y == len(s)-1) {
					buffer.WriteString("#")
				} else {
					turnedOn := 0
					if y > 0 && x > 0 && s[y-1][x-1] == '#' {
						turnedOn++
					}
					if y < len(s)-1 && x < len(s[y])-1 && s[y+1][x+1] == '#' {
						turnedOn++
					}
					if y < len(s)-1 && x > 0 && s[y+1][x-1] == '#' {
						turnedOn++
					}
					if y > 0 && x < len(s[y])-1 && s[y-1][x+1] == '#' {
						turnedOn++
					}
					if y > 0 && s[y-1][x] == '#' {
						turnedOn++
					}
					if y < len(s)-1 && s[y+1][x] == '#' {
						turnedOn++
					}
					if x > 0 && s[y][x-1] == '#' {
						turnedOn++
					}
					if x < len(s[y])-1 && s[y][x+1] == '#' {
						turnedOn++
					}
					if s[y][x] == '.' {
						if turnedOn == 3 {
							buffer.WriteString("#")
						} else {
							buffer.WriteString(".")
						}
					} else {
						if turnedOn == 3 || turnedOn == 2 {
							buffer.WriteString("#")
						} else {
							buffer.WriteString(".")
						}
					}
				}
			}
			m = append(m, buffer.String())
		}
		s = m
	}
	return getLightsOn(m)
}

func main() {
	absPathName, _ := filepath.Abs("src/day18/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumLightOn(output, 100))
	fmt.Println(getNumLightOn2(output, 100))
}
