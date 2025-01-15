package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	"advent2015/pkg/file"
)

type Replacement struct {
	a, b string
}

func parseInput(s []string) ([]Replacement, string) {
	var (
		a, b, str    string
		change       = false
		replacements []Replacement
	)
	for _, line := range s {
		if len(line) < 2 {
			change = true
		} else {
			if change == false {
				_, _ = fmt.Sscanf(line, "%s => %s", &a, &b)
				replacements = append(replacements, Replacement{a: a, b: b})
			} else {
				str = line
			}
		}
	}
	return replacements, str
}

func getNumMolecules(s []string) int {
	replacements, str := parseInput(s)
	molecules := map[string]struct{}{}
	for _, replacement := range replacements {
		nextIdx := 0
		for {
			if idx := strings.Index(str[nextIdx:], replacement.a); idx > -1 {
				newStr := str[:idx+nextIdx] + strings.Replace(str[idx+nextIdx:], replacement.a, replacement.b, 1)
				nextIdx += idx + len(replacement.a)
				molecules[newStr] = struct{}{}
				if nextIdx > len(str) {
					break
				}
				continue
			}
			break
		}
	}
	return len(molecules)
}

// Reverse algorithm, staring from final molecule to the initial "e"
func getNumSteps(s []string) int {
	var (
		a, b, str    string
		change       = false
		replacements = map[string]string{}
		keys         []string
		deadEnd      bool
	)
	for _, line := range s {
		if len(line) < 2 {
			change = true
		} else {
			if change == false {
				_, _ = fmt.Sscanf(line, "%s => %s", &a, &b)
				replacements[b] = a
			} else {
				str = line
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	for k, _ := range replacements {
		keys = append(keys, k)
	}
	result := 0
	molecule := str

	for molecule != "e" {
		deadEnd = true
		for _, key := range keys {
			// if replacement is not in the current molecule stage don't check it
			if count := strings.Count(molecule, key); count > 0 {
				result += count
				molecule = strings.ReplaceAll(molecule, key, replacements[key])
				deadEnd = false
				break
			}
		}
		// if deadEnd -> Try again
		if deadEnd {
			rand.Shuffle(len(keys), func(i, j int) {
				keys[i], keys[j] = keys[j], keys[i]
			})
			molecule = str
			result = 0
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day19/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumMolecules(output))
	fmt.Println(getNumSteps(output))
}
