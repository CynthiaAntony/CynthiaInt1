package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	a := make(map[string]int)
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("-------String list manager-------")
	for {
		fmt.Print("\n\n\nChoose: \nList\nAdd\nRemove\nCheck\nQuit: ")
		c, e := takeinp(sc)
		if e != nil {
			fmt.Println(e)
			return
		}
		c = strings.ToLower(c)
		fmt.Println()
		switch c {
		case "add":
			fmt.Print("Enter new string: ")
			s, e := takeinp(sc)
			if e != nil {
				fmt.Println(e)
				return
			}
			add(s, a)
		case "remove":
			fmt.Print("Enter string to remove: ")
			s, e := takeinp(sc)
			if e != nil {
				fmt.Println(e)
				return
			}
			_, t := remove(s, a)
			if t == nil {
				fmt.Println("String Removed")
			} else {
				fmt.Println(t)
			}
		case "check":
			fmt.Print("Enter string to check: ")
			s, e := takeinp(sc)
			if e != nil {
				fmt.Println(e)
				return
			}
			_, t := check(s, a)
			if t == nil {
				fmt.Println("String is in the list")
			} else {
				fmt.Println(t)
			}
		case "list":
			if len(a) == 0 {
				fmt.Println("No strings in the list")
				continue
			}
			fmt.Println("Strings:")
			for i := range a {
				fmt.Printf("%v\n", i)
			}
		case "quit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Wrong, try again")
		}
	}
}

func add(s string, a map[string]int) {
	if s != "" {
		a[s] = 1
		fmt.Println("String added")
	} else {
		fmt.Println("String cannot be empty, cannot add")
	}
}

func remove(s string, a map[string]int) (bool, error) {
	if _, ok := a[s]; ok {
		delete(a, s)
		return true, nil
	}
	return false, errors.New("String not found")
}

func check(s string, a map[string]int) (bool, error) {
	if _, ok := a[s]; ok {
		return true, nil
	}
	return false, errors.New("String not found")
}

func takeinp(sc *bufio.Scanner) (string, error) {
	if sc.Scan() {
		return strings.TrimSpace(sc.Text()), nil
	}
	return "", errors.New("Failed to read input")
}
