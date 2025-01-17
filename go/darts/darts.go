package darts

import "math"

func Score(x, y float64) int {
	x = math.Abs(x)
	y = math.Abs(y)

	dementions := math.Sqrt(x*x + y*y)

	if dementions <= 1 {
		return 10
	}
	if dementions <= 5 {
		return 5
	}
	if dementions <= 10 {
		return 1
	}
	return 0
}
