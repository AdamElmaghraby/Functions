package main

import (
	"fmt"
	"strings"

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
	var FavNum int
	fmt.Scanln(&FavNum)

	fmt.Println("Wierd choice but okay.. \n What about your 2nd favorite number?")
	var FavNum2 int
	fmt.Scanln(&FavNum2)

	var StringAnswer string

	fmt.Println("Moving on..")

	for {
		fmt.Println("Moving on..\nAnswer this problem (enter numbers separated by spaces): What are the first 5 digits of pi?")
		// Read the entire line as a string
		fmt.Scanln(&StringAnswer)

		// Split the input into individual numbers
		numbers := strings.Fields(StringAnswer)

		// Check if the numbers match
		if len(numbers) == 5 &&
			numbers[0] == "3" &&
			numbers[1] == "1" &&
			numbers[2] == "4" &&
			numbers[3] == "1" &&
			numbers[4] == "5" {
			fmt.Println("Correct! Moving on.")
			break
		} else {
			fmt.Println("Wrong. Try again.")
			fmt.Scanln()
		}
	}

	fmt.Println("Finally.. \n What is your name?")
	var name string
	fmt.Scanln(&name)

	numbers := strings.Fields(StringAnswer)

	var SliceAnswer []int
	for _, numStr := range numbers {
		var num int
		_, err := fmt.Sscan(numStr, &num)
		if err != nil {
			fmt.Println("Error converting input to numbers:", err)
			return
		}
		SliceAnswer = append(SliceAnswer, num)
	}

	fmt.Println("Based off your results, I want to tell you these fun facts!")

	randomNumbers := utils.RandNumber(FavNum, FavNum2)
	fmt.Println("Random numbers using your two favorite numbers:", randomNumbers)
	shuffledNumbers := utils.Shuffle(SliceAnswer)
	fmt.Println("Shuffled Random Numbers using pi:", shuffledNumbers)

	reversedIntSlice := utils.ReverseArr(SliceAnswer)
	fmt.Println("Reversed Digits of pi:", reversedIntSlice)

	reversedString := utils.ReverseString(name)
	fmt.Println("Your name reversed:", reversedString)

	if utils.IsEven(FavNum) {
		fmt.Printf("%d is even.\n", FavNum)
	} else {
		fmt.Printf("%d is odd.\n", FavNum)
	}

	min := FavNum
	max := FavNum2
	evenNumbers := utils.GenEven(FavNum, FavNum2)
	oddNumbers := utils.GenOdd(FavNum, FavNum2)
	fmt.Printf("Even numbers between %d and %d: %v\n", min, max, evenNumbers)
	fmt.Printf("Odd numbers between %d and %d: %v\n", min, max, oddNumbers)

	if utils.HasElement(SliceAnswer, FavNum) {
		fmt.Printf("your Favorite number, %d , does exist in first 5 digits of pi.\n", FavNum)
	} else {
		fmt.Printf("your Favorite number, %d , does not exist in first 5 digits of pi.\n", FavNum)
	}

}
