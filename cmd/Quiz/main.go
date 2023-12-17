package main

import (
	"fmt"
	"strings"
)

func main() {

	questions := []string{
		"How many employees work at Vitality South",
		"When was Vitality South established",
		"What town is Vitality South based in",
	}

	answers := []string{"13", "2019", "Tupelo"}

	correctAnswers := 0

	fmt.Println("Welcome to the Quiz! Answer the following questions:")

	for i, question := range questions {
		fmt.Println("\nQuestion", i+1, ":", question)
		var userAnswer string
		fmt.Print("Your answer: ")
		_, _ = fmt.Scanln(&userAnswer)

		userAnswer = strings.TrimSpace(userAnswer)
		if strings.EqualFold(userAnswer, answers[i]) {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Println("Incorrect! The correct answer is:", answers[i])
		}
	}

	fmt.Println("\nQuiz completed!")
	fmt.Println("You scored", correctAnswers, "Out of", len(questions), "Questions.")
	if correctAnswers == len(questions) {
		fmt.Println("Great job! You got all the answers right :D")
	} else if correctAnswers >= len(questions)/2 {
		fmt.Println("Decent job! You answered most of the questions correct")
	} else {
		fmt.Println("Your going to want to keep practicing, you didn't do too hot...")
	}
}
