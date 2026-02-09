package main

import "fmt"

func main() {
	a := []string{}
	fmt.Println("-------String list manager-------")
	for {
		fmt.Print("\n\n\nChoose: \nList\nAdd\nRemove\nCheck\nQuit: ")
		var c, s string
		fmt.Scan(&c)
		switch c {
		case "Add":
			fmt.Print("Enter new string: ")
			fmt.Scan(&s)
			a = append(a, s)
		case "Remove":
			fmt.Print("Enter string to remove: ")
			d := 0
			fmt.Scan(&s)
			for i, v := range a {
				if v == s {
					a = append(a[:i], a[i+1:]...)
					fmt.Println("String removed")
					d = 1
					break
				}
			}
			if d == 0 {
				fmt.Println("String not found")
			}
		case "Check":
			fmt.Print("Enter string to check: ")
			d := 0
			fmt.Scan(&s)
			for i, v := range a {
				if v == s {
					fmt.Println("Exists! String found at index", i)
					d = 1
					break
				}
			}
			if d == 0 {
				fmt.Println("String not found")
			}
		case "List":
			fmt.Println("String list:")
			if len(a) == 0 {
				fmt.Println("No strings in the list")
			}
			for i, v := range a {
				fmt.Printf("%v: %v\n", i, v)
			}
		case "Quit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Wrong, try again")
		}
	}
}
