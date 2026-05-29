package main

import (
	"errors"
	"fmt"
)

var (
	addition       = createOperation("+")
	subtraction    = createOperation("-")
	multiplication = createOperation("*")
	division       = createOperation("/")
)

func createOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 { return a + b }
	case "-":
		return func(a, b float64) float64 { return a - b }
	case "*":
		return func(a, b float64) float64 { return a * b }
	case "/":
		return func(a, b float64) float64 { return a / b }
	default:
		return nil
	}
}

func operate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return addition(a, b), nil
	case "-":
		return subtraction(a, b), nil
	case "*":
		return multiplication(a, b), nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return division(a, b), nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
}

func main() {
	var a, b float64
	var op string

	for {
		fmt.Print("Enter two numbers and an operation (+, -, *, /, quit): ")
		_, err := fmt.Scan(&a, &b, &op)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}

		if op == "quit" {
			fmt.Println("Bye!")
			break
		}

		result, err := operate(a, b, op)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %.0f\n", result)
	}
}
