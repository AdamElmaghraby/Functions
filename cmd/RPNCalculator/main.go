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
				return 0, fmt.Errorf("not enough operands for operator %s", token)
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				if operand2 == 0 {
					return 0, fmt.Errorf("Division by 0")
				}
				result = operand1 / operand2
			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}

			stack = append(stack, result)
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("Invalid Expression")
	}
	return stack[0], nil
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
