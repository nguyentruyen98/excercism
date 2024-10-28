package isogram

import "strings"

func IsIsogram(word string) bool {
	isIsoGram := true

	for i := 0; i < len(word); i++ {
		for j := i + 1; j < len(word); j++ {

			if word[i] == '-' || word[i] == ' ' {
				continue
			}
			if strings.EqualFold(string(word[i]), string(word[j])) {
				isIsoGram = false
			}
		}
	}
	return isIsoGram

}
