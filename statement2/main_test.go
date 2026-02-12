package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	a := make(map[string]int)
	t.Run("Add", func(t *testing.T) {
		add("hello", a)
		if a["hello"] != 1 {
			t.Fatalf("Expected count of 'hello' to be 1, got %d", a["hello"])
		}
	})
	t.Run("Check", func(t *testing.T) {
		if !check("hello", a) {
			t.Fatalf("Expected 'hello' to be in list")
		}
		if check("world", a) {
			t.Fatalf("Expected 'world' to not be in list")
		}
	})
	t.Run("Remove", func(t *testing.T) {
		remove("hello", a)
		if a["hello"] != 0 {
			t.Fatalf("Expected count of 'hello' to be 0, got %d", a["hello"])
		}
	})
}
