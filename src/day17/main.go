package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"advent2015/pkg/file"
)

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

func getCombinations(containers []int, init, liters int, currentCombination int, combinations *int) {
	if currentCombination == liters {
		*combinations++
	} else if currentCombination < liters {
		for i := init; i < len(containers); i++ {
			getCombinations(containers, i+1, liters, currentCombination+containers[i], combinations)
		}
	}
}

func getNumCombinations(s []string, liters int) int {
	containers := make([]int, len(s))
	for idx, container := range s {
		containers[idx], _ = strconv.Atoi(container)
	}
	var combinations int
	getCombinations(containers, 0, liters, 0, &combinations)
	return combinations
}

func getCombinations2(containers []int, init, liters int, currentCombinations []int, combinations *[][]int) {
	currentCombination := 0
	for _, v := range currentCombinations {
		currentCombination += v
	}
	if currentCombination == liters {
		*combinations = append(*combinations, currentCombinations)
	} else if currentCombination < liters {
		for i := init; i < len(containers); i++ {
			getCombinations2(containers, i+1, liters, append(currentCombinations, containers[i]), combinations)
		}
	}
}

func getNumMinimumWays(s []string, liters int) int {
	containers := make([]int, len(s))
	for idx, container := range s {
		containers[idx], _ = strconv.Atoi(container)
	}
	var combinations [][]int
	getCombinations2(containers, 0, liters, []int{}, &combinations)
	minimum := MaxInt
	for _, combination := range combinations {
		if len(combination) < minimum {
			minimum = len(combination)
		}
	}
	result := 0
	for _, combination := range combinations {
		if len(combination) == minimum {
			result++
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day17/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumCombinations(output, 150))
	fmt.Println(getNumMinimumWays(output, 150))
}
