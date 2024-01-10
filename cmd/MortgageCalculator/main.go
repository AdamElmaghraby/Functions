package main

import (
	"fmt"
	"math"
)

func main() {
	var loanAmount float64
	var annualInterestRate float64
	var loanTermInYears float64

	fmt.Println("Welcome to the Mortgage Calculator! \n where everything will be done in only int64!")
	fmt.Println("Please enter the loan amount")
	fmt.Scanln(&loanAmount)

	fmt.Println("Please enter the anual interset rate (as a percentage)")
	fmt.Scanln(&annualInterestRate)

	fmt.Println("Please enter the loan term in years")
	fmt.Scanln(&loanTermInYears)

	monthlyInterestRate := annualInterestRate / 100 / 12

	loanTermInMonths := loanTermInYears * 12

	monthlyPayment := (loanAmount * monthlyInterestRate) / (1 - math.Pow(1+monthlyInterestRate, -loanTermInMonths))

	fmt.Printf("Your monthly mortgage payment is: $%.2f\n", monthlyPayment)
}