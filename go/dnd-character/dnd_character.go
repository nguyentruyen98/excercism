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

func min(arr []int) int {
	min := arr[0]

	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

func sum(arr []int) int {
	sum := 0

	for _, value := range arr {
		sum = sum + value
	}
	return sum
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	// Chuyển score thành float64 để phép chia được chính xác
	return int(math.Floor(float64(score-10) / 2.0))

}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolling := []int{rand.Intn(6) + 1, rand.Intn(6) + 1, rand.Intn(6) + 1, rand.Intn(6) + 1}

	return sum(rolling) - min(rolling)
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	str := Ability()
	dex := Ability()
	con := Ability()
	intel := Ability()
	wis := Ability()
	cha := Ability()
	hit := 10 + Modifier(con)
	return Character{str, dex, con, intel, wis, cha, hit}

}
