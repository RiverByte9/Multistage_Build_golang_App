package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: ./calculator <num1> <operator> <num2>")
		return
	}

	num1, err1 := strconv.Atoi(os.Args[1])
	operator := os.Args[2]
	num2, err2 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid number format")
		return
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Error: Unsupported operator")
		return
	}

	fmt.Printf("Result: %d\n", result)
}
