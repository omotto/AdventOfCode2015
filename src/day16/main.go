package main

import (
	"advent2015/pkg/file"
	"fmt"
	"path/filepath"
)

func parseInput(s []string) []map[string]int {
	var (
		sueNum, q1, q2, q3     int
		param1, param2, param3 string
	)
	sues := make([]map[string]int, len(s))
	for _, line := range s {
		_, _ = fmt.Sscanf(line, "Sue %d: %s %d, %s %d, %s %d ", &sueNum, &param1, &q1, &param2, &q2, &param3, &q3)
		sues[sueNum-1] = map[string]int{param1[:len(param1)-1]: q1, param2[:len(param2)-1]: q2, param3[:len(param3)-1]: q3}
	}
	return sues
}

func getSueNumber(s []string, chars map[string]int) int {
	sues := parseInput(s)
	for idx, sue := range sues {
		isSue := true
		for k, v := range sue {
			if vv, ok := chars[k]; ok && vv != v {
				isSue = false
				break
			}
		}
		if isSue {
			return idx + 1
		}
	}
	return 0
}

func getSueNumber2(s []string, chars map[string]int) int {
	sues := parseInput(s)
	for idx, sue := range sues {
		isSue := true
		for k, v := range sue {
			switch k {
			case "cats", "trees":
				if vv, ok := chars[k]; ok && vv >= v {
					isSue = false
				}
			case "pomeranians", "goldfish":
				if vv, ok := chars[k]; ok && vv <= v {
					isSue = false
				}
			default:
				if vv, ok := chars[k]; ok && vv != v {
					isSue = false
				}
			}
			if !isSue {
				break
			}
		}
		if isSue {
			return idx + 1
		}
	}
	return 0
}

func main() {
	absPathName, _ := filepath.Abs("src/day16/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSueNumber(output, map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}))

	fmt.Println(getSueNumber2(output, map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}))
}
