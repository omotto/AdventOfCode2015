package main

import (
	"fmt"
	"path/filepath"

	"advent2015/pkg/file"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

// -- Player

const playerHitPoints = 100

// -- SHOP 			Cost  Damage  Armor

const (
	COST    = 0
	DAMAGE  = 1
	DEFENSE = 2
)

var (
	Weapons = map[string][3]int{
		"Dagger":     {8, 4, 0},
		"Shortsword": {10, 5, 0},
		"Warhammer":  {25, 6, 0},
		"Longsword":  {40, 7, 0},
		"Greataxe":   {74, 8, 0},
	}

	Armor = map[string][3]int{
		"NoArmor":    {0, 0, 0},
		"Leather":    {13, 0, 1},
		"Chainmail":  {31, 0, 2},
		"Splintmail": {53, 0, 3},
		"Bandedmail": {75, 0, 4},
		"Platemail":  {102, 0, 5},
	}

	Rings = map[string][3]int{
		"NoRing":     {0, 0, 0},
		"Damage +1":  {25, 1, 0},
		"Damage +2":  {50, 2, 0},
		"Damage +3":  {100, 3, 0},
		"Defense +1": {20, 0, 1},
		"Defense +2": {40, 0, 2},
		"Defense +3": {80, 0, 3},
	}
)

func play(playerHitPoints, playerDamage, playerArmor, enemyHitPoints, enemyDamage, enemyArmor int) bool {
	enemyAttack := enemyDamage - playerArmor
	if enemyAttack < 1 {
		enemyAttack = 1
	}
	playerAttack := playerDamage - enemyArmor
	if playerAttack < 1 {
		playerAttack = 1
	}
	for {
		enemyHitPoints -= playerAttack
		playerHitPoints -= enemyAttack
		if enemyHitPoints <= 0 {
			return true
		}
		if playerHitPoints <= 0 {
			return false
		}
	}
}

func getMinCost(s []string) int {
	var enemyHitPoints, enemyDamage, enemyArmor int
	_, _ = fmt.Sscanf(s[0], "Hit Points: %d", &enemyHitPoints)
	_, _ = fmt.Sscanf(s[1], "Damage: %d", &enemyDamage)
	_, _ = fmt.Sscanf(s[2], "Armor: %d", &enemyArmor)
	// Get all possible combinations
	var combinations [][]int
	for _, weapon := range Weapons {
		for _, armor := range Armor {
			for _, ring1 := range Rings {
				for _, ring2 := range Rings {
					cost := weapon[COST] + armor[COST] + ring1[COST] + ring2[COST]
					damage := weapon[DAMAGE] + armor[DAMAGE] + ring1[DAMAGE] + ring2[DAMAGE]
					defense := weapon[DEFENSE] + armor[DEFENSE] + ring1[DEFENSE] + ring2[DEFENSE]
					combinations = append(combinations, []int{cost, damage, defense})
				}
			}
		}
	}
	// Check each combination to get min cost to win
	minCost := MaxInt
	for _, combination := range combinations {
		playerDamage := combination[DAMAGE]
		playerArmor := combination[DEFENSE]
		winner := play(playerHitPoints, playerDamage, playerArmor, enemyHitPoints, enemyDamage, enemyArmor)
		if winner && combination[COST] < minCost {
			minCost = combination[COST]
		}
	}
	return minCost
}

func getMaxCost(s []string) int {
	var enemyHitPoints, enemyDamage, enemyArmor int
	_, _ = fmt.Sscanf(s[0], "Hit Points: %d", &enemyHitPoints)
	_, _ = fmt.Sscanf(s[1], "Damage: %d", &enemyDamage)
	_, _ = fmt.Sscanf(s[2], "Armor: %d", &enemyArmor)
	// Get all possible combinations
	var combinations [][]int
	for _, weapon := range Weapons {
		for _, armor := range Armor {
			for _, ring1 := range Rings {
				for _, ring2 := range Rings {
					cost := weapon[COST] + armor[COST] + ring1[COST] + ring2[COST]
					damage := weapon[DAMAGE] + armor[DAMAGE] + ring1[DAMAGE] + ring2[DAMAGE]
					defense := weapon[DEFENSE] + armor[DEFENSE] + ring1[DEFENSE] + ring2[DEFENSE]
					combinations = append(combinations, []int{cost, damage, defense})
				}
			}
		}
	}
	// Check each combination to get max cost to lose
	maxCost := 0
	for _, combination := range combinations {
		playerDamage := combination[DAMAGE]
		playerArmor := combination[DEFENSE]
		winner := play(playerHitPoints, playerDamage, playerArmor, enemyHitPoints, enemyDamage, enemyArmor)
		if !winner && combination[COST] > maxCost {
			maxCost = combination[COST]
		}

	}
	return maxCost
}

func main() {
	absPathName, _ := filepath.Abs("src/day21/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getMinCost(output))
	fmt.Println(getMaxCost(output))
}
