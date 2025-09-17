package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10.0) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	r := rand.Intn(6) + 1
	dice := []int{r, r, r, r}

	minIndex := 0
	for i := 1; i < 4; i++ {
		if dice[i] < dice[minIndex] {
			minIndex = i
		}
	}
	sum := 0
	for i := range dice {
		if i != minIndex {
			sum += dice[i]
		}
	}
	return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	var char Character
	char.Strength = Ability()
	char.Dexterity = Ability()
	char.Constitution = Ability()
	char.Intelligence = Ability()
	char.Wisdom = Ability()
	char.Charisma = Ability()
	char.Hitpoints = 10 + Modifier(char.Constitution)

	return char
}
