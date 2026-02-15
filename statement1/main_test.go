package main

import "testing"

func Test1(t *testing.T) {
	tests := []struct {
		name    string
		x, y    float64
		c       rune
		want    float64
		wanterr bool
	}{
		{"Add", 10, 5, '+', 15, false},
		{"Sub", 10, 8, '-', 2, false},
		{"Mul", 10, 3, '*', 30, false},
		{"Div", 20, 5, '/', 4, false},
		{"Divzero", 90, 0, '/', 0, true},
		{"Otherop", 10, 5, '%', 0, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, e := calc(tc.x, tc.y, tc.c)
			if tc.wanterr && e == nil {
				t.Fatalf("Expected error but got no error")
			}
			if !tc.wanterr && e != nil {
				t.Fatalf("Unexpected error occured: %v", e)
			}
			if r != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, r)
			}
		})
	}
}
