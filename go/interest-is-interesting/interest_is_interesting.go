package interest

import "fmt"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	}
	if 0 <= balance && balance < 1000 {
		return 0.5
	}
	if 1000 <= balance && balance < 5000 {
		return 1.621
	}
	return 2.475

}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	intersRate := InterestRate(balance)
	return float64(intersRate) * balance / 100

}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	year := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		fmt.Printf("Balance: %f\n", balance)
		year++
	}
	return year
}
