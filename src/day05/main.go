package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dlclark/regexp2"

	"advent2015/pkg/file"
)

func checkString(s string) int {
	if strings.Contains(s, "ab") ||
		strings.Contains(s, "cd") ||
		strings.Contains(s, "pq") ||
		strings.Contains(s, "xy") {
		return 0
	}
	if strings.Count(s, "a")+
		strings.Count(s, "e")+
		strings.Count(s, "i")+
		strings.Count(s, "o")+
		strings.Count(s, "u") > 2 {
		for idx := 0; idx < len(s)-1; idx++ {
			if s[idx] == s[idx+1] {
				return 1
			}
		}
	}
	return 0
}

func getNiceStringNumber(s []string) int {
	niceString := 0
	for _, line := range s {
		niceString += checkString(line)
	}
	return niceString
}

func getNiceStringNumber2(s []string) int {
	niceString := 0
	re1 := regexp2.MustCompile(`(.).\1`, 0)
	re2 := regexp2.MustCompile(`(..).*\1`, 0)
	for _, line := range s {
		// Using RegEx
		match1, _ := re1.MatchString(line)
		match2, _ := re2.MatchString(line)
		if match1 && match2 {
			niceString++
		}
	}
	return niceString
}

func main() {
	absPathName, _ := filepath.Abs("src/day05/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNiceStringNumber(output))
	fmt.Println(getNiceStringNumber2(output))
}
