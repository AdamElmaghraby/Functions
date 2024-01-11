package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the input analysis program!")
	fmt.Print("Enter a string:")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	for _, char := range input {
		switch {
		case char >= 'a' && char <= 'z':
			fmt.Printf("%c is a lowercase letter\n", char)
		case char >= 'A' && char <= 'Z':
			fmt.Printf("%c is a uppercase letter\n", char)
		case char >= '0' && char <= '9':
			fmt.Printf("%c is a digit\n", char)
		case char == ' ' || char == '\t' || char == '\n':
			fmt.Printf("%c is a whitespace character\n", char)
		default:
			fmt.Printf("%c is a special character\n", char)
		}
	}
}
