package trie

import (
	"bufio"
	"os"
)

type Node struct {
	Children map[rune]*Node
	IsEnd    bool
}

type Trie struct {
	Root *Node
}

func InitTrie() *Trie {
	return &Trie{Root: &Node{Children: make(map[rune]*Node), IsEnd: false}}
}

func (t *Trie) Add(s string) {
	current := t.Root
	for _, letter := range s {
		if current.Children[letter] == nil {
			current.Children[letter] = &Node{Children: make(map[rune]*Node), IsEnd: false}
		}
		current = current.Children[letter]
	}
	current.IsEnd = true
}

func (t *Trie) Check(s string) bool {
	current := t.Root
	for _, letter := range s {
		if current.Children[letter] == nil {
			return false
		}
		current = current.Children[letter]
	}
	return current.IsEnd
}

func (t *Trie) Remove(s string) bool {
	if len(s) == 0 {
		return false
	}
	var deleted bool
	var deleter func(current *Node, index int) bool
	deleter = func(current *Node, index int) bool {
		if index == len(s) {
			if !current.IsEnd {
				return false
			}
			current.IsEnd = false
			deleted = true
			return len(current.Children) == 0
		}
		ch := current.Children[rune(s[index])]
		if ch == nil {
			return false
		}
		canDelete := deleter(ch, index+1)
		if canDelete {
			delete(current.Children, rune(s[index]))
		}
		return !current.IsEnd && len(current.Children) == 0
	}
	deleter(t.Root, 0)
	return deleted
}

func (t *Trie) List() []string {
	res := []string{}
	var dfs func(current *Node, path string)
	dfs = func(current *Node, path string) {
		if current.IsEnd {
			res = append(res, path)
		}
		for letter, child := range current.Children {
			dfs(child, path+string(letter))
		}
	}
	dfs(t.Root, "")
	return res
}

func (t *Trie) Save(path string) error {
	words := t.List()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, w := range words {
		file.WriteString(w + "\n")
	}
	return nil
}

func (t *Trie) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t.Add(sc.Text())
	}
	return nil
}
