package main

import (
	"fmt"
)

func main() {

	var num1, num2 float64
	var operator string
	var continueCalc string

	fmt.Println("Welcome to the Golangulator")

	for {
		fmt.Print("Please, inform the first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Now, chosse the operation: (+, -, *, /): ")
		fmt.Scanln(&operator)
		fmt.Print("Please inform the second number: ")
		fmt.Scanln(&num2)

		switch operator {
		case "+":
			fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, num1+num2)
		case "-":
			fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, num1-num2)
		case "*":
			fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, num1*num2)
		case "/":
			if num2 != 0 {
				fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, num1/num2)
			} else {
				fmt.Println("Error: Division by zero.")
			}
		default:
			fmt.Println("Error: Invalid operator.")
		}
		fmt.Print("Do you want to perform another calculation? (y/n): ")
		fmt.Scanln(&continueCalc)
		if continueCalc != "y" && continueCalc != "Y" {
			break
		}
	}
	fmt.Println("Thank you for using Golangulator!")
}
