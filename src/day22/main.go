package main

import (
	"fmt"
	"math"
)

type Spell struct {
	Name   string
	Cost   int
	Damage int
	Heal   int
	Effect Effect
}

type Effect struct {
	Name     string
	Duration int
	Damage   int
	Armor    int
	Mana     int
}

type State struct {
	PlayerHP   int
	PlayerMana int
	BossHP     int
	BossDamage int
	Effects    map[string]Effect
	ManaSpent  int
	PlayerTurn bool
}

var spells = []Spell{
	{"Magic Missile", 53, 4, 0, Effect{}},
	{"Drain", 73, 2, 2, Effect{}},
	{"Shield", 113, 0, 0, Effect{"Shield", 6, 0, 7, 0}},
	{"Poison", 173, 0, 0, Effect{"Poison", 6, 3, 0, 0}},
	{"Recharge", 229, 0, 0, Effect{"Recharge", 5, 0, 0, 101}},
}

func main() {
	initialState := State{
		PlayerHP:   50,
		PlayerMana: 500,
		BossHP:     55, // Cambia esto según el input del problema
		BossDamage: 8,  // Cambia esto según el input del problema
		Effects:    make(map[string]Effect),
		ManaSpent:  0,
		PlayerTurn: true,
	}

	minMana := math.MaxInt32
	dfs(initialState, &minMana)
	fmt.Println("Mínimo mana gastado para ganar:", minMana)
}

func dfs(state State, minMana *int) {
	if state.ManaSpent >= *minMana {
		return
	}

	applyEffects(&state)

	if state.BossHP <= 0 {
		if state.ManaSpent < *minMana {
			*minMana = state.ManaSpent
		}
		return
	}

	if state.PlayerHP <= 0 {
		return
	}

	if !state.PlayerTurn {
		// Turno del jefe
		damage := state.BossDamage - getArmor(state)
		if damage < 1 {
			damage = 1
		}
		state.PlayerHP -= damage
		state.PlayerTurn = true
		dfs(state, minMana)
		return
	}

	// Turno del jugador
	for _, spell := range spells {
		if canCastSpell(spell, state) {
			newState := state
			newState.PlayerTurn = false
			newState.ManaSpent += spell.Cost
			newState.PlayerMana -= spell.Cost

			// Aplicar el efecto del hechizo
			if spell.Effect.Duration > 0 {
				newState.Effects[spell.Effect.Name] = spell.Effect
			}

			// Aplicar daño o curación instantánea
			newState.BossHP -= spell.Damage
			newState.PlayerHP += spell.Heal

			dfs(newState, minMana)
		}
	}
}

func applyEffects(state *State) {
	for name, effect := range state.Effects {
		if effect.Duration > 0 {
			state.BossHP -= effect.Damage
			state.PlayerMana += effect.Mana
			effect.Duration--
			state.Effects[name] = effect
		}
		if effect.Duration == 0 {
			delete(state.Effects, name)
		}
	}
}

func canCastSpell(spell Spell, state State) bool {
	if spell.Cost > state.PlayerMana {
		return false
	}
	if spell.Effect.Duration > 0 {
		_, exists := state.Effects[spell.Effect.Name]
		return !exists
	}
	return true
}

func getArmor(state State) int {
	armor := 0
	for _, effect := range state.Effects {
		armor += effect.Armor
	}
	return armor
}

