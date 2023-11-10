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

	fmt.Println("Weird choice but okay.. \n What about your 2nd favorite number?")
	var FavNum2 int
	fmt.Scanln(&FavNum2)

	var StringAnswer int

	fmt.Println("Moving on..")

	for {
		fmt.Println("Answer this problem (enter the first 5 digits of pi as a single integer):")
		fmt.Scanln(&StringAnswer)

		if StringAnswer != 31415 {
			fmt.Println("Wrong. Try again.")
		} else {
			fmt.Println("Correct! Moving on.")
			break
		}
	}

	fmt.Println("Finally.. \n What is your name?")
	var name string
	fmt.Scanln(&name)

	numbers := strings.Fields(fmt.Sprint(StringAnswer))

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
