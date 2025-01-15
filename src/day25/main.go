package main

import (
	"fmt"
	"path/filepath"

	"advent2015/pkg/file"
)

func getCode(s string) int {
	var row, col int
	_, _ = fmt.Sscanf(s, "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.", &row, &col)
	var position int
	var xInit int
	for y := 1; y < row+1; y++ {
		if y == 1 {
			xInit = 1
		} else {
			xInit = xInit + (y - 1)
		}
		var v int
		for x := 1; x < col+1; x++ {
			if x == 1 {
				v = xInit
			} else {
				v = v + y - 1 + x
			}
			if y == row && x == col {
				position = v
			}
		}
	}
	next := 20151125
	for idx := 2; idx <= position; idx++ {
		next = (next * 252533) % 33554393
	}
	return next
}

func main() {
	absPathName, _ := filepath.Abs("src/day25/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getCode(output[0]))
}
