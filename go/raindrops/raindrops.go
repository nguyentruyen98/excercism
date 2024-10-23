package raindrops

import "fmt"

func Convert(n int) string {
	var result string

	if n%3 == 0 {
		result += "Pling"
	}
	if n%5 == 0 {
		result += "Plang"
	}
	if n%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return fmt.Sprint(n)
	}

	return result
}
