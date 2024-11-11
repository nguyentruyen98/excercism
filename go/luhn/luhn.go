package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	if len(strings.TrimSpace(id)) <= 1 {
		return false
	}
	sum := 0
	nDigits := 0
	for i := len(id) - 1; i >= 0; i-- {
		s := rune(id[i])
		if unicode.IsSpace(s) {
			continue
		}
		if !unicode.IsDigit(s) {
			return false
		}
		num, _ := strconv.Atoi(string(s))
		if nDigits%2 == 1 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
		nDigits++
	}
	return sum%10 == 0
}
