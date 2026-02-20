package main

import "testing"

func TestAll(t *testing.T) {
	trie := InitTrie()
	trie.Add("hello")
	tests := []struct {
		name   string
		action string
		input  string
		want   bool
	}{
		{"Check existing string", "check", "hello", true},
		{"Remove string", "remove", "hello", true},
		{"Check non-existing string", "check", "hello", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.action {
			case "check":
				got := trie.Check(tc.input)
				if got != tc.want {
					t.Fatalf("Got = %v, want %v", got, tc.want)
				}
			case "remove":
				got := trie.Remove(tc.input)
				if got != tc.want {
					t.Fatalf("Got = %v, want %v", got, tc.want)
				}
			}
		})
	}
}
