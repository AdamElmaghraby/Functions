package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to RPN Calculator!")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter RPN Expression (or 'exit' to quit): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		result, err := evaluateRPN(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}

func evaluateRPN(expr string) (float64, error) {
	stack := make([]float64, 0)

	tokens := strings.Fields(expr)

	for _, token := range tokens {
		if isNumber(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("not enough operands for operator %s")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			swithc token {}
		}
	}
}