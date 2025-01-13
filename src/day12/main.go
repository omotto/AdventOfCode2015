package main

import (
	"advent2015/pkg/file"
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
)

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

func getSumNumbers(s []string) int {
	result := 0
	re, _ := regexp.Compile("-?\\d+")
	for _, line := range s {
		idx := re.FindAllStringIndex(line, -1)
		for _, i := range idx {
			if (line[i[0]-1] == ':' || line[i[0]-1] == '[' || line[i[0]-1] == ',') &&
				(line[i[1]] == '}' || line[i[1]] == ']' || line[i[1]] == ',') {
				v, _ := strconv.Atoi(line[i[0]:i[1]])
				result += v
			}
		}

	}
	return result
}

func getSumNumbers2(line string) int {
	if !regexp.MustCompile("[{}]").MatchString(line) ||
		!regexp.MustCompile("red").MatchString(line) {
		return getSumNumbers([]string{":" + line + ","})
	}
	var data map[string]interface{}
	err := json.Unmarshal([]byte(line), &data)
	// error if it's not a map[string]interface -> That means it's an array
	if err != nil {
		var array []interface{}
		_ = json.Unmarshal([]byte(line), &array)
		var total int
		for _, v := range array {
			str, _ := json.Marshal(v)
			total += getSumNumbers2(string(str))
		}
		return total
	}
	// if there is a "red" value un the object return 0 (no analyze the integers on it)
	for _, v := range data {
		if str, ok := v.(string); ok && str == "red" {
			return 0
		}
	}
	var total int
	for _, v := range data {
		str, _ := json.Marshal(v)
		total += getSumNumbers2(string(str))
	}
	return total
}

func main() {
	absPathName, _ := filepath.Abs("src/day12/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumNumbers(output))
	fmt.Println(getSumNumbers2(output[0]))
}
