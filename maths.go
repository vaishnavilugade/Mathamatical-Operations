package main

import "fmt"

func main() {
	fmt.Print("\n1.Addition\n2.Substraction\n3.Multiplication\n4.Division\n")
	fmt.Print("Enter your choice:")
	var c int
	fmt.Scanln(&c)
	fmt.Print("Enter first number:")
	var first int

	fmt.Scanln(&first)
	fmt.Print("Enter second number:")
	var second int
	fmt.Scanln(&second)

	var result int
	switch c {
	case 1:
		result = first + second
		fmt.Print("Addition: ", result)
	case 2:
		result = first - second
		fmt.Print("Substraction: ", result)
	case 3:
		result = first * second
		fmt.Print("Multiplication: ", result)
	case 4:
		result = first / second
		fmt.Print("division: ", result)
	}

}
