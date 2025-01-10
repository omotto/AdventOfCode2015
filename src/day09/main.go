package main

import (
	"fmt"
	"path/filepath"
	"slices"

	"advent2015/pkg/file"
)

func parseInput(s []string) map[string]map[string]int {
	var (
		cityA, cityB string
		distance     int
	)
	data := make(map[string]map[string]int)
	for _, line := range s {
		_, _ = fmt.Sscanf(line, "%s to %s = %d", &cityA, &cityB, &distance)
		if v, ok := data[cityA]; ok {
			v[cityB] = distance
			data[cityA] = v
		} else {
			data[cityA] = map[string]int{cityB: distance}
		}
		if v, ok := data[cityB]; ok {
			v[cityA] = distance
			data[cityB] = v
		} else {
			data[cityB] = map[string]int{cityA: distance}
		}
	}
	return data
}

func getDistance(city string, data map[string]map[string]int, visited []string, distance int, distances *[]int, depth int) {
	if slices.IndexFunc(visited, func(s string) bool { return s == city }) >= 0 {
		return
	}
	if depth == len(data)-1 {
		*distances = append(*distances, distance)
		return
	}
	visited = append(visited, city)
	destinations := data[city]
	for k, v := range destinations {
		getDistance(k, data, visited, v+distance, distances, depth+1)
	}
}

func getDistances(data map[string]map[string]int) []int {
	var distances []int
	for k, _ := range data {
		visited := []string{}
		getDistance(k, data, visited, 0, &distances, 0)
	}
	return distances
}

func getMinimumDistance(s []string) int {
	data := parseInput(s)
	distances := getDistances(data)
	return slices.Min(distances)
}

func getMaximumDistance(s []string) int {
	data := parseInput(s)
	distances := getDistances(data)
	return slices.Max(distances)
}

func main() {
	absPathName, _ := filepath.Abs("src/day09/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getMinimumDistance(output))
	fmt.Println(getMaximumDistance(output))
}
