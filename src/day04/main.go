package main

import (
	"advent2015/pkg/file"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getLowestPositiveNumber(s string, numZeros int) int {
	idx := 1
	value := strings.Repeat("0", numZeros)
	for {
		strIdx := strconv.Itoa(idx)
		hash := getMD5Hash(s + strIdx)
		if hash[:numZeros] == value {
			return idx
		}
		idx++
	}
}

func main() {
	absPathName, _ := filepath.Abs("src/day04/input.txt")
	output, _ := file.ReadInput(absPathName)
	fmt.Println(getLowestPositiveNumber(output[0], 5))
	fmt.Println(getLowestPositiveNumber(output[0], 6))
}
