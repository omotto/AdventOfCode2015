package main

import (
	"fmt"
	"path/filepath"
	"slices"

	"advent2015/pkg/file"
)

func parseInput(s []string) map[string]map[string]int {
	var (
		personA, personB, loseGain string
		happiness                  int
	)
	data := make(map[string]map[string]int)
	for _, line := range s {
		_, _ = fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &personA, &loseGain, &happiness, &personB)
		personB = personB[:len(personB)-1]
		if loseGain == "lose" {
			happiness *= -1
		}
		if v, ok := data[personA]; ok {
			v[personB] = happiness
			data[personA] = v
		} else {
			data[personA] = map[string]int{personB: happiness}
		}
	}
	return data
}

func getHappiness(first, person string, data map[string]map[string]int, visited []string, happiness int, happinesses *[]int, depth int) {
	if depth == len(data) && person == first {
		*happinesses = append(*happinesses, happiness)
		return
	}
	if slices.IndexFunc(visited, func(s string) bool { return s == person }) >= 0 {
		return
	}
	visited = append(visited, person)
	people := data[person]
	for k, v := range people {
		vv := data[k][person]
		getHappiness(first, k, data, visited, vv+v+happiness, happinesses, depth+1)
	}
}

func getHappinesses(data map[string]map[string]int) []int {
	var happinesses []int
	for k, _ := range data {
		visited := []string{}
		getHappiness(k, k, data, visited, 0, &happinesses, 0)
	}
	return happinesses
}

func getMaximumHappiness(s []string) int {
	data := parseInput(s)
	happinesses := getHappinesses(data)
	return slices.Max(happinesses)
}

func main() {
	absPathName, _ := filepath.Abs("src/day13/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getMaximumHappiness(output))

	output = append(output, []string{
		"Me would gain 0 happiness units by sitting next to Bob.",
		"Me would gain 0 happiness units by sitting next to Carol.",
		"Me would gain 0 happiness units by sitting next to David.",
		"Me would gain 0 happiness units by sitting next to Eric.",
		"Me would gain 0 happiness units by sitting next to Frank.",
		"Me would gain 0 happiness units by sitting next to George.",
		"Me would gain 0 happiness units by sitting next to Mallory.",
		"Me would gain 0 happiness units by sitting next to Alice.",
		"Bob would gain 0 happiness units by sitting next to Me.",
		"Carol would gain 0 happiness units by sitting next to Me.",
		"David would gain 0 happiness units by sitting next to Me.",
		"Eric would gain 0 happiness units by sitting next to Me.",
		"Frank would gain 0 happiness units by sitting next to Me.",
		"George would gain 0 happiness units by sitting next to Me.",
		"Mallory would gain 0 happiness units by sitting next to Me.",
		"Alice would gain 0 happiness units by sitting next to Me.",
	}...)
	fmt.Println(getMaximumHappiness(output))
}
