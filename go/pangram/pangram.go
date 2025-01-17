package pangram

import (
	"strings"
)

func generateAlphabet() string {
	alphabet := ""
	for i := 'a'; i <= 'z'; i++ {
		alphabet += string(i)
	}
	return alphabet
}

func IsPangram(input string) bool {

	fommatedStrings := strings.ToLower(strings.ReplaceAll(input, " ", ""))

	alphabet := generateAlphabet()

	charMap := make(map[rune]bool)
	for _, char := range alphabet {
		charMap[char] = false
	}

	result := true

	for _, char := range fommatedStrings {
		charMap[char] = true
	}

	for char := range charMap {
		if !charMap[char] {
			result = false
			break
		}
	}
	return result

}
