package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2015/pkg/file"
)

func getRegister(s []string, register string, visited map[string]int) int {
	if v, ok := visited[register]; ok {
		return v
	}
	var op, dst string
	result := 0
	for _, line := range s {
		p := strings.Split(line, " -> ")
		op, dst = p[0], p[1]
		if dst == register {
			parts := strings.Split(op, " ")
			switch len(parts) {
			case 1:
				if v, err := strconv.Atoi(parts[0]); err != nil { // is String
					result = getRegister(s, parts[0], visited)
				} else { // is Int
					result = v
				}
			case 2:
				result = 65536 + ^getRegister(s, parts[1], visited)
			case 3:
				var a, b int
				if v, err := strconv.Atoi(parts[0]); err != nil { // is String
					a = getRegister(s, parts[0], visited)
				} else { // is Int
					a = v
				}
				if v, err := strconv.Atoi(parts[2]); err != nil { // is String
					b = getRegister(s, parts[2], visited)
				} else { // is Int
					b = v
				}
				switch parts[1] {
				case "AND":
					result = a & b
				case "OR":
					result = a | b
				case "RSHIFT":
					result = a >> b
				case "LSHIFT":
					result = a << b
				}
			}
			break
		}
	}
	visited[dst] = result
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day07/input.txt")
	output, _ := file.ReadInput(absPathName)

	visited := map[string]int{}
	a := getRegister(output, "a", visited)
	fmt.Println(a)

	visited = map[string]int{"b": a}
	a = getRegister(output, "a", visited)
	fmt.Println(a)
}
