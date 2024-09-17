package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, time int) int {

	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time

}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
	noodle := 0
	sauce := 0

	for _, str := range layers {
		if str == "noodles" {
			noodle++
		}
		if str == "sauce" {
			sauce++
		}
	}

	return noodle * 50, float64(sauce) * 0.2

}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(layers1, layers2 []string) {
	layers2[len(layers2)-1] = layers1[len(layers1)-1]

}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(slice []float64, positions int) []float64 {
	var scaledAmounts []float64

	for _, recipeQty := range slice {
		scaledAmounts = append(scaledAmounts, recipeQty/2*float64(positions))
	}
	return scaledAmounts

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
