package main

import (
	"bytes"
	"fmt"
	"path/filepath"

	"advent2015/pkg/file"
)

func calculateMemLength(line string) int {
	c := 0
	for idx := 1; idx < len(line)-1; idx++ {
		char := string(line[idx])
		if char == `\` {
			nextChar := string(line[idx+1])
			switch nextChar {
			case `\`, `"`:
				idx++
				c++
				continue
			case `x`:
				idx += 3
				c++
				continue
			}
		}
		c++
	}
	return c
}

func getSumLengths(s []string) int {
	result := 0
	for _, line := range s {
		strLen := len(line)
		memLen := calculateMemLength(line)
		result += strLen - memLen
	}
	return result
}

func generateLine(line string) string {
	var strBuffer bytes.Buffer
	strBuffer.WriteString(`"`)
	for _, char := range line {
		strChar := string(char)
		switch strChar {
		case `"`:
			strBuffer.WriteString(`\"`)
		case `\`:
			strBuffer.WriteString(`\\`)
		default:
			strBuffer.WriteString(strChar)
		}
	}
	strBuffer.WriteString(`"`)
	return strBuffer.String()
}

func getSumLengths2(s []string) int {
	result := 0
	for _, line := range s {
		strLen := len(line)
		newStrLen := len(generateLine(line))
		result += newStrLen - strLen
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day08/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumLengths(output))
	fmt.Println(getSumLengths2(output))
}
