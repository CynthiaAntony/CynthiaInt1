package main

import "testing"

func TestAll(t *testing.T) {
	trie := InitTrie()
	tests := []struct {
		name   string
		action string
		input  string
		want   bool
	}{
		{"Add string", "add", "hello", true},
		{"Check existing string", "check", "hello", true},
		{"Remove string", "remove", "hello", true},
		{"Check non-existing string", "check", "hello", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.action {
			case "add":
				got := trie.Add(tc.input)
				if got != tc.want {
					t.Fatalf("Got = %v, want %v", got, tc.want)
				}
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
