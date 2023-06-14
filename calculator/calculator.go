package calculator

import (
	"errors"
	"fmt"
)

func Calc() {

	var op string
	var num1 int
	var num2 int
	//handles error
	zd_err:=errors.New("Divisor Can't be Zero (ZERO DIVISION ERROR)")

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
			fmt.Println(zd_err)
		} else {
			fmt.Println(num1 / num2)
		}
	case "*":
		fmt.Println(num1 * num2)
	default:
		fmt.Println("Invalid")
	}
}
