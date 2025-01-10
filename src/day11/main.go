package main

import (
	"fmt"
	"path/filepath"

	"github.com/dlclark/regexp2"

	"advent2015/pkg/file"
)

func validateIncreasingStraight(password string) bool {
	for idx := 0; idx < len(password)-2; idx++ {
		a := password[idx] + 2
		b := password[idx+1] + 1
		c := password[idx+2]
		if a == b && c == b {
			return true
		}
	}
	return false
}

func validateHasTwoPairs(password string) bool {
	re := regexp2.MustCompile(`(\w)\1.*(\w)\2`, 0)
	match, _ := re.MatchString(password)
	return match
}

func validateInvalidChars(password string) bool {
	invalidChars := map[uint8]struct{}{'i': {}, 'l': {}, 'o': {}}
	for idx := 0; idx < len(password); idx++ {
		if _, ok := invalidChars[password[idx]]; ok {
			return false
		}
	}
	return true
}

func isValidPassword(password string) bool {
	return validateIncreasingStraight(password) && validateHasTwoPairs(password) && validateInvalidChars(password)
}

func incrementString(input string) string {
	var (
		output   string
		nextChar uint8
	)
	char := input[len(input)-1]
	if char == 'z' {
		nextChar = 'a'
	} else {
		nextChar = char + 1
	}
	if nextChar == 'a' {
		output = incrementString(input[:len(input)-1]) + string(nextChar)
	} else {
		output = input[:len(input)-1] + string(nextChar)
	}
	return output
}

func getNextPassword(password string) string {
	for {
		password = incrementString(password)
		if isValidPassword(password) {
			break
		}
	}
	return password
}

func main() {
	absPathName, _ := filepath.Abs("src/day11/input.txt")
	output, _ := file.ReadInput(absPathName)

	nextPassword := getNextPassword(output[0])
	fmt.Println(nextPassword)
	nextPassword = getNextPassword(nextPassword)
	fmt.Println(nextPassword)
}
