package main

import (
	"advent2015/pkg/file"
	"bytes"
	"fmt"
	"path/filepath"
	"strconv"
)

func lookAndSay(input string) string {
	var (
		output  bytes.Buffer
		counter int = 1
		prev    int32
	)
	for idx, char := range input {
		if idx == 0 {
			prev = char
			if idx == len(input)-1 {
				output.WriteString(strconv.Itoa(counter))
				output.WriteString(string(char))
			}
		} else if idx > 0 {
			if char == prev {
				counter++
				if idx == len(input)-1 {
					output.WriteString(strconv.Itoa(counter))
					output.WriteString(string(char))
				}
			} else {
				output.WriteString(strconv.Itoa(counter))
				output.WriteString(string(prev))
				prev = char
				counter = 1
				if idx == len(input)-1 {
					output.WriteString(strconv.Itoa(counter))
					output.WriteString(string(char))
				}
			}
		}
	}
	return output.String()
}

func getLookAndSayLength(s string, times int) int {
	strFinalNumber := s
	for idx := 0; idx < times; idx++ {
		strFinalNumber = lookAndSay(strFinalNumber)
	}
	return len(strFinalNumber)
}

func main() {
	absPathName, _ := filepath.Abs("src/day10/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getLookAndSayLength(output[0], 40))
	fmt.Println(getLookAndSayLength(output[0], 50))
}
