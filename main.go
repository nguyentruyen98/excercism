package main

func CollatzConjecture(n int) (int, error) {

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

func main() {
	sum, er := CollatzConjecture(1000000)
	println(sum)
	println(er)
}