/*
import (
	"fmt"
	"path/filepath"

	"advent2015/pkg/file"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

// -- Player

const (
	playerMana      = 500
	playerHitPoints = 50
)

const (
	MANA = 0
	DAMAGE = 2
)

// -- Spells -> mana cost - effect turns - damage - heals - temporal armor - poison - recharge mana

var Spells = map[string][]int{
	"Magic Missile": {53,  0, 4, 0, 0, 0, 0},
	"Drain":         {73,  0, 2, 2, 0, 0, 0},
	"Shield":        {113, 6, 0, 0, 7, 0, 0},
	"Poison":        {173, 6, 0, 0, 0, 3, 0},
	"Recharge":      {229, 5, 0, 0, 0, 0, 101},
}

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

func getMinMana(s []string) int {
	var enemyHitPoints, enemyDamage int
	_, _ = fmt.Sscanf(s[0], "Hit Points: %d", &enemyHitPoints)
	_, _ = fmt.Sscanf(s[1], "Damage: %d", &enemyDamage)
	type TimeShot struct {
		enemyHitPoints,
		enemyDamage,
		playerHitPoints,
		playerMana,
		playerArmor int
		currentSpell string
		spellTurn int
	}
	queue := make([]TimeShot, 0)
	queue = append(queue, TimeShot{
		enemyHitPoints:  enemyHitPoints,
		enemyDamage:     enemyDamage,
		playerHitPoints: playerHitPoints,
		playerMana:      playerMana,
		playerArmor: 	 0,
		currentSpell:    "",
		spellTurn: 		 0,
	})
	for len(queue) > 0 {
		timeShot := queue[0] // Get first
		queue = queue[1:]    // Remove it
		if timeShot.enemyHitPoints <= 0 {
			return mana
		}
		for name, effect := range Spells {
			switch name {
			case "Magic Missile":
				enemyAttack := timeShot.enemyDamage - timeShot.playerArmor
				if enemyAttack < 1 {
					enemyAttack = 1
				}
				playerAttack := effect[DAMAGE]
				timeShot.enemyHitPoints -= playerAttack
				timeShot.playerHitPoints -= enemyAttack
				timeShot.playerMana -= effect[MANA]
			case "Drain":


				{73,  0, 2, 2, 0, 0, 0},
			case "Shield":        {113, 6, 0, 0, 7, 0, 0},
			case "Poison":        {173, 6, 0, 0, 0, 3, 0},
			case "Recharge":      {229, 5, 0, 0, 0, 0, 101},
			}
		}

		if tile.x == ex && tile.y == ey {
			return visited
		}
		newScore := tile.score + 1
		for _, direction := range directions {
			newX := tile.x + direction[0]
			newY := tile.y + direction[1]
			if _, ok := visited[coords{newX, newY}]; !ok && room[newY][newX] != '#' {
				visited[coords{newX, newY}] = newScore
				queue = append(queue, Tile{
					x:     newX,
					y:     newY,
					score: newScore,
				})
			}
		}
	}
	return minMana
}

func main() {
	absPathName, _ := filepath.Abs("src/day22/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getMinMana(output))
}
*/
/*
import (
"advent2015/pkg/file"
"fmt"
"math"
"path/filepath"
"strconv"
"strings"
)

func main() {
	var part int = 2

	absPathName, _ := filepath.Abs("src/day22/input.txt")
	output, _ := file.ReadInput(absPathName)

	ans := wizardSimulator(strings.Join(output, "\n"), 50, 500, part)
	fmt.Println("Output:", ans)
}

func wizardSimulator(input string, myHP, myMana, part int) int {
	lines := strings.Split(input, "\n")
	bossHP, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	bossDamage, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])

	initState := newBattleState(myHP, myMana, bossHP, bossDamage, [5]int{}, true, 0)

	return simBattle(initState, map[string]int{}, part)
}

// Spell struct is used to generalize all spell types by leveraging zero values.
// The zero value for ints is 0 (which can be added with no effect)
type spell struct {
	name          string // redundant, for debugging
	index         int    // for indexing in an array (which is easily passed by value)
	cost          int
	effectLength  int
	instantDamage int
	instantHeal   int
	effectDamage  int
	heal          int
	armorBuff     int
	manaRecharge  int
}

var spellsMap = map[string]spell{
	"Magic Missile": {
		name:          "Magic Missile",
		index:         0,
		cost:          53,
		instantDamage: 4,
	},
	"Drain": {
		name:          "Drain",
		index:         1,
		cost:          73,
		instantDamage: 2,
		instantHeal:   2,
	},
	"Shield": {
		name:         "Shield",
		index:        2,
		cost:         113,
		effectLength: 6,
		armorBuff:    7, // does not stack for each turn
	},
	"Poison": {
		name:         "Poison",
		index:        3,
		cost:         173,
		effectLength: 6,
		effectDamage: 3,
	},
	"Recharge": {
		name:         "Recharge",
		index:        4,
		cost:         229,
		effectLength: 5,
		manaRecharge: 101,
	},
}

type battleState struct {
	myHP            int
	myMana          int
	bossHP          int
	bossDamage      int
	effectDurations [5]int
	isMyTurn        bool
	depth           int // recursive branch depth, for debugging
}

func newBattleState(myHP, myMana, bossHP, bossDamage int, effectDurations [5]int, isMyTurn bool, depth int) battleState {
	return battleState{
		myHP:            myHP,
		myMana:          myMana,
		bossHP:          bossHP,
		bossDamage:      bossDamage,
		effectDurations: effectDurations,
		isMyTurn:        isMyTurn,
		depth:           depth,
	}
}

func (s battleState) hashKey() string {
	return fmt.Sprintf("%d_%d_%d_%v_%v", s.myHP, s.myMana, s.bossHP, s.effectDurations, s.isMyTurn)
}

func simBattle(state battleState, memo map[string]int, part int) (minMana int) {
	// check cache
	hash := state.hashKey()
	if val, ok := memo[hash]; ok {
		return val
	}

	if part == 2 && state.isMyTurn {
		state.myHP--
	}

	// check myHP after a potential part 2 HP loss, if player dies, then return
	// a huge number which will essentially be ignored by a mathy.MinInt comparison
	if state.myHP <= 0 {
		return math.MaxInt32
	}

	// apply any active spell effects
	var myArmor int
	for _, sp := range spellsMap {
		if state.effectDurations[sp.index] > 0 {
			state.effectDurations[sp.index]--
			// many of values will be zero for any given spell
			state.bossHP -= sp.effectDamage
			state.myHP += sp.heal
			myArmor += sp.armorBuff
			state.myMana += sp.manaRecharge
		}
	}

	// check bossHP after effects take place, it could die form poison
	if state.bossHP <= 0 {
		return 0
	}

	// get minMana from the current state, to a player win
	minMana = math.MaxInt32
	if state.isMyTurn {
		// iterate through spells, create a recursive call for each spell that
		// can be called (i.e. its effectDuration index is zero)
		var spellCasted bool
		for _, sp := range spellsMap {
			if state.effectDurations[sp.index] == 0 {
				if state.myMana >= sp.cost {
					spellCasted = true
					// make new durations array & add effect duration for this spell
					newDurations := state.effectDurations
					newDurations[sp.index] += sp.effectLength

					nextState := newBattleState(state.myHP+sp.instantHeal,
						state.myMana-sp.cost,
						state.bossHP-sp.instantDamage,
						state.bossDamage,
						newDurations,
						false,
						state.depth+1,
					)

					castResult := sp.cost + simBattle(nextState, memo, part)

					minMana = min(minMana, castResult)
				}
			}
		}
		// if cannot cast spell, player loses
		if !spellCasted {
			return math.MaxInt32
		}
	} else {
		// boss's turn, boss attacks w/ a minimum damage of 1
		attackDamage := max(1, state.bossDamage-myArmor)

		// recurse w/ next state
		nextState := newBattleState(state.myHP-attackDamage,
			state.myMana,
			state.bossHP,
			state.bossDamage,
			state.effectDurations,
			true,
			state.depth+1,
		)
		bossAttackResult := simBattle(nextState, memo, part)

		minMana = min(minMana, bossAttackResult)
	}

	// add to memoized to prevent unnecessary recursive branches, then return
	memo[hash] = minMana
	return minMana
}
*/
