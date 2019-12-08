package main

import "testing"

var cases = []struct {
	mass, want int
}{
	{mass: 12, want: 2},
	{mass: 14, want: 2},
	{mass: 1969, want: 654},
	{mass: 100756, want: 33583},
}

func TestCalcFuel(t *testing.T) {
	for _, c := range cases {
		got := calcFuel(c.mass)
		if got != c.want {
			t.Fatalf("mismatch: got %d; want %d", got, c.want)
		}
	}
}
