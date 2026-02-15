package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("-----String Manager-----")
	sc := bufio.NewScanner(os.Stdin)
	t := InitTrie()
	for {
		fmt.Print("\n\nEnter:\nAdd\nRemove\nCheck\nList\nExit: ")
		choice, e := takeInput(sc)
		if e != nil {
			fmt.Println(e)
			return
		}
		choice = strings.ToLower(choice)
		switch choice {
		case "add":
			fmt.Print("Enter string to add: ")
			str, e := takeInput(sc)
			if e != nil {
				fmt.Println(e)
				return
			}
			res := t.Add(str)
			if res {
				fmt.Println("String added")
			} else {
				fmt.Println(errors.New("Failed to add string"))
			}
			fmt.Printf("Added: %s\n", str)
		case "remove":
			fmt.Print("Enter string to remove: ")
			str, e := takeInput(sc)
			if e != nil {
				fmt.Println(e)
				return
			}
			res := t.Remove(str)
			if res {
				fmt.Printf("Removed: %s\n", str)
			} else {
				fmt.Println(errors.New("String not found"))
			}
		case "check":
			fmt.Print("Enter string to check: ")
			str, e := takeInput(sc)
			if e != nil {
				fmt.Println(e)
				return
			}
			res := t.Check(str)
			if res {
				fmt.Printf("String exists")
			} else {
				fmt.Printf("Not found: %s\n", str)
			}
		case "list":
			fmt.Println("Stored strings:")
			str := t.List()
			fmt.Println(str)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
func takeInput(sc *bufio.Scanner) (string, error) {
	if sc.Scan() {
		return strings.TrimSpace(sc.Text()), nil
	}
	return "", errors.New("Failed to read input")
}

type Node struct {
	children map[rune]*Node
	isEnd    bool
}
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{root: &Node{make(map[rune]*Node), false}}
}
func (t *Trie) Add(s string) bool {
	current := t.root
	for _, letter := range s {
		if current.children[letter] == nil {
			current.children[letter] = &Node{make(map[rune]*Node), false}
		}
		current = current.children[letter]
	}
	current.isEnd = true
	return true
}
func (t *Trie) Remove(s string) bool {
	if len(s) == 0 {
		return false
	}
	var deleter func(current *Node, index int) bool
	deleter = func(current *Node, index int) bool {
		if len(s) == index {
			if !current.isEnd {
				return false
			}
			current.isEnd = false
			return len(current.children) == 0
		}
		if current.children[rune(s[index])] == nil {
			return false
		}
		canDelete := deleter(current.children[rune(s[index])], index+1)
		if canDelete {
			delete(current.children, rune(s[index]))
		}
		return !current.isEnd && len(current.children) == 0
	}
	deleter(t.root, 0)
	return true
}
func (t *Trie) Check(s string) bool {
	current := t.root
	for _, letter := range s {
		if current.children[letter] == nil {
			return false
		}
		current = current.children[letter]
	}
	return current.isEnd
}
func (t *Trie) List() []string {
	res := []string{}
	var dfs func(current *Node, path string)
	dfs = func(current *Node, path string) {
		if current.isEnd {
			res = append(res, path)
		}
		for letter, child := range current.children {
			if child != nil {
				dfs(child, path+string(letter))
			}
		}
	}
	dfs(t.root, "")
	return res
}
