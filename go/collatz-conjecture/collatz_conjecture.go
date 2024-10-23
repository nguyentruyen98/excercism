package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return n, fmt.Errorf("Err")
	}

	result := 0

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		result += 1
	}

	return result, nil

}
