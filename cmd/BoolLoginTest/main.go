package main

import "fmt"

func main() {
	fmt.Println("Example 1:")
	age := 25
	isStudent := true

	if age >= 18 && isStudent {
		fmt.Println("You are an adult student.")
	} else {
		fmt.Println("You are either not an adult or not a student.")
	}

	fmt.Println("\nExample 2:")
	isWorking := false
	hoursWorked := 40

	if isWorking || hoursWorked > 30 {
		fmt.Println("You have either been working or worked more than 30 hours.")
	} else {
		fmt.Println("You haven't been working much.")
	}

	fmt.Println("\nExample 3:")
	isSunny := true

	if !isSunny {
		fmt.Println("It's not sunny today.")
	} else {
		fmt.Println("It's sunny today.")
	}

	

}
