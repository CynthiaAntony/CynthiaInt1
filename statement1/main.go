package main

import "fmt"

func main() {
	var x, y float32
	var c string
	fmt.Println("------Cynthia's Calculator------")
	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)
	fmt.Print("Enter the choice of operation (+,-,*,/): ")
	fmt.Scan(&c)
	switch c {
	case "+":
		fmt.Printf("%v+%v=%v\n", x, y, x+y)
	case "-":
		fmt.Printf("%v-%v=%v\n", x, y, x-y)
	case "*":
		fmt.Printf("%v*%v=%v\n", x, y, x*y)
	case "/":
		if y == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Printf("%v/%v=%v\n", x, y, x/y)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
