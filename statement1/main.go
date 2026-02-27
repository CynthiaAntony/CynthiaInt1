package main

import "fmt"

func main() {
	var x, y float64
	var c rune
	fmt.Println("------Cynthia's Calculator------")
	fmt.Print("Enter first number: ")
	_, e1 := fmt.Scanln(&x)
	if e1 != nil {
		fmt.Println("Error: Invalid input. Please enter valid number")
		return
	}
	fmt.Print("Enter second number: ")
	_, e2 := fmt.Scanln(&y)
	if e2 != nil {
		fmt.Println("Error: Invalid input. Please enter valid number")
		return
	}
	fmt.Print("Enter the choice of operation (+,-,*,/): ")
	_, e3 := fmt.Scanf("%c", &c)
	if e3 != nil {
		fmt.Println("Error: Invalid input. Please enter valid operator")
		return
	}
	r, e := calc(x, y, c)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Result: %v\n", r)
	}
}
func calc(x, y float64, c rune) (float64, error) {
	switch c {
	case '+':
		return x + y, nil
	case '-':
		return x - y, nil
	case '*':
		return x * y, nil
	case '/':
		if y == 0 {
			return 0, fmt.Errorf("Division by zero is not allowed.")
		} else {
			return x / y, nil
		}
	default:
		return 0, fmt.Errorf("Invalid operation")
	}
}
