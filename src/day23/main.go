package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2015/pkg/file"
)

func getRegisterB(s []string, registerA int) int {
	registers := map[string]int{"a": registerA, "b": 0}
	pc := 0
	for pc >= 0 && pc < len(s) {
		// Get
		instruction := s[pc]
		// Decode + Execute
		parts := strings.Split(instruction, " ")
		switch parts[0] {
		case "hlf":
			registers[parts[1]] /= 2
			pc++
		case "tpl":
			registers[parts[1]] *= 3
			pc++
		case "inc":
			registers[parts[1]] += 1
			pc++
		case "jmp":
			num, _ := strconv.Atoi(parts[1])
			pc += num
		case "jie":
			register := registers[parts[1][:len(parts[1])-1]]
			num, _ := strconv.Atoi(parts[2])
			if register%2 == 0 {
				pc += num
			} else {
				pc++
			}
		case "jio":
			register := registers[parts[1][:len(parts[1])-1]]
			num, _ := strconv.Atoi(parts[2])
			if register == 1 {
				pc += num
			} else {
				pc++
			}
		}
	}
	return registers["b"]
}

func main() {
	absPathName, _ := filepath.Abs("src/day23/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getRegisterB(output, 0))
	fmt.Println(getRegisterB(output, 1))
}
