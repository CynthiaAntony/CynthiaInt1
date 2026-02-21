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
			remove(s, a)
		case "check":
			fmt.Print("Enter string to check: ")
			s, e := takeinp(sc)
			if e != nil {
				fmt.Println(e)
				return
			}
			t := check(s, a)
			if t == true {
				fmt.Printf("String '%v' is in the list\n", s)
			} else {
				fmt.Printf("String '%v' is not in the list\n", s)
			}
		case "list":
			if len(a) == 0 {
				fmt.Println("No strings in the list")
				continue
			}
			fmt.Println("String: Count")
			for i, v := range a {
				fmt.Printf("%v: %v\n", i, v)
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
		a[s] += 1
		fmt.Println("String added")
	} else {
		fmt.Println("String cannot be empty, cannot add")
	}
}

func remove(s string, a map[string]int) {
	if s != "" && a[s] > 0 {
		a[s] -= 1
		fmt.Println("String removed")
		if a[s] == 0 {
			delete(a, s)
		}
	} else {
		fmt.Println("String not found, cannot remove")
	}
}

func check(s string, a map[string]int) bool {
	if s != "" && a[s] > 0 {
		return true
	}
	return false
}

func takeinp(sc *bufio.Scanner) (string, error) {
	if sc.Scan() {
		return strings.TrimSpace(sc.Text()), nil
	}
	return "", errors.New("Failed to read input")
}
