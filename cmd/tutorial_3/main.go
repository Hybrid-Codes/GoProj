package main

import (
	"fmt"
	"os"
)

func main() {
	var choice string
	fmt.Println("Enter your operator of choice: \n +, -, *, / or exit")
	fmt.Scan(&choice)

	var num1, num2 float64
	fmt.Println("Enter two numbers")
	fmt.Scan(&num1, &num2)

	switch choice {
		case "+":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)

		case "*":
			fmt.Println(num1 * num2)

		case "/":
			fmt.Println(num1 / num2)
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Invalid operator")
	}
}



