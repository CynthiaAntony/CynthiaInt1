package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	a := make(map[string]int)
	t.Run("Add", func(t *testing.T) {
		add("hello", a)
		if _, e := a["hello"]; !e {
			t.Fatalf("Expected 'hello' to be added to list")
		}
	})
	t.Run("Check", func(t *testing.T) {
		if _, e := check("hello", a); e != nil {
			t.Fatalf("Expected 'hello' to be in list")
		}
		if _, e := check("world", a); e == nil {
			t.Fatalf("Expected 'world' to not be in list")
		}
	})
	t.Run("Remove", func(t *testing.T) {
		if _, e := remove("hello", a); e != nil {
			t.Fatalf("Expected 'hello' to be removed without error")
		}
	})
}
