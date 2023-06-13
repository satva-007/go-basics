package calculator

import (
	"fmt"
)

func Calc() {

	var op string
	var num1 int
	var num2 int
	fmt.Println("Enter First Number ")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second Number ")
	fmt.Scanln(&num2)
	fmt.Println("Enter operator ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "%":
		fmt.Println(num1 % num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Divisor cant be 0 ")
		} else {
			fmt.Println(num1 / num2)
		}
	case "*":
		fmt.Println(num1 * num2)
	default:
		fmt.Println("Invalid")
	}
}
