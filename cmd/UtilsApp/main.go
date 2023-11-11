package main

import (
	"fmt"
	"strconv"

	"github.com/AdamElmaghraby/Functions/pkg/utils"
)

func main() {
	fmt.Println("Welcome to the Utils Test Application!\nAre you ready to get started? (Y/N)")

	var response string
	fmt.Scanln(&response)

	if response == "N" || response == "n" {
		fmt.Println("Exiting the application. Goodbye!")
		return
	} else if response != "Y" && response != "y" {
		fmt.Println("Invalid response. Please enter Y or N.")
		return
	}

	fmt.Println("Great! Let's get started. \n What's your favorite number?")
	var favNum int
	fmt.Scanln(&favNum)

	fmt.Println("Weird choice but okay.. \n What about your 2nd favorite number?")
	var favNum2 int
	fmt.Scanln(&favNum2)

	var stringAnswer int

	fmt.Println("Moving on..")

	for {
		fmt.Println("Answer this problem (enter the first 5 digits of pi as a single integer):")
		fmt.Scanln(&stringAnswer)

		if stringAnswer != 31415 {
			fmt.Println("Wrong. Try again.")
		} else {
			fmt.Println("Correct! Moving on.")
			break
		}
	}

	fmt.Println("Finally.. \n What is your name?")
	var name string
	fmt.Scanln(&name)

	var sliceAnswer []int

	for _, numStr := range strconv.Itoa(stringAnswer) {
		i, err := strconv.Atoi(string(numStr))

		if err != nil {
			continue
		}
		sliceAnswer = append(sliceAnswer, i)
	}

	fmt.Println("Based off your results, I want to tell you these fun facts!")

	randomNumbers := utils.RandNumber(favNum, favNum2)
	fmt.Println("Random numbers using your two favorite numbers:", randomNumbers)
	shuffledNumbers := utils.Shuffle(sliceAnswer)
	fmt.Println("Shuffled Random Numbers using pi:", shuffledNumbers)

	reversedIntSlice := utils.ReverseArr(sliceAnswer)
	fmt.Println("Reversed Digits of pi:", reversedIntSlice)

	reversedString := utils.ReverseString(name)
	fmt.Println("Your name reversed:", reversedString)

	if utils.IsEven(favNum) {
		fmt.Printf("%d is even.\n", favNum)
	} else {
		fmt.Printf("%d is odd.\n", favNum)
	}

	min := favNum
	max := favNum2
	evenNumbers := utils.GenEven(favNum, favNum2)
	oddNumbers := utils.GenOdd(favNum, favNum2)
	fmt.Printf("Even numbers between %d and %d: %v\n", min, max, evenNumbers)
	fmt.Printf("Odd numbers between %d and %d: %v\n", min, max, oddNumbers)

	if utils.HasElement(sliceAnswer, favNum) {
		fmt.Printf("your Favorite number, %d , does exist in first 5 digits of pi.\n", favNum)
	} else {
		fmt.Printf("your Favorite number, %d , does not exist in first 5 digits of pi.\n", favNum)
	}

}
