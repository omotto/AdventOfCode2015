package main

import (
	"fmt"
	"path/filepath"
	"slices"

	"advent2015/pkg/file"
)

type Specs struct {
	speed, during, rest int
}

func (s Specs) GetDistanceAt(seconds int) int {
	time := s.during + s.rest
	times := seconds / time
	rest := seconds % time
	if rest > s.during {
		return (times + 1) * s.speed * s.during
	}
	return times*s.speed*s.during + s.speed*rest
}

func parseInput(s []string) map[string]Specs {
	var (
		name                string
		speed, during, rest int
	)
	reindeer := make(map[string]Specs)
	for _, line := range s {
		_, _ = fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &during, &rest)
		reindeer[name] = Specs{
			speed:  speed,
			during: during,
			rest:   rest,
		}
	}
	return reindeer
}

func getMaxDistanceAt(s []string, seconds int) int {
	reindeer := parseInput(s)
	distances := make([]int, len(reindeer))
	idx := 0
	for _, specs := range reindeer {
		distances[idx] = specs.GetDistanceAt(seconds)
		idx++
	}
	return slices.Max(distances)
}

func getMaxPointsAt(s []string, seconds int) int {
	reindeer := parseInput(s)
	distances := make([]int, len(reindeer))
	points := make([]int, len(reindeer))
	keys := make([]string, 0, len(reindeer))
	for k, _ := range reindeer {
		keys = append(keys, k)
	}
	for time := 1; time < seconds+1; time++ {
		for idx, k := range keys {
			distances[idx] = reindeer[k].GetDistanceAt(time)
		}
		maxDistance := slices.Max(distances)
		for i := 0; i < len(reindeer); i++ {
			if distances[i] == maxDistance {
				points[i]++
			}
		}
	}
	return slices.Max(points)
}

func main() {
	absPathName, _ := filepath.Abs("src/day14/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getMaxDistanceAt(output, 2503))
	fmt.Println(getMaxPointsAt(output, 2503))
}
