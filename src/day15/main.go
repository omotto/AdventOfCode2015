package main

import (
	"advent2015/pkg/file"
	"fmt"
	"path/filepath"
)

type Specs struct {
	capacity, durability, flavor, texture, calories int
}

func parseInput(s []string) []Specs {
	var (
		name string
		capacity,
		durability,
		flavor,
		texture,
		calories int
	)
	ingredients := make([]Specs, len(s))
	for idx, line := range s {
		_, _ = fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories)
		ingredients[idx] = Specs{
			capacity:   capacity,
			durability: durability,
			flavor:     flavor,
			texture:    texture,
			calories:   calories,
		}
	}
	return ingredients
}

func getScore(spoons []int, depth int, ingredients []Specs, maxScore *int, calories int) {
	if depth == len(ingredients) {
		totalSpoons := 0
		for _, spoon := range spoons {
			totalSpoons += spoon
		}
		if totalSpoons == 100 {
			score := [5]int{}
			for idx, ingredient := range ingredients {
				score[0] += ingredient.capacity * spoons[idx]   // total capacity
				score[1] += ingredient.durability * spoons[idx] // total durability
				score[2] += ingredient.texture * spoons[idx]    // total texture
				score[3] += ingredient.flavor * spoons[idx]     // total flavor
				score[4] += ingredient.calories * spoons[idx]   // total calories
			}
			if score[4] == calories || calories == -1 {
				if score[0] < 0 {
					score[0] = 0
				}
				if score[1] < 0 {
					score[1] = 0
				}
				if score[2] < 0 {
					score[2] = 0
				}
				if score[3] < 0 {
					score[3] = 0
				}
				totalScore := score[0] * score[1] * score[2] * score[3]
				if *maxScore < totalScore {
					*maxScore = totalScore
				}
			}
		}
	} else {
		for spoon := 0; spoon < 101; spoon++ {
			spoons[depth] = spoon
			getScore(spoons, depth+1, ingredients, maxScore, calories)
		}
	}
}

func getBestScore(s []string, calories int) int {
	ingredients := parseInput(s)
	var maxScore int
	spoons := make([]int, len(ingredients))
	getScore(spoons, 0, ingredients, &maxScore, calories)
	return maxScore
}

func main() {
	absPathName, _ := filepath.Abs("src/day15/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getBestScore(output, -1))
	fmt.Println(getBestScore(output, 500))
}
