package main

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"

	"advent2015/pkg/file"
)

func sum(values []int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func product(values []int) int {
	result := 1
	for _, v := range values {
		result *= v
	}
	return result
}

func removeFrom(src []int, rm []int) []int {
	var result []int
	for _, srcValue := range src {
		if slices.Index(rm, srcValue) == -1 {
			result = append(result, srcValue)
		}
	}
	return result
}

func getCombinations(values []int, current, depth int, currentCombinations []int, combinations *[][]int) {
	if len(currentCombinations) == depth {
		*combinations = append(*combinations, currentCombinations)
	} else {
		for i := current; i < len(values); i++ {
			getCombinations(values, i+1, depth, append(currentCombinations, values[i]), combinations)
		}
	}
}

func getMiQuantum(s []string) int {
	values := make([]int, len(s))
	total := 0
	for idx, line := range s {
		values[idx], _ = strconv.Atoi(line)
		total += values[idx]
	}
	avg := total / 3
	exit := false
	minQE := -1
	firstSize := (len(values) / 3) - (len(values)/3)/2
	for sizeA := firstSize; sizeA <= len(values)/3; sizeA++ {
		var groupsA [][]int
		getCombinations(values, 0, sizeA, []int{}, &groupsA)
		for _, groupA := range groupsA {
			if sum(groupA) == avg {
				qe := product(groupA)
				if minQE == -1 || qe <= minQE {
					newValues := removeFrom(values, groupA)
					for sizeB := firstSize; sizeB <= len(newValues)-firstSize; sizeB++ {
						var groupsB [][]int
						getCombinations(newValues, 0, sizeB, []int{}, &groupsB)
						for _, groupB := range groupsB {
							if sum(groupB) == avg {
								exit = true
								if minQE == -1 || qe < minQE {
									minQE = qe
								}
							}
						}
					}
					if exit {
						break
					}
				}
			}
		}
		if exit {
			break
		}
	}
	return minQE
}

func getMiQuantum2(s []string) int {
	values := make([]int, len(s))
	total := 0
	for idx, line := range s {
		values[idx], _ = strconv.Atoi(line)
		total += values[idx]
	}
	avg := total / 4
	exit := false
	minQE := -1
	firstSize := (len(values) / 3) - (len(values)/3)/2 - 1
	for sizeA := firstSize; sizeA < len(values); sizeA++ {
		var groupsA [][]int
		getCombinations(values, 0, sizeA, []int{}, &groupsA)
		for _, groupA := range groupsA {
			if sum(groupA) == avg {
				qe := product(groupA)
				if minQE == -1 || qe <= minQE {
					newValuesA := removeFrom(values, groupA)
					for sizeB := firstSize + 1; sizeB <= len(newValuesA)-firstSize; sizeB++ {
						var groupsB [][]int
						getCombinations(newValuesA, 0, sizeB, []int{}, &groupsB)
						for _, groupB := range groupsB {
							if sum(groupB) == avg {
								newValuesB := removeFrom(newValuesA, groupB)
								for sizeC := firstSize + 1; sizeC <= len(newValuesB)-firstSize; sizeC++ {
									var groupsC [][]int
									getCombinations(newValuesB, 0, sizeC, []int{}, &groupsC)
									if sum(groupB) == avg {
										exit = true
										if minQE == -1 || qe <= minQE {
											minQE = qe
										}
									}
								}
							}
						}
					}
				}
			}
		}
		if exit {
			break
		}
	}
	return minQE
}

func main() {
	absPathName, _ := filepath.Abs("src/day24/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getMiQuantum(output))
	fmt.Println(getMiQuantum2(output))
}
