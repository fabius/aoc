package main

import (
	"testing"
)

func TestA(t *testing.T) {
	want := 24000
	have := solveA("../data/1_test.txt")
	if have != want {
		t.Fatalf("want %d != have %d\n", want, have)
	}
}

func TestB(t *testing.T) {
	want := 24000 + 11000 + 10000
	have := solveB("../data/1_test.txt")
	if want != have {
		t.Fatalf("want %v != have %v\n", want, have)
	}
}
