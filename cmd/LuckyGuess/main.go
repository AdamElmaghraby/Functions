package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var guess int

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	fmt.Println("Guess the number!")

	for {
		fmt.Println("Please input your guess:")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		diff := guess - secretNumber
		if diff == 0 {
			fmt.Println("You win!! Congratulations :D")
			break
		} else if diff == 1 || diff == -1 {
			fmt.Println("LAVA!!.. you're one number off!")
		} else if diff == 3 || diff == -3 {
			fmt.Println("HOT!.. you're 3 numbers off!")
		} else if diff == 4 || diff == -4 {
			fmt.Println("Warm!.. you're 4 numbers off!")
		} else {
			fmt.Println("Cold! You're more than four digits off.")
		}
	}
}
