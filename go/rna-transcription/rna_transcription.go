package strand

func ToRNA(dna string) string {
	result := ""

	for _, rune := range dna {
		if string(rune) == "G" {
			result += "C"
		}
		if string(rune) == "C" {
			result += "G"
		}
		if string(rune) == "T" {
			result += "A"
		}
		if string(rune) == "A" {
			result += "U"
		}
	}
	return result
}
