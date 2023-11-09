package main

import (
	"fmt"

	"github.com/AdamElmaghraby/Functions/utils"
)

func main() {
	randomNumbers := utils.RandNumber(10, 100)
	fmt.Println("Original Random Numbers:", randomNumbers)
	shuffledNumbers := utils.Shuffle(randomNumbers)
	fmt.Println("Shuffled Random Numbers:", shuffledNumbers)

	intSlice := []int{1, 2, 3, 4, 5}
	reversedIntSlice := utils.ReverseArr(intSlice)
	fmt.Println("Reversed Integer Slice:", reversedIntSlice)

	inputString := "Hello, World!"
	reversedString := utils.ReverseString(inputString)
	fmt.Println("Reversed String:", reversedString)

	num := 7
	if utils.IsEven(num) {
		fmt.Printf("%d is even.\n", num)
	} else {
		fmt.Printf("%d is odd.\n", num)
	}

	min := 10
	max := 20
	evenNumbers := utils.GenEven(min, max)
	oddNumbers := utils.GenOdd(min, max)
	fmt.Printf("Even numbers between %d and %d: %v\n", min, max, evenNumbers)
	fmt.Printf("Odd numbers between %d and %d: %v\n", min, max, oddNumbers)

	slice1 := []int{1, 2, 3, 4, 5}
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

	numbers := []int{1, 2, 3, 4, 5}
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
