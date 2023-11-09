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

	fmt.Println("That's alright.. \n What about your 2nd favorite number?")
	var FavNum2 int
	fmt.Scanln(&FavNum2)

	fmt.Println("Moving on..\nAnswer this problem (enter numbers separated by spaces):")
	var StringAnswer string
	fmt.Scanln(&StringAnswer)
	if StringAnswer == "3 1 4 1 5" {
		fmt.Println("Correct! Moving on.")
		break
	} else {
		fmt.Println("Wrong. Try again.")
	}
	
	fmt.Println("Finally.. \n What is your name?")
	var name string
	fmt.Scanln(&FavNum2)

	numbers := strings.Fields(StringAnswer)

	var SliceAnswer []int
	for _, numStr := range numbers {
		num, err := fmt.Sscan(numStr, &SliceAnswer)
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

    intSlice := SliceAnswer
    reversedIntSlice := utils.ReverseArr(intSlice)
    fmt.Println("Reversed Integer Slice:", reversedIntSlice)

    inputString := name
    reversedString := utils.ReverseString(inputString)
    fmt.Println("Reversed String:", reversedString)

    num := FavNum
    if utils.IsEven(num) {
        fmt.Printf("%d is even.\n", num)
    } else {
        fmt.Printf("%d is odd.\n", num)
    }

    min := FavNum
    max := FavNum2
    evenNumbers := utils.GenEven(min, max)
    oddNumbers := utils.GenOdd(min, max)
    fmt.Printf("Even numbers between %d and %d: %v\n", min, max, evenNumbers)
    fmt.Printf("Odd numbers between %d and %d: %v\n", min, max, oddNumbers)

    slice1 := SliceAnswer
    slice2 := []int{1, 2, 3, 4, 6}
    if utils.ArrEqual(slice1, slice2) {
        fmt.Println("Slice1 and Slice2 are equal.")
    } else {
        fmt.Println("Slice1 and Slice2 are not equal.")
    }

    strSlice1 := []string{"apple", "banana", "cherry"}
    strSlice2 := []string{"apple", "banana", "cherry"}
    if utils.ArrEqualStr(strSlice1, strSlice2) {
        fmt.Println("String slices are equal.")
    } else {
        fmt.Println("String slices are not equal.")
    }

    numbers = SliceAnswer
    element := 3
    if utils.HasElement(numbers, element) {
        fmt.Printf("%d exists in the slice.\n", element)
    } else {
        fmt.Printf("%d does not exist in the slice.\n", element)
    }

    duplicateStrings := []string{"apple", "banana", "cherry", "banana", "apple"}
    uniqueStrings := utils.Deduplicate(duplicateStrings)
    fmt.Println("Unique strings:", uniqueStrings)
}